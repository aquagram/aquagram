package aquagram

import (
	"context"
)

type SendDocumentParams struct {
	BusinessConnectionID        string           `json:"business_connection_id,omitempty"`
	MessageThreadID             int              `json:"message_thread_id,omitempty"`
	Document                    *InputFile       `json:"document"`
	Thumbnail                   *InputFile       `json:"thumbnail,omitempty"`
	Caption                     string           `json:"caption,omitempty"`
	ParseMode                   ParseMode        `json:"parse_mode,omitempty"`
	CaptionEntities             []MessageEntity  `json:"caption_entities,omitempty"`
	DisableContentTypeDetection bool             `json:"disable_content_type_detection,omitempty"`
	DisableNotification         bool             `json:"disable_notification,omitempty"`
	ProtectContent              bool             `json:"protect_content,omitempty"`
	MessageEffectID             string           `json:"message_effect_id,omitempty"`
	ReplyParameters             *ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup                 ReplyMarkup      `json:"reply_markup,omitempty"`
}

func (bot *Bot) SendDocument(chatID string, document *InputFile, params *SendDocumentParams) (*Message, error) {
	return bot.SendDocumentWithContext(bot.stopContext, chatID, document, params)
}

func (bot *Bot) SendDocumentWithContext(ctx context.Context, chatID string, document *InputFile, params *SendDocumentParams) (*Message, error) {
	if params == nil {
		params = new(SendDocumentParams)
	}

	sendParams := CommonSendParams{
		ChatID:               chatID,
		BusinessConnectionID: params.BusinessConnectionID,
		MessageThreadID:      params.MessageThreadID,
		Caption:              params.Caption,
		ParseMode:            params.ParseMode,
		CaptionEntities:      params.CaptionEntities,
		DisableNotification:  params.DisableNotification,
		ProtectContent:       params.ProtectContent,
		MessageEffectID:      params.MessageEffectID,
		ReplyParameters:      params.ReplyParameters,
		ReplyMarkup:          params.ReplyMarkup,
	}

	files := Files{}
	files["document"] = document

	if params.Thumbnail != nil {
		files["thumbnail"] = params.Thumbnail
	}

	paramsMap, err := sendParams.ToParams()
	if err != nil {
		return nil, err
	}

	data, err := bot.RawFile(ctx, "sendDocument", paramsMap, files)
	if err != nil {
		return nil, err
	}

	message, err := ParseRawResult[Message](bot, data)
	if err != nil {
		return nil, err
	}

	return message, nil
}

type SendPhotoParams struct {
	BusinessConnectionID        string           `json:"business_connection_id,omitempty"`
	ChatID                      string           `json:"chat_id"`
	MessageThreadID             int              `json:"message_thread_id,omitempty"`
	Photo                       *InputFile       `json:"photo"`
	Thumbnail                   *InputFile       `json:"thumbnail,omitempty"`
	Caption                     string           `json:"caption,omitempty"`
	ParseMode                   ParseMode        `json:"parse_mode,omitempty"`
	CaptionEntities             []MessageEntity  `json:"caption_entities,omitempty"`
	DisableContentTypeDetection bool             `json:"disable_content_type_detection"`
	DisableNotification         bool             `json:"disable_notification,omitempty"`
	ProtectContent              bool             `json:"protect_content,omitempty"`
	MessageEffectID             string           `json:"message_effect_id,omitempty"`
	ReplyParameters             *ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup                 ReplyMarkup      `json:"reply_markup,omitempty"`
}

func (bot *Bot) SendPhoto(chatID string, photo *InputFile, params *SendPhotoParams) (*Message, error) {
	return bot.SendPhotoWithContext(bot.stopContext, chatID, photo, params)
}

func (bot *Bot) SendPhotoWithContext(ctx context.Context, chatID string, photo *InputFile, params *SendPhotoParams) (*Message, error) {
	if params == nil {
		params = new(SendPhotoParams)
	}

	sendParams := CommonSendParams{
		BusinessConnectionID:        params.BusinessConnectionID,
		ChatID:                      chatID,
		MessageThreadID:             params.MessageThreadID,
		Caption:                     params.Caption,
		ParseMode:                   params.ParseMode,
		CaptionEntities:             params.CaptionEntities,
		DisableContentTypeDetection: params.DisableContentTypeDetection,
		ProtectContent:              params.ProtectContent,
		MessageEffectID:             params.MessageEffectID,
		ReplyParameters:             params.ReplyParameters,
		ReplyMarkup:                 params.ReplyMarkup,
	}

	paramsMap, err := sendParams.ToParams()
	if err != nil {
		return nil, err
	}

	files := Files{}
	files["photo"] = photo

	if params.Thumbnail != nil {
		files["thumbnail"] = params.Thumbnail
	}

	data, err := bot.RawFile(ctx, "sendPhoto", paramsMap, files)
	if err != nil {
		return nil, err
	}

	message, err := ParseRawResult[Message](bot, data)
	if err != nil {
		return nil, err
	}

	return message, nil
}
