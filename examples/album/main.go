package main

import (
	"fmt"
	"os"

	"github.com/aquagram/aquagram"
)

func main() {
	bot := aquagram.NewBot(os.Getenv("TOKEN"))

	documentPath := "examples/document/sample.txt"

	mediaGroup := []aquagram.InputMedia{
		&aquagram.InputMediaDocument{
			Media:   aquagram.InputFileFromPath(documentPath),
			Caption: "sample caption",
		},
		&aquagram.InputMediaDocument{
			Media: aquagram.InputFileFromPath(documentPath),
		},
	}

	messages, err := bot.SendMediaGroup(os.Getenv("CHAT_ID"), mediaGroup, nil)
	if err != nil {
		panic(err)
	}

	for _, message := range messages {
		if message.Document != nil {
			fmt.Println("message_id", message.MessageID)
		}
	}
}
