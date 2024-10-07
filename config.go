package aquagram

import (
	"time"
)

type Config struct {
	// API URL, by default is https://api.telegram.org
	API string

	// A list of the update types you want your bot to receive.
	AllowedUpdates []UpdateType

	// Time to wait between errors.
	//
	// By default is 1s
	RetriesInterval time.Duration
}

func DefaultConfig() *Config {
	config := new(Config)

	config.API = "https://api.telegram.org"
	config.RetriesInterval = time.Second

	return config
}
