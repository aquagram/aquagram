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

func ParseRawResult[T any](bot *Bot, data []byte) (T, error) {
	var res struct {
		Ok          bool   `json:"ok"`
		Result      T      `json:"result"`
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
	}

	if err := json.Unmarshal(data, &res); err != nil {
		return res.Result, err
	}

	if !res.Ok {
		if res.ErrorCode != 0 {
			return res.Result, errTgBadRequest(res.ErrorCode, res.Description)
		}

		return res.Result, errors.New("unknown error parsing raw result: " + string(data))
	}

	if message, ok := any(&res.Result).(*Message); ok {
		message.process(bot)
	}

	return res.Result, nil
}

func ParseChatID(chatID string) string {
	at := "@"

	if !strings.HasPrefix(chatID, at) {
		_, err := strconv.ParseInt(chatID, 10, 0)
		if err != nil {
			chatID = at + chatID
		}
	}

	return chatID
}
