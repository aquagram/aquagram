package aquagram

import (
	"context"
	"time"
)

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

// https://core.telegram.org/bots/api#answercallbackquery
func (callback *CallbackQuery) Answer(bot *Bot, params *AnswerCallbackQueryParams) (bool, error) {
	return bot.AnswerCallbackQuery(callback.ID, params)
}

type AnswerCallbackQueryParams struct {
	CallbackQueryID string        `json:"callback_query_id"`
	Text            string        `json:"text,omitempty"`
	ShowAlert       bool          `json:"show_alert,omitempty"`
	URL             string        `json:"url,omitempty"`
	CacheTime       time.Duration `json:"-"`
	CacheTimeRaw    int64         `json:"cache_time,omitempty"`
}

// https://core.telegram.org/bots/api#answercallbackquery
func (bot *Bot) AnswerCallbackQuery(callbackQueryID string, params *AnswerCallbackQueryParams) (bool, error) {
	return bot.AnswerCallbackQueryWithContext(bot.stopContext, callbackQueryID, params)
}

func (bot *Bot) AnswerCallbackQueryWithContext(ctx context.Context, callbackQueryID string, params *AnswerCallbackQueryParams) (bool, error) {
	if params == nil {
		params = new(AnswerCallbackQueryParams)
	}

	params.CallbackQueryID = callbackQueryID
	params.CacheTimeRaw = int64(params.CacheTime.Seconds())

	data, err := bot.Raw(ctx, "answerCallbackQuery", params)
	if err != nil {
		return false, err
	}

	success, err := ParseRawResult[bool](bot, data)
	if err != nil {
		return false, err
	}

	return *success, nil
}

func (callback *CallbackQuery) GetMessage() *Message {
	return nil
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
	return callback
}

func (callback *CallbackQuery) GetEntities() []*MessageEntity {
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
