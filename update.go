package aquagram

import (
	"errors"
)

type UpdateType string

const (
	// telegram
	OnMessage                UpdateType = "message"
	OnEditedMessage          UpdateType = "edited_message"
	OnChannelPost            UpdateType = "channel_post"
	OnEditedChannelPost      UpdateType = "edited_channel_post"
	OnBusinessConnection     UpdateType = "business_connection"
	OnBusinessMessage        UpdateType = "business_message"
	OnEditedBusinessMessage  UpdateType = "edited_business_message"
	OnDeletedBusinessMessage UpdateType = "deleted_business_messages"
	OnMessageReaction        UpdateType = "message_reaction"
	OnMessageReactionCount   UpdateType = "message_reaction_count"
	OnInlineQuery            UpdateType = "inline_query"
	OnChosenInlineResult     UpdateType = "chosen_inline_result"
	OnCallbackQuery          UpdateType = "callback_query"
	OnShippingQuery          UpdateType = "shipping_query"
	OnPreCheckoutQuery       UpdateType = "pre_checkout_query"
	OnPoll                   UpdateType = "poll"
	OnPollAnswer             UpdateType = "poll_answer"
	OnMyChatMember           UpdateType = "my_chat_member"
	OnChatMember             UpdateType = "chat_member"
	OnChatJoinRequest        UpdateType = "chat_join_request"
	OnChatBoost              UpdateType = "chat_boost"
	OnRemovedChatBoost       UpdateType = "removed_chat_boost"

	// custom
	OnAnimation UpdateType = "animation"
	OnAudio     UpdateType = "audio"
	OnDocument  UpdateType = "document"
	OnPhoto     UpdateType = "photo"
	OnVideo     UpdateType = "video"
	OnVoice     UpdateType = "voice"
)

type Update struct {
	UpdateID          int            `json:"update_id"`
	Message           *Message       `json:"message,omitempty"`
	EditedMessage     *Message       `json:"edited_message,omitempty"`
	ChannelPost       *Message       `json:"channel_post,omitempty"`
	EditedChannelPost *Message       `json:"edited_channel_post,omitempty"`
	CallbackQuery     *CallbackQuery `json:"callback_query,omitempty"`
}

type Handlers = map[UpdateType][]*Handler

type Handler struct {
	Middlewares []Middleware
	Callback    func(bot *Bot, update any) error
}

func (bot *Bot) HandleUpdate(updateType UpdateType, update Event) {
	handlers, ok := bot.handlers[updateType]
	if !ok {
		return
	}

	for _, handler := range handlers {
		for _, middleware := range handler.Middlewares {
			err := middleware(update)

			if err != nil {
				if errors.Is(err, ErrStopPropagation) {
					return
				}

				bot.Logger.Printf("skiping handler execution due to a middleware error: %v\n", err)
				return
			}
		}

		if err := handler.Callback(bot, update); err != nil {
			if errors.Is(err, ErrStopPropagation) {
				return
			}

			bot.Logger.Printf("handler error: %v", err)
		}
	}
}

func (bot *Bot) DispatchUpdate(update *Update) {
	if update.Message != nil {
		message := update.Message
		message.process()

		bot.HandleUpdate(OnMessage, message)

		if message.Animation != nil {
			bot.HandleUpdate(OnAnimation, message)
		}

		if message.Audio != nil {
			bot.HandleUpdate(OnAudio, message)
		}

		if message.Document != nil {
			bot.HandleUpdate(OnDocument, message)
		}

		if message.Photo != nil {
			bot.HandleUpdate(OnPhoto, message)
		}

		if message.Video != nil {
			bot.HandleUpdate(OnVideo, message)
		}

		if message.Voice != nil {
			bot.HandleUpdate(OnVoice, message)
		}
	}

	if update.EditedMessage != nil {
		bot.HandleUpdate(OnEditedMessage, update.EditedMessage)
	}

	if update.ChannelPost != nil {
		bot.HandleUpdate(OnChannelPost, update.ChannelPost)
	}

	if update.EditedChannelPost != nil {
		bot.HandleUpdate(OnEditedChannelPost, update.ChannelPost)
	}

	if update.CallbackQuery != nil {
		bot.HandleUpdate(OnCallbackQuery, update.CallbackQuery)
	}
}
