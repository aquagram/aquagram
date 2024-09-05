package aquagram

func (message *Message) process(bot *Bot) {
	message.Bot = bot

	for _, entity := range message.Entities {
		entity.Message = message
	}

	for _, entity := range message.CaptionEntities {
		entity.Message = message
	}

	if message.ReplyToMessage != nil {
		message.ReplyToMessage.process(bot)
	}
}

func (callback *CallbackQuery) process(bot *Bot) {
	callback.Bot = bot

	if callback.Message != nil {
		callback.Message.process(bot)
	}
}
