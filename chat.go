package aquagram

import (
	"context"
	"fmt"
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
	return fmt.Sprintf("%d", id)
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

func (bot *Bot) GetChatMember(chatID string, userID int) (*ChatMember, error) {
	return bot.GetChatMemberWithContext(bot.stopContext, chatID, userID)
}

func (bot *Bot) GetChatMemberWithContext(ctx context.Context, chatID string, userID int) (*ChatMember, error) {
	params := map[string]any{
		"chat_id": chatID,
		"user_id": userID,
	}

	data, err := bot.Raw(ctx, "getChatMember", params)
	if err != nil {
		return nil, err
	}

	member, err := ParseRawResult[ChatMember](bot, data)
	if err != nil {
		return nil, err
	}

	return member, nil
}
