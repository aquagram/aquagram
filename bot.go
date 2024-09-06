package aquagram

import (
	"context"
	"log"
	"net/http"
	"sync"
)

const (
	DefaultApiUrl = "https://api.telegram.org"
)

type Bot struct {
	ApiUrl string
	Me     *User // only available before call getMe

	token string

	Middlewares []Middleware
	handlers    Handlers

	commands        []string
	commandHandlers map[string][]*MessageHandler

	Logger *log.Logger
	Client *http.Client

	mux         sync.Mutex
	stopContext context.Context
	stopFunc    context.CancelFunc

	LastUpdateID   int
	AllowedUpdates []MessageEntity
}

func NewBot(token string) *Bot {
	bot := new(Bot)
	bot.ApiUrl = DefaultApiUrl
	bot.token = token

	bot.handlers = make(Handlers)
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

func (bot *Bot) StartPolling(dropPendingUpdates bool) error {
	options := new(PollingOptions)
	options.DropPendingUpdates = dropPendingUpdates

	return bot.StartPollingWithOptions(options)
}

func (bot *Bot) StartPollingWithOptions(options *PollingOptions) error {
	if err := bot.start(); err != nil {
		return err
	}

	updater := NewPollingUpdater(bot)
	updater.Options = options
	updater.Start()

	return nil
}

func (bot *Bot) StartWebhook(addr string, secretToken string) error {
	if err := bot.start(); err != nil {
		return err
	}

	updater := NewWebhookUpdater(bot)
	updater.secretToken = secretToken

	return updater.Start(addr)
}

func (bot *Bot) Stop() {
	bot.stopFunc()
}

/*
A simple method for testing your bot's authentication token.

https://core.telegram.org/bots/api#getme
*/
func (bot *Bot) GetMe() (User, error) {
	return bot.GetMeWithContext(bot.stopContext)
}

func (bot *Bot) GetMeWithContext(ctx context.Context) (user User, err error) {
	var data []byte
	if data, err = bot.Raw(ctx, "getMe", nil); err != nil {
		return
	}

	if user, err = ParseRawResult[User](bot, data); err != nil {
		return
	}

	bot.Me = &user
	return
}

/*
Use this method to log out from the cloud Bot API server before launching the bot locally.

https://core.telegram.org/bots/api#logout
*/
func (bot *Bot) LogOut() error {
	return bot.LogOutWithContext(bot.stopContext)
}

func (bot *Bot) LogOutWithContext(ctx context.Context) error {
	_, err := bot.Raw(ctx, "logOut", nil)
	return err
}

/*
Use this method to close the bot instance before moving it from one local server to another.

https://core.telegram.org/bots/api#close
*/
func (bot *Bot) Close() error {
	return bot.CloseWithContext(bot.stopContext)
}

func (bot *Bot) CloseWithContext(ctx context.Context) error {
	_, err := bot.Raw(ctx, "close", nil)
	return err
}
