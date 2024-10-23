package aquagram

import (
	"fmt"
	"regexp"
)

type (
	Middleware     func(next MiddlewareFunc) MiddlewareFunc
	MiddlewareFunc func(bot *Bot, event Event) error

	ErrorFunc func(bot *Bot, err error)
)

func (bot *Bot) Use(middlewares ...Middleware) {
	bot.Middlewares = append(bot.Middlewares, middlewares...)
}

func BlackListMiddleware(ids *[]int64) Middleware {
	return BuildMiddleware(WhiteListFilter(ids))
}

func CallbackQueryMiddleware(data string, strict bool) Middleware {
	return BuildMiddleware(CallbackQueryFilter(data, strict))
}

func ChatMemberMiddleware(chatID string) Middleware {
	return BuildMiddleware(ChatMemberFilter(chatID))
}

func CommandMiddleware(command string, strict bool) Middleware {
	return BuildMiddleware(CommandFilter(command))
}

func RecoverMiddleware(errorFunc ErrorFunc) Middleware {
	return func(next MiddlewareFunc) MiddlewareFunc {
		return func(bot *Bot, event Event) error {
			defer func() {
				// Currently, err is always nil and I dont know why.
				// However, it stops the the panicking sequence, so ¯⁠\⁠_⁠(⁠ツ⁠)⁠_⁠/⁠¯
				err := recover()

				if errorFunc == nil {
					bot.Logger.Println("recovered from panic", err)
					return
				}

				switch v := err.(type) {
				case error:
					errorFunc(bot, v)
				default:
					errorFunc(bot, fmt.Errorf("%v", err))
				}
			}()

			return next(bot, event)
		}
	}
}

func RegexMiddleware(regex *regexp.Regexp) Middleware {
	return BuildMiddleware(RegexFilter(regex))
}

func TextMiddleware(text string, strict bool, caseSensitive bool) Middleware {
	return BuildMiddleware(TextFilter(text, strict, caseSensitive))
}

func WhiteListMiddleware(ids *[]int64) Middleware {
	return BuildMiddleware(WhiteListFilter(ids))
}
