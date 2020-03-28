package models

// This object represents an incoming inline query. When the user sends an empty query, your bot could return
// some default or trending results.
type InlineQuery struct {
	// Unique identifier for this query
	ID string `json:"id"`

	// Sender
	From *User `json:"from"`

	// Optional. Sender location, only for bots that request user location
	Location *Location `json:"location,omitempty"`

	// Text of the query (up to 256 characters)
	Query string `json:"query"`

	// Offset of the results to be returned, can be controlled by the bot
	Offset string `json:"offset"`
}

// Use this method to send answers to an inline query. On success, True is returned.
// No more than 50 results per query are allowed.
type AnswerInlineQuery struct {
	// Unique identifier for the answered query
	InlineQueryID string `json:"inline_query_id"`

	// A JSON-serialized array of results for the inline query
	Results []InlineQueryResultInterface `json:"results"`

	// Optional. The maximum amount of time in seconds that the result of the inline query may be cached on the server.
	// Defaults to 300.
	CacheTime int64 `json:"cache_time,omitempty"`

	// Optional. Pass True, if results may be cached on the server side only for the user that sent the query.
	// By default, results may be returned to any user who sends the same query
	IsPersonal bool `json:"is_personal,omitempty"`

	// Optional. Pass the offset that a client should send in the next query with the same text to receive more results.
	// Pass an empty string if there are no more results or if you don‘t support pagination. Offset length can’t exceed 64 bytes.
	NextOffset string `json:"next_offset,omitempty"`

	// Optional. If passed, clients will display a button with specified text that switches the user to a private chat
	// with the bot and sends the bot a start message with the parameter switch_pm_parameter
	SwitchPmText string `json:"switch_pm_text,omitempty"`

	// Optional. [Deep-linking](https://core.telegram.org/bots#deep-linking) parameter for the /start message sent
	// to the bot when user presses the switch button. 1-64 characters, only A-Z, a-z, 0-9, _ and - are allowed.
	//
	// Example: An inline bot that sends YouTube videos can ask the user to connect the bot to their YouTube a
	// ccount to adapt search results accordingly. To do this, it displays a ‘Connect your YouTube account’ button
	// above the results, or even before showing any. The user presses the button, switches to a private chat with
	// the bot and, in doing so, passes a start parameter that instructs the bot to return an oauth link. Once done,
	// the bot can offer a switch_inline button so that the user can easily return to the chat where they wanted
	// to use the bot's inline capabilities.
	SwitchPmParameter string `json:"switch_pm_parameter,omitempty"`
}

const (
	InlineQueryResultTypeArticle  InlineQueryResultType = "article"
	InlineQueryResultTypePhoto    InlineQueryResultType = "photo"
	InlineQueryResultTypeGif      InlineQueryResultType = "gif"
	InlineQueryResultTypeVideo    InlineQueryResultType = "video"
	InlineQueryResultTypeMpeg4Gif InlineQueryResultType = "mpeg4_gif"
	InlineQueryResultTypeAudio    InlineQueryResultType = "audio"
	InlineQueryResultTypeVoice    InlineQueryResultType = "voice"
	InlineQueryResultTypeDocument InlineQueryResultType = "document"
	InlineQueryResultTypeLocation InlineQueryResultType = "location"
	InlineQueryResultTypeVenue    InlineQueryResultType = "venue"
	InlineQueryResultTypeContact  InlineQueryResultType = "contact"
	InlineQueryResultTypeGame     InlineQueryResultType = "game"
	InlineQueryResultTypeSticker  InlineQueryResultType = "sticker"
)

type InlineQueryResultType string

type InlineQueryResultInterface interface {
	GetType() InlineQueryResultType
}

type InlineQueryResult struct {
	// Type of the result
	Type InlineQueryResultType

	// Unique identifier for this result, 1-64 Bytes
	ID string `json:"id"`

	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (r InlineQueryResult) GetType() InlineQueryResultType {
	return r.Type
}

// Represents a link to an article or web page.
type InlineQueryResultArticle struct {
	InlineQueryResult

	// Title of the result
	Title string `json:"title"`

	// Content of the message to be sent
	InputMessageContent InputMessageContentInterface `json:"input_message_content"`

	// Optional. URL of the result
	URL string `json:"url,omitempty"`

	// Optional. Pass True, if you don't want the URL to be shown in the message
	HideURL bool `json:"hide_url,omitempty"`

	// Optional. Short description of the result
	Description string `json:"description,omitempty"`

	// Optional. Url of the thumbnail for the result
	ThumbURL string `json:"thumb_url,omitempty"`

	// Optional. Thumbnail width
	ThumbWidth int `json:"thumb_width,omitempty"`

	// Optional. Thumbnail height
	ThumbHeight int `json:"thumb_height,omitempty"`
}

// Represents a link to a photo. By default, this photo will be sent by the user with optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
type InlineQueryResultPhoto struct {
	InlineQueryResult

	// A valid file identifier of the photo
	FileID string `json:"photo_file_id,omitempty"`

	// A valid URL of the photo. Photo must be in jpeg format. Photo size must not exceed 5MB
	URL string `json:"photo_url,omitempty"`

	// URL of the thumbnail for the photo
	ThumbURL string `json:"thumb_url,omitempty"`

	// Optional. Width of the photo
	Width int `json:"photo_width,omitempty"`

	// Optional. Height of the photo
	Height int `json:"photo_height,omitempty"`

	// Optional. Title for the result
	Title string `json:"title,omitempty"`

	// Optional. Short description of the result
	Description string `json:"description,omitempty"`

	// Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline
	// URLs in the media caption.
	ParseMode ParseMode `json:"parse_mode,omitempty"`

	// Optional. Content of the message to be sent instead of the photo
	InputMessageContent InputMessageContentInterface `json:"input_message_content,omitempty"`
}

// Represents a link to an animated GIF file. By default, this animated GIF file will be sent by the user with
// optional caption. Alternatively, you can use input_message_content to send a message with the specified content
// instead of the animation.
type InlineQueryResultGif struct {
	InlineQueryResult

	// A valid file identifier for the GIF file
	FileID string `json:"gif_file_id,omitempty"`

	// A valid URL for the GIF file. File size must not exceed 1MB
	URL string `json:"gif_url,omitempty"`

	// Optional. Width of the GIF
	Width int `json:"gif_width,omitempty"`

	// Optional. Height of the GIF
	Height int `json:"gif_height,omitempty"`

	// Optional. Duration of the GIF
	Duration int `json:"gif_duration,omitempty"`

	// URL of the thumbnail for the GIF
	ThumbURL string `json:"thumb_url,omitempty"`

	// Optional. Title for the result
	Title string `json:"title,omitempty"`

	// Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline
	// URLs in the media caption.
	ParseMode ParseMode `json:"parse_mode,omitempty"`

	// Optional. Content of the message to be sent instead of the GIF animation
	InputMessageContent InputMessageContentInterface `json:"input_message_content,omitempty"`
}

// Represents a link to a video animation (H.264/MPEG-4 AVC video without sound). By default, this animated MPEG-4 file
// will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message
// with the specified content instead of the animation.
type InlineQueryResultMpeg4Gif struct {
	InlineQueryResult

	// A valid file identifier for the MP4 file
	Mpeg4FileID string `json:"mpeg_4_file_id,omitempty"`

	// A valid URL for the MP4 file. File size must not exceed 1MB
	URL string `json:"mpeg4_url,omitempty"`

	// Optional. Width of the MP4
	Width int `json:"mpeg4_width,omitempty"`

	// Optional. Height of the MP4
	Height int `json:"mpeg4_height,omitempty"`

	// Optional. Duration of the MP4
	Duration int `json:"mpeg4_duration,omitempty"`

	// URL of the thumbnail for the MP4
	ThumbURL string `json:"thumb_url,omitempty"`

	// Optional. Title for the result
	Title string `json:"title,omitempty"`

	// Optional. Caption of the photo to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline
	// URLs in the media caption.
	ParseMode ParseMode `json:"parse_mode,omitempty"`

	// Optional. Content of the message to be sent instead of the video animation
	InputMessageContent InputMessageContentInterface `json:"input_message_content,omitempty"`
}

// Represents a link to a page containing an embedded video player or a video file. By default, this video file will
// be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message
// with the specified content instead of the video.
// If an InlineQueryResultVideo message contains an embedded video (e.g., YouTube), you must replace its content using
// input_message_content.
type InlineQueryResultVideo struct {
	InlineQueryResult

	// A valid file identifier for the video file
	FileID string `json:"video_file_id,omitempty"`

	// A valid URL for the embedded video player or video file
	URL string `json:"video_url,omitempty"`

	// Mime type of the content of video url, “text/html” or “video/mp4”
	MimeType string `json:"mime_type,omitempty"`

	// URL of the thumbnail (jpeg only) for the video
	ThumbURL string `json:"thumb_url,omitempty"`

	// Optional. Title for the result
	Title string `json:"title,omitempty"`

	// Optional. Caption of the video to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline
	// URLs in the media caption.
	ParseMode ParseMode `json:"parse_mode,omitempty"`

	// Optional. Width of the MP4
	Width int `json:"video_width,omitempty"`

	// Optional. Height of the MP4
	Height int `json:"video_height,omitempty"`

	// Optional. Video duration in seconds
	Duration int `json:"video_duration,omitempty"`

	// Optional. Short description of the result
	Description string `json:"description,omitempty"`

	// Optional. Content of the message to be sent instead of the video. This field is required
	// if InlineQueryResultVideo is used to send an HTML-page as a result (e.g., a YouTube vi	deo).
	InputMessageContent InputMessageContentInterface `json:"input_message_content,omitempty"`
}

// Represents a link to an MP3 audio file. By default, this audio file will be sent by the user. Alternatively,
// you can use input_message_content to send a message with the specified content instead of the audio.
// Note: This will only work in Telegram versions released after 9 April, 2016. Older clients will ignore them.
type InlineQueryResultAudio struct {
	InlineQueryResult

	// A valid file identifier for the audio file
	FileID string `json:"audio_file_id,omitempty"`

	// A valid URL for the audio file
	URL string `json:"audio_url,omitempty"`

	// Optional. Title for the result
	Title string `json:"title,omitempty"`

	// Optional. Caption, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline
	// URLs in the media caption.
	ParseMode ParseMode `json:"parse_mode,omitempty"`

	// Optional. Performer
	Performer string `json:"performer,omitempty"`

	// Optional. Audio duration in seconds
	Duration int `json:"audio_duration,omitempty"`

	// Optional. Content of the message to be sent instead of the file
	InputMessageContent InputMessageContentInterface `json:"input_message_content,omitempty"`
}

// Represents a link to a voice recording in an .OGG container encoded with OPUS. By default, this voice recording will
// be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content
// instead of the the voice message.
// Note: This will only work in Telegram versions released after 9 April, 2016. Older clients will ignore them.
type InlineQueryResultVoice struct {
	InlineQueryResult

	// A valid file identifier for the voice message
	FileID string `json:"voice_file_id,omitempty"`

	// A valid URL for the voice recording
	URL string `json:"voice_url,omitempty"`

	// Recording title
	Title string `json:"title,omitempty"`

	// Optional. Caption, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline
	// URLs in the media caption.
	ParseMode ParseMode `json:"parse_mode,omitempty"`

	// Optional. Recording duration in seconds
	Duration int `json:"voice_duration,omitempty"`

	// Optional. Content of the message to be sent instead of the file
	InputMessageContent InputMessageContentInterface `json:"input_message_content,omitempty"`
}

// Represents a link to a file. By default, this file will be sent by the user with an optional caption. Alternatively,
// you can use input_message_content to send a message with the specified content instead of the file. Currently,
// only .PDF and .ZIP files can be sent using this method.
// Note: This will only work in Telegram versions released after 9 April, 2016. Older clients will ignore them.
type InlineQueryResultDocument struct {
	InlineQueryResult

	// A valid file identifier for the file
	FileID string `json:"document_file_id,omitempty"`

	// Title for the result
	Title string `json:"title,omitempty"`

	// Optional. Caption of the document to be sent, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline
	// URLs in the media caption.
	ParseMode ParseMode `json:"parse_mode,omitempty"`

	// A valid URL for the file
	URL string `json:"document_url,omitempty"`

	// Mime type of the content of the file, either “application/pdf” or “application/zip”
	MimeType string `json:"mime_type,omitempty"`

	// Optional. Short description of the result
	Description string `json:"description,omitempty"`

	// Optional. Content of the message to be sent instead of the file
	InputMessageContent InputMessageContentInterface `json:"input_message_content,omitempty"`

	// Optional. URL of the thumbnail (jpeg only) for the file
	ThumbURL string `json:"thumb_url,omitempty"`

	// Optional. Thumbnail width
	ThumbWidth int `json:"thumb_width,omitempty"`

	// Optional. Thumbnail height
	ThumbHeight int `json:"thumb_height,omitempty"`
}

// Represents a location on a map. By default, the location will be sent by the user. Alternatively, you can use
// input_message_content to send a message with the specified content instead of the location.
type InlineQueryResultLocation struct {
	InlineQueryResult

	// Location latitude in degrees
	Latitude float64 `json:"latitude"`

	// Location longitude in degrees
	Longitude float64 `json:"longitude"`

	// Location title
	Title string `json:"title"`

	// Optional. Period in seconds for which the location can be updated, should be between 60 and 86400.
	LivePeriod int64 `json:"live_period,omitempty"`

	// Optional. Content of the message to be sent instead of the location
	InputMessageContent InputMessageContentInterface `json:"input_message_content,omitempty"`

	// Optional. URL of the thumbnail (jpeg only) for the file
	ThumbURL string `json:"thumb_url,omitempty"`

	// Optional. Thumbnail width
	ThumbWidth int `json:"thumb_width,omitempty"`

	// Optional. Thumbnail height
	ThumbHeight int `json:"thumb_height,omitempty"`
}

// Represents a venue. By default, the venue will be sent by the user. Alternatively, you can use input_message_content
// to send a message with the specified content instead of the venue.
// Note: This will only work in Telegram versions released after 9 April, 2016. Older clients will ignore them.
type InlineQueryResultVenue struct {
	InlineQueryResult

	// Latitude of the venue location in degrees
	Latitude float64 `json:"latitude"`

	// Longitude of the venue location in degrees
	Longitude float64 `json:"longitude"`

	// Title of the venue
	Title string `json:"title"`

	// Address of the venue
	Address string `json:"address"`

	// Optional. Foursquare identifier of the venue if known
	FoursquareID string `json:"foursquare_id,omitempty"`

	// Optional. Foursquare type of the venue, if known. (For example, “arts_entertainment/default”,
	// “arts_entertainment/aquarium” or “food/icecream”.)
	FoursquareType string `json:"foursquare_type,omitempty"`

	// Optional. Content of the message to be sent instead of the venue
	InputMessageContent InputMessageContentInterface `json:"input_message_content,omitempty"`

	// Optional. URL of the thumbnail (jpeg only) for the file
	ThumbURL string `json:"thumb_url,omitempty"`

	// Optional. Thumbnail width
	ThumbWidth int `json:"thumb_width,omitempty"`

	// Optional. Thumbnail height
	ThumbHeight int `json:"thumb_height,omitempty"`
}

// Represents a contact with a phone number. By default, this contact will be sent by the user. Alternatively,
// you can use input_message_content to send a message with the specified content instead of the contact.
// Note: This will only work in Telegram versions released after 9 April, 2016. Older clients will ignore them.
type InlineQueryResultContact struct {
	InlineQueryResult

	// Contact's phone number
	PhoneNumber string `json:"phone_number"`

	// Contact's first name
	FirstName string `json:"first_name"`

	// Optional. Contact's last name
	LastName string `json:"last_name,omitempty"`

	// Optional. Additional data about the contact in the form of a vCard, 0-2048 bytes
	VCard string `json:"vcard,omitempty"`

	// Optional. Content of the message to be sent instead of the contact
	InputMessageContent InputMessageContentInterface `json:"input_message_content,omitempty"`

	// Optional. URL of the thumbnail (jpeg only) for the file
	ThumbURL string `json:"thumb_url,omitempty"`

	// Optional. Thumbnail width
	ThumbWidth int `json:"thumb_width,omitempty"`

	// Optional. Thumbnail height
	ThumbHeight int `json:"thumb_height,omitempty"`
}

// Represents a Game.
// Note: This will only work in Telegram versions released after October 1, 2016. Older clients will not display any
// inline results if a game result is among them.
type InlineQueryResultGame struct {
	InlineQueryResult

	// Short name of the game
	GameShortName string `json:"game_short_name"`
}

// Represents a link to a sticker stored on the Telegram servers. By default, this sticker will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the sticker.
type InlineQueryResultSticker struct {
	InlineQueryResult

	// A valid file identifier of the sticker
	StickerFileID string `json:"sticker_file_id"`

	// Optional. Content of the message to be sent instead of the sticker
	InputMessageContent InputMessageContentInterface `json:"input_message_content,omitempty"`
}

type InputMessageContentInterface interface {
	GetType() InputMessageContentType
}

type InputMessageContentType string

type InputMessageContent struct {
	t InputMessageContentType
}

func (c InputMessageContent) GetType() InputMessageContentType {
	return c.t
}

// Represents the content of a text message to be sent as the result of an inline query.
type InputTextMessageContent struct {
	InputMessageContent

	// Text of the message to be sent, 1-4096 characters
	MessageText string `json:"message_text"`

	// Optional. Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline
	// URLs in your bot's message.
	ParseMode ParseMode `json:"parse_mode,omitempty"`

	// Optional. Disables link previews for links in the sent message
	DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`
}

// Represents the content of a location message to be sent as the result of an inline query.
type InputLocationMessageContent struct {
	InputMessageContent

	// Latitude of the location in degrees
	Latitude float64 `json:"latitude"`

	// Longitude of the location in degrees
	Longitude float64 `json:"longitude"`

	// Optional. Period in seconds for which the location can be updated, should be between 60 and 86400.
	LivePeriod int32 `json:"live_period,omitempty"`
}

// Represents the content of a venue message to be sent as the result of an inline query.
type InputVenueMessageContent struct {
	InputMessageContent

	// Latitude of the location in degrees
	Latitude float64 `json:"latitude"`

	// Longitude of the location in degrees
	Longitude float64 `json:"longitude"`

	// Name of the venue
	Title string `json:"title"`

	// Address of the venue
	Address string `json:"address"`

	// Optional. Foursquare identifier of the venue, if known
	FoursquareID string `json:"foursquare_id"`

	// Optional. Foursquare type of the venue, if known. (For example, “arts_entertainment/default”,
	// “arts_entertainment/aquarium” or “food/icecream”.)
	FoursquareType string `json:"foursquare_type"`
}

// Represents the content of a contact message to be sent as the result of an inline query.
type InputContactMessageContent struct {
	InputMessageContent

	// Contact's phone number
	PhoneNumber string `json:"phone_number"`

	// Contact's first name
	FirstName string `json:"first_name"`

	// Optional. Contact's last name
	LastName string `json:"last_name,omitempty"`

	// Optional. Additional data about the contact in the form of a vCard, 0-2048 bytes
	VCard string `json:"v_card,omitempty"`
}

// Represents a result of an inline query that was chosen by the user and sent to their chat partner.
// Note: It is necessary to enable inline feedback via @Botfather in order to receive these objects in updates.
type ChosenInlineResult struct {
	// The unique identifier for the result that was chosen
	ID string `json:"result_id"`

	// The user that chose the result
	From *User `json:"from"`

	// Optional. Sender location, only for bots that require user location
	Location *Location `json:"location,omitempty"`

	// Optional. Identifier of the sent inline message. Available only if there is an inline keyboard attached
	// to the message. Will be also received in callback queries and can be used to edit the message.
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// The query that was used to obtain the result
	Query string `json:"query"`
}
