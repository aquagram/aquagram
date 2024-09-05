package aquagram

import (
	"context"
)

type MessageID struct {
	MessageID int64 `json:"message_id"`
}

// TODO implement all fields
//
// https://core.telegram.org/bots/api#message
type Message struct {
	Bot *Bot `json:"-"`

	MessageID             int64               `json:"message_id"`
	MessageThreadID       int64               `json:"message_thread_id,omitempty"`
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
	MessageID int64 `json:"message_id"`
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
	MessageThreadID      int64               `json:"message_thread_id,omitempty"`
	Text                 string              `json:"text"`
	ParseMode            ParseMode           `json:"parse_mode,omitempty"`
	Entities             []MessageEntity     `json:"entities,omitempty"`
	LinkPreviewOptions   *LinkPreviewOptions `json:"link_preview_options,omitempty"`
	DisableNotification  bool                `json:"disable_notification,omitempty"`
	ProtectContent       bool                `json:"protect_content,omitempty"`
	MessageEffectID      string              `json:"message_effect_id,omitempty"`
	ReplyParameters      *ReplyParameters    `json:"reply_parameters,omitempty"`
	ReplyMarkup          ReplyMarkup         `json:"reply_markup,omitempty"`
}

/*
[Reply] is an alias for [SendMessage].
*/
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

/*
[SendMessage] wraps [SendMessageWithContext] using the default bot context.
*/
func (bot *Bot) SendMessage(chatID string, text string, params *SendMessageParams) (*Message, error) {
	return bot.SendMessageWithContext(bot.stopContext, chatID, text, params)
}

/*
[sendMessage] - Use this method to send text messages.

[sendMessage]: https://core.telegram.org/bots/api#sendmessage
*/
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

	return ParseRawResult[Message](bot, data)
}

type ForwardParams struct {
	ChatID              string `json:"chat_id"`
	MessageThreadID     int64  `json:"message_thread_id,omitempty"`
	FromChatID          string `json:"from_chat_id"`
	DisableNotification bool   `json:"disable_notification,omitempty"`
	ProtectContent      bool   `json:"protect_content,omitempty"`
}

type ForwardMessageParams struct {
	ForwardParams
	MessageID int64 `json:"message_id"`
}

/*
[Forward] is an alias for [ForwardMessage].
*/
func (message *Message) Forward(chatID string, params *ForwardMessageParams) (*Message, error) {
	return message.Bot.ForwardMessage(chatID, ChatID(message.Chat.ID), message.MessageID, params)
}

/*
[ForwardMessage] wraps [ForwardMessageWithContext] using the default bot context.
*/
func (bot *Bot) ForwardMessage(chatID string, fromChatID string, messageID int64, params *ForwardMessageParams) (*Message, error) {
	return bot.ForwardMessageWithContext(bot.stopContext, chatID, fromChatID, messageID, params)
}

/*
[forwardMessage] - Use this method to forward messages of any kind.

Service messages and messages with protected content can't be forwarded.

On success, the sent Message is returned.

[forwardMessage]: https://core.telegram.org/bots/api#forwardmessage
*/
func (bot *Bot) ForwardMessageWithContext(ctx context.Context, chatID string, fromChatID string, messageID int64, params *ForwardMessageParams) (*Message, error) {
	if params == nil {
		params = new(ForwardMessageParams)
	}

	params.ChatID = ParseChatID(chatID)
	params.FromChatID = ParseChatID(fromChatID)
	params.MessageID = messageID

	data, err := bot.Raw(ctx, "forwardMessage", params)
	if err != nil {
		return nil, err
	}

	return ParseRawResult[Message](bot, data)
}

type ForwardMessagesParams struct {
	ForwardParams
	MessageIDs []int64
}

/*
[ForwardMessages] wraps [ForwardMessagesWithContext] using the default bot context.
*/
func (bot *Bot) ForwardMessages(chatID string, fromChatID string, messageIDs []int64, params *ForwardMessagesParams) ([]int64, error) {
	return bot.ForwardMessagesWithContext(bot.stopContext, chatID, fromChatID, messageIDs, params)
}

/*
[forwardMessage] - Use this method to forward multiple messages of any kind.

If some of the specified messages can't be found or forwarded, they are skipped.
Service messages and messages with protected content can't be forwarded.
Album grouping is kept for forwarded messages.

On success, an array of MessageID of the sent messages is returned.

[forwardMessage]: https://core.telegram.org/bots/api#forwardmessages
*/
func (bot *Bot) ForwardMessagesWithContext(ctx context.Context, chatID string, fromChatID string, messageIDs []int64, params *ForwardMessagesParams) ([]int64, error) {
	if params == nil {
		params = new(ForwardMessagesParams)
	}

	params.ChatID = ParseChatID(chatID)
	params.FromChatID = ParseChatID(fromChatID)
	params.MessageIDs = messageIDs

	data, err := bot.Raw(ctx, "forwardMessage", params)
	if err != nil {
		return nil, err
	}

	ids, err := ParseRawResult[[]MessageID](bot, data)

	var intArr []int64

	for _, result := range *ids {
		intArr = append(intArr, result.MessageID)
	}

	return intArr, err
}

type CopyMessageParams struct {
	ChatID                string           `json:"chat_id"`
	MessageThreadID       int64            `json:"message_thread_id,omitempty"`
	FromChatID            string           `json:"from_chat_id"`
	MessageID             int64            `json:"message_id"`
	Caption               string           `json:"caption,omitempty"`
	ParseMode             ParseMode        `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity  `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool             `json:"show_caption_above_media,omitempty"`
	DisableNotification   bool             `json:"disable_notification,omitempty"`
	ProtectContent        bool             `json:"protect_content,omitempty"`
	ReplyParameters       *ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup           ReplyMarkup      `json:"reply_markup,omitempty"`
}

/*
[Copy] is an alias for [CopyMessage].
*/
func (message *Message) Copy(chatID string, params *CopyMessageParams) (int64, error) {
	return message.Bot.CopyMessage(chatID, ChatID(message.Chat.ID), message.MessageID, params)
}

/*
[CopyMessage] wraps [CopyMessageWithContext] using the default bot context.
*/
func (bot *Bot) CopyMessage(chatID string, fromChatID string, messageID int64, params *CopyMessageParams) (int64, error) {
	return bot.CopyMessageWithContext(bot.stopContext, chatID, fromChatID, messageID, params)
}

/*
[copyMessage] - Use this method to copy messages of any kind.

Service messages, paid media messages, giveaway messages, giveaway winners messages, and invoice messages can't be copied.
A quiz poll can be copied only if the value of the field correct_option_id is known to the bot.

The method is analogous to the method [forwardMessage], but the copied message doesn't have a link to the original message.

Returns the MessageID of the sent message on success.

[copyMessage]: https://core.telegram.org/bots/api#copymessage
[forwardMessage]: https://core.telegram.org/bots/api#forwardmessages
*/
func (bot *Bot) CopyMessageWithContext(ctx context.Context, chatID string, fromChatID string, messageID int64, params *CopyMessageParams) (int64, error) {
	if params == nil {
		params = new(CopyMessageParams)
	}

	params.ChatID = chatID
	params.FromChatID = fromChatID
	params.MessageID = messageID

	data, err := bot.Raw(ctx, "copyMessage", params)
	if err != nil {
		return 0, err
	}

	id, err := ParseRawResult[MessageID](bot, data)

	return id.MessageID, err
}

type CopyMessagesParams struct {
	ChatID          string  `json:"chat_id"`
	MessageThreadID int64   `json:"message_thread_id,omitempty"`
	FromChatID      string  `json:"from_chat_id"`
	MessageIDs      []int64 `json:"message_ids"`
}

func (bot *Bot) CopyMessages(chatID string, fromChatID string, messageIDs []int64, params *CopyMessagesParams) ([]int64, error) {
	return bot.CopyMessagesWithContext(bot.stopContext, chatID, fromChatID, messageIDs, params)
}

/*
[copyMessages] - Use this method to copy messages of any kind.

If some of the specified messages can't be found or copied, they are skipped.
Service messages, paid media messages, giveaway messages, giveaway winners messages, and invoice messages can't be copied.
A quiz poll can be copied only if the value of the field correct_option_id is known to the bot.

The method is analogous to the method forwardMessages, but the copied messages don't have a link to the original message.
Album grouping is kept for copied messages.

On success, an array of MessageID of the sent messages is returned.

[copyMessages]: https://core.telegram.org/bots/api#copymessages
*/
func (bot *Bot) CopyMessagesWithContext(ctx context.Context, chatID string, fromChatID string, messageIDs []int64, params *CopyMessagesParams) ([]int64, error) {
	if params == nil {
		params = new(CopyMessagesParams)
	}

	params.ChatID = ParseChatID(chatID)
	params.FromChatID = ParseChatID(fromChatID)
	params.MessageIDs = messageIDs

	data, err := bot.Raw(ctx, "copyMessages", params)
	if err != nil {
		return nil, err
	}

	ids, err := ParseRawResult[[]MessageID](bot, data)

	var intArr []int64

	for _, result := range *ids {
		intArr = append(intArr, result.MessageID)
	}

	return intArr, err
}

type EditMessageParams struct {
	BusinessConnectionID string              `json:"business_connection_id,omitempty"`
	ChatID               string              `json:"chat_id"`
	MessageID            int64               `json:"message_id,omitempty"`
	InlineMessageID      int64               `json:"inline_message_id,omitempty"`
	Text                 string              `json:"text"`
	ParseMode            ParseMode           `json:"parse_mode,omitempty"`
	Entities             []MessageEntity     `json:"entities,omitempty"`
	LinkPreviewOptions   *LinkPreviewOptions `json:"link_preview_options,omitempty"`
	ReplyMarkup          ReplyMarkup         `json:"reply_markup,omitempty"`
}

func (message *Message) EditText(text string, params *EditMessageParams) (*Message, error) {
	return message.Bot.EditMessageText(ChatID(message.Chat.ID), message.MessageID, text, params)
}

/*
[CopyMessage] wraps [CopyMessageWithContext] using the default bot context.
*/
func (bot *Bot) EditMessageText(chatID string, messageID int64, text string, params *EditMessageParams) (*Message, error) {
	return bot.EditMessageTextWithContext(bot.stopContext, chatID, messageID, text, params)
}

/*
[editMessageText] - Use this method to edit text and game messages.

On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.

[editMessageText]: https://core.telegram.org/bots/api#editmessagetext
*/
func (bot *Bot) EditMessageTextWithContext(ctx context.Context, chatID string, messageID int64, text string, params *EditMessageParams) (*Message, error) {
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

/*
[Delete] is an alias for [DeleteMessage]
*/
func (message *Message) Delete() error {
	return message.Bot.DeleteMessage(ChatID(message.Chat.ID), message.MessageID)
}

/*
[DeleteMessage] wraps [DeleteMessageWithContext] using the default bot context.
*/
func (bot *Bot) DeleteMessage(chatID string, messageID int64) error {
	return bot.DeleteMessageWithContext(bot.stopContext, chatID, messageID)
}

/*
[deleteMessage] - Use this method to delete a message, including service messages, with the following limitations:
  - A message can only be deleted if it was sent less than 48 hours ago.
  - Service messages about a supergroup, channel, or forum topic creation can't be deleted.
  - A dice message in a private chat can only be deleted if it was sent more than 24 hours ago.
  - Bots can delete outgoing messages in private chats, groups, and supergroups.
  - Bots can delete incoming messages in private chats.
  - Bots granted can_post_messages permissions can delete outgoing messages in channels.
  - If the bot is an administrator of a group, it can delete any message there.
  - If the bot has can_delete_messages permission in a supergroup or a channel, it can delete any message there.

Returns a nil error on success.

[deleteMessage]: https://core.telegram.org/bots/api#deletemessage
*/
func (bot *Bot) DeleteMessageWithContext(ctx context.Context, chatID string, messageID int64) error {
	params := map[string]any{
		"chat_id":    chatID,
		"message_id": messageID,
	}

	data, err := bot.Raw(ctx, "deleteMessage", params)
	if err != nil {
		return err
	}

	success, err := ParseRawResult[bool](bot, data)
	if err != nil {
		return err
	}

	if !*success {
		return ErrExpectedTrue
	}

	return nil
}

/*
[DeleteMessages] wraps [DeleteMessagesWithContext] using the default bot context.
*/
func (bot *Bot) DeleteMessages(chatID string, messageIDs []int) error {
	return bot.DeleteMessagesWithContext(bot.stopContext, chatID, messageIDs)
}

/*
[deleteMessages] - Use this method to delete multiple messages simultaneously.

If some of the specified messages can't be found, they are skipped.

Returns a nil error on success.

[deleteMessages]: https://core.telegram.org/bots/api#deletemessages
*/
func (bot *Bot) DeleteMessagesWithContext(ctx context.Context, chatID string, messageIDs []int) error {
	params := map[string]any{
		"chat_id":     chatID,
		"message_ids": messageIDs,
	}

	data, err := bot.Raw(bot.stopContext, "deleteMessages", params)
	if err != nil {
		return err
	}

	success, err := ParseRawResult[bool](bot, data)
	if err != nil {
		return err
	}

	if !*success {
		return ErrExpectedTrue
	}

	return nil
}
