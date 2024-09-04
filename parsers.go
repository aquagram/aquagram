package aquagram

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

const (
	EmptyString   string = ""
	TrueAsString  string = "true"
	FalseAsString string = "false"
)

func ParseRawResult[T any](data []byte) (*T, error) {
	var res struct {
		Ok          bool   `json:"ok"`
		Result      *T     `json:"result"`
		ErrorCode   int    `json:"error_code"`
		Description string `json:"description"`
	}

	if err := json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	if !res.Ok {
		if res.ErrorCode != 0 {
			return nil, errTgBadRequest(res.ErrorCode, res.Description)
		}

		return nil, fmt.Errorf("unknown error parsing raw result: %s", string(data))
	}

	if message, ok := any(res.Result).(*Message); ok {
		message.process()
	}

	return res.Result, nil
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
