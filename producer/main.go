package main

import (
	"bufio"
	"context"
	"fmt"
	"github.com/SyaibanAhmadRamadhan/go-pub-sub-kafka/producer/infra"
	"github.com/rs/zerolog/log"
	"github.com/segmentio/kafka-go"
	"os"
	"os/signal"
	"regexp"
	"syscall"

	"github.com/SyaibanAhmadRamadhan/go-pub-sub-kafka/producer/internal"
)

func main() {
	infra.Init()
	infra.InitLogger()

	w := internal.KafkaWriter()
	scanner := bufio.NewScanner(os.Stdin)
	ctx, cancel := context.WithCancel(context.Background())

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		log.Info().Msgf("received signal %v, initiating graceful shutdown", sig)
		if err := w.Close(); err != nil {
			log.Err(err).Msg("failed to close writer kafka")
		}
		cancel()
		os.Exit(1)
	}()

	producer(ctx, w, scanner)
}

func isValidEmail(email string) bool {
	const emailRegex = `^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}

func producer(ctx context.Context, w *kafka.Writer, scanner *bufio.Scanner) {
	for {
		fmt.Print("Enter email address ('exit' to exit):")
		scanner.Scan()
		email := scanner.Text()

		if email == "exit" {
			break
		}
		if !isValidEmail(email) {
			log.Warn().Msgf("invalid email address: %s", email)
			continue
		}

		err := internal.WriteMsg(ctx, email, w)
		if err != nil {
			log.Err(err).Msg("failed to write message to kafka")
			continue
		}

		log.Info().Msgf("successfully published email address '%s'", email)
	}
}
