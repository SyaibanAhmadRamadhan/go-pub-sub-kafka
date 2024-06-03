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
			Broker                    string        `mapstructure:"BROKER"`
			Topic                     string        `mapstructure:"TOPIC"`
			User                      string        `mapstructure:"USER"`
			Pass                      string        `mapstructure:"PASSWORD"`
			MaxRetryWriteMsg          int           `mapstructure:"MAX_RETRY_WRITE_MSG"`
			TimeIntervalRetryWriteMsg time.Duration `mapstructure:"TIME_INTERVAL_RETRY_WRITE_MSG"`
		} `mapstructure:"KAFKA"`
	} `mapstructure:"APPLICATION"`
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
