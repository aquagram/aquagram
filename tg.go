package aquagram

// https://core.telegram.org/bots/api#callbackquery
type CallbackQuery struct {
	ID              string   `json:"id"`
	From            *User    `json:"from"`
	Message         *Message `json:"callback"` // https://core.telegram.org/bots/api#maybeinaccessiblemessage
	InlineMessageID string   `json:"inline_callback_id"`
	ChatInstance    string   `json:"chat_instance"`
	Data            string   `json:"data"`
	GameShortName   string   `json:"game_short_name"`
}

func (callback *CallbackQuery) GetFrom() *User {
	return callback.From
}

func (callback *CallbackQuery) GetChat() *Chat {
	if callback.Message != nil {
		return callback.Message.Chat
	}

	return nil
}

func (callback *CallbackQuery) GetCallbackQuery() *CallbackQuery {
	return nil
}

func (callback *CallbackQuery) GetText() string {
	return EmptyString
}

func (callback *CallbackQuery) GetEntities() []MessageEntity {
	return nil
}

// https://core.telegram.org/bots/api#linkpreviewoptions
type LinkPreviewOptions struct {
	IsDisabled       bool   `json:"is_disabled,omitempty"`
	Url              string `json:"url,omitempty"`
	PreferSmallMedia bool   `json:"prefer_small_media,omitempty"`
	PreferLargeMedia bool   `json:"prefer_large_media,omitempty"`
	ShowAboveText    bool   `json:"show_above_text,omitempty"`
}

// https://core.telegram.org/bots/api#textquote
type TextQuote struct {
	Text     string          `json:"text"`
	Entities []MessageEntity `json:"entities"`
	Position int             `json:"position,omitempty"`
	IsManual bool            `json:"is_manual,omitempty"`
}

// TODO implement all field
// https://core.telegram.org/bots/api#paidmediainfo
type PaidMediaInfo struct{}

// TODO implement all fields
// https://core.telegram.org/bots/api#sticker
type Sticker struct{}
