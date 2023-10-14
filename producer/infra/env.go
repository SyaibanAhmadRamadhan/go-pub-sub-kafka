package infra

import (
	"fmt"
	"log"
	"sync"

	"github.com/spf13/viper"
)

var (
	cfg    Config
	doOnce sync.Once
)

type Config struct {
	Application struct {
		Kafka struct {
			Broker string `mapstructure:"BROKER"`
			Topic  string `mapstructure:"TOPIC"`
			User   string `mapstructure:"USER"`
			Pass   string `mapstructure:"PASSWORD"`
		} `mapstructure:"KAFKA"`
	} `mapstructure:"APPLICATION"`
}

func Get() Config {
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf(fmt.Sprintf("cannot read file env : %v", err))
	}

	doOnce.Do(func() {
		if err := viper.Unmarshal(&cfg); err != nil {
			log.Fatalln("cannot unmarsahl config")
		}
	})

	return cfg
}
