package proc

import (
	nproc "github.com/niean/gotools/proc"
	"log"
)

// trace
var (
	trace *nproc.DataTrace
)

// counter
var (
	HttpRequestCnt = nproc.NewSCounterQps("HttpRequestCnt")

	MailSendCnt    = nproc.NewSCounterQps("MailSendCnt")
	MailSendOkCnt  = nproc.NewSCounterQps("MailSendOkCnt")
	MailSendErrCnt = nproc.NewSCounterQps("MailSendErrCnt")
)

func Start() {
	log.Println("proc.Start, ok")
}

func GetAll() []interface{} {
	ret := make([]interface{}, 0)

	ret = append(ret, HttpRequestCnt)

	ret = append(ret, MailSendCnt)
	ret = append(ret, MailSendOkCnt)
	ret = append(ret, MailSendErrCnt)

	return ret
}
