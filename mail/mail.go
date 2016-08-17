package mail

import (
	"errors"
	"log"

	"github.com/gunsluo/blog-mail/common"
	mailgun "github.com/mailgun/mailgun-go"
)

func SendMail(subject string, content string, from string, tos ...string) error {

	c := common.Config().MailGun
	if c == nil {
		return errors.New("mailgun config is nil.")
	}

	gun := mailgun.NewMailgun(c.Domain, c.Key, c.Pubkey)

	m := mailgun.NewMessage(from, subject, content, tos...)
	_, id, err := gun.Send(m)
	if err != nil {
		return err
	}

	log.Printf("send mail Response ID: %s\n", id)
	return nil
}
