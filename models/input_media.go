package models

const (
	InputMediaTypePhoto     InputMediaType = "photo"
	InputMediaTypeVideo     InputMediaType = "video"
	InputMediaTypeAnimation InputMediaType = "animation"
	InputMediaTypeAudio     InputMediaType = "audio"
	InputMediaTypeDocument  InputMediaType = "document"
)

type InputMediaType string

type InputMediaInterface interface {
	GetType() InputMediaType
}

// This object represents the content of a media message to be sent.
type inputMedia struct {
	// Type of the result
	Type InputMediaType `json:"type"`
}

func (t inputMedia) GetType() InputMediaType {
	return t.Type
}

// Represents a photo to be sent.
type InputMediaPhoto struct {
	inputMedia

	// File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL
	// for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using
	// multipart/form-data under <file_attach_name> name.
	// [More info on Sending Files](https://core.telegram.org/bots/api#sending-files)
	Media string `json:"media"`

	// Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text
	// or inline URLs in the media caption.
	ParseMode ParseMode `json:"parse_mode,omitempty"`
}

// Represents a video to be sent.
type InputMediaVideo struct {
	// File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>”
	// to upload a new one using multipart/form-data under <file_attach_name> name.
	// [More info on Sending Files](https://core.telegram.org/bots/api#sending-files)
	Media string `json:"media"`

	// Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported
	// server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height
	// should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused
	// and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was
	// uploaded using multipart/form-data under <file_attach_name>.
	// [More info on Sending Files](https://core.telegram.org/bots/api#sending-files)
	Thumb string `json:"thumb,omitempty"`

	// Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text
	// or inline URLs in the media caption.
	ParseMode ParseMode `json:"parse_mode,omitempty"`

	// Optional. Video width
	Width int `json:"width,omitempty"`

	// Optional. Video height
	Height int `json:"height,omitempty"`

	// Optional. Video duration
	DurationInSeconds int `json:"duration,omitempty"`

	// Optional. Pass True, if the uploaded video is suitable for streaming
	SupportsStreaming bool `json:"supports_streaming,omitempty"`
}

// Represents an animation file (GIF or H.264/MPEG-4 AVC video without sound) to be sent.
type InputMediaAnimation struct {
	inputMedia

	// File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload
	// a new one using multipart/form-data under <file_attach_name> name.
	// [More info on Sending Files](https://core.telegram.org/bots/api#sending-files)
	Media string `json:"media"`

	// Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported
	// server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height
	// should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused
	// and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was
	// uploaded using multipart/form-data under <file_attach_name>.
	// [More info on Sending Files](https://core.telegram.org/bots/api#sending-files)
	Thumb string `json:"thumb,omitempty"`

	// Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text
	// or inline URLs in the media caption.
	ParseMode ParseMode `json:"parse_mode,omitempty"`

	// Optional. Video width
	Width int `json:"width,omitempty"`

	// Optional. Video height
	Height int `json:"height,omitempty"`

	// Optional. Video duration
	DurationInSeconds int `json:"duration,omitempty"`
}

// Represents an audio file to be treated as music to be sent.
type InputMediaAudio struct {
	inputMedia

	// File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>”
	// to upload a new one using multipart/form-data under <file_attach_name> name.
	// [More info on Sending Files](https://core.telegram.org/bots/api#sending-files)
	Media string `json:"media"`

	// Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported
	// server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height
	// should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused
	// and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was
	// uploaded using multipart/form-data under <file_attach_name>.
	// [More info on Sending Files](https://core.telegram.org/bots/api#sending-files)
	Thumb string `json:"thumb,omitempty"`

	// Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text
	// or inline URLs in the media caption.
	ParseMode ParseMode `json:"parse_mode,omitempty"`

	// Optional. Video duration
	DurationInSeconds int `json:"duration,omitempty"`

	// Optional. Performer of the audio
	Performer string `json:"performer,omitempty"`

	// Optional. Title of the audio
	Title string `json:"title,omitempty"`
}

// Represents a general file to be sent.
type InputMediaDocument struct {
	inputMedia

	// File to send. Pass a file_id to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>”
	// to upload a new one using multipart/form-data under <file_attach_name> name.
	// [More info on Sending Files](https://core.telegram.org/bots/api#sending-files)
	Media string `json:"media"`

	// Optional. Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported
	// server-side. The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height
	// should not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused
	// and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was
	// uploaded using multipart/form-data under <file_attach_name>.
	// [More info on Sending Files](https://core.telegram.org/bots/api#sending-files)
	Thumb string `json:"thumb,omitempty"`

	// Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text
	// or inline URLs in the media caption.
	ParseMode ParseMode `json:"parse_mode,omitempty"`
}
