package aquagram

import (
	"encoding/json"
	"strconv"
)

type Params map[string]string

type CommonSendParams struct {
	Type                        MediaType        `json:"type,omitempty"`
	ChatID                      string           `json:"chat_id,omitempty"`
	Text                        string           `json:"text,omitempty"`
	BusinessConnectionID        string           `json:"business_connection_id,omitempty"`
	MessageThreadID             int64            `json:"message_thread_id,omitempty"`
	ParseMode                   ParseMode        `json:"parse_mode,omitempty"`
	Entities                    []MessageEntity  `json:"entities,omitempty"`
	DisableNotification         bool             `json:"disable_notification,omitempty"`
	Media                       []Params         `json:"media,omitempty"`
	Caption                     string           `json:"caption,omitempty"`
	CaptionEntities             []MessageEntity  `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia       bool             `json:"show_caption_above_media,omitempty"`
	Width                       int              `json:"width,omitempty"`
	Height                      int              `json:"height,omitempty"`
	Duration                    int              `json:"duration,omitempty"`
	HasSpoiler                  bool             `json:"has_spoiler,omitempty"`
	DisableContentTypeDetection bool             `json:"disable_content_type_detection,omitempty"`
	Performer                   string           `json:"performer,omitempty"`
	Title                       string           `json:"title,omitempty"`
	SupportStreaming            bool             `json:"supports_streaming,omitempty"`
	ProtectContent              bool             `json:"protect_content,omitempty"`
	MessageEffectID             string           `json:"message_effect_id,omitempty"`
	ReplyParameters             *ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup                 ReplyMarkup      `json:"reply_markup,omitempty"`
}

func (p *CommonSendParams) Params(bot *Bot) (Params, error) {
	params := make(Params)

	if p.Type != mediaTypeUnknown {
		params["type"] = p.Type.String()
	}

	params["chat_id"] = ParseChatID(p.ChatID)

	if p.Text != EmptyString {
		params["text"] = p.Text
	}

	if p.BusinessConnectionID != EmptyString {
		params["business_connection_id"] = p.BusinessConnectionID
	}

	if p.MessageThreadID != 0 {
		params["message_thread_id"] = strconv.FormatInt(p.MessageThreadID, 10)
	}

	if p.ParseMode != ParseModeDisabled {
		params["parse_mode"] = bot.ParseMode(p.ParseMode)
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

	if p.ShowCaptionAboveMedia {
		params["show_caption_above_media"] = TrueAsString
	}

	if p.Width != 0 {
		params["width"] = strconv.Itoa(p.Width)
	}

	if p.Height != 0 {
		params["height"] = strconv.Itoa(p.Height)
	}

	if p.Duration != 0 {
		params["duration"] = strconv.Itoa(p.Duration)
	}

	if p.HasSpoiler {
		params["has_spoiler"] = TrueAsString
	}

	if p.DisableContentTypeDetection {
		params["disable_content_type_detection"] = TrueAsString
	}

	if p.Performer != EmptyString {
		params["performer"] = p.Performer
	}

	if p.Title != EmptyString {
		params["title"] = p.Title
	}

	if p.SupportStreaming {
		params["support_streaming"] = TrueAsString
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
