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

// Type of the entity.
type MessageEntityType string
