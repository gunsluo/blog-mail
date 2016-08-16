package http

import (
	"net/http"

	"github.com/gunsluo/blog-mail/mail"
)

func configRoutes() {

	http.HandleFunc("/mail", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		/*
			var tos, subject, content string
			if r.Method == "GET" {
				tos = r.Form["tos"][0]
				subject = r.Form["subject"][0]
				content = r.Form["content"][0]
			} else if r.Method == "POST" {
				tos = r.PostFormValue("tos")
				subject = r.PostFormValue("subject")
				content = r.PostFormValue("content")
			}
		*/

		mail.SendMail("", "", "", "")
		RenderDataJson(w, "")
	})

}
