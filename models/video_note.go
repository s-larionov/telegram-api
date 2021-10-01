package models

// VideoNote This object represents a video message (available in Telegram apps as of v.4.0).
type VideoNote struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`

	// Unique identifier for this file, which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Video width and height (diameter of the video message) as defined by sender
	Length int `json:"length"`

	// Duration of the video in seconds as defined by sender
	DurationInSeconds int `json:"duration"`

	// Optional. Document thumbnail as defined by sender
	Thumb *PhotoSize `json:"thumb,omitempty"`

	// Optional. File size
	FileSize int64 `json:"file_size,omitempty"`
}
