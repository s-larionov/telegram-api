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

// Use this entity to send static .WEBP or animated .TGS stickers.
type SendStickerRequest struct {
	MessageRequestBase

	// Sticker to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL as a String for Telegram to get a .WEBP file from the Internet, or upload a new one
	// using multipart/form-data.
	Sticker InputFile `json:"sticker"`
}

// Use this entity to create a new sticker set owned by a user. The bot will be able to edit the sticker
// set thus created. You must use exactly one of the fields png_sticker or tgs_sticker.
type NewStickerSetRequest struct {
	// User identifier of created sticker set owner
	UserID int64 `json:"user_id"`

	// Short name of sticker set, to be used in t.me/addstickers/ URLs (e.g., animals). Can contain only english
	// letters, digits and underscores. Must begin with a letter, can't contain consecutive underscores and must
	// end in “_by_<bot username>”. <bot_username> is case insensitive. 1-64 characters.
	Name string `json:"name"`

	// Sticker set title, 1-64 characters
	Title string `json:"title"`

	// Optional. PNG image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px,
	// and either width or height must be exactly 512px. Pass a file_id as a String to send a file that already exists
	// on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload
	// a new one using multipart/form-data.
	PngSticker InputFile `json:"png_sticker,omitempty"`

	// Optional	TGS animation with the sticker, uploaded using multipart/form-data.
	// See https://core.telegram.org/animated_stickers#technical-requirements for technical requirements
	TgsSticker InputFile `json:"tgs_sticker,omitempty"`

	// One or more emoji corresponding to the sticker
	Emojis string `json:"emojis"`

	// Optional. Pass True, if a set of mask stickers should be created
	ContainsMasks bool `json:"contains_masks,omitempty"`

	// Optional. A JSON-serialized object for position where the mask should be placed on faces
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
}

// Use this entity to add a new sticker to a set created by the bot. You must use exactly one of the fields png_sticker
// or tgs_sticker. Animated stickers can be added to animated sticker sets and only to them. Animated sticker sets
// can have up to 50 stickers. Static sticker sets can have up to 120 stickers.
type AddStickerToSetSetRequest struct {
	// User identifier of sticker set owner
	UserID int64 `json:"user_id"`

	// Sticker set name
	Name string `json:"name"`

	// Optional. PNG image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px,
	// and either width or height must be exactly 512px. Pass a file_id as a String to send a file that already exists
	// on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload
	// a new one using multipart/form-data.
	PngSticker InputFile `json:"png_sticker,omitempty"`

	// Optional	TGS animation with the sticker, uploaded using multipart/form-data.
	// See https://core.telegram.org/animated_stickers#technical-requirements for technical requirements
	TgsSticker InputFile `json:"tgs_sticker,omitempty"`

	// One or more emoji corresponding to the sticker
	Emojis string `json:"emojis"`

	// Optional. A JSON-serialized object for position where the mask should be placed on faces
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
}

// Use this entity to set the thumbnail of a sticker set. Animated thumbnails can be set for animated sticker sets only.
type StickerSetThumbRequest struct {
	// Sticker set name
	Name string `json:"name"`

	// User identifier of the sticker set owner
	UserID int64 `json:"user_id"`

	// A PNG image with the thumbnail, must be up to 128 kilobytes in size and have width and height exactly 100px,
	// or a TGS animation with the thumbnail up to 32 kilobytes in size;
	// see https://core.telegram.org/animated_stickers#technical-requirements for animated sticker technical
	// requirements. Pass a file_id as a String to send a file that already exists on the Telegram servers,
	// pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using
	// multipart/form-data. More info on Sending Files ». Animated sticker set thumbnail can't be uploaded via HTTP URL.
	Thumb InputFile `json:"thumb,omitempty"`
}
