package main

import (
	"context"
	"fmt"
	"log"

	"github.com/SyaibanAhmadRamadhan/go-pub-sub-kafka/consumer/internal"
)

func main() {
	r := internal.KafkaReader()
	defer func() {
		if err := r.Close(); err != nil {
			log.Fatalf("failed close reader kafka | err %v", err)
		}
	}()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			log.Fatalf("failed read message from topic | err %v", err)
		}

		internal.SendWithGomail(m.Value)

		formatMsg := fmt.Sprintf("message at topic/partition/offset %v/%v/%v: %s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		log.Println(formatMsg)
	}
}
