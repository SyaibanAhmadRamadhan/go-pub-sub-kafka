package internal

import (
	"log"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"

	"github.com/SyaibanAhmadRamadhan/go-pub-sub-kafka/consumer/infra"
)

func KafkaReader() *kafka.Reader {
	dialer := &kafka.Dialer{
		Timeout:   10 * time.Second,
		DualStack: true,
		SASLMechanism: plain.Mechanism{
			Username: infra.Get().Application.Kafka.User,
			Password: infra.Get().Application.Kafka.Pass,
		},
	}

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{infra.Get().Application.Kafka.Broker},
		GroupID:  infra.Get().Application.Kafka.Group,
		Topic:    infra.Get().Application.Kafka.Topic,
		Dialer:   dialer,
		MaxWait:  time.Second,
		MinBytes: 1,
		MaxBytes: 10e6,
	})

	log.Println("init kafka successfully")
	return r
}
