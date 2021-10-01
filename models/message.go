package models

const (
	MessageEntityTypeMention       MessageEntityType = "mention"       // @username
	MessageEntityTypeHashTag       MessageEntityType = "hashtag"       // #hashtag
	MessageEntityTypeCashTag       MessageEntityType = "cashtag"       // $USD
	MessageEntityTypeBotCommand    MessageEntityType = "bot_command"   // /start@jobs_bot
	MessageEntityTypeURL           MessageEntityType = "url"           // https://telegram.org
	MessageEntityTypeEmail         MessageEntityType = "email"         // do-not-reply@telegram.org
	MessageEntityTypePhoneNumber   MessageEntityType = "phone_number"  // +1-212-555-0123
	MessageEntityTypeBold          MessageEntityType = "bold"          // bold text
	MessageEntityTypeItalic        MessageEntityType = "italic"        // italic text
	MessageEntityTypeUnderline     MessageEntityType = "underline"     // underline text
	MessageEntityTypeStrikethrough MessageEntityType = "strikethrough" // strikethrough text
	MessageEntityTypeCode          MessageEntityType = "code"          // monowidth string
	MessageEntityTypePre           MessageEntityType = "pre"           // monowidth string
	MessageEntityTypeTextLink      MessageEntityType = "text_link"     // for clickable text URLs
	MessageEntityTypeTextMention   MessageEntityType = "text_mention"  // for users without usernames
)

// MessageEntityType Type of the message entity.
type MessageEntityType string

// Message This object represents a message.
type Message struct {
	// Unique message identifier inside this chat
	ID int64 `json:"message_id"`

	// Optional. Sender, empty for messages sent to channels
	From *User `json:"from,omitempty"`

	// Date the message was sent in Unix time
	Timestamp int64 `json:"date"`

	// Optional. Date the message was last edited in Unix time
	EditTimestamp int64 `json:"edit_date,omitempty"`

	// Conversation the message belongs to
	Chat *Chat `json:"chat"`

	// Optional. For forwarded messages, sender of the original message
	ForwardFrom *User `json:"forward_from,omitempty"`

	// Optional. For messages forwarded from channels, information about the original channel
	ForwardFromChat *Chat `json:"forward_from_chat,omitempty"`

	// Optional. For messages forwarded from channels, identifier of the original message in the channel
	ForwardFromMessageID int64 `json:"forward_from_message_id,omitempty"`

	// Optional. For messages forwarded from channels, signature of the post author if present
	ForwardSignature string `json:"forward_signature,omitempty"`

	// Optional. Sender's name for messages forwarded from users who disallow adding a link
	// to their account in forwarded messages
	ForwardSenderName string `json:"forward_sender_name.omitempty"`

	// Optional. For forwarded messages, date the original message was sent in Unix time
	ForwardTimestamp int64 `json:"forward_date,omitempty"`

	// Optional. For replies, the original message. Note that the Message object in this field will not contain
	// further reply_to_message fields even if it itself is a reply.
	ReplyToMessage *Message `json:"reply_to_message,omitempty"`

	// Optional. The unique identifier of a media message group this message belongs to
	MediaGroupID string `json:"media_group_id,omitempty"`

	// Optional. Signature of the post author for messages in channels
	AuthorSignature string `json:"author_signature,omitempty"`

	// Optional. For text messages, the actual UTF-8 text of the message, 0-4096 characters
	Text string `json:"text,omitempty"`

	// Optional. For text messages, special entities like usernames, URLs, bot commands, etc. that appear in the text
	Entities []*MessageEntity `json:"entities,omitempty"`

	// Optional. For messages with a caption, special entities like usernames, URLs, bot commands, etc. that
	// appear in the caption
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`

	// Optional. Message is an audio file, information about the file
	Audio *Audio `json:"audio,omitempty"`

	// Optional. Message is a general file, information about the file
	Document *Document `json:"document,omitempty"`

	// Optional. Message is an animation, information about the animation. For backward compatibility,
	// when this field is set, the document field will also be set
	Animation *Animation `json:"animation,omitempty"`

	// FIXME
	// Optional. Message is a game, information about the game.
	// https://core.telegram.org/bots/api#games
	// Game *Game `json:"game,omitempty"`

	// Optional. Message is a photo, available sizes of the photo
	Photo []*PhotoSize `json:"photo,omitempty"`

	// Optional. Message is a sticker, information about the sticker
	Sticker *Sticker `json:"sticker,omitempty"`

	// Optional. Message is a video, information about the video
	Video *Video `json:"video,omitempty"`

	// Optional. Message is a voice message, information about the file
	Voice *Voice `json:"voice,omitempty"`

	// Optional. Message is a video note, information about the video message
	VideoNote *VideoNote `json:"video_note,omitempty"`

	// Optional. Caption for the animation, audio, document, photo, video or voice, 0-1024 characters
	Caption string `json:"caption,omitempty"`

	// Optional. Message is a shared contact, information about the contact
	Contact *Contact `json:"contact,omitempty"`

	// Optional. Message is a shared location, information about the location
	Location *Location `json:"location,omitempty"`

	// Optional. Message is a venue, information about the venue
	Venue *Venue `json:"venue,omitempty"`

	// Optional. Message is a native poll, information about the poll
	Poll *Poll `json:"poll,omitempty"`

	// Optional. New members that were added to the group or supergroup and information about them
	// (the bot itself may be one of these members)
	NewChatMembers []*User `json:"new_chat_members,omitempty"`

	// Optional. A member was removed from the group, information about them (this member may be the bot itself)
	LeftChatMember *User `json:"left_chat_member,omitempty"`

	// Optional. A chat title was changed to this value
	NewChatTitle string `json:"new_chat_title,omitempty"`

	// Optional. A chat photo was change to this value
	NewChatPhoto []*PhotoSize `json:"new_chat_photo,omitempty"`

	// Optional. Service message: the chat photo was deleted
	DeleteChatPhoto bool `json:"delete_chat_photo,omitempty"`

	// Optional. Service message: the group has been created
	GroupChatCreated bool `json:"group_chat_created,omitempty"`

	// Optional. Service message: the supergroup has been created. This field can‘t be received in a message
	// coming through updates, because bot can’t be a member of a supergroup when it is created. It can only be found
	// in reply_to_message if someone replies to a very first message in a directly created supergroup.
	SupergroupChatCreated bool `json:"supergroup_chat_created,omitempty"`

	// Optional. Service message: the channel has been created. This field can‘t be received in a message coming
	// through updates, because bot can’t be a member of a channel when it is created. It can only be found
	// in reply_to_message if someone replies to a very first message in a channel.
	ChannelChatCreated bool `json:"channel_chat_created,omitempty"`

	// Optional. The group has been migrated to a supergroup with the specified identifier. This number may be greater
	// than 32 bits and some programming languages may have difficulty/silent defects in interpreting it.
	// But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe
	// for storing this identifier.
	MigrateToChatID int64 `json:"migrate_to_chat_id,omitempty"`

	// Optional. The supergroup has been migrated from a group with the specified identifier. This number may be
	// greater than 32 bits and some programming languages may have difficulty/silent defects in interpreting it.
	// But it is smaller than 52 bits, so a signed 64 bit integer or double-precision float type are safe for
	// storing this identifier.
	MigrateFromChatID int64 `json:"migrate_from_chat_id,omitempty"`

	// Optional. Specified message was pinned. Note that the Message object in this field will not contain further
	// reply_to_message fields even if it is itself a reply.
	PinnedMessage *Message `json:"pinned_message,omitempty"`

	// FIXME
	// Optional. Message is an invoice for a payment, information about the invoice.
	// https://core.telegram.org/bots/api#payments
	// Invoice *Invoice `json:"invoice,omitempty"`

	// FIXME
	// Optional. Message is a service message about a successful payment, information about the payment.
	// https://core.telegram.org/bots/api#payments
	// SuccessfulPayment *SuccessfulPayment `json:"successful_payment,omitempty"`

	// Optional. The domain name of the website on which the user has logged in.
	// https://core.telegram.org/widgets/login
	ConnectedWebsite string `json:"connected_website,omitempty"`

	// FIXME
	// Optional. Telegram Passport data
	// PassportData *PassportData `json:"passport_data,omitempty"`

	// Optional. Inline keyboard attached to the message. login_url buttons are represented as ordinary url buttons.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// MessageEntity This object represents one special entity in a text message. For example, hashtags, usernames, URLs, etc.
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

// ForwardMessageRequest Use this entity to forward messages of any kind.
type ForwardMessageRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatID string `json:"chat_id"`

	// Unique identifier for the chat where the original message was sent (or channel username
	// in the format @channelusername)
	FromChatID string `json:"from_chat_id"`

	// Message identifier in the chat specified in from_chat_id
	MessageID int64 `json:"message_id"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
}

// MessageRequestBase Use this entity to send text messages.
type MessageRequestBase struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatID string `json:"chat_id"`

	// Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs
	// in the media caption.
	ParseMode ParseMode `json:"parse_mode,omitempty"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`

	// If the message is a reply, ID of the original message
	ReplyToMessageID int64 `json:"reply_to_message_id,omitempty"`

	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard,
	// instructions to remove reply keyboard or to force a reply from the user.
	// Can be one of types: InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

// MessageRequest Use this entity to send text messages.
type MessageRequest struct {
	MessageRequestBase

	// Text of the message to be sent, 1-4096 characters after entities parsing
	Text string `json:"text"`

	// Disables link previews for links in this message
	DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`
}

// PhotoMessageRequest Use this entity to send photos
type PhotoMessageRequest struct {
	MessageRequestBase

	// Photo to send. Pass a file_id as String to send a photo that exists on the Telegram servers (recommended),
	// pass an HTTP URL as a String for Telegram to get a photo from the Internet, or upload a new photo using
	// multipart/form-data.
	Photo InputFile `json:"photo"`

	// Photo caption (may also be used when resending photos by file_id), 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`
}

// AudioMessageRequest Use this entity to send audio files, if you want Telegram clients to display them in the music player.
// Your audio must be in the .MP3 or .M4A format.
type AudioMessageRequest struct {
	MessageRequestBase

	// Audio file to send. Pass a file_id as String to send an audio file that exists on the Telegram servers
	// (recommended), pass an HTTP URL as a String for Telegram to get an audio file from the Internet, or upload
	// a new one using multipart/form-data.
	Audio InputFile `json:"audio"`

	// Photo caption (may also be used when resending photos by file_id), 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Duration of the audio in seconds
	Duration int `json:"duration,omitempty"`

	// Performer
	Performer string `json:"performer,omitempty"`

	// Track name
	Title string `json:"title,omitempty"`

	// Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side.
	// The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should
	// not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused
	// and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail
	// was uploaded using multipart/form-data under <file_attach_name>.
	Thumb InputFile `json:"thumb,omitempty"`
}

// DocumentMessageRequest Use this entity to send general files.
type DocumentMessageRequest struct {
	MessageRequestBase

	// File to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended),
	// pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using
	// multipart/form-data.
	Document InputFile `json:"document"`

	// Document caption (may also be used when resending documents by file_id), 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side.
	// The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should
	// not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused
	// and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail
	// was uploaded using multipart/form-data under <file_attach_name>
	Thumb InputFile `json:"thumb,omitempty"`
}

// VideoMessageRequest Use this entity to send video files.
type VideoMessageRequest struct {
	MessageRequestBase

	// Video to send. Pass a file_id as String to send a video that exists on the Telegram servers (recommended),
	// pass an HTTP URL as a String for Telegram to get a video from the Internet, or upload a new video using
	// multipart/form-data.
	Video InputFile `json:"video"`

	// Video caption (may also be used when resending documents by file_id), 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Duration of sent video in seconds
	Duration int `json:"duration,omitempty"`

	// Video width
	Width int `json:"width,omitempty"`

	// Video height
	Height int `json:"height,omitempty"`

	// Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side.
	// The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should
	// not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused
	// and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail
	// was uploaded using multipart/form-data under <file_attach_name>
	Thumb InputFile `json:"thumb,omitempty"`

	// Pass True, if the uploaded video is suitable for streaming
	SupportsStreaming bool `json:"supports_streaming,omitempty"`
}

// AnimationMessageRequest Use this entity to send animation files (GIF or H.264/MPEG-4 AVC video without sound).
type AnimationMessageRequest struct {
	MessageRequestBase

	// Animation to send. Pass a file_id as String to send an animation that exists on the Telegram
	// servers (recommended), pass an HTTP URL as a String for Telegram to get an animation from the Internet,
	// or upload a new animation using multipart/form-data.
	Animation InputFile `json:"animation"`

	// Animation caption (may also be used when resending documents by file_id), 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Duration of sent animation in seconds
	Duration int `json:"duration,omitempty"`

	// Animation width
	Width int `json:"width,omitempty"`

	// Animation height
	Height int `json:"height,omitempty"`

	// Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side.
	// The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should
	// not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused
	// and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail
	// was uploaded using multipart/form-data under <file_attach_name>
	Thumb InputFile `json:"thumb,omitempty"`
}

// VideoNoteMessageRequest As of v.4.0, Telegram clients support rounded square mp4 videos of up to 1 minute long.
// Use this entity to send video messages.
type VideoNoteMessageRequest struct {
	MessageRequestBase

	// Video note to send. Pass a file_id as String to send a video note that exists on the Telegram
	// servers (recommended) or upload a new video using multipart/form-data.
	// Sending video notes by a URL is currently unsupported
	VideoNote InputFile `json:"video_note"`

	// Video caption (may also be used when resending documents by file_id), 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Duration of sent video in seconds
	Duration int `json:"duration,omitempty"`

	// Video width and height, i.e. diameter of the video message
	Length int `json:"length,omitempty"`

	// Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side.
	// The thumbnail should be in JPEG format and less than 200 kB in size. A thumbnail‘s width and height should
	// not exceed 320. Ignored if the file is not uploaded using multipart/form-data. Thumbnails can’t be reused
	// and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail
	// was uploaded using multipart/form-data under <file_attach_name>
	Thumb InputFile `json:"thumb,omitempty"`
}

// VoiceMessageRequest Use this method to send audio files, if you want Telegram clients to display the file as a playable voice message.
// For this to work, your audio must be in an .OGG file encoded with OPUS (other formats may be sent as Audio
// or Document)
type VoiceMessageRequest struct {
	MessageRequestBase

	// Audio file to send. Pass a file_id as String to send an audio file that exists on the Telegram servers
	// (recommended), pass an HTTP URL as a String for Telegram to get an audio file from the Internet, or upload
	// a new one using multipart/form-data.
	Voice InputFile `json:"voice"`

	// Photo caption (may also be used when resending photos by file_id), 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Duration of the audio in seconds
	Duration int `json:"duration,omitempty"`
}

// MediaGroupMessageRequest Use this entity to send a group of photos or videos as an album.
type MediaGroupMessageRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatID string `json:"chat_id"`

	// A JSON-serialized array describing photos and videos to be sent, must include 2–10 items
	// Array of InputMediaPhoto and InputMediaVideo
	Media []InputMediaInterface `json:"media"`

	// If the message is a reply, ID of the original message
	ReplyToMessageID int64 `json:"reply_to_message_id,omitempty"`

	// Sends the message silently. Users will receive a notification with no sound.
	DisableNotification bool `json:"disable_notification,omitempty"`
}

// LocationMessageRequest Use this entity to send point on the map.
type LocationMessageRequest struct {
	MessageRequestBase

	// Latitude of the location
	Latitude float64 `json:"latitude"`

	// Longitude of the location
	Longitude float64 `json:"longitude"`

	// Period in seconds for which the location will be updated.
	// See [Live Locations](https://telegram.org/blog/live-locations), should be between 60 and 86400.
	LivePeriod int32 `json:"live_period,omitempty"`
}

// EditMessageLiveLocation Use this entity to edit live location messages. A location can be edited until its live_period expires
// or editing is explicitly disabled by a call to stopMessageLiveLocation.
type EditMessageLiveLocation struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatID string `json:"chat_id"`

	// Latitude of the location
	Latitude float64 `json:"latitude"`

	// Longitude of the location
	Longitude float64 `json:"longitude"`

	// Required if inline_message_id is not specified. Identifier of the message to edit
	MessageID int64 `json:"message_id,omitempty"`

	// Required if chat_id and message_id are not specified. Identifier of the inline message
	InlineMessageID int64 `json:"inline_message_id,omitempty"`

	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard,
	// instructions to remove reply keyboard or to force a reply from the user.
	// Can be one of types: InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

// StopMessageLiveLocation Use this entity to edit live location messages. A location can be edited until its live_period expires
// or editing is explicitly disabled by a call to stopMessageLiveLocation.
type StopMessageLiveLocation struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatID string `json:"chat_id"`

	// Required if inline_message_id is not specified. Identifier of the message to edit
	MessageID int64 `json:"message_id,omitempty"`

	// Required if chat_id and message_id are not specified. Identifier of the inline message
	InlineMessageID int64 `json:"inline_message_id,omitempty"`

	// Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard,
	// instructions to remove reply keyboard or to force a reply from the user.
	// Can be one of types: InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

// VenueMessageRequest Use this entity to send information about a venue.
type VenueMessageRequest struct {
	MessageRequestBase

	// Latitude of the location
	Latitude float64 `json:"latitude"`

	// Longitude of the location
	Longitude float64 `json:"longitude"`

	// Name of the venue
	Title string `json:"title"`

	// Address of the venue
	Address string `json:"address"`

	// Foursquare identifier of the venue
	FoursquareID string `json:"foursquare_id,omitempty"`

	// Foursquare type of the venue, if known. (For example, “arts_entertainment/default”,
	// “arts_entertainment/aquarium” or “food/icecream”.)
	FoursquareType string `json:"foursquare_type,omitempty"`
}

// ContactMessageRequest Use this entity to send phone contacts.
type ContactMessageRequest struct {
	MessageRequestBase

	// Contact's phone number
	PhoneNumber string `json:"phone_number"`

	// Contact's first name
	FirstName string `json:"first_name"`

	// Optional. Contact's last name
	LastName string `json:"last_name,omitempty"`

	// Additional data about the contact in the form of a vCard, 0-2048 bytes
	VCard string `json:"vcard,omitempty"`
}

// PollMessageRequest Use this entity to send a native poll.
type PollMessageRequest struct {
	MessageRequestBase

	// Poll question, 1-255 characters
	Question string `json:"question"`

	// A JSON-serialized list of answer options, 2-10 strings 1-100 characters each
	Options []string `json:"options"`

	// True, if the poll needs to be anonymous, defaults to True
	IsAnonymous bool `json:"is_anonymous"`

	// Poll type, “quiz” or “regular”, defaults to “regular”
	Type PollType `json:"type,omitempty"`

	// True, if the poll allows multiple answers, ignored for polls in quiz mode, defaults to False
	AllowsMultipleAnswers bool `json:"allows_multiple_answers,omitempty"`

	// 0-based identifier of the correct answer option, required for polls in quiz mode
	CorrectOptionID int `json:"correct_option_id"`

	// Pass True, if the poll needs to be immediately closed. This can be useful for poll preview.
	IsClosed bool `json:"is_closed,omitempty"`
}

// StopPollRequest Use this entity to stop a poll which was sent by the bot.
type StopPollRequest struct {
	// Unique identifier for the target chat or username of the target channel (in the format @channelusername)
	ChatID string `json:"chat_id"`

	// Identifier of the original message with the poll
	MessageID int64 `json:"message_id"`

	// Optional. A JSON-serialized object for a new message inline keyboard.
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

// DiceMessageRequest Use this entity to send a dice, which will have a random value from 1 to 6.
type DiceMessageRequest struct {
	MessageRequestBase
}

type EditMessageRequest struct {
	// Required if inline_message_id is not specified. Unique identifier for the target chat or username of the target
	// channel (in the format @channelusername)
	ChatID string `json:"chat_id,omitempty"`

	// Required if inline_message_id is not specified. Identifier of the message to edit
	MessageID int64 `json:"message_id,omitempty"`

	// Required if chat_id and message_id are not specified. Identifier of the inline message
	InlineMessageID int64 `json:"inline_message_id,omitempty"`

	// Send Markdown or HTML, if you want Telegram apps to show bold, italic, fixed-width text or inline URLs
	// in your bot's message.
	ParseMode ParseMode `json:"parse_mode,omitempty"`

	// A JSON-serialized object for an inline keyboard.
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

// EditMessageTextRequest Use this entity to edit text and game messages.
type EditMessageTextRequest struct {
	EditMessageRequest

	// New text of the message, 1-4096 characters after entities parsing
	Text string `json:"text"`

	// Optional. Disables link previews for links in this message
	DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`
}

// EditMessageCaptionRequest Use this entity to edit text and game messages.
type EditMessageCaptionRequest struct {
	EditMessageRequest

	// Optional. New caption of the message, 0-1024 characters after entities parsing
	Caption string `json:"caption"`
}

// EditMessageMediaRequest Use this entity to edit animation, audio, document, photo, or video messages.
type EditMessageMediaRequest struct {
	EditMessageRequest

	// Optional. New caption of the message, 0-1024 characters after entities parsing
	Media InputMediaInterface `json:"media"`
}

// EditMessageReplyMarkupRequest Use this entity to edit only the reply markup of messages
type EditMessageReplyMarkupRequest struct {
	EditMessageRequest
}
