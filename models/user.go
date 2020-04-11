package models

// This object represents a Telegram user or bot.
type User struct {
	// Unique identifier for this user or bot
	ID int64 `json:"id"`

	// True, if this user is a bot
	IsBot bool `json:"is_bot"`

	// User‘s or bot’s first name
	FirstName string `json:"first_name"`

	// Optional. User‘s or bot’s last name
	LastName string `json:"last_name,omitempty"`

	// Optional. User‘s or bot’s username
	Username string `json:"username,omitempty"`

	// Optional. IETF language tag of the user's language
	LanguageCode string `json:"language_code,omitempty"`

	// Optional. True, if the bot can be invited to groups. Returned only in getMe.
	CanJoinGroups bool `json:"can_join_groups,omitempty"`

	// Optional. True, if privacy mode is disabled for the bot. Returned only in getMe.
	CanReadAllGroupMessages bool `json:"can_read_all_group_messages,omitempty"`

	// Optional. True, if the bot supports inline queries. Returned only in getMe.
	SupportsInlineQueries bool `json:"supports_inline_queries,omitempty"`
}

// This object represent a user's profile pictures.
type UserProfilePhotos struct {
	// Total number of profile pictures the target user has
	TotalCount int `json:"total_count"`

	// Requested profile pictures (in up to 4 sizes each)
	Photos [][]*PhotoSize `json:"photos"`
}

// Use this entity to get a list of profile pictures for a user.
type UserProfilePhotosRequest struct {
	// Unique identifier of the target user
	UserID int64 `json:"user_id"`

	// Sequential number of the first photo to be returned. By default, all photos are returned.
	Offset int `json:"offset,omitempty"`

	// Limits the number of photos to be retrieved. Values between 1—100 are accepted. Defaults to 100.
	Limit int `json:"limit,omitempty"`
}
