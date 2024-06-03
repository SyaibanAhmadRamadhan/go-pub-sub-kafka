package internal

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
	"time"

	"github.com/SyaibanAhmadRamadhan/go-pub-sub-kafka/consumer/infra"
)

func KafkaReader() *kafka.Reader {
	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
		SASLMechanism: plain.Mechanism{
			Username: infra.Conf.Application.Kafka.User,
			Password: infra.Conf.Application.Kafka.Pass,
		},
	}

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{infra.Conf.Application.Kafka.Broker},
		GroupID:  infra.Conf.Application.Kafka.Group,
		Topic:    infra.Conf.Application.Kafka.Topic,
		Dialer:   dialer,
		MaxWait:  time.Second,
		MinBytes: 1,
		MaxBytes: 10e6,
	})

	log.Info().Msg("init kafka successfully")
	return r
}

func RetryCommitMsgMechanisme(ctx context.Context, r *kafka.Reader, m kafka.Message) error {
	var err error
	for i := 0; i < infra.Conf.Application.Kafka.MaxRetryCommit; i++ {
		if err = r.CommitMessages(ctx, m); err == nil {
			return nil
		}
		log.Warn().Err(err).Msgf("failed to commit message (attempt %d/%d)", i+1, infra.Conf.Application.Kafka.MaxRetryCommit)
		time.Sleep(infra.Conf.Application.Kafka.RetryIntervalCommit)
	}

	return fmt.Errorf("commit failed after %d retries: %w", infra.Conf.Application.Kafka.MaxRetryCommit, err)
}
