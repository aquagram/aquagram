package main

import (
	"fmt"
	"os"

	"github.com/aquagram/aquagram"
)

func StartCommandHandler(bot *aquagram.Bot, message *aquagram.Message) error {
	message.Reply(bot, "Hello from Aquagram!", nil)
	return nil
}

func HelloCommandHandler(bot *aquagram.Bot, message *aquagram.Message) error {
	mention := message.From.TextMention(aquagram.ParseModeHtml)
	text := fmt.Sprintf("Hi %s!", mention)

	// a non-nil message is always returned when error is nil.
	msg, err := message.Reply(bot, text, &aquagram.SendMessageParams{
		ParseMode: aquagram.ParseModeHtml,
	})

	fmt.Println(msg.Text)

	return err
}

func main() {
	bot := aquagram.NewBot(os.Getenv("TOKEN"))

	// the command can start with "/" or not
	bot.OnCommand("start", StartCommandHandler)
	bot.OnCommand("/hello", HelloCommandHandler)

	if err := bot.StartPolling(); err != nil {
		panic(err)
	}
}
