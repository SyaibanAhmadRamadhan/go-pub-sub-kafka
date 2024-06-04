package main

import (
	"context"
	"fmt"
	"github.com/SyaibanAhmadRamadhan/go-pub-sub-kafka/consumer/infra"
	"github.com/SyaibanAhmadRamadhan/go-pub-sub-kafka/consumer/internal"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"github.com/wneessen/go-mail"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	infra.Init()
	infra.InitLogger()
	r := internal.KafkaReader()
	mailing := internal.NewMailtrap()
	ctx, cancel := context.WithCancel(context.Background())

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		log.Info().Msgf("received signal %v, initiating graceful shutdown", sig)
		if err := r.Close(); err != nil {
			log.Err(err).Msg("failed to close reader")
		}
		if err := mailing.Close(); err != nil {
			log.Err(err).Msg("failed to close mailing connection")
		}
		cancel()
		os.Exit(1)
	}()

	consume(ctx, mailing, r)
}

func consume(ctx context.Context, client *mail.Client, r *kafka.Reader) {
	for {
		m, err := r.FetchMessage(ctx)
		if err != nil {
			log.Err(err).Msg("failed to fetch message")
			continue
		}
		formatMsg := fmt.Sprintf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		log.Info().Msg(formatMsg)

		err = internal.RetrySendMailMechanism(client, m.Value)
		if err != nil {
			log.Err(err).Msgf("failed to send mail")
			// DLQ process or any action
			continue
		}

		if err = internal.RetryCommitMsgMechanisme(ctx, r, m); err != nil {
			log.Err(err).Msg("failed to commit message to Kafka")
			// DLQ process or any action
			continue
		}

		log.Info().Msgf("successfully committed message to Kafka and send mail")
	}
}
