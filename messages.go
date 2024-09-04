package aquagram

import (
	"context"
)

// TODO implement all fields
//
// https://core.telegram.org/bots/api#message
type Message struct {
	MessageID             int                 `json:"message_id"`
	MessageThreadID       int                 `json:"message_thread_id,omitempty"`
	From                  *User               `json:"from,omitempty"`
	SenderChat            *Chat               `json:"sender_chat,omitempty"`
	SenderBoostCount      int                 `json:"sender_boost_count,omitempty"`
	SenderBusinessBot     *User               `json:"sender_business_bot,omitempty"`
	Date                  int64               `json:"date"`
	BusinessConnectionID  string              `json:"business_connection_id,omitempty"`
	Chat                  *Chat               `json:"chat"`
	ForwardOrigin         MessageOrigin       `json:"forward_origin,omitempty"`
	IsTopicMessage        bool                `json:"is_topic_message,omitempty"`
	IsAutomaticMessage    bool                `json:"is_automatic_forward,omitempty"`
	ReplyToMessage        *Message            `json:"reply_to_message,omitempty"`
	ExternalReply         *ExternalReply      `json:"external_reply,omitempty"`
	Quote                 *TextQuote          `json:"quote,omitempty"`
	ViaBot                *User               `json:"via_bot,omitempty"`
	EditDate              int64               `json:"edit_date,omitempty"`
	HasProtectedContent   bool                `json:"has_protected_content,omitempty"`
	IsFromOffline         bool                `json:"is_from_offline,omitempty"`
	MediaGroupID          string              `json:"media_group_id,omitempty"`
	AuthorSignature       string              `json:"author_signature,omitempty"`
	Text                  string              `json:"text"`
	Entities              []*MessageEntity    `json:"entities,omitempty"`
	LinkPreviewOptions    *LinkPreviewOptions `json:"link_preview_options,omitempty"`
	EffectID              string              `json:"effect_id,omitempty"`
	Animation             *Animation          `json:"animation,omitempty"`
	Audio                 *Audio              `json:"audio,omitempty"`
	Document              *Document           `json:"document,omitempty"`
	PaidMedia             *PaidMediaInfo      `json:"paid_media_info,omitempty"`
	Photo                 []PhotoSize         `json:"photo,omitempty"`
	Sticker               *Sticker            `json:"sticker,omitempty"`
	Story                 *Story              `json:"story,omitempty"`
	Video                 *Video              `json:"video,omitempty"`
	VideoNote             *VideoNote          `json:"video_note,omitempty"`
	Voice                 *Voice              `json:"voice,omitempty"`
	Caption               string              `json:"caption,omitempty"`
	CaptionEntities       []*MessageEntity    `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool                `json:"show_caption_above_media,omitempty"`
	HasMediaSpoiler       bool                `json:"has_media_spoiler,omitempty"`
}

type MessageOrigin struct{}
type ExternalReply struct{}

type ReplyParameters struct {
	MessageID int `json:"message_id"`
}

type SendMessageParams struct {
	ChatID               string           `json:"chat_id"`
	Text                 string           `json:"text"`
	BusinessConnectionID string           `json:"business_connection_id,omitempty"`
	MessageThreadID      int              `json:"message_thread_id,omitempty"`
	ParseMode            ParseMode        `json:"parse_mode,omitempty"`
	Entities             []MessageEntity  `json:"entities,omitempty"`
	DisableNotification  bool             `json:"disable_notification,omitempty"`
	ReplyParameters      *ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup          *ReplyMarkup     `json:"reply_markup,omitempty"`
}

func (bot *Bot) SendMessage(chatID string, text string, params *SendMessageParams) (*Message, error) {
	return bot.SendMessageWithContext(bot.stopContext, chatID, text, params)
}

func (bot *Bot) SendMessageWithContext(ctx context.Context, chatID string, text string, params *SendMessageParams) (*Message, error) {
	if params == nil {
		params = new(SendMessageParams)
	}

	params.ChatID = ParseChatID(chatID)
	params.Text = text

	data, err := bot.Raw(ctx, "sendMessage", params)
	if err != nil {
		return nil, err
	}

	message, err := ParseRawResult[Message](data)
	if err != nil {
		return nil, err
	}

	return message, nil
}

func (message *Message) Reply(bot *Bot, text string, params *SendMessageParams) (*Message, error) {
	if params == nil {
		params = new(SendMessageParams)
	}

	if params.ReplyParameters == nil {
		params.ReplyParameters = new(ReplyParameters)
	}

	params.ReplyParameters.MessageID = message.MessageID

	return bot.SendMessage(ChatID(message.Chat.ID), text, params)
}

func (message *Message) GetMessage() *Message {
	return message
}

func (message *Message) GetFrom() *User {
	return message.From
}

func (message *Message) GetChat() *Chat {
	return message.Chat
}

func (message *Message) GetCallbackQuery() *CallbackQuery {
	return nil
}

func (message *Message) GetEntities() []*MessageEntity {
	return message.Entities
}

func (message *Message) process() {
	for _, entity := range message.Entities {
		entity.Message = message
	}

	for _, entity := range message.CaptionEntities {
		entity.Message = message
	}
}
