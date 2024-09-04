package aquagram

import (
	"context"
)

// TODO implement all fields
//
// https://core.telegram.org/bots/api#message
type Message struct {
	Bot *Bot `json:"-"`

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

func (message *Message) process(bot *Bot) {
	message.Bot = bot

	for _, entity := range message.Entities {
		entity.Message = message
	}

	for _, entity := range message.CaptionEntities {
		entity.Message = message
	}
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

type SendMessageParams struct {
	BusinessConnectionID string              `json:"business_connection_id,omitempty"`
	ChatID               string              `json:"chat_id"`
	MessageThreadID      int                 `json:"message_thread_id,omitempty"`
	Text                 string              `json:"text"`
	ParseMode            ParseMode           `json:"parse_mode,omitempty"`
	Entities             []MessageEntity     `json:"entities,omitempty"`
	LinkPreviewOptions   *LinkPreviewOptions `json:"link_preview_options,omitempty"`
	DisableNotification  bool                `json:"disable_notification,omitempty"`
	ReplyParameters      *ReplyParameters    `json:"reply_parameters,omitempty"`
	ReplyMarkup          ReplyMarkup         `json:"reply_markup,omitempty"`
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

	message, err := ParseRawResult[Message](bot, data)
	if err != nil {
		return nil, err
	}

	return message, nil
}

func (message *Message) Reply(text string, params *SendMessageParams) (*Message, error) {
	if params == nil {
		params = new(SendMessageParams)
	}

	if params.ReplyParameters == nil {
		params.ReplyParameters = new(ReplyParameters)
	}

	params.ReplyParameters.MessageID = message.MessageID

	return message.Bot.SendMessage(ChatID(message.Chat.ID), text, params)
}

type EditMessageParams struct {
	BusinessConnectionID string              `json:"business_connection_id,omitempty"`
	ChatID               string              `json:"chat_id"`
	MessageID            int                 `json:"message_id,omitempty"`
	InlineMessageID      int                 `json:"inline_message_id,omitempty"`
	Text                 string              `json:"text"`
	ParseMode            ParseMode           `json:"parse_mode,omitempty"`
	Entities             []MessageEntity     `json:"entities,omitempty"`
	LinkPreviewOptions   *LinkPreviewOptions `json:"link_preview_options,omitempty"`
	ReplyMarkup          ReplyMarkup         `json:"reply_markup,omitempty"`
}

func (message *Message) EditText(text string, params *EditMessageParams) (*Message, error) {
	return message.Bot.EditMessageText(ChatID(message.Chat.ID), message.MessageID, text, params)
}

// https://core.telegram.org/bots/api#editmessagetext
func (bot *Bot) EditMessageText(chatID string, messageID int, text string, params *EditMessageParams) (*Message, error) {
	return bot.EditMessageTextWithContext(bot.stopContext, chatID, messageID, text, params)
}

func (bot *Bot) EditMessageTextWithContext(ctx context.Context, chatID string, messageID int, text string, params *EditMessageParams) (*Message, error) {
	if params == nil {
		params = new(EditMessageParams)
	}

	params.ChatID = chatID
	params.MessageID = messageID
	params.Text = text

	data, err := bot.Raw(ctx, "editMessageText", params)
	if err != nil {
		return nil, err
	}

	return ParseRawResult[Message](bot, data)
}

func (message *Message) Delete() (bool, error) {
	return message.Bot.DeleteMessage(ChatID(message.Chat.ID), message.MessageID)
}

// https://core.telegram.org/bots/api#deletemessage
func (bot *Bot) DeleteMessage(chatID string, messageID int) (bool, error) {
	return bot.DeleteMessageWithContext(bot.stopContext, chatID, messageID)
}

func (bot *Bot) DeleteMessageWithContext(ctx context.Context, chatID string, messageID int) (bool, error) {
	params := map[string]any{
		"chat_id":    chatID,
		"message_id": messageID,
	}

	data, err := bot.Raw(ctx, "deleteMessage", params)
	if err != nil {
		return false, err
	}

	success, err := ParseRawResult[bool](bot, data)
	if err != nil {
		return false, err
	}

	return *success, nil
}

// https://core.telegram.org/bots/api#deletemessages
func (bot *Bot) DeleteMessages(chatID string, messageIDs []int) (bool, error) {
	return bot.DeleteMessagesWithContext(bot.stopContext, chatID, messageIDs)
}

func (bot *Bot) DeleteMessagesWithContext(ctx context.Context, chatID string, messageIDs []int) (bool, error) {
	params := map[string]any{
		"chat_id":     chatID,
		"message_ids": messageIDs,
	}

	data, err := bot.Raw(bot.stopContext, "deleteMessages", params)
	if err != nil {
		return false, err
	}

	success, err := ParseRawResult[bool](bot, data)
	if err != nil {
		return false, err
	}

	return *success, nil
}
