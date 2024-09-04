package aquagram

// https://core.telegram.org/bots/api#messageentity
type EntityType string

const (
	EntityTypeMention              EntityType = "mention"               // @username
	EntityTypeHashtag              EntityType = "hashtag"               // #hashtag
	EntityTypeCashtag              EntityType = "cashtag"               // $USD
	EntityTypeBotCommand           EntityType = "bot_command"           // /start@jobs_bot
	EntityTypeUrl                  EntityType = "url"                   // https://telegram.org
	EntityTypeEmail                EntityType = "email"                 // do-not-reply@telegram.org
	EntityTypePhoneNumber          EntityType = "phone_number"          // +1-212-555-0123
	EntityTypeBold                 EntityType = "bold"                  // bold text
	EntityTypeItalic               EntityType = "italic"                // italic text
	EntityTypeUnderline            EntityType = "underline"             // underlined text
	EntityTypeStrikethrough        EntityType = "strikethrough"         // strikethrough text
	EntityTypeSpoiler              EntityType = "spoiler"               // spoiler message
	EntityTypeBlockquote           EntityType = "blockquote"            // block quotation
	EntityTypeExpandableBlockquote EntityType = "expandable_blockquote" // collapsed-by-default block quotation
	EntityTypeCode                 EntityType = "code"                  // monowidth string
	EntityTypePre                  EntityType = "pre"                   // monowidth block
	EntityTypeTextLink             EntityType = "text_link"             // for clickable text URLs
	EntityTypeTextMention          EntityType = "text_mention"          // for users without usernames
	EntityTypeCustomEmoji          EntityType = "custom_emoji"          // for inline custom emoji stickers
)

// https://core.telegram.org/bots/api#messageentity
type MessageEntity struct {
	Message *Message `json:"-"`

	Type          EntityType `json:"type"`
	Offset        int        `json:"offset"`
	Length        int        `json:"length"`
	Url           string     `json:"url,omitempty"`
	User          *User      `json:"user,omitempty"`
	Language      string     `json:"language,omitempty"`
	CustomEmojiID string     `json:"custom_emoji_id,omitempty"`
}

func (entity *MessageEntity) Text() (text string) {
	var str string

	if entity.Message == nil {
		return
	}

	if entity.Message.Text != EmptyString {
		str = entity.Message.Text
	} else if entity.Message.Caption != EmptyString {
		str = entity.Message.Caption
	} else {
		return
	}

	stopIndex := entity.Offset + entity.Length

	for runeIndex, runeValue := range []rune(str) {
		if runeIndex < entity.Offset {
			continue
		}

		if runeIndex == stopIndex {
			break
		}

		text += string(runeValue)
	}

	return
}
