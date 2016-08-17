package http

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gunsluo/blog-mail/common"
	"github.com/gunsluo/blog-mail/mail"
	"github.com/pquerna/ffjson/ffjson"
)

type MailRequest struct {
	Subject string `json:"subject"`
	Content string `json:"content"`
	From    string `json:"from"`
}

func get(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	var subject, content, from string
	if len(r.Form["subject"]) > 0 {
		subject = r.Form["subject"][0]
	}
	if len(r.Form["content"]) > 0 {
		content = r.Form["content"][0]
	}
	if len(r.Form["from"]) > 0 {
		from = r.Form["from"][0]
	}

	err := handleMail(subject, content, from)
	if err != nil {
		RenderFailedDataJson(w, err.Error())
	} else {
		RenderSuccessDataJson(w)
	}

}

func post(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		RenderFailedDataJson(w, err.Error())
		return
	}

	var mailParam MailRequest
	err = ffjson.Unmarshal(body, &mailParam)
	if err != nil {
		RenderFailedDataJson(w, err.Error())
		return
	}

	err = handleMail(mailParam.Subject, mailParam.Content, mailParam.From)
	if err != nil {
		RenderFailedDataJson(w, err.Error())
	} else {
		RenderSuccessDataJson(w)
	}

}

func handleMail(subject string, content string, from string) error {

	if subject == "" || content == "" || from == "" {
		return errors.New("param is incorrect")
	}

	c := common.Config().MailGun
	err := mail.SendMail(subject, content, from, c.To)
	if err != nil {
		return err
	}

	return nil
}
