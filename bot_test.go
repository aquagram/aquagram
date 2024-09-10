package aquagram_test

import (
	"net/http"
	"os"
	"testing"

	"github.com/aquagram/aquagram"
)

func TestGetMe(t *testing.T) {
	bot := aquagram.NewBot(os.Getenv("TOKEN"))

	me, err := bot.GetMe()
	if err != nil {
		t.Error(err)
		return
	}

	if bot.Me == nil || me == nil || me.ID == 0 {
		t.Error(http.StatusText(http.StatusTeapot))
	}
}

func TestGetMyName(t *testing.T) {
	bot := aquagram.NewBot(os.Getenv("TOKEN"))

	botName, err := bot.GetMyName("")
	if err != nil {
		t.Error(err)
		return
	}

	if botName == nil || botName.Name == "" {
		t.Error("bot name should be not empty")
	}
}

func TestSetMyName(t *testing.T) {
	bot := aquagram.NewBot(os.Getenv("TOKEN"))

	botName, err := bot.GetMyName("")
	if err != nil {
		t.Error(err)
		return
	}

	if err := bot.SetMyName(botName.Name, ""); err != nil {
		t.Error(err)
	}
}
