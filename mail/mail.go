package mail

import (
	"fmt"

	mailgun "github.com/mailgun/mailgun-go"
)

func SendMail(subject string, content string, from string, tos ...string) error {

	//postmaster@jerrylou.me
	gun := mailgun.NewMailgun("jerrylou.me", "key-535755ffb130fb94d36c88f19a67e32c", "pubkey-31911d1d3ccd3d143c2b4e91d0a7bc9a")

	m := mailgun.NewMessage(from, subject, content, tos...)
	response, id, err := gun.Send(m)
	fmt.Printf("err: %v\n", err)
	fmt.Printf("Response ID: %s\n", id)
	fmt.Printf("Message from server: %s\n", response)
}
