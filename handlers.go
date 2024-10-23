package aquagram

import (
	"errors"
	"regexp"
)

type Handlers = map[UpdateType][]*Handler
type HandlerFunc[T any] func(bot *Bot, update T) error

type Handler struct {
	Middlewares []Middleware
	Callback    HandlerFunc[any]
}

func (handler *Handler) Use(middlewares ...Middleware) {
	handler.Middlewares = append(handler.Middlewares, middlewares...)
}

func Register(bot *Bot, updateType UpdateType, handler *Handler) *Handler {
	eventHandlers, _ := bot.handlers[updateType]
	eventHandlers = append(eventHandlers, handler)

	bot.handlers[updateType] = eventHandlers
	return handler
}

func handlerFunc[T any](fn HandlerFunc[T]) HandlerFunc[any] {
	return func(bot *Bot, update any) error {
		event, ok := update.(T)
		if !ok {
			return errors.New("handler func: can not convert update (type any) to generic type T")
		}

		return fn(bot, event)
	}
}

func (bot *Bot) OnMessage(handler HandlerFunc[*Message], middlewares ...Middleware) *Handler {
	msgHandler := new(Handler)
	msgHandler.Middlewares = middlewares
	msgHandler.Callback = handlerFunc(handler)

	return Register(bot, OnMessage, msgHandler)
}

func (bot *Bot) OnCommand(command string, handler HandlerFunc[*Message], middlewares ...Middleware) *Handler {
	return bot.OnMessage(handler, CommandMiddleware(command, false))
}

func (bot *Bot) OnRegex(regex *regexp.Regexp, handler HandlerFunc[*Message]) *Handler {
	return bot.OnMessage(handler, RegexMiddleware(regex))
}

func (bot *Bot) OnText(text string, strict bool, caseSensitive bool, handler HandlerFunc[*Message]) *Handler {
	return bot.OnMessage(handler, TextMiddleware(text, strict, caseSensitive))
}

func (bot *Bot) OnCallbackQuery(data string, strict bool, handler HandlerFunc[*CallbackQuery], middlewares ...Middleware) *Handler {
	callbackHandler := new(Handler)
	callbackHandler.Use(middlewares...)
	callbackHandler.Use(CallbackQueryMiddleware(data, strict))
	callbackHandler.Callback = handlerFunc(handler)

	return Register(bot, OnCallbackQuery, callbackHandler)
}
