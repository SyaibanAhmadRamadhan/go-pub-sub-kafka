package internal

import (
	"github.com/SyaibanAhmadRamadhan/go-pub-sub-kafka/consumer/infra"
	"github.com/rs/zerolog/log"
	"github.com/wneessen/go-mail"
)

func NewMailtrap() *mail.Client {
	client, err := mail.NewClient(infra.Conf.Mail.Host,
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithPort(infra.Conf.Mail.Port),
		mail.WithUsername(infra.Conf.Mail.Username),
		mail.WithPassword(infra.Conf.Mail.Pass),
		mail.WithTLSPortPolicy(mail.TLSMandatory),
	)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create mail client")
	}

	return client
}
