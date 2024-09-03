package aquagram

import (
	"slices"
	"strings"
)

type Middleware func(event Event) error

func (bot *Bot) Use(middlewares ...Middleware) {
	for _, middleware := range middlewares {
		bot.Middlewares = append(bot.Middlewares, middleware)
	}
}

func UsersMiddleware(ids ...int64) Middleware {
	return func(event Event) error {
		from := event.GetFrom()
		if from == nil {
			return ErrStopPropagation
		}

		if !slices.Contains(ids, from.ID) {
			return ErrStopPropagation
		}

		return nil
	}
}

func CommandMiddleware(bot *Bot, command string) Middleware {
	slash := "/"

	if !strings.HasPrefix(command, slash) {
		command = slash + command
	}

	return func(event Event) error {
		message := event.GetMessage()
		if message == nil {
			return ErrStopPropagation
		}

		if message.Text == "" {
			return ErrStopPropagation
		}

		var commandText string

		entities := event.GetEntities()
		for _, entity := range entities {
			if entity.Type != EntityTypeBotCommand || entity.Offset != 0 {
				continue
			}

			commandText = message.Text[entity.Offset:entity.Length]
			break
		}

		if commandText == "" {
			return ErrStopPropagation
		}

		atIndex := strings.Index(commandText, "@")
		if atIndex != -1 {
			commandText = commandText[:atIndex]
		}

		if commandText != command {
			return ErrStopPropagation
		}

		return nil
	}
}

func CallbackQueryMiddleware(callback string) Middleware {
	return func(event Event) error {
		callbackQuery := event.GetCallbackQuery()

		if callbackQuery == nil {
			return ErrStopPropagation
		}

		if callbackQuery.Data != callback {
			return ErrStopPropagation
		}

		return nil
	}
}
