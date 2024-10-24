package aquagram

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
	UpdateID              int            `json:"update_id"`
	Message               *Message       `json:"message,omitempty"`
	EditedMessage         *Message       `json:"edited_message,omitempty"`
	ChannelPost           *Message       `json:"channel_post,omitempty"`
	EditedChannelPost     *Message       `json:"edited_channel_post,omitempty"`
	BusinessMessage       *Message       `json:"business_message,omitempty"`
	EditedBusinessMessage *Message       `json:"edited_business_message,omitempty"`
	CallbackQuery         *CallbackQuery `json:"callback_query,omitempty"`
}

func (bot *Bot) DispatchUpdate(update *Update) {
	if update.Message != nil {
		message := update.Message
		message.process(bot)

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

	if update.BusinessMessage != nil {
		bot.HandleUpdate(OnBusinessMessage, update.BusinessMessage)
	}

	if update.EditedBusinessMessage != nil {
		bot.HandleUpdate(OnBusinessMessage, update.EditedBusinessMessage)
	}

	if update.CallbackQuery != nil {
		update.CallbackQuery.process(bot)
		bot.HandleUpdate(OnCallbackQuery, update.CallbackQuery)
	}
}

func (bot *Bot) HandleUpdate(updateType UpdateType, update Event) {
	next := func(bot *Bot, event Event) error {
		return bot.runHandlers(bot.handlers, updateType, event)
	}

	err := bot.runMiddlewares(bot.Middlewares, update, next)
	if err != nil {
		bot.Logger.Printf("handling update: %v", err)
	}
}

func (bot *Bot) runHandlers(handlers Handlers, updateType UpdateType, event Event) error {
	for _, handler := range handlers[updateType] {
		err := bot.runHandlerMiddlewares(handler, event)
		if err == StopPropagation {
			break
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func (bot *Bot) runHandlerMiddlewares(handler *Handler, event Event) error {
	next := func(bot *Bot, event Event) error {
		return handler.Callback(bot, event)
	}

	return bot.runMiddlewares(handler.Middlewares, event, next)
}

func (bot *Bot) runMiddlewares(middlewares []Middleware, event Event, next MiddlewareFunc) error {
	length := len(middlewares) - 1

	for i := length; i > -1; i-- {
		middleware := middlewares[i]
		next = middleware(next)
	}

	return next(bot, event)
}
