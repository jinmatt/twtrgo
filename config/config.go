package config

import (
	"os"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

type Config struct {
	Environment              string `env:"TWTRGO_ENV" envDefault:"default"`
	Port                     string `env:"PORT" envDefault:"8080"`
	TwitterConsumerKey       string `env:"TWITTER_CONSUMER_KEY"`
	TwitterConsumerSecret    string `env:"TWITTER_CONSUMER_SECRET"`
	TwitterAccessToken       string `env:"TWITTER_ACCESS_TOKEN"`
	TwitterAccessTokenSecret string `env:"TWITTER_ACCESS_TOKEN_SECRET"`
}

func NewConfig() (*Config, error) {

	// set exports from .env file if running local or with default environment
	if os.Getenv("TWTRGO_ENV") == "" || os.Getenv("TWTRGO_ENV") == "default" {
		err := godotenv.Load()
		if err != nil {
			return nil, err
		}
	}

	config := &Config{}
	err := env.Parse(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
