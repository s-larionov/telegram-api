package models

// Video This object represents a video file.
type Video struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`

	// Unique identifier for this file, which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Video width as defined by sender
	Width int `json:"width"`

	// Video height as defined by sender
	Height int `json:"height"`

	// Duration of the video in seconds as defined by sender
	DurationInSeconds int `json:"duration"`

	// Optional. Document thumbnail as defined by sender
	Thumb *PhotoSize `json:"thumb,omitempty"`

	// Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type,omitempty"`

	// Optional. File size
	FileSize int64 `json:"file_size,omitempty"`
}
