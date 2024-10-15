package aquagram

import "context"

type ChatMemberStatus string

const (
	ChatMemberStatusCreator       ChatMemberStatus = "creator"
	ChatMemberStatusAdministrator ChatMemberStatus = "administrator"
	ChatMemberStatusMember        ChatMemberStatus = "member"
	ChatMemberStatusRestricted    ChatMemberStatus = "restricted"
	ChatMemberStatusLeft          ChatMemberStatus = "left"
	ChatMemberStatusKicked        ChatMemberStatus = "kicked"
)

/*
[ChatMember] - This object contains information about one member of a chat.

[ChatMember]: https://core.telegram.org/bots/api#chatmember
*/
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
	ChatPermissions
}

/*
[ChatPermissions] - Describes actions that a non-administrator user is allowed to take in a chat.

[ChatPermissions]: https://core.telegram.org/bots/api#chatpermissions
*/
type ChatPermissions struct {
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
	CanChangeInfo         bool `json:"can_change_info,omitempty"`
	CanInviteUsers        bool `json:"can_invite_users,omitempty"`
	CanPinMessages        bool `json:"can_pin_messages,omitempty"`
	CanManageTopics       bool `json:"can_manage_topics,omitempty"`
}

type BanChatMemberParams struct {
	ChatID         int64 `json:"chat_id"`
	UserID         int64 `json:"user_id"`
	UntilDate      int64 `json:"until_date,omitempty"`
	RevokeMessages bool  `json:"revoke_messages,omitempty"`
}

/*
[getChatMember] - Use this method to get information about a member of a chat.

The method is only guaranteed to work for other users if the bot is an administrator in the chat.

[getChatMember]: https://core.telegram.org/bots/api#getchatmember
*/
func (bot *Bot) GetChatMember(chatID string, userID int64) (*ChatMember, error) {
	return bot.GetChatMemberWithContext(bot.stopContext, chatID, userID)
}

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
[banChatMember] - Use this method to ban a user in a group, a supergroup or a channel.

[banChatMember]: https://core.telegram.org/bots/api#banchatmember
*/
func (bot *Bot) BanChatMember(chatID int64, userID int64, params *BanChatMemberParams) error {
	return bot.BanChatMemberWithContext(bot.Context(), chatID, userID, params)
}

func (bot *Bot) BanChatMemberWithContext(ctx context.Context, chatID int64, userID int64, params *BanChatMemberParams) error {
	if params == nil {
		params = new(BanChatMemberParams)
	}

	params.ChatID = chatID
	params.UserID = userID

	data, err := bot.Raw(ctx, "banChatMember", params)
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

type UnbanChatMemberParams struct {
	ChatID       int64 `json:"chat_id"`
	UserID       int64 `json:"user_id"`
	OnlyIfBanned bool  `json:"only_if_banned,omitempty"`
}

/*
[unbanChatMember] - Use this method to unban a previously banned user in a supergroup or channel.

[unbanChatMember]: https://core.telegram.org/bots/api#unbanchatmember
*/
func (bot *Bot) UnbanChatMember(chatID int64, userID int64, params *UnbanChatMemberParams) error {
	return bot.UnbanChatMemberWithContext(bot.Context(), chatID, userID, params)
}

func (bot *Bot) UnbanChatMemberWithContext(ctx context.Context, chatID int64, userID int64, params *UnbanChatMemberParams) error {
	if params == nil {
		params = new(UnbanChatMemberParams)
	}

	params.ChatID = chatID
	params.UserID = userID

	data, err := bot.Raw(ctx, "unbanChatMember", params)
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

type RestrictChatMemberParams struct {
	ChatID                        int64           `json:"chat_id"`
	UserID                        int64           `json:"user_id"`
	Permissions                   ChatPermissions `json:"permissions"`
	UseIndependentChatPermissions bool            `json:"use_independent_chat_permissions,omitempty"`
}

/*
[restrictChatMember] - Use this method to restrict a user in a supergroup.

[restrictChatMember]: https://core.telegram.org/bots/api#restrictchatmember
*/
func (bot *Bot) RestrictChatMember(chatID, userID int64, permissions ChatPermissions, params *RestrictChatMemberParams) error {
	return bot.RestrictChatMemberWithContext(bot.Context(), chatID, userID, permissions, params)
}

func (bot *Bot) RestrictChatMemberWithContext(ctx context.Context, chatID, userID int64, permissions ChatPermissions, params *RestrictChatMemberParams) error {
	if params == nil {
		params = new(RestrictChatMemberParams)
	}

	params.ChatID = chatID
	params.UserID = userID
	params.Permissions = permissions

	data, err := bot.Raw(ctx, "restrictChatMember", params)
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
Use this method to remove all user permissions in a group.
*/
func (bot *Bot) MuteChatMember(chatID, userID int64) error {
	return bot.MuteChatMemberWithContext(bot.Context(), chatID, userID)
}

func (bot *Bot) MuteChatMemberWithContext(ctx context.Context, chatID, userID int64) error {
	permissions := ChatPermissions{}
	return bot.RestrictChatMemberWithContext(ctx, chatID, userID, permissions, nil)
}

type PromoteChatMemberParams struct {
	ChatMemberAdministratorPermissions

	ChatID int64 `json:"chat_id"`
	UserID int64 `json:"user_id"`
}

/*
[promoteChatMember] - Use this method to promote or demote a user in a supergroup or a channel.

[promoteChatMember]: https://core.telegram.org/bots/api#promotechatmember
*/
func (bot *Bot) PromoteChatMember(chatID, userID int64, params *PromoteChatMemberParams) error {
	return bot.PromoteChatMemberWithContext(bot.Context(), chatID, userID, params)
}

func (bot *Bot) PromoteChatMemberWithContext(ctx context.Context, chatID, userID int64, params *PromoteChatMemberParams) error {
	if params == nil {
		params = new(PromoteChatMemberParams)
	}

	params.ChatID = chatID
	params.UserID = userID

	data, err := bot.Raw(ctx, "promoteChatMember", params)
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

type SetChatAdministratorCustomTitleParams struct {
	ChatID      int64  `json:"chat_id"`
	UserID      int64  `json:"user_id"`
	CustomTitle string `json:"custom_title"`
}

/*
[setChatAdministratorCustomTitle] - Use this method to set a custom title for an administrator in a supergroup promoted by the bot.

[setChatAdministratorCustomTitle]: https://core.telegram.org/bots/api#setchatadministratorcustomtitle
*/
func (bot *Bot) SetChatAdministratorCustomTitle(chatID, userID int64, customTitle string) error {
	return bot.SetChatAdministratorCustomTitleWithContext(bot.Context(), chatID, userID, customTitle)
}

func (bot *Bot) SetChatAdministratorCustomTitleWithContext(ctx context.Context, chatID, userID int64, customTitle string) error {
	params := SetChatAdministratorCustomTitleParams{
		ChatID:      chatID,
		UserID:      userID,
		CustomTitle: customTitle,
	}

	data, err := bot.Raw(ctx, "setChatAdministratorCustomTitle", params)
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
