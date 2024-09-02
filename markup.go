package aquagram

import "encoding/json"

type ReplyMarkup interface {
	GetInlineKeyboardMarkup() *InlineKeyboardMarkup
	GetReplyKeyboardMarkup() *ReplyKeyboardMarkup
	GetReplyKeyboardRemove() *ReplyKeyboardRemove
	GetForceReply() *ForceReply
}

func ParseReplyMarkup(markup ReplyMarkup) ([]byte, error) {
	if inlineKeyboardMarkup := markup.GetInlineKeyboardMarkup(); inlineKeyboardMarkup != nil {
		return json.Marshal(inlineKeyboardMarkup)
	}

	if replyKeyboardMarkup := markup.GetReplyKeyboardMarkup(); replyKeyboardMarkup != nil {
		return json.Marshal(replyKeyboardMarkup)
	}

	if keyboardRemove := markup.GetReplyKeyboardRemove(); keyboardRemove != nil {
		return json.Marshal(keyboardRemove)
	}

	if forceReply := markup.GetForceReply(); forceReply != nil {
		return json.Marshal(forceReply)
	}

	return nil, ErrUnknownMarkup
}

// https://core.telegram.org/bots/api#inlinekeyboardmarkup
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]*InlineKeyboardButton `json:"inline_keyboard"`
}

func (markup *InlineKeyboardMarkup) GetInlineKeyboardMarkup() *InlineKeyboardMarkup {
	return markup
}

func (markup *InlineKeyboardMarkup) GetReplyKeyboardMarkup() *ReplyKeyboardMarkup {
	return nil
}

func (markup *InlineKeyboardMarkup) GetReplyKeyboardRemove() *ReplyKeyboardRemove {
	return nil
}

func (markup *InlineKeyboardMarkup) GetForceReply() *ForceReply {
	return nil
}

// TODO implement all fields
//
// https://core.telegram.org/bots/api#inlinekeyboardbutton
type InlineKeyboardButton struct {
	Text                         string `json:"text"`
	Url                          string `json:"url,omitempty"`
	CallbackData                 string `json:"callback_data,omitempty"` // 1-64 bytes
	SwitchInlineQuery            string `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat string `json:"switch_inline_query_current_chat,omitempty"`
	Pay                          bool   `json:"pay,omitempty"`
}

// https://core.telegram.org/bots/api#replykeyboardmarkup
type ReplyKeyboardMarkup struct {
	Keyboard              [][]*KeyboardButton `json:"keyboard"`
	IsPersistent          bool                `json:"is_persistent,omitempty"`
	ResizeKeyboard        bool                `json:"resize_keyboard,omitempty"`
	OneTimeKeyboard       bool                `json:"one_time_keyboard,omitempty"`
	InputFieldPlaceholder string              `json:"input_field_placeholder,omitempty"`
	Selective             bool                `json:"selective,omitempty"`
}

func (markup *ReplyKeyboardMarkup) GetInlineKeyboardMarkup() *InlineKeyboardMarkup {
	return nil
}

func (markup *ReplyKeyboardMarkup) GetReplyKeyboardMarkup() *ReplyKeyboardMarkup {
	return markup
}

func (markup *ReplyKeyboardMarkup) GetReplyKeyboardRemove() *ReplyKeyboardRemove {
	return nil
}

func (markup *ReplyKeyboardMarkup) GetForceReply() *ForceReply {
	return nil
}

// TODO implement all fields
//
// https://core.telegram.org/bots/api#keyboardbutton
type KeyboardButton struct {
	Text string `json:"text"`
}

// https://core.telegram.org/bots/api#replykeyboardremove
type ReplyKeyboardRemove struct {
	RemoveKeyboard bool `json:"remove_keyboard"`
	Selective      bool `json:"selective,omitempty"`
}

func (markup *ReplyKeyboardRemove) GetInlineKeyboardMarkup() *InlineKeyboardMarkup {
	return nil
}

func (markup *ReplyKeyboardRemove) GetReplyKeyboardMarkup() *ReplyKeyboardMarkup {
	return nil
}

func (markup *ReplyKeyboardRemove) GetReplyKeyboardRemove() *ReplyKeyboardRemove {
	return markup
}

func (markup *ReplyKeyboardRemove) GetForceReply() *ForceReply {
	return nil
}

// https://core.telegram.org/bots/api#forcereply
type ForceReply struct {
	ForceReply            bool   `json:"force_reply"`
	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"`
	Selective             bool   `json:"selective,omitempty"`
}

func (markup *ForceReply) GetInlineKeyboardMarkup() *InlineKeyboardMarkup {
	return nil
}

func (markup *ForceReply) GetReplyKeyboardMarkup() *ReplyKeyboardMarkup {
	return nil
}

func (markup *ForceReply) GetReplyKeyboardRemove() *ReplyKeyboardRemove {
	return nil
}

func (markup *ForceReply) GetForceReply() *ForceReply {
	return markup
}
