package main

import (
	"fmt"
	"os"

	"github.com/aquagram/aquagram"
)

func StartCommandHandler(bot *aquagram.Bot, message *aquagram.Message) error {
	message.Reply("Hello from Aquagram!", nil)
	return nil
}

func HelloCommandHandler(bot *aquagram.Bot, message *aquagram.Message) error {
	mention := message.From.TextMention(aquagram.ParseModeHTML)
	text := fmt.Sprintf("Hi %s!", mention)

	// a non-nil message is always returned when error is nil.
	msg, err := message.Reply(text, &aquagram.SendMessageParams{
		ParseMode: aquagram.ParseModeHTML,
	})

	fmt.Println(msg.Text)

	return err
}

func main() {
	bot := aquagram.NewBot(os.Getenv("TOKEN"))

	// the command can start with "/" or not
	bot.OnCommand("start", StartCommandHandler)
	bot.OnCommand("/hello", HelloCommandHandler)

	if err := bot.StartPolling(true); err != nil {
		panic(err)
	}
}
