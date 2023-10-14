package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/SyaibanAhmadRamadhan/go-pub-sub-kafka/producer/internal"
)

func main() {
	w := internal.KafkaWriter()
	defer func() {
		if err := w.Close(); err != nil {
			log.Printf("failed close kafka writer | err %v", err)
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	for {
		fmt.Print("Enter email address ('exit' to exit):")
		scanner.Scan()
		email := scanner.Text()

		if email == "exit" {
			break
		}

		internal.WriteMsg(ctx, email, w)

	}

}
