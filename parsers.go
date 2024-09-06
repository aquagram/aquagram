package aquagram

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

const (
	EmptyString   string = ""
	TrueAsString  string = "true"
	FalseAsString string = "false"
)

func ParseRawResult[T any](bot *Bot, data []byte) (result T, err error) {
	var res struct {
		Ok          bool   `json:"ok"`
		Result      T      `json:"result"`
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
	}

	if err = json.Unmarshal(data, &res); err != nil {
		return
	}

	if !res.Ok {
		if res.ErrorCode != 0 {
			err = errTgBadRequest(res.ErrorCode, res.Description)
			return
		}

		err = errors.New("unknown error parsing raw result: " + string(data))
		return
	}

	if message, ok := any(&res.Result).(*Message); ok {
		message.process(bot)
	}

	result = res.Result
	return
}

func ParseChatID(chatID string) string {
	at := "@"

	if !strings.HasPrefix(chatID, at) {
		_, err := strconv.ParseInt(chatID, 0, 0)
		if err != nil {
			chatID = at + chatID
		}
	}

	return chatID
}
