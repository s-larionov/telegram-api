package models

const (
	ChatTypePrivate    ChatType = "private"
	ChatTypeGroup      ChatType = "group"
	ChatTypeSuperGroup ChatType = "supergroup"
	ChatTypeChannel    ChatType = "channel"

	ChatMemberStatusCreator       ChatMemberStatus = "creator"
	ChatMemberStatusAdministrator ChatMemberStatus = "administrator"
	ChatMemberStatusMember        ChatMemberStatus = "member"
	ChatMemberStatusRestricted    ChatMemberStatus = "restricted"
	ChatMemberStatusLeft          ChatMemberStatus = "left"
	ChatMemberStatusKicked        ChatMemberStatus = "kicked"

	ChatActionTyping          ChatAction = "typing"
	ChatActionUploadPhoto     ChatAction = "upload_photo"
	ChatActionRecordVideo     ChatAction = "record_video"
	ChatActionUploadVideo     ChatAction = "upload_video"
	ChatActionRecordAudio     ChatAction = "record_audio"
	ChatActionUploadAudio     ChatAction = "upload_audio"
	ChatActionUploadDocument  ChatAction = "upload_document"
	ChatActionFindLocation    ChatAction = "find_location"
	ChatActionRecordVideoNote ChatAction = "record_video_note"
	ChatActionUploadVideoNote ChatAction = "upload_video_note"
)

// ChatType of chat, can be either “private”, “group”, “supergroup” or “channel”
type ChatType string

// ChatMemberStatus The member's status in the chat.
// Can be “creator”, “administrator”, “member”, “restricted”, “left” or “kicked”
type ChatMemberStatus string

type ChatAction string

// Chat This object represents a chat.
type Chat struct {
	// Unique identifier for this chat.
	// This number may be greater than 32 bits and some programming languages may have difficulty/silent defects
	// in interpreting it. But it is smaller than 52 bits, so a signed 64 bit integer or double-precision
	// float type are safe for storing this identifier.
	ID int64 `json:"id"`

	// ChatType of chat, can be either “private”, “group”, “supergroup” or “channel”
	Type ChatType `json:"type"`

	// Optional. Title, for supergroups, channels and group chats
	Title string `json:"title,omitempty"`

	// Optional. Username, for private chats, supergroups and channels if available
	Username string `json:"username,omitempty"`

	// Optional. First name of the other party in a private chat
	FirstName string `json:"first_name,omitempty"`

	// Optional. Last name of the other party in a private chat
	LastName string `json:"last_name,omitempty"`

	// Optional. Chat photo.
	// Returned only in getChat.
	Photo *ChatPhoto `json:"photo,omitempty"`

	// Optional. Description, for groups, supergroups and channel chats.
	// Returned only in getChat.
	Description string `json:"description,omitempty"`

	// Optional. Chat invite link, for groups, supergroups and channel chats. Each administrator in a chat generates
	// their own invite links, so the bot must first generate the link using exportChatInviteLink.
	// Returned only in getChat.
	InviteLink string `json:"invite_link,omitempty"`

	// Optional. Pinned message, for groups, supergroups and channels.
	// Returned only in getChat.
	PinnedMessage *Message `json:"pinned_message,omitempty"`

	// Optional. Default chat member permissions, for groups and supergroups.
	// Returned only in getChat.
	Permissions *ChatPermissions `json:"permissions,omitempty"`

	// Optional. For supergroups, the minimum allowed delay between consecutive messages sent by each unpriviledged user.
	// Returned only in getChat.
	SlowModeDelay int `json:"slow_mode_delay,omitempty"`

	// Optional. For supergroups, name of group sticker set.
	// Returned only in getChat.
	StickerSetName string `json:"sticker_set_name,omitempty"`

	// Optional. True, if the bot can change the group sticker set.
	// Returned only in getChat.
	CanSetStickerSet bool `json:"can_set_sticker_set,omitempty"`
}

// ChatPermissions Describes actions that a non-administrator user is allowed to take in a chat.
type ChatPermissions struct {
	// Optional. True, if the user is allowed to send text messages, contacts, locations and venues
	CanSendMessages bool `json:"can_send_messages"`

	// Optional. True, if the user is allowed to send audios, documents, photos, videos, video notes and voice notes,
	// implies can_send_messages
	CanSendMediaMessages bool `json:"can_send_media_messages,omitempty"`

	// Optional. True, if the user is allowed to send polls, implies can_send_messages
	CanSendPolls bool `json:"can_send_polls,omitempty"`

	// Optional. True, if the user is allowed to send animations, games, stickers and use inline bots,
	// implies can_send_media_messages
	CanSendOtherMessages bool `json:"can_send_other_messages,omitempty"`

	// Optional. True, if the user is allowed to add web page previews to their messages, implies can_send_media_messages
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`

	// Optional. True, if the user is allowed to change the chat title, photo and other settings.
	// Ignored in public supergroups
	CanChangeInfo bool `json:"can_change_info,omitempty"`

	// Optional. True, if the user is allowed to invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users,omitempty"`

	// Optional. True, if the user is allowed to pin messages. Ignored in public supergroups
	CanPinMessages bool `json:"can_pin_messages,omitempty"`
}

// ChatPhoto This object represents a chat photo.
type ChatPhoto struct {
	// File identifier of small (160x160) chat photo. This file_id can be used only for photo download
	// and only for as long as the photo is not changed.
	SmallFileID string `json:"small_file_id"`

	// Unique file identifier of small (160x160) chat photo, which is supposed to be the same over time
	// and for different bots. Can't be used to download or reuse the file.
	SmallFileUniqueID string `json:"small_file_unique_id"`

	// File identifier of big (640x640) chat photo. This file_id can be used only for photo download and
	// only for as long as the photo is not changed.
	BigFileID string `json:"big_file_id"`

	// Unique file identifier of big (640x640) chat photo, which is supposed to be the same over time
	// and for different bots. Can't be used to download or reuse the file.
	BigFileUniqueID string `json:"big_file_unique_id"`
}

// ChatMember This object contains information about one member of a chat.
type ChatMember struct {
	// Information about the user
	User *User `json:"user"`

	// The member's status in the chat.
	Status ChatMemberStatus `json:"status"`

	// Optional. Owner and administrators only. Custom title for this user
	CustomTitle string `json:"custom_title,omitempty"`

	// Optional. Restricted and kicked only. Date when restrictions will be lifted for this user; unix time
	UntilTimestamp int64 `json:"until_date"`

	// Optional. Administrators only. True, if the bot is allowed to edit administrator privileges of that user
	CanBeEdited bool `json:"can_be_edited,omitempty"`

	// Optional. Administrators only. True, if the administrator can post in the channel; channels only
	CanPostMessages bool `json:"can_post_messages,omitempty"`

	// Optional. Administrators only. True, if the administrator can edit messages of other users and can
	// pin messages; channels only
	CanEditMessages bool `json:"can_edit_messages,omitempty"`

	// Optional. Administrators only. True, if the administrator can delete messages of other users
	CanDeleteMessages bool `json:"can_delete_messages,omitempty"`

	// Optional. Administrators only. True, if the administrator can restrict, ban or unban chat members
	CanRestrictMembers bool `json:"can_restrict_members,omitempty"`

	// Optional. Administrators only. True, if the administrator can add new administrators with a subset
	// of his own privileges or demote administrators that he has promoted, directly or indirectly (promoted by
	// administrators that were appointed by the user)
	CanPromoteMembers bool `json:"can_promote_members,omitempty"`

	// Optional. Administrators and restricted only. True, if the user is allowed to change the chat title,
	// photo and other settings
	CanChangeInfo bool `json:"can_change_info,omitempty"`

	// Optional. Administrators and restricted only. True, if the user is allowed to invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users,omitempty"`

	// Optional. Administrators and restricted only. True, if the user is allowed to pin messages;
	// groups and supergroups only
	CanPinMessages bool `json:"can_pin_messages,omitempty"`

	// Optional. Restricted only. True, if the user is a member of the chat at the moment of the request
	IsMember bool `json:"is_member,omitempty"`

	// Optional. Restricted only. True, if the user is allowed to send text messages, contacts, locations and venues
	CanSendMessages bool `json:"can_send_messages,omitempty"`

	// Optional. Restricted only. True, if the user is allowed to send audios, documents, photos, videos, video notes
	// and voice notes
	CanSendMediaMessages bool `json:"can_send_media_messages,omitempty"`

	// Optional. Restricted only. True, if the user is allowed to send polls
	CanSendPolls bool `json:"can_send_polls,omitempty"`

	// Optional. Restricted only. True, if the user is allowed to send animations, games, stickers and use inline bots
	CanSendOtherMessages bool `json:"can_send_other_messages,omitempty"`

	// Optional. Restricted only. True, if the user is allowed to add web page previews to their messages
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`
}

// ChatMemberRestrictionsRequest Use this entity to restrict a user in a supergroup.
type ChatMemberRestrictionsRequest struct {
	// Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatID string `json:"chat_id"`

	// Unique identifier of the target user
	UserID int64 `json:"user_id"`

	// New user permissions
	Permissions ChatPermissions `json:"permissions"`

	// Date when restrictions will be lifted for the user, unix time. If user is restricted for more than 366 days
	// or less than 30 seconds from the current time, they are considered to be restricted forever
	UntilTimestamp int64 `json:"until_date,omitempty"`
}

// ChatMemberPromotionRequest Use this entity to promote or demote a user in a supergroup or a channel
type ChatMemberPromotionRequest struct {
	// Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
	ChatID string `json:"chat_id"`

	// Unique identifier of the target user
	UserID int64 `json:"user_id"`

	// Optional. Pass True, if the administrator can change chat title, photo and other settings
	CanChangeInfo bool `json:"can_change_info,omitempty"`

	// Optional. Pass True, if the administrator can create channel posts, channels only
	CanPostMessages bool `json:"can_post_messages,omitempty"`

	// Optional. Pass True, if the administrator can edit messages of other users and can pin messages, channels only
	CanEditMessages bool `json:"can_edit_messages,omitempty"`

	// Optional. Pass True, if the administrator can delete messages of other users
	CanDeleteMessages bool `json:"can_delete_messages,omitempty"`

	// Optional. Pass True, if the administrator can invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users,omitempty"`

	// Optional. Pass True, if the administrator can restrict, ban or unban chat members
	CanRestrictMembers bool `json:"can_restrict_members,omitempty"`

	// Optional. Pass True, if the administrator can pin messages, supergroups only
	CanPinMessages bool `json:"can_pin_messages,omitempty"`

	// Optional. Pass True, if the administrator can add new administrators with a subset of his own privileges
	// or demote administrators that he has promoted, directly or indirectly (promoted by administrators
	// that were appointed by him)
	CanPromoteMembers bool `json:"can_promote_members,omitempty"`
}

// ChatSetPhotoRequest Use this entity to set a new profile photo for the chat.
type ChatSetPhotoRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatID string `json:"chat_id"`

	// New chat photo, uploaded using multipart/form-data
	Photo InputFile `json:"photo"`
}
