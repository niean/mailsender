package http

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/niean/mailsender/g"
	"github.com/niean/mailsender/proc"
	"github.com/niean/mailsender/sender"
)

func configMailSenderApiRoutes() {
	http.HandleFunc("/mail/sender", func(w http.ResponseWriter, req *http.Request) {
		// statistics
		proc.HttpRequestCnt.Incr()

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if req.Method != "POST" {
			RenderDataJson(w, fmt.Sprintf("%s not supported", req.Method))
			return
		}

		cfg := g.GetConfig()
		req.ParseForm()
		params := req.Form
		content, exist := params["content"]
		if !exist || len(content[0]) < 1 {
			if len(cfg.Mail.Content) > 0 {
				content = append(content, cfg.Mail.Content)
			} else {
				RenderDataJson(w, "bad content")
				return
			}
		}

		subject, exist := params["subject"]
		if !exist || len(subject[0]) < 1 {
			if len(cfg.Mail.Subject) > 0 {
				subject = append(subject, cfg.Mail.Subject)
			} else {
				RenderDataJson(w, "bad subject")
				return
			}
		}

		tos, exist := params["tos"]
		if !exist || len(tos[0]) < 1 {
			if len(cfg.Mail.Tos) > 0 {
				tos = append(tos, cfg.Mail.Tos)
			} else {
				RenderDataJson(w, "bad tos")
				return
			}
		}

		from := []string{}
		fromUser, exist := params["user"]
		if exist && len(fromUser[0]) > 0 {
			from = append(from, fromUser[0])
		} else {
			if len(cfg.Mail.User) > 0 {
				from = append(from, cfg.Mail.User)
			}
		}
		if ok := sender.AddMail(strings.Split(tos[0], ","), subject[0], content[0], from...); !ok {
			RenderDataJson(w, "error, service busy")
			return
		}
		RenderDataJson(w, "ok")
	})
}
