package aquagram

import (
	"regexp"
	"slices"
	"strings"
)

func CallbackQueryFilter(data string, strict bool) FilterFunc {
	return func(bot *Bot, event Event) (bool, error) {
		callbackQuery := event.GetCallbackQuery()
		if callbackQuery == nil {
			return false, nil
		}

		if strict && callbackQuery.Data == data {
			return true, nil
		}

		return strings.HasPrefix(callbackQuery.Data, data), nil
	}
}

func ChatMemberFilter(chatID string) FilterFunc {
	return func(bot *Bot, event Event) (bool, error) {
		from := event.GetFrom()
		if from == nil {
			return false, nil
		}

		member, err := bot.GetChatMember(chatID, from.ID)
		if err != nil {
			return false, err
		}

		if member.IsOwner() ||
			member.IsAdministrator() ||
			member.IsMember() ||
			member.IsRestricted() {
			return true, nil
		}

		return false, nil
	}
}

func CommandFilter(command string) FilterFunc {
	slash := "/"

	if !strings.HasPrefix(command, slash) {
		command = slash + command
	}

	return func(bot *Bot, event Event) (bool, error) {
		message := event.GetMessage()
		if message == nil {
			return false, nil
		}

		if message.Text == "" {
			return false, nil
		}

		var text string

		entities := event.GetEntities()
		for _, entity := range entities {
			if entity.Type != EntityTypeBotCommand || entity.Offset != 0 {
				continue
			}

			text = message.Text[entity.Offset:entity.Length]
			break
		}

		if text == "" {
			return false, nil
		}

		atIndex := strings.Index(text, "@")
		if atIndex != -1 {
			text = text[:atIndex]
		}

		if text != command {
			return false, nil
		}

		return true, nil
	}
}

func BlackListFilter(ids *[]int64) FilterFunc {
	return func(bot *Bot, event Event) (bool, error) {
		from := event.GetFrom()
		if from == nil {
			return false, nil
		}

		return !slices.Contains(*ids, from.ID), nil
	}
}

func RegexFilter(regex *regexp.Regexp) FilterFunc {
	return func(bot *Bot, event Event) (bool, error) {
		message := event.GetMessage()
		if message == nil {
			return false, nil
		}

		if !regex.MatchString(message.Text) {
			return false, nil
		}

		return true, nil
	}
}

func TextFilter(text string, strict bool, caseSensitive bool) FilterFunc {
	if !caseSensitive {
		text = strings.ToLower(text)
	}

	return func(bot *Bot, event Event) (bool, error) {
		message := event.GetMessage()
		if message == nil {
			return false, nil
		}

		messageText := message.Text
		if !caseSensitive {
			messageText = strings.ToLower(messageText)
		}

		if strict && messageText == text {
			return true, nil
		}

		return strings.Contains(messageText, text), nil
	}
}

func WhiteListFilter(ids *[]int64) FilterFunc {
	return func(bot *Bot, event Event) (bool, error) {
		from := event.GetFrom()
		if from == nil {
			return false, nil
		}

		return slices.Contains(*ids, from.ID), nil
	}
}
