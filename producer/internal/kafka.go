package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/SyaibanAhmadRamadhan/go-pub-sub-kafka/producer/infra"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"
	"time"
)

func KafkaWriter() *kafka.Writer {
	w := &kafka.Writer{
		Addr:  kafka.TCP(infra.Conf.Application.Kafka.Broker),
		Topic: infra.Conf.Application.Kafka.Topic,
		Transport: &kafka.Transport{
			SASL: plain.Mechanism{
				Username: infra.Conf.Application.Kafka.User,
				Password: infra.Conf.Application.Kafka.Pass,
			},
		},
		AllowAutoTopicCreation: false,
	}

	return w
}

type WriteMsgInput struct {
	Data any
}

func WriteMsg(ctx context.Context, input WriteMsgInput, w *kafka.Writer) error {

	kafkaMsg, err := json.Marshal(input.Data)

	if err != nil {
		return fmt.Errorf("marshalling message: %w", err)
	}

	for i := 0; i < infra.Conf.Application.Kafka.MaxRetryWriteMsg; i++ {
		ctxTimeOut, cancel := context.WithTimeout(ctx, 10*time.Second)

		if err = w.WriteMessages(ctxTimeOut, kafka.Message{
			Value: kafkaMsg,
		}); err == nil {
			cancel()
			return nil
		} else {
			log.Warn().Err(err).Msgf("failed to write message (attempt %d/%d)", i+1, infra.Conf.Application.Kafka.MaxRetryWriteMsg)
			time.Sleep(infra.Conf.Application.Kafka.TimeIntervalRetryWriteMsg)
		}

		cancel()
	}

	return fmt.Errorf("failed to write message after %d retries", infra.Conf.Application.Kafka.MaxRetryWriteMsg)
}
