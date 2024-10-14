package aquagram

type MessageHandler = func(bot *Bot, message *Message) error
type CallbackQueryHandlerCallback = func(bot *Bot, callback *CallbackQuery) error

func (bot *Bot) OnEvent(updateType UpdateType, handler *Handler) {
	eventHandlers, _ := bot.handlers[updateType]
	eventHandlers = append(eventHandlers, handler)

	bot.handlers[updateType] = eventHandlers
}

func AddHandlerMiddlewares(handler *Handler, userDefined []Middleware, priorized ...Middleware) {
	for _, middleware := range priorized {
		handler.Middlewares = append(handler.Middlewares, middleware)
	}

	handler.Middlewares = append(handler.Middlewares, userDefined...)
}

func (bot *Bot) OnMessage(callback MessageHandler, middlewares ...Middleware) *Handler {
	handler := new(Handler)
	handler.Middlewares = middlewares

	handler.Callback = func(bot *Bot, update any) error {
		message, ok := update.(*Message)
		if !ok {
			return nil
		}

		return callback(bot, message)
	}

	bot.OnEvent(OnMessage, handler)
	return handler
}

func (bot *Bot) OnCommand(command string, callback MessageHandler, middlewares ...Middleware) *Handler {
	handler := new(Handler)

	AddHandlerMiddlewares(handler, middlewares, CommandMiddleware(command))

	handler.Callback = func(bot *Bot, update any) error {
		message, ok := update.(*Message)
		if !ok {
			return nil
		}

		return callback(bot, message)
	}

	bot.OnEvent(OnMessage, handler)

	return handler
}

func (bot *Bot) OnCallbackQuery(callbackData string, callback CallbackQueryHandlerCallback, middlewares ...Middleware) *Handler {
	handler := new(Handler)

	callbackMiddleware := CallbackQueryMiddleware(callbackData)
	AddHandlerMiddlewares(handler, middlewares, callbackMiddleware)

	handler.Callback = func(bot *Bot, update any) error {
		callbackQuery, ok := update.(*CallbackQuery)
		if !ok {
			return nil
		}

		return callback(bot, callbackQuery)
	}

	bot.OnEvent(OnCallbackQuery, handler)

	return handler
}
