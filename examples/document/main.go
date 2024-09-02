package main

import (
	"fmt"
	"os"

	"github.com/aquagram/aquagram"
)

func main() {
	bot := aquagram.NewBot(os.Getenv("TOKEN"))

	document := new(aquagram.InputFile)
	document.FromPath = "examples/document/sample.txt"

	message, err := bot.SendDocument(os.Getenv("CHAT_ID"), document, nil)
	if err != nil {
		panic(err)
	}

	if message.Document != nil {
		fmt.Println("file_id:", message.Document.FileID)
	}
}
