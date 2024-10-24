package aquagram

import (
	"log"
	"net/http"
	"time"
)

type (
	ErrorFunc func(bot *Bot, err error)
	StartFunc func(bot *Bot)
)

type Config struct {
	// API URL, by default is https://api.telegram.org
	API string

	// A list of the update types you want your bot to receive.
	AllowedUpdates []UpdateType

	// HTTP Client using to perform API requests
	Client *http.Client

	// ParseMode that the bot will use wherever
	// necessary unless specified otherwise.
	//
	// NOTE: It is not fully implemented yet
	DefaultParseMode ParseMode

	Logger *log.Logger

	// Function called when an error occurs in the bot
	OnErrorFunc ErrorFunc

	// Function called when the bot is started
	//
	// It is called after [GetMe] execution,
	// before starting the updater
	OnStartFunc StartFunc

	// Time to wait between errors.
	//
	// By default is 1s
	RetriesInterval time.Duration
}

func DefaultConfig() *Config {
	config := new(Config)

	config.API = "https://api.telegram.org"
	config.Client = new(http.Client)
	config.DefaultParseMode = ParseModeDisabled

	config.Logger = log.Default()
	config.Logger.SetPrefix("[aquagram]: ")

	config.OnErrorFunc = func(bot *Bot, err error) {
		bot.Config.Logger.Println(err)
	}

	config.RetriesInterval = time.Second

	return config
}
