package internal

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	mail "github.com/wneessen/go-mail"
	"strings"
	"time"

	"github.com/SyaibanAhmadRamadhan/go-pub-sub-kafka/consumer/infra"
	"github.com/SyaibanAhmadRamadhan/go-pub-sub-kafka/consumer/template"
)

func SendWithGomail(client *mail.Client, payload []byte) (err error) {
	var mailMap map[string]string

	if err = json.Unmarshal(payload, &mailMap); err != nil {
		return fmt.Errorf("failed unmarshal payload | err : %v", err)
	}

	templateBuffer := template.TemplateSendMail(mailMap)

	msg := mail.NewMsg()
	err = msg.FromFormat(infra.Conf.Mail.Name, infra.Conf.Mail.Sender)
	if err != nil {
		return fmt.Errorf("FromFormat: %v", err)
	}

	email := mailMap["to"]
	mailName := strings.Split(email, "@")[0]
	err = msg.AddToFormat(mailName, email)
	if err != nil {
		return fmt.Errorf("AddToFormat: %v", err)
	}

	msg.Subject("subject")
	msg.SetBodyString(mail.TypeTextHTML, templateBuffer.String())

	err = client.DialAndSend(msg)
	if err != nil {
		return fmt.Errorf("DialAndSend: %v", err)
	}

	log.Info().Msg("send mailMap successfully")
	return
}

func RetrySendMailMechanisme(m *mail.Client, message []byte) error {
	var err error
	for i := 0; i < infra.Conf.Mail.MaxRetrySendMail; i++ {
		if err = SendWithGomail(m, message); err == nil {
			return nil
		}
		log.Warn().Err(err).Msgf("failed to send email (attempt %d/%d)", i+1, infra.Conf.Mail.MaxRetrySendMail)
		time.Sleep(infra.Conf.Mail.TimeIntervalRetrySendMail)
	}

	return fmt.Errorf("send email failed after %d retries: %w", infra.Conf.Mail.MaxRetrySendMail, err)
}
