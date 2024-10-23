package aquagram

import (
	"context"
	"log"
	"net/http"
)

type Bot struct {
	// User-configurable properties can be found here.
	Config *Config

	// Updated on each execution of GetMe, by default is a nil-pointer.
	Me *User

	token string

	Middlewares []Middleware
	handlers    Handlers

	Logger *log.Logger
	Client *http.Client

	stopContext context.Context
	stopFunc    context.CancelFunc

	LastUpdateID int
}

func NewBot(token string) *Bot {
	bot := new(Bot)
	bot.Config = DefaultConfig()

	bot.token = token

	bot.handlers = make(Handlers)
	bot.Client = new(http.Client)

	bot.Logger = log.Default()
	bot.Logger.SetPrefix("[aquagram]: ")

	bot.stopContext, bot.stopFunc = context.WithCancel(context.Background())

	return bot
}

func (bot *Bot) Context() context.Context {
	return bot.stopContext
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
func (bot *Bot) GetMe() (*User, error) {
	return bot.GetMeWithContext(bot.stopContext)
}

func (bot *Bot) GetMeWithContext(ctx context.Context) (*User, error) {
	data, err := bot.Raw(ctx, "getMe", nil)
	if err != nil {
		return nil, err
	}

	user, err := ParseRawResult[*User](bot, data)
	if err != nil {
		return nil, err
	}

	bot.Me = user

	return user, nil
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

func (bot *Bot) SetMyName(name string, languageCode string) error {
	return bot.SetMyNameWithContext(bot.stopContext, name, languageCode)
}

func (bot *Bot) SetMyNameWithContext(ctx context.Context, name string, languageCode string) error {
	params := map[string]string{
		"name":          name,
		"language_code": languageCode,
	}

	data, err := bot.Raw(ctx, "setMyName", params)
	if err != nil {
		return err
	}

	success, err := ParseRawResult[bool](bot, data)
	if err != nil {
		return err
	}

	if !success {
		return ErrExpectedTrue
	}

	return nil
}

func (bot *Bot) GetMyName(languageCode string) (*BotName, error) {
	return bot.GetMyNameWithContext(bot.stopContext, languageCode)
}

func (bot *Bot) GetMyNameWithContext(ctx context.Context, languageCode string) (*BotName, error) {
	data, err := bot.Raw(ctx, "getMyName", nil)
	if err != nil {
		return nil, err
	}

	return ParseRawResult[*BotName](bot, data)
}
