package aquagram

import (
	"context"
	"time"
)

type BotName struct {
	Name string `json:"name"`
}

type MaybeInaccessibleMessage = Message

// https://core.telegram.org/bots/api#callbackquery
type CallbackQuery struct {
	Bot *Bot `json:"-"`

	ID              string                    `json:"id"`
	From            *User                     `json:"from"`
	Message         *MaybeInaccessibleMessage `json:"message"`
	InlineMessageID string                    `json:"inline_callback_id"`
	ChatInstance    string                    `json:"chat_instance"`
	Data            string                    `json:"data"`
	GameShortName   string                    `json:"game_short_name"`
}

func (callback *CallbackQuery) IsMessageInaccessible() bool {
	return callback.Message == nil
}

type AnswerCallbackQueryParams struct {
	CallbackQueryID string        `json:"callback_query_id"`
	Text            string        `json:"text,omitempty"`
	ShowAlert       bool          `json:"show_alert,omitempty"`
	URL             string        `json:"url,omitempty"`
	CacheTime       time.Duration `json:"-"`
	CacheTimeRaw    int64         `json:"cache_time,omitempty"`
}

/*
[Answer] is an alias for [AnswerCallbackQuery]
*/
func (callback *CallbackQuery) Answer(params *AnswerCallbackQueryParams) error {
	return callback.Bot.AnswerCallbackQuery(callback.ID, params)
}

// https://core.telegram.org/bots/api#answercallbackquery
/*

 */
func (bot *Bot) AnswerCallbackQuery(callbackQueryID string, params *AnswerCallbackQueryParams) error {
	return bot.AnswerCallbackQueryWithContext(bot.stopContext, callbackQueryID, params)
}

/*
[answerCallbackQuery] - Use this method to send answers to callback queries sent from inline keyboards.

The answer will be displayed to the user as a notification at the top of the chat screen or as an alert.

On success, a nil error is returned.

[answerCallbackQuery]: https://core.telegram.org/bots/api#answercallbackquery
*/
func (bot *Bot) AnswerCallbackQueryWithContext(ctx context.Context, callbackQueryID string, params *AnswerCallbackQueryParams) error {
	if params == nil {
		params = new(AnswerCallbackQueryParams)
	}

	params.CallbackQueryID = callbackQueryID
	params.CacheTimeRaw = int64(params.CacheTime.Seconds())

	data, err := bot.Raw(ctx, "answerCallbackQuery", params)
	if err != nil {
		return err
	}

	success, err := ParseRawResult[bool](bot, data)
	if err != nil {
		return err
	}

	if !success {
		return ErrExpectedTrue
	}

	return nil
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
