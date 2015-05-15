package sender

import (
	nsema "github.com/niean/gotools/concurrent/semaphore"
	nlist "github.com/niean/gotools/container/list"
	gomail "github.com/niean/mail/gomail"
	"github.com/niean/mailsender/g"
	"github.com/niean/mailsender/proc"
	"log"
	"time"
)

var (
	mailqueue *nlist.SafeListLimited
	sendsema  *nsema.Semaphore
)

func Start() {
	sendsema = nsema.NewSemaphore(g.GetConfig().Mail.SendConcurrent)
	mailqueue = nlist.NewSafeListLimited(g.GetConfig().Mail.MaxQueueSize)
	go startSender()
}

// try pushing one mail into sender queue, maybe failed
func AddMail(r []string, subject string, content string, from ...string) bool {
	mcfg := g.GetConfig().Mail
	fromUserName := mcfg.FromUser
	if len(from) == 1 {
		fromUserName = from[0]
	}

	nm := NewMailObject(r, subject, content, fromUserName)
	return mailqueue.PushFront(nm)
}

// sender cron
func startSender() {
	for {
		raw := mailqueue.PopBack()
		if raw == nil {
			time.Sleep(time.Duration(10) * time.Millisecond)
			continue
		}
		// control sending concurrents
		sendsema.Acquire()
		go func(mailObject *MailObject) {
			defer sendsema.Release()
			sendMail(mailObject)
		}(raw.(*MailObject))
	}
}

func sendMail(mo *MailObject) {
	mcfg := g.GetConfig().Mail
	msg := gomail.NewMessage()
	// from
	msg.SetAddressHeader("From", mcfg.MailServerAccount, mo.FromUser)
	// receivers
	for i, to := range mo.Receivers {
		if i == 0 {
			msg.SetHeader("To", to)
		} else {
			msg.AddHeader("To", to)
		}
	}
	// subject
	msg.SetHeader("Subject", mo.Subject)
	// content
	msg.SetBody("text/plain", mo.Content)

	// statistics
	proc.MailSendCnt.Incr()

	m := gomail.NewMailer(mcfg.MailServerHost, mcfg.MailServerAccount, mcfg.MailServerPasswd, mcfg.MailServerPort)
	if err := m.Send(msg); err != nil {
		// statistics
		proc.MailSendErrCnt.Incr()
		log.Println(err, ", mailObject:", mo)
	} else {
		// statistics
		proc.MailSendOkCnt.Incr()
	}
}

// Mail Content Struct
type MailObject struct {
	Receivers []string
	Subject   string
	Content   string
	FromUser  string
}

func NewMailObject(receivers []string, subject string, content string, fromUserName string) *MailObject {
	return &MailObject{Receivers: receivers, Subject: subject, Content: content,
		FromUser: fromUserName}
}
