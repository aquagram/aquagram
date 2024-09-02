package aquagram

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type UpdatesMethod string

const (
	DefaultApiUrl = "https://api.telegram.org"
	MethodPolling = "polling"
	MethodWebhook = "webhook"
)

type Bot struct {
	ApiUrl string
	Me     *User // only available before call getMe

	token string

	Middlewares     []Middleware
	handlers        Handlers
	messageHandlers map[UpdateType][]*MessageHandler

	commands        []string
	commandHandlers map[string][]*MessageHandler

	Logger *log.Logger
	Client *http.Client

	mux         sync.Mutex
	stopContext context.Context
	stopFunc    context.CancelFunc
}

func NewBot(token string) *Bot {
	bot := new(Bot)
	bot.ApiUrl = DefaultApiUrl

	bot.token = token

	bot.handlers = make(Handlers)
	bot.messageHandlers = make(map[UpdateType][]*MessageHandler)

	bot.Client = new(http.Client)

	bot.Logger = log.Default()
	bot.Logger.SetPrefix("[aquagram]: ")

	bot.stopContext, bot.stopFunc = context.WithCancel(context.Background())

	return bot
}

func (bot *Bot) start() error {
	if bot.token == EmptyString {
		return ErrEmptyToken
	}

	_, err := bot.GetMe()
	if err != nil {
		return err
	}

	return nil
}

func (bot *Bot) StartPolling() error {
	if err := bot.start(); err != nil {
		return err
	}

	updater := NewPollingUpdater(bot)
	updater.Start(bot.stopContext)

	return nil
}

func (bot *Bot) StartWebhook() error {
	if err := bot.start(); err != nil {
		return err
	}

	return nil
}

func (bot *Bot) Stop() {
	bot.stopFunc()
}

func (bot *Bot) GetMe() (*User, error) {
	return bot.GetMeWithContext(bot.stopContext)
}

func (bot *Bot) GetMeWithContext(ctx context.Context) (*User, error) {
	data, err := bot.Raw(ctx, "getMe", nil)
	if err != nil {
		return nil, err
	}

	user, err := ParseRawResult[User](data)
	if err != nil {
		return nil, err

	}

	bot.Me = user

	return user, nil
}

func (bot *Bot) methodUrl(method string) string {
	return fmt.Sprintf("%s/bot%s/%s", bot.ApiUrl, bot.token, method)
}
