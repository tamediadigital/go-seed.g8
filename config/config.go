package config

import (
	"log"

	"github.com/caarlos0/env"
)

type Config struct {
	KafkaBrokers       []string `env:"KAFKA_BROKERS,required" envSeparator:","`
	KafkaTopic         string   `env:"KAFKA_TOPIC,required"`
	KafkaConsumerGroup string   `env:"KAFKA_CONSUMER_GROUP,required"`
	PrometheusEndpoint string   `env:"PROMETHEUS_ENDPOINT,required"`
	RedisHost          string   `env:"REDIS_HOST,required"`
	RedisPort          int      `env:"REDIS_PORT,required"`
	RedisDB            int      `env:"REDIS_DB,required"`
	RedisUsersKey      string   `env:"REDIS_USERS_KEY,required"`
}

func LoadConfiguration() Config {
	config := Config{}

	if err := env.Parse(&config); err != nil {
		log.Fatal("Couldn't load environment variables! ", err)
	}
	return config
}
