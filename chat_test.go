package aquagram_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/aquagram/aquagram"
)

func TestGetChatAdministrators(t *testing.T) {
	bot := aquagram.NewBot(os.Getenv("TOKEN"))

	admins, err := bot.GetChatAdministrators(os.Getenv("CHAT_ID"))
	if err != nil {
		t.Error(err)
		return
	}

	if len(admins) == 0 {
		t.Error("chat administrators array with length 0")
	}
}

func TestGetChatMemberCount(t *testing.T) {
	bot := aquagram.NewBot(os.Getenv("TOKEN"))

	count, err := bot.GetChatMemberCount(os.Getenv("CHAT_ID"))
	if err != nil {
		t.Error(err)
		return
	}

	if count == 0 {
		t.Error("chat members count should be greater than 0")
	}
}

func TestGetChatMember(t *testing.T) {
	bot := aquagram.NewBot(os.Getenv("TOKEN"))

	userIDRaw := os.Getenv("USER_ID")
	userID, _ := strconv.ParseInt(userIDRaw, 10, 64)

	member, err := bot.GetChatMember(os.Getenv("CHAT_ID"), userID)
	if err != nil {
		t.Error(err)
		return
	}

	if member.User.ID <= 0 {
		t.Error("members ID should be greater than 0")
	}
}
