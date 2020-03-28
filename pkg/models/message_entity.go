package models

// This object represents one special entity in a text message. For example, hashtags, usernames, URLs, etc.
type MessageEntity struct {
	// Type of the entity.
	Type MessageEntityType `json:"type"`

	// Offset in UTF-16 code units to the start of the entity
	Offset int32 `json:"offset"`

	// Length of the entity in UTF-16 code units
	Length int32 `json:"length"`

	// Optional. For “text_link” only, url that will be opened after user taps on the text
	URL string `json:"url,omitempty"`

	// Optional. For “text_mention” only, the mentioned user
	User *User `json:"user,omitempty"`

	// Optional. For “pre” only, the programming language of the entity text
	Language string `json:"language,omitempty"`
}
