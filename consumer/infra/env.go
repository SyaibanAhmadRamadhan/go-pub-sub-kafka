package infra

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/spf13/viper"
)

var (
	Conf   *Config
	doOnce sync.Once
)

type Config struct {
	Application struct {
		Kafka struct {
			Group               string        `mapstructure:"GROUP"`
			Broker              string        `mapstructure:"BROKER"`
			Topic               string        `mapstructure:"TOPIC"`
			User                string        `mapstructure:"USER"`
			Pass                string        `mapstructure:"PASSWORD"`
			MaxRetryCommit      int           `mapstructure:"MAX_RETRY_COMMIT"`
			RetryIntervalCommit time.Duration `mapstructure:"RETRY_INTERVAL_COMMIT"`
		} `mapstructure:"KAFKA"`
	} `mapstructure:"APPLICATION"`
	Mail struct {
		Host                      string        `mapstructure:"HOST"`
		Port                      int           `mapstructure:"PORT"`
		Username                  string        `mapstructure:"USERNAME"`
		Sender                    string        `mapstructure:"SENDER"`
		Pass                      string        `mapstructure:"PASS"`
		Name                      string        `mapstructure:"NAME"`
		MaxRetrySendMail          int           `mapstructure:"MAX_RETRY_SEND_MAIL"`
		TimeIntervalRetrySendMail time.Duration `mapstructure:"TIME_INTERVAL_RETRY_SEND_MAIL"`
	} `mapstructure:"MAIL"`
}

func Init() {
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf(fmt.Sprintf("cannot read file env : %v", err))
	}

	doOnce.Do(func() {
		if err := viper.Unmarshal(&Conf); err != nil {
			log.Fatalln("cannot unmarsahl config")
		}
	})
}
