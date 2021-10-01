package models

// CallbackQuery This object represents an incoming callback query from a callback button in an inline keyboard. If the button that
// originated the query was attached to a message sent by the bot, the field message will be present. If the button
// was attached to a message sent via the bot (in inline mode), the field inline_message_id will be present. Exactly
// one of the fields data or game_short_name will be present.
//
// NOTE: After the user presses a callback button, Telegram clients will display a progress bar until you call
//       answerCallbackQuery. It is, therefore, necessary to react by calling answerCallbackQuery even if
//       no notification to the user is needed (e.g., without specifying any of the optional parameters).
type CallbackQuery struct {
	// Unique identifier for this query
	ID string `json:"id"`

	// Sender
	From *User `json:"from"`

	// Optional. Message with the callback button that originated the query. Note that message content and message
	// date will not be available if the message is too old
	Message *Message `json:"message,omitempty"`

	// Global identifier, uniquely corresponding to the chat to which the message with the callback button was sent.
	// Useful for high scores in games.
	ChatInstance string `json:"chat_instance"`

	// Optional. Data associated with the callback button. Be aware that a bad client can send arbitrary
	// data in this field.
	Data string `json:"data,omitempty"`

	// Optional. Short name of a Game to be returned, serves as the unique identifier for the game
	GameShortName string `json:"game_short_name,omitempty"`
}

// AnswerCallbackQuery Use this entity to send answers to callback queries sent from inline keyboards.
type AnswerCallbackQuery struct {
	// Unique identifier for the query to be answered
	CallbackQueryID string `json:"callback_query_id"`

	// Optional. Text of the notification. If not specified, nothing will be shown to the user, 0-200 characters
	Text string `json:"text,omitempty"`

	// Optional. If true, an alert will be shown by the client instead of a notification at the top of the chat screen. Defaults to false.
	ShowAlert bool `json:"show_alert,omitempty"`

	// Optional. URL that will be opened by the user's client. If you have created a Game and accepted the conditions
	// via @Botfather, specify the URL that opens your game – note that this will only work if the query comes
	// from a callback_game button.
	//
	// Otherwise, you may use links like t.me/your_bot?start=XXXX that open your bot with a parameter.
	URL string `json:"url,omitempty"`

	// Optional. The maximum amount of time in seconds that the result of the callback query may be cached client-side.
	// Telegram apps will support caching starting in version 3.14. Defaults to 0.
	CacheTime int `json:"cache_time,omitempty"`
}
