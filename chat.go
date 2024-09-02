package aquagram

import (
	"fmt"
)

type ChatType string

const (
	ChatTypePrivate    ChatType = "private"
	ChatTypeGroup      ChatType = "group"
	ChatTypeSuperGroup ChatType = "supergroup"
	ChatTypeChannel    ChatType = "channel"
)

type Chat struct {
	ID        int64    `json:"id"`
	Type      ChatType `json:"type"`
	Title     string   `json:"title"`
	Username  string   `json:"username"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	IsForum   bool     `json:"is_forum"`
}

func ChatID(id int64) string {
	return fmt.Sprintf("%d", id)
}

func (chat *Chat) IsPrivate() bool {
	return chat.Type == ChatTypePrivate
}

func (chat *Chat) IsGroup() bool {
	return chat.Type == ChatTypeGroup
}

func (chat *Chat) IsSuperGroup() bool {
	return chat.Type == ChatTypeSuperGroup
}

func (chat *Chat) IsChannel() bool {
	return chat.Type == ChatTypeChannel
}
