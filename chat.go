package aquagram

import (
	"context"
	"strconv"
)

type ChatType string

const (
	ChatTypePrivate    ChatType = "private"
	ChatTypeGroup      ChatType = "group"
	ChatTypeSuperGroup ChatType = "supergroup"
	ChatTypeChannel    ChatType = "channel"
)

type ChatMemberStatus string

const (
	ChatMemberStatusCreator       ChatMemberStatus = "creator"
	ChatMemberStatusAdministrator ChatMemberStatus = "administrator"
	ChatMemberStatusMember        ChatMemberStatus = "member"
	ChatMemberStatusRestricted    ChatMemberStatus = "restricted"
	ChatMemberStatusLeft          ChatMemberStatus = "left"
	ChatMemberStatusKicked        ChatMemberStatus = "kicked"
)

type Chat struct {
	ID        int64    `json:"id"`
	Type      ChatType `json:"type"`
	Title     string   `json:"title"`
	Username  string   `json:"username"`
	FirstName string   `json:"first_name"`
	LastName  string   `json:"last_name"`
	IsForum   bool     `json:"is_forum"`
}

func ChatID(id int64) string {
	return strconv.FormatInt(id, 10)
}

func (chat *Chat) IsPrivate() bool {
	return chat.Type == ChatTypePrivate
}

func (chat *Chat) IsGroup() bool {
	return chat.Type == ChatTypeGroup
}

func (chat *Chat) IsSuperGroup() bool {
	return chat.Type == ChatTypeSuperGroup
}

func (chat *Chat) IsChannel() bool {
	return chat.Type == ChatTypeChannel
}

type ChatMember struct {
	ChatMemberAdministratorPermissions
	ChatMemberRestrictedPermissions

	Status      ChatMemberStatus `json:"status"`
	User        *User            `json:"user"`
	IsAnonymus  bool             `json:"is_anonymus,omitempty"`
	CustomTitle string           `json:"custom_title,omitempty"`
	UntilDate   int64            `json:"until_date,omitempty"`
}

func (member *ChatMember) IsOwner() bool {
	return member.Status == ChatMemberStatusCreator
}

func (member *ChatMember) IsAdministrator() bool {
	return member.Status == ChatMemberStatusAdministrator
}

func (member *ChatMember) IsMember() bool {
	return member.Status == ChatMemberStatusMember
}

func (member *ChatMember) IsRestricted() bool {
	return member.Status == ChatMemberStatusRestricted
}

func (member *ChatMember) IsLeft() bool {
	return member.Status == ChatMemberStatusLeft
}

func (member *ChatMember) IsKicked() bool {
	return member.Status == ChatMemberStatusKicked
}

type ChatMemberAdministratorPermissions struct {
	CanBeEdited         bool `json:"can_be_edited,omitempty"`
	CanManageChat       bool `json:"can_manage_chat,omitempty"`
	CanDeleteMessages   bool `json:"can_delete_messages,omitempty"`
	CanManageVideoChats bool `json:"can_manage_video_chats,omitempty"`
	CanRestrictMembers  bool `json:"can_restrict_members,omitempty"`
	CanPromoteMembers   bool `json:"can_promote_members,omitempty"`
	CanChangeInfo       bool `json:"can_change_info,omitempty"`
	CanInviteUsers      bool `json:"can_invite_users,omitempty"`
	CanPostStories      bool `json:"can_post_stories,omitempty"`
	CanEditStories      bool `json:"can_edit_stories,omitempty"`
	CanDeleteStories    bool `json:"can_delete_stories,omitempty"`
	CanPostMessages     bool `json:"can_post_messages,omitempty"`
	CanEditMessages     bool `json:"can_edit_messages,omitempty"`
	CanPinMessages      bool `json:"can_pin_messages,omitempty"`
	CanManageTopics     bool `json:"can_manage_topics,omitempty"`
}

type ChatMemberRestrictedPermissions struct {
	CanSendMessages       bool `json:"can_send_messages,omitempty"`
	CanSendAudios         bool `json:"can_send_audios,omitempty"`
	CanSendDocuments      bool `json:"can_send_documents,omitempty"`
	CanSendPhotos         bool `json:"can_send_photos,omitempty"`
	CanSendVideos         bool `json:"can_send_videos,omitempty"`
	CanSendVideoNotes     bool `json:"can_send_video_notes,omitempty"`
	CanSendVoiceNotes     bool `json:"can_send_voice_notes,omitempty"`
	CanSendPolls          bool `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  bool `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`

	// duplicated json tags
	//CanPinMessages        bool `json:"can_pin_messages"`
	//CanInviteUsers        bool `json:"can_invite_users"`
	//CanChangeInfo         bool `json:"can_change_info"`
}

/*
[SetChatTitle] wraps [SetChatTitleWithContext] using the default bot context.
*/
func (bot *Bot) SetChatTitle(chatID string, title string) error {
	return bot.SetChatTitleWithContext(bot.stopContext, chatID, title)
}

/*
[setChatTitle] - Use this method to change the title of a chat.

[setChatTitle]: https://core.telegram.org/bots/api#getchatadministrators
*/
func (bot *Bot) SetChatTitleWithContext(ctx context.Context, chatID string, title string) error {
	params := map[string]string{
		"chat_id": ParseChatID(chatID),
		"title":   title,
	}

	data, err := bot.Raw(ctx, "setChatTitle", params)
	if err != nil {
		return err
	}

	success, err := ParseRawResult[bool](bot, data)
	if err != nil {
		return err
	}

	if !success {
		return ErrExpectedTrue
	}

	return nil
}

/*
[SetChatDescription] wraps [SetChatDescriptionWithContext] using the default bot context.
*/
func (bot *Bot) SetChatDescription(chatID string, description string) error {
	return bot.SetChatDescriptionWithContext(bot.stopContext, chatID, description)
}

/*
[setChatDescription] - Use this method to change the description of a group, a supergroup or a channel.

[setChatDescription]: https://core.telegram.org/bots/api#setchatdescription
*/
func (bot *Bot) SetChatDescriptionWithContext(ctx context.Context, chatID string, description string) error {
	params := map[string]string{
		"chat_id":     ParseChatID(chatID),
		"description": description,
	}

	data, err := bot.Raw(ctx, "setChatDescription", params)
	if err != nil {
		return err
	}

	success, err := ParseRawResult[bool](bot, data)
	if err != nil {
		return err
	}

	if !success {
		return ErrExpectedTrue
	}

	return nil
}

/*
[LeaveChat] wraps [LeaveChatWithContext] using the default bot context.
*/
func (bot *Bot) LeaveChat(chatID string) error {
	return bot.LeaveChatWithContext(bot.stopContext, chatID)
}

/*
[leaveChat] - Use this method for your bot to leave a group, supergroup or channel.

[leaveChat]: https://core.telegram.org/bots/api#leavechat
*/
func (bot *Bot) LeaveChatWithContext(ctx context.Context, chatID string) error {
	params := map[string]string{
		"chat_id": ParseChatID(chatID),
	}

	data, err := bot.Raw(ctx, "leaveChat", params)
	if err != nil {
		return err
	}

	success, err := ParseRawResult[bool](bot, data)
	if err != nil {
		return err
	}

	if !success {
		return ErrExpectedTrue
	}

	return nil
}

/*
[GetChatAdministrators] wraps [GetChatAdministratorsWithContext] using the default bot context.
*/
func (bot *Bot) GetChatAdministrators(chatID string) ([]*ChatMember, error) {
	return bot.GetChatAdministratorsWithContext(bot.stopContext, chatID)
}

/*
[getChatAdministrators] - Use this method to get a list of administrators in a chat, which aren't bots.

[getChatAdministrators]: https://core.telegram.org/bots/api#getchatadministrators
*/
func (bot *Bot) GetChatAdministratorsWithContext(ctx context.Context, chatID string) ([]*ChatMember, error) {
	params := map[string]string{
		"chat_id": ParseChatID(chatID),
	}

	data, err := bot.Raw(ctx, "getChatAdministrators", params)
	if err != nil {
		return nil, err
	}

	return ParseRawResult[[]*ChatMember](bot, data)
}

/*
[GetChatMemberCount] wraps [GetChatMemberCountWithContext] using the default bot context.
*/
func (bot *Bot) GetChatMemberCount(chatID string) (int, error) {
	return bot.GetChatMemberCountWithContext(bot.stopContext, chatID)
}

/*
[getChatMemberCount] - Use this method to get the number of members in a chat.

[getChatMemberCount]: https://core.telegram.org/bots/api#getchatmembercount
*/
func (bot *Bot) GetChatMemberCountWithContext(ctx context.Context, chatID string) (int, error) {
	params := map[string]string{
		"chat_id": ParseChatID(chatID),
	}

	data, err := bot.Raw(ctx, "getChatMemberCount", params)
	if err != nil {
		return 0, err
	}

	return ParseRawResult[int](bot, data)
}

/*
[GetChatMember] wraps [GetChatMemberWithContext] using the default bot context.
*/
func (bot *Bot) GetChatMember(chatID string, userID int64) (*ChatMember, error) {
	return bot.GetChatMemberWithContext(bot.stopContext, chatID, userID)
}

/*
[getChatMember] - Use this method to get information about a member of a chat.

The method is only guaranteed to work for other users if the bot is an administrator in the chat.

[getChatMember]: https://core.telegram.org/bots/api#getchatmember
*/
func (bot *Bot) GetChatMemberWithContext(ctx context.Context, chatID string, userID int64) (*ChatMember, error) {
	params := map[string]any{
		"chat_id": ParseChatID(chatID),
		"user_id": userID,
	}

	data, err := bot.Raw(ctx, "getChatMember", params)
	if err != nil {
		return nil, err
	}

	return ParseRawResult[*ChatMember](bot, data)
}

/*
[SetChatStickerSet] wraps [SetChatStickerSetWithContext] using the default bot context.
*/
func (bot *Bot) SetChatStickerSet(chatID string, stickerSetName string) error {
	return bot.SetChatStickerSetWithContext(bot.stopContext, chatID, stickerSetName)
}

/*
[setChatStickerSet] - Use this method to set a new group sticker set for a supergroup.

The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method.

[setChatStickerSet]: https://core.telegram.org/bots/api#setchatstickerset
*/
func (bot *Bot) SetChatStickerSetWithContext(ctx context.Context, chatID string, stickerSetName string) error {
	params := map[string]string{
		"chat_id":          ParseChatID(chatID),
		"sticker_set_name": stickerSetName,
	}

	data, err := bot.Raw(ctx, "setChatStickerSet", params)
	if err != nil {
		return err
	}

	success, err := ParseRawResult[bool](bot, data)
	if err != nil {
		return err
	}

	if !success {
		return ErrExpectedTrue
	}

	return nil
}

/*
[DeleteChatStickerSet] wraps [DeleteChatStickerSetWithContext] using the default bot context.
*/
func (bot *Bot) DeleteChatStickerSet(chatID string) error {
	return bot.DeleteChatStickerSetWithContext(bot.stopContext, chatID)
}

/*
[deleteChatStickerSet] - Use this method to delete a group sticker set from a supergroup.

The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights.
Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method.

[deleteChatStickerSet]: https://core.telegram.org/bots/api#deletechatstickerset
*/
func (bot *Bot) DeleteChatStickerSetWithContext(ctx context.Context, chatID string) error {
	params := map[string]string{
		"chat_id": ParseChatID(chatID),
	}

	data, err := bot.Raw(ctx, "deleteChatStickerSet", params)
	if err != nil {
		return err
	}

	success, err := ParseRawResult[bool](bot, data)
	if err != nil {
		return err
	}

	if !success {
		return ErrExpectedTrue
	}

	return nil
}
