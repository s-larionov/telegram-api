package models

// This object represents an inline keyboard that appears right next to the message it belongs to.
//
// Note: This will only work in Telegram versions released after 9 April, 2016. Older clients
//       will display unsupported message.
type InlineKeyboardMarkup struct {
	// Array of button rows, each represented by an Array of InlineKeyboardButton objects
	InlineKeyboard [][]*InlineKeyboardButton `json:"inline_keyboard"`
}
