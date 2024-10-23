package aquagram

import "context"

type MediaGroup []InputMedia

type InputMedia interface {
	InputMediaParams() InputMediaParams
}

type MediaType int8

const (
	mediaTypeUnknown MediaType = iota
	MediaTypeAnimation
	MediaTypeAudio
	MediaTypeDocument
	MediaTypePhoto
	MediaTypeVideo
)

func (media MediaType) String() string {
	switch media {
	case MediaTypeAnimation:
		return "animation"
	case MediaTypeAudio:
		return "audio"
	case MediaTypeDocument:
		return "document"
	case MediaTypePhoto:
		return "photo"
	case MediaTypeVideo:
		return "video"
	default:
		return ""
	}
}

type InputMediaParams struct {
	Type                        MediaType       `json:"type"`
	Media                       *InputFile      `json:"media"`
	Thumbnail                   *InputFile      `json:"thumbnail,omitempty"`
	Caption                     string          `json:"caption,omitempty"`
	ParseMode                   ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities             []MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia       bool            `json:"show_caption_above_media,omitempty"`
	Width                       int             `json:"width,omitempty"`
	Height                      int             `json:"height,omitempty"`
	Duration                    int             `json:"duration,omitempty"`
	HasSpoiler                  bool            `json:"has_spoiler,omitempty"`
	DisableContentTypeDetection bool            `json:"disable_content_type_detection,omitempty"`
	Performer                   string          `json:"performer,omitempty"`
	Title                       string          `json:"title,omitempty"`
	SupportStreaming            bool            `json:"supports_streaming,omitempty"`
}

func (params InputMediaParams) Params(bot *Bot) (Params, error) {
	common := CommonSendParams{
		Type:                        params.Type,
		Caption:                     params.Caption,
		ParseMode:                   params.ParseMode,
		CaptionEntities:             params.CaptionEntities,
		ShowCaptionAboveMedia:       params.ShowCaptionAboveMedia,
		Width:                       params.Width,
		Height:                      params.Height,
		Duration:                    params.Duration,
		HasSpoiler:                  params.HasSpoiler,
		DisableContentTypeDetection: params.DisableContentTypeDetection,
		Performer:                   params.Performer,
		Title:                       params.Title,
		SupportStreaming:            params.SupportStreaming,
	}

	return common.Params(bot)
}

// This object represents one size of a photo or a file / sticker thumbnail.
//
// https://core.telegram.org/bots/api#photosize
type PhotoSize struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	FileSize     int    `json:"file_size"`
}

// This object represents an animation file (GIF or H.264/MPEG-4 AVC video without sound).
//
// https://core.telegram.org/bots/api#animation
type Animation struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Duration     int        `json:"duration"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileName     string     `json:"file_name,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	FileSize     int64      `json:"file_size,omitempty"`
}

// This object represents an audio file to be treated as music by the Telegram clients.
//
// https://core.telegram.org/bots/api#audio
type Audio struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Duration     int        `json:"duration"`
	Performer    string     `json:"performer,omitempty"`
	Title        string     `json:"title,omitempty"`
	FileName     string     `json:"file_name,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	FileSize     int64      `json:"file_size,omitempty"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
}

// This object represents a general file (as opposed to photos, voice messages and audio files).
//
// https://core.telegram.org/bots/api#document
type Document struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileName     string     `json:"file_name,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	FileSize     int64      `json:"file_size,omitempty"`
}

// This object represents a story.
//
// https://core.telegram.org/bots/api#story
type Story struct {
	Chat *Chat `json:"chat"`
	ID   int   `json:"id"`
}

// This object represents a video file.
//
// https://core.telegram.org/bots/api#video
type Video struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Width        int        `json:"width"`
	Height       int        `json:"height"`
	Duration     int        `json:"duration"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileName     string     `json:"file_name,omitempty"`
	MimeType     string     `json:"mime_type,omitempty"`
	FileSize     int64      `json:"file_size,omitempty"`
}

// This object represents a video message (available in Telegram apps as of v.4.0).
//
// https://core.telegram.org/bots/api#videonote
type VideoNote struct {
	FileID       string     `json:"file_id"`
	FileUniqueID string     `json:"file_unique_id"`
	Lenght       int        `json:"length"`
	Duration     int        `json:"duration"`
	Thumbnail    *PhotoSize `json:"thumbnail,omitempty"`
	FileSize     int64      `json:"file_size,omitempty"`
}

// This object represents a voice note.
//
// https://core.telegram.org/bots/api#voice
type Voice struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Duration     int    `json:"duration"`
	MimeType     string `json:"mime_type,omitempty"`
	FileSize     int64  `json:"file_size,omitempty"`
}

// Represents an animation file (GIF or H.264/MPEG-4 AVC video without sound) to be sent.
//
// https://core.telegram.org/bots/api#inputmediaanimation
type InputMediaAnimation struct {
	Type                  string          `json:"type"`
	Media                 *InputFile      `json:"media"`
	Thumbnail             *InputFile      `json:"thumbnail,omitempty"`
	Caption               string          `json:"caption,omitempty"`
	ParseMode             ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool            `json:"show_caption_above_media,omitempty"`
	Width                 int             `json:"width,omitempty"`
	Height                int             `json:"height,omitempty"`
	Duration              int             `json:"duration,omitempty"`
	HasSpoiler            bool            `json:"has_spoiler,omitempty"`
}

func (media *InputMediaAnimation) InputMediaParams() InputMediaParams {
	return InputMediaParams{
		Type:                  MediaTypeAnimation,
		Media:                 media.Media,
		Thumbnail:             media.Thumbnail,
		Caption:               media.Caption,
		ParseMode:             media.ParseMode,
		CaptionEntities:       media.CaptionEntities,
		ShowCaptionAboveMedia: media.ShowCaptionAboveMedia,
		Width:                 media.Width,
		Height:                media.Height,
		Duration:              media.Duration,
		HasSpoiler:            media.HasSpoiler,
	}
}

// Represents a general file to be sent.
//
// https://core.telegram.org/bots/api#inputmediadocument
type InputMediaDocument struct {
	Type                        string          `json:"type"`
	Media                       *InputFile      `json:"media"`
	Thumbnail                   *InputFile      `json:"thumbnail,omitempty"`
	Caption                     string          `json:"caption,omitempty"`
	ParseMode                   ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities             []MessageEntity `json:"caption_entities,omitempty"`
	DisableContentTypeDetection bool            `json:"disable_content_type_detection,omitempty"`
}

func (media *InputMediaDocument) InputMediaParams() InputMediaParams {
	return InputMediaParams{
		Type:                        MediaTypeDocument,
		Media:                       media.Media,
		Thumbnail:                   media.Thumbnail,
		Caption:                     media.Caption,
		ParseMode:                   media.ParseMode,
		CaptionEntities:             media.CaptionEntities,
		DisableContentTypeDetection: media.DisableContentTypeDetection,
	}
}

// Represents an audio file to be treated as music to be sent.
//
// https://core.telegram.org/bots/api#inputmediaaudio
type InputMediaAudio struct {
	Type            string          `json:"type"`
	Media           *InputFile      `json:"media"`
	Thumbnail       *InputFile      `json:"thumbnail,omitempty"`
	Caption         string          `json:"caption,omitempty"`
	ParseMode       ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	Duration        int             `json:"duration,omitempty"`
	Performer       string          `json:"performer,omitempty"`
	Title           string          `json:"title,omitempty"`
}

func (media *InputMediaAudio) InputMediaParams() InputMediaParams {
	return InputMediaParams{
		Type:            MediaTypeAudio,
		Media:           media.Media,
		Thumbnail:       media.Thumbnail,
		Caption:         media.Caption,
		ParseMode:       media.ParseMode,
		CaptionEntities: media.CaptionEntities,
		Duration:        media.Duration,
		Performer:       media.Performer,
		Title:           media.Title,
	}
}

// Represents a photo to be sent.
//
// https://core.telegram.org/bots/api#inputmediaphoto
type InputMediaPhoto struct {
	Type                  string          `json:"type"`
	Media                 *InputFile      `json:"media"`
	Caption               string          `json:"caption,omitempty"`
	ParseMode             ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool            `json:"show_caption_above_media,omitempty"`
	HasSpoiler            bool            `json:"has_spoiler,omitempty"`
}

func (media *InputMediaPhoto) InputMediaParams() InputMediaParams {
	return InputMediaParams{
		Type:                  MediaTypePhoto,
		Media:                 media.Media,
		Caption:               media.Caption,
		ParseMode:             media.ParseMode,
		CaptionEntities:       media.CaptionEntities,
		ShowCaptionAboveMedia: media.ShowCaptionAboveMedia,
		HasSpoiler:            media.HasSpoiler,
	}
}

// Represents a video to be sent.
//
// https://core.telegram.org/bots/api#inputmediavideo
type InputMediaVideo struct {
	Type                  string          `json:"type"`
	Media                 *InputFile      `json:"media"`
	Thumbnail             *InputFile      `json:"thumbnail,omitempty"`
	Caption               string          `json:"caption,omitempty"`
	ParseMode             ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities       []MessageEntity `json:"caption_entities,omitempty"`
	ShowCaptionAboveMedia bool            `json:"show_caption_above_media,omitempty"`
	Width                 int             `json:"width,omitempty"`
	Height                int             `json:"height,omitempty"`
	Duration              int             `json:"duration,omitempty"`
	SupportStreaming      bool            `json:"supports_streaming,omitempty"`
	HasSpoiler            bool            `json:"has_spoiler,omitempty"`
}

func (media *InputMediaVideo) InputMediaParams() InputMediaParams {
	return InputMediaParams{
		Type:                  MediaTypeVideo,
		Media:                 media.Media,
		Thumbnail:             media.Thumbnail,
		Caption:               media.Caption,
		ParseMode:             media.ParseMode,
		CaptionEntities:       media.CaptionEntities,
		ShowCaptionAboveMedia: media.ShowCaptionAboveMedia,
		Width:                 media.Width,
		Height:                media.Height,
		Duration:              media.Duration,
		SupportStreaming:      media.SupportStreaming,
		HasSpoiler:            media.HasSpoiler,
	}
}

type SendAudioParams struct {
	BusinessConnectionID string           `json:"business_connection_id,omitempty"`
	MessageThreadID      int64            `json:"message_thread_id,omitempty"`
	Audio                *InputFile       `json:"audio"`
	Caption              string           `json:"caption,omitempty"`
	ParseMode            ParseMode        `json:"parse_mode,omitempty"`
	CaptionEntities      []MessageEntity  `json:"caption_entities,omitempty"`
	Duration             int              `json:"duration,omitempty"`
	Performer            string           `json:"performer,omitempty"`
	Title                string           `json:"title,omitempty"`
	Thumbnail            *InputFile       `json:"thumbnail,omitempty"`
	DisableNotification  bool             `json:"disable_notification,omitempty"`
	ProtectContent       bool             `json:"protect_content,omitempty"`
	MessageEffectID      string           `json:"message_effect_id,omitempty"`
	ReplyParameters      *ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup          ReplyMarkup      `json:"reply_markup,omitempty"`
}

/*
[sendAudio] - Use this method to send audio files.

If you want Telegram clients to display them in the music player.
Your audio must be in the .MP3 or .M4A format.

[sendAudio]: https://core.telegram.org/bots/api#sendaudio
*/
func (bot *Bot) SendAudio(chatID string, audio *InputFile, params *SendAudioParams) (*Message, error) {
	return bot.SendAudioWithContext(bot.Context(), chatID, audio, params)
}

func (bot *Bot) SendAudioWithContext(ctx context.Context, chatID string, audio *InputFile, params *SendAudioParams) (*Message, error) {
	if params == nil {
		params = new(SendAudioParams)
	}

	sendParams := CommonSendParams{
		ChatID:               chatID,
		BusinessConnectionID: params.BusinessConnectionID,
		MessageThreadID:      params.MessageThreadID,
		Caption:              params.Caption,
		ParseMode:            params.ParseMode,
		CaptionEntities:      params.CaptionEntities,
		Duration:             params.Duration,
		Performer:            params.Performer,
		Title:                params.Title,
		DisableNotification:  params.DisableNotification,
		ProtectContent:       params.ProtectContent,
		MessageEffectID:      params.MessageEffectID,
		ReplyParameters:      params.ReplyParameters,
		ReplyMarkup:          params.ReplyMarkup,
	}

	files := Files{}
	files["audio"] = audio

	if params.Thumbnail != nil {
		files["thumbnail"] = params.Thumbnail
	}

	paramsMap, err := sendParams.Params(bot)
	if err != nil {
		return nil, err
	}

	data, err := bot.RawFile(ctx, "sendAudio", paramsMap, files)
	if err != nil {
		return nil, err
	}

	return ParseRawResult[*Message](bot, data)
}

type SendDocumentParams struct {
	BusinessConnectionID        string           `json:"business_connection_id,omitempty"`
	MessageThreadID             int64            `json:"message_thread_id,omitempty"`
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

/*
[sendDocument] - Use this method to send general files.

[sendDocument]: https://core.telegram.org/bots/api#senddocument
*/
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

	paramsMap, err := sendParams.Params(bot)
	if err != nil {
		return nil, err
	}

	data, err := bot.RawFile(ctx, "sendDocument", paramsMap, files)
	if err != nil {
		return nil, err
	}

	return ParseRawResult[*Message](bot, data)
}

type SendPhotoParams struct {
	BusinessConnectionID        string           `json:"business_connection_id,omitempty"`
	ChatID                      string           `json:"chat_id"`
	MessageThreadID             int64            `json:"message_thread_id,omitempty"`
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

/*
[sendPhoto] - Use this method to send photos.

[sendPhoto]: https://core.telegram.org/bots/api#sendphoto
*/
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

	paramsMap, err := sendParams.Params(bot)
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

	return ParseRawResult[*Message](bot, data)
}

type SendVideoParams struct {
	BusinessConnectionID        string           `json:"business_connection_id,omitempty"`
	MessageThreadID             int64            `json:"message_thread_id,omitempty"`
	Video                       *InputFile       `json:"video"`
	Duration                    int              `json:"duration,omitempty"`
	Width                       int              `json:"width,omitempty"`
	Height                      int              `json:"height,omitempty"`
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

func (bot *Bot) SendVideo(chatID string, video *InputFile, params *SendVideoParams) (*Message, error) {
	return bot.SendVideoWithContext(bot.Context(), chatID, video, params)
}

func (bot *Bot) SendVideoWithContext(ctx context.Context, chatID string, video *InputFile, params *SendVideoParams) (*Message, error) {
	if params == nil {
		params = new(SendVideoParams)
	}

	sendParams := CommonSendParams{
		BusinessConnectionID: params.BusinessConnectionID,
		ChatID:               chatID,
		MessageThreadID:      params.MessageThreadID,
		Duration:             params.Duration,
		Width:                params.Width,
		Height:               params.Height,
		Caption:              params.Caption,
		ParseMode:            params.ParseMode,
		CaptionEntities:      params.CaptionEntities,

		DisableContentTypeDetection: params.DisableContentTypeDetection,
		ProtectContent:              params.ProtectContent,
		MessageEffectID:             params.MessageEffectID,
		ReplyParameters:             params.ReplyParameters,
		ReplyMarkup:                 params.ReplyMarkup,
	}

	paramsMap, err := sendParams.Params(bot)
	if err != nil {
		return nil, err
	}

	files := Files{}
	files["video"] = video

	if params.Thumbnail != nil {
		files["thumbnail"] = params.Thumbnail
	}

	data, err := bot.RawFile(ctx, "sendVideo", paramsMap, files)
	if err != nil {
		return nil, err
	}

	return ParseRawResult[*Message](bot, data)
}
