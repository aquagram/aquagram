package aquagram

import (
	"encoding/json"
	"strconv"
)

type Params map[string]string

type CommonSendParams struct {
	ChatID                      string           `json:"chat_id"`
	Text                        string           `json:"text"`
	BusinessConnectionID        string           `json:"business_connection_id,omitempty"`
	MessageThreadID             int              `json:"message_thread_id,omitempty"`
	ParseMode                   ParseMode        `json:"parse_mode,omitempty"`
	Entities                    []MessageEntity  `json:"entities,omitempty"`
	DisableNotification         bool             `json:"disable_notification,omitempty"`
	Media                       []Params         `json:"media"`
	Caption                     string           `json:"caption,omitempty"`
	CaptionEntities             []MessageEntity  `json:"caption_entities,omitempty"`
	DisableContentTypeDetection bool             `json:"disable_content_type_detection,omitempty"`
	ProtectContent              bool             `json:"protect_content,omitempty"`
	MessageEffectID             string           `json:"message_effect_id,omitempty"`
	ReplyParameters             *ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup                 ReplyMarkup      `json:"reply_markup,omitempty"`
}

func (p *CommonSendParams) ToParams() (Params, error) {
	params := make(Params)

	params["chat_id"] = ParseChatID(p.ChatID)

	if p.Text != EmptyString {
		params["text"] = p.Text
	}

	if p.BusinessConnectionID != EmptyString {
		params["business_connection_id"] = p.BusinessConnectionID
	}

	if p.MessageThreadID != 0 {
		params["message_thread_id"] = strconv.Itoa(p.MessageThreadID)
	}

	if string(p.ParseMode) != EmptyString {
		params["parse_mode"] = string(p.ParseMode)
	}

	if p.Entities != nil {
		data, err := json.Marshal(p.Entities)
		if err != nil {
			return nil, err
		}

		params["entities"] = string(data)
	}

	if p.DisableNotification {
		params["disable_notification"] = TrueAsString
	}

	if p.Media != nil {
		data, err := json.Marshal(p.Media)
		if err != nil {
			return nil, err
		}

		params["media"] = string(data)
	}

	if p.Caption != EmptyString {
		params["caption"] = p.Caption
	}

	if p.CaptionEntities != nil {
		data, err := json.Marshal(p.CaptionEntities)
		if err != nil {
			return nil, err
		}

		params["caption_entities"] = string(data)
	}

	if p.DisableContentTypeDetection {
		params["disable_content_type_detection"] = TrueAsString
	}

	if p.ProtectContent {
		params["protect_content"] = TrueAsString
	}

	if p.MessageEffectID != EmptyString {
		params["message_effect_id"] = p.MessageEffectID
	}

	if p.ReplyParameters != nil {
		data, err := json.Marshal(p.ReplyParameters)
		if err != nil {
			return nil, err
		}

		params["reply_parameters"] = string(data)
	}

	if p.ReplyMarkup != nil {
		data, err := ParseReplyMarkup(p.ReplyMarkup)
		if err != nil {
			return nil, err
		}

		params["reply_markup"] = string(data)
	}

	return params, nil
}
