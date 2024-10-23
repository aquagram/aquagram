package aquagram

import "fmt"

// This object represents a Telegram user or bot.
//
// https://core.telegram.org/bots/api#user
type User struct {
	ID                      int64  `json:"id"`
	IsBot                   bool   `json:"is_bool,omitempty"`
	Username                string `json:"username"`
	FirstName               string `json:"first_name"`
	LastName                string `json:"last_name,omitempty"`
	LanguageCode            string `json:"language_code,omitempty"`
	IsPremium               bool   `json:"is_premium,omitempty"`
	AddedToAttachmentMenu   bool   `json:"added_to_attachment_menu,omitempty"`
	CanJoinGroups           bool   `json:"can_join_groups,omitempty"`
	CanReadAllGroupMessages bool   `json:"can_read_all_group_messages,omitempty"`
	SupportsInlineQueries   bool   `json:"supports_inline_queries,omitempty"`
	CanConnectToBusiness    bool   `json:"can_connect_to_business,omitempty"`
	HasMainWebApp           bool   `json:"has_main_web_app,omitempty"`
}

func (user *User) TextMention(mode ParseMode) string {
	if mode == ParseModeHTML {
		return fmt.Sprintf(`<a href="tg://user?id=%d">%s</a>`, user.ID, user.FirstName)
	}

	return fmt.Sprintf("[%s](tg://user?id=%d)", user.FirstName, user.ID)
}
