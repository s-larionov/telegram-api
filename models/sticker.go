package models

const (
	MaskPositionPointForehead MaskPositionPoint = "forehead"
	MaskPositionPointEyes     MaskPositionPoint = "eyes"
	MaskPositionPointMouth    MaskPositionPoint = "mouth"
	MaskPositionPointChin     MaskPositionPoint = "chin"
)

// The part of the face relative to which the mask should be placed.
type MaskPositionPoint string

// This object represents a sticker set.
type StickerSet struct {
	// Sticker set name
	Name string `json:"name"`

	// Sticker set title
	Title string `json:"title"`

	// True, if the sticker set contains animated stickers
	IsAnimated bool `json:"is_animated"`

	// True, if the sticker set contains masks
	ContainsMasks bool `json:"contains_masks"`

	// List of all set stickers
	Stickers []Sticker `json:"stickers"`
}

// This object represents a sticker.
type Sticker struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`

	// Unique identifier for this file, which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Sticker width
	Width int `json:"width"`

	// Sticker height
	Height int `json:"height"`

	// True, if the sticker is animated
	IsAnimated bool `json:"is_animated"`

	// Optional. Sticker thumbnail in the .WEBP or .JPG format
	Thumb *PhotoSize `json:"thumb,omitempty"`

	// Optional. Emoji associated with the sticker
	Emoji string `json:"emoji,omitempty"`

	// Optional. Name of the sticker set to which the sticker belongs
	SetName string `json:"set_name,omitempty"`

	// Optional. For mask stickers, the position where the mask should be placed
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`

	// Optional. File size
	FileSize int `json:"file_size,omitempty"`
}

// This object describes the position on faces where a mask should be placed by default.
type MaskPosition struct {
	// The part of the face relative to which the mask should be placed. One of “forehead”, “eyes”, “mouth”, or “chin”.
	Point MaskPositionPoint `json:"point"`

	// Shift by X-axis measured in widths of the mask scaled to the face size, from left to right.
	// For example, choosing -1.0 will place mask just to the left of the default mask position.
	XShift float32 `json:"x_shift"`

	// Shift by Y-axis measured in heights of the mask scaled to the face size, from top to bottom.
	// For example, 1.0 will place the mask just below the default mask position.
	YShift float32 `json:"y_shift"`

	// Mask scaling coefficient. For example, 2.0 means double size.
	Scale float32 `json:"scale"`
}
