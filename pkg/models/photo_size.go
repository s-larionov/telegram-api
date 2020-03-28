package models

// This object represents one size of a photo or a file / sticker thumbnail.
type PhotoSize struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`

	// Unique identifier for this file, which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// ChatPhoto width
	Width int `json:"width"`

	// ChatPhoto height
	Height int `json:"height"`

	// Optional. File size
	FileSize int `json:"file_size,omitempty"`
}
