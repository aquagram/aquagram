package aquagram

import (
	"context"
	"fmt"
	"strconv"
)

type SendMediaGroupParams struct {
	BusinessConnectionID string           `json:"business_connection_id,omitempty"`
	ChatID               string           `json:"chat_id"`
	MessageThreadID      int64            `json:"message_thread_id,omitempty"`
	Media                []*InputFile     `json:"media"`
	DisableNotification  bool             `json:"disable_notification,omitempty"`
	ProtectContent       bool             `json:"protect_content,omitempty"`
	MessageEffectID      string           `json:"message_effect_id,omitempty"`
	ReplyParameters      *ReplyParameters `json:"reply_parameters,omitempty"`
}

func (bot *Bot) SendMediaGroup(chatID string, media MediaGroup, params *SendMediaGroupParams) ([]*Message, error) {
	return bot.SendMediaGroupWithContext(bot.stopContext, chatID, media, params)
}

func (bot *Bot) SendMediaGroupWithContext(ctx context.Context, chatID string, media MediaGroup, params *SendMediaGroupParams) ([]*Message, error) {
	if params == nil {
		params = new(SendMediaGroupParams)
	}

	files := make(Files)
	mediaFiles := make([]Params, 0)

	for index, item := range media {
		itemParams := item.GetParams()

		itemParamsMap, err := itemParams.ToParams()
		if err != nil {
			return nil, err
		}

		itemMedia := item.GetMedia()
		fieldname := strconv.Itoa(index)

		if itemMedia.FromReader != nil || itemMedia.FromPath != EmptyString {
			itemParamsMap["media"] = fmt.Sprintf("attach://%s", fieldname)
			files[fieldname] = itemMedia

		} else if str := itemMedia.FromFileID; str != EmptyString {
			itemParamsMap["media"] = str

		} else if str := itemMedia.FromURL; str != EmptyString {
			itemParamsMap["media"] = str

		} else {
			return nil, ErrUnknownFileSource
		}

		mediaFiles = append(mediaFiles, itemParamsMap)
	}

	sendParams := CommonSendParams{
		BusinessConnectionID: params.BusinessConnectionID,
		ChatID:               chatID,
		Media:                mediaFiles,
		MessageThreadID:      params.MessageThreadID,
		DisableNotification:  params.DisableNotification,
		ProtectContent:       params.ProtectContent,
		MessageEffectID:      params.MessageEffectID,
		ReplyParameters:      params.ReplyParameters,
	}

	paramsMap, err := sendParams.ToParams()
	if err != nil {
		return nil, err
	}

	if len(files) == 0 {
		data, err := bot.Raw(ctx, "sendMediaGroup", paramsMap)
		if err != nil {
			return nil, err
		}

		return ParseRawResult[[]*Message](bot, data)
	}

	data, err := bot.RawFile(ctx, "sendMediaGroup", paramsMap, files)
	if err != nil {
		return nil, err
	}

	return ParseRawResult[[]*Message](bot, data)
}
