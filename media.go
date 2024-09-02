package aquagram

import (
	"encoding/json"
	"strconv"
)

type MediaGroup []InputMedia

type InputMedia interface {
	GetParams() InputMediaParams
	GetMedia() *InputFile
	GetThumbnail() *InputFile
}

type InputMediaParams struct {
	Type                        string          `json:"type"`
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

func (p *InputMediaParams) ToParams() (Params, error) {
	params := make(Params)

	params["type"] = p.Type

	if p.Caption != EmptyString {
		params["caption"] = p.Caption
	}

	if string(p.ParseMode) != EmptyString {
		params["parse_mode"] = string(p.ParseMode)
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

	return params, nil
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

func (media *InputMediaAnimation) GetParams() InputMediaParams {
	return InputMediaParams{
		Type:                  "animation",
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

func (media *InputMediaAnimation) GetMedia() *InputFile {
	return media.Media
}

func (media *InputMediaAnimation) GetThumbnail() *InputFile {
	return media.Thumbnail
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

func (media *InputMediaDocument) GetParams() InputMediaParams {
	return InputMediaParams{
		Type:                        "document",
		Caption:                     media.Caption,
		ParseMode:                   media.ParseMode,
		CaptionEntities:             media.CaptionEntities,
		DisableContentTypeDetection: media.DisableContentTypeDetection,
	}
}

func (media *InputMediaDocument) GetMedia() *InputFile {
	return media.Media
}

func (media *InputMediaDocument) GetThumbnail() *InputFile {
	return media.Thumbnail
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

func (media *InputMediaAudio) GetParams() InputMediaParams {
	return InputMediaParams{
		Type:            "audio",
		Caption:         media.Caption,
		ParseMode:       media.ParseMode,
		CaptionEntities: media.CaptionEntities,
		Duration:        media.Duration,
		Performer:       media.Performer,
		Title:           media.Title,
	}
}

func (media *InputMediaAudio) GetMedia() *InputFile {
	return media.Media
}

func (media *InputMediaAudio) GetThumbnail() *InputFile {
	return media.Thumbnail
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

func (media *InputMediaPhoto) GetParams() InputMediaParams {
	return InputMediaParams{
		Type:                  "photo",
		Caption:               media.Caption,
		ParseMode:             media.ParseMode,
		CaptionEntities:       media.CaptionEntities,
		ShowCaptionAboveMedia: media.ShowCaptionAboveMedia,
		HasSpoiler:            media.HasSpoiler,
	}
}

func (media *InputMediaPhoto) GetMedia() *InputFile {
	return media.Media
}

func (media *InputMediaPhoto) GetThumbnail() *InputFile {
	return nil
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

func (media *InputMediaVideo) GetParams() InputMediaParams {
	return InputMediaParams{
		Type:                  "video",
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

func (media *InputMediaVideo) GetMedia() *InputFile {
	return media.Media
}

func (media *InputMediaVideo) GetThumbnail() *InputFile {
	return nil
}
