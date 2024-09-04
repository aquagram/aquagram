package main

import (
	"os"

	"github.com/aquagram/aquagram"
)

var (
	WhiteList = []int64{12345678, 12345678, 12345678}
)

func StartCommandHandler(bot *aquagram.Bot, message *aquagram.Message) error {
	_, err := message.Reply(bot, "Hello from Aquagram!", nil)
	return err
}

func main() {
	bot := aquagram.NewBot(os.Getenv("TOKEN"))

	// global middleware
	bot.Use(aquagram.UsersMiddleware(WhiteList...))

	// handler scoped middleware
	bot.OnCommand("start", StartCommandHandler, aquagram.UsersMiddleware(WhiteList...))

	if err := bot.StartPolling(true); err != nil {
		panic(err)
	}
}
