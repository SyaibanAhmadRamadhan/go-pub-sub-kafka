package internal

import (
	"encoding/json"
	"log"

	"gopkg.in/gomail.v2"

	"github.com/SyaibanAhmadRamadhan/go-pub-sub-kafka/consumer/infra"
	"github.com/SyaibanAhmadRamadhan/go-pub-sub-kafka/consumer/template"
)

func SendWithGomail(payload []byte) {
	var mail map[string]string

	if err := json.Unmarshal(payload, &mail); err != nil {
		log.Fatalf("failed unmarshal payload | err : %v", err)
	}

	templateBuffer := template.TemplateSendMail(mail)

	mailer := gomail.NewMessage()
	mailer.SetHeader("From", infra.Get().Mail.Sender)
	mailer.SetAddressHeader("Cc", infra.Get().Mail.Address, infra.Get().Mail.Address)
	mailer.SetHeader("Subject", "hello")
	mailer.SetBody("text/html", templateBuffer.String())

	dialer := gomail.NewDialer(
		infra.Get().Mail.Host,
		infra.Get().Mail.Port,
		infra.Get().Mail.Address,
		infra.Get().Mail.Pass,
	)

	err := dialer.DialAndSend(mailer)
	if err != nil {
		log.Fatalf("failed dial and send mail | err %v", err)
	}

	log.Println("send mail successfully")
}
