package models

// This object contains information about one member of a chat.
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
