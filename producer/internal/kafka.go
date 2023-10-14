package internal

import (
	"context"
	"encoding/json"
	"log"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/sasl/plain"

	"github.com/SyaibanAhmadRamadhan/go-pub-sub-kafka/producer/infra"
)

func KafkaWriter() *kafka.Writer {
	w := &kafka.Writer{
		Addr:  kafka.TCP(infra.Get().Application.Kafka.Broker),
		Topic: infra.Get().Application.Kafka.Topic,
		Transport: &kafka.Transport{
			SASL: plain.Mechanism{
				Username: infra.Get().Application.Kafka.User,
				Password: infra.Get().Application.Kafka.Pass,
			},
		},
	}

	return w
}

func WriteMsg(ctx context.Context, mail string, w *kafka.Writer) {
	msg := map[string]string{
		"value": "test-value",
		"to":    mail,
	}

	kafkaMsg, err := json.Marshal(msg)

	if err != nil {
		log.Fatal("failed marshal map")
	}

	if err = w.WriteMessages(ctx, kafka.Message{
		Value: kafkaMsg,
	}); err != nil {
		log.Fatalf("failed write message to kafka | err : %v", err)
	}

	log.Println("success publish your message | mail : ", mail)
}
