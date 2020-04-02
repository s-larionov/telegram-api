package models

const (
	ReplyMarkupTypeForceReply ReplyMarkupType = iota
	ReplyMarkupTypeRemoveKeyboard
	ReplyMarkupTypeKeyboardMarkup
	ReplyMarkupTypeInlineKeyboardMarkup
)

type ReplyMarkupType int

type ReplyMarkupInterface interface {
	GetType() ReplyMarkupType
}

type ReplyMarkup struct {
	t ReplyMarkupType
}

func (m ReplyMarkup) GetType() ReplyMarkupType {
	return m.t
}

// Upon receiving a message with this object, Telegram clients will display a reply interface to the user (act as if
// the user has selected the bot‘s message and tapped ’Reply'). This can be extremely useful if you want to create
// user-friendly step-by-step interfaces without having to sacrifice privacy mode.
//
// Example: A [poll bot](https://t.me/PollBot) for groups runs in privacy mode (only receives commands, replies to its
// messages and mentions). There could be two ways to create a new poll:
// - Explain the user how to send a command with parameters (e.g. /newpoll question answer1 answer2). May be appealing
//   for hardcore users but lacks modern day polish.
// - Guide the user through a step-by-step process. ‘Please send me your question’, ‘Cool, now let’s add the first
//   answer option‘, ’Great. Keep adding answer options, then send /done when you‘re ready’.
// The last option is definitely more attractive. And if you use ForceReply in your bot‘s questions, it will receive the user’s answers even if it only receives replies, commands and mentions — without any extra work for the user.
type ForceReply struct {
	ReplyMarkup

	// Shows reply interface to the user, as if they manually selected the bot‘s message and tapped ’Reply'
	ForceReply bool `json:"force_reply"`

	// Optional. Use this parameter if you want to force reply from specific users only. Targets: 1) users that
	// are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id),
	// sender of the original message.
	Selective bool `json:"selective,omitempty"`
}

func NewForceReply(selective ...bool) ForceReply {
	isSelective := false
	if len(selective) > 0 {
		isSelective = selective[0]
	}

	reply := ForceReply{
		ReplyMarkup: ReplyMarkup{
			t: ReplyMarkupTypeForceReply,
		},
		ForceReply: true,
		Selective:  isSelective,
	}

	return reply
}

// Upon receiving a message with this object, Telegram clients will remove the current custom keyboard and display
// the default letter-keyboard. By default, custom keyboards are displayed until a new keyboard is sent by a bot.
// An exception is made for one-time keyboards that are hidden immediately after the user presses a button
// (see ReplyKeyboardMarkup).
type ReplyKeyboardRemove struct {
	ReplyMarkup

	// Requests clients to remove the custom keyboard (user will not be able to summon this keyboard;
	// if you want to hide the keyboard from sight but keep it accessible, use one_time_keyboard in ReplyKeyboardMarkup)
	RemoveKeyboard bool `json:"remove_keyboard"`

	// Optional. Use this parameter if you want to remove the keyboard for specific users only. Targets: 1) users
	// that are @mentioned in the text of the Message object; 2) if the bot's message is a reply
	// (has reply_to_message_id), sender of the original message.
	//
	// Example: A user votes in a poll, bot returns confirmation message in reply to the vote and removes the keyboard
	//          for that user, while still showing the keyboard with poll options to users who haven't voted yet.
	Selective bool `json:"selective,omitempty"`
}

func NewKeyboardRemoveReply(selective ...bool) ReplyKeyboardRemove {
	isSelective := false
	if len(selective) > 0 {
		isSelective = selective[0]
	}

	reply := ReplyKeyboardRemove{
		ReplyMarkup: ReplyMarkup{
			t: ReplyMarkupTypeRemoveKeyboard,
		},
		RemoveKeyboard: true,
		Selective:      isSelective,
	}

	return reply
}

// This object represents a custom keyboard with reply options (see Introduction to bots for details and examples).
type ReplyKeyboardMarkup struct {
	ReplyMarkup

	// Array of button rows, each represented by an Array of KeyboardButton objects
	Keyboard [][]KeyboardButton `json:"keyboard"`

	// Optional. Requests clients to resize the keyboard vertically for optimal fit (e.g., make the keyboard smaller
	// if there are just two rows of buttons). Defaults to false, in which case the custom keyboard is always
	// of the same height as the app's standard keyboard.
	ResizeKeyboard bool `json:"resize_keyboard,omitempty"`

	// Optional. Requests clients to hide the keyboard as soon as it's been used. The keyboard will still be available,
	// but clients will automatically display the usual letter-keyboard in the chat – the user can press a special
	// button in the input field to see the custom keyboard again. Defaults to false.
	OneTimeKeyboard bool `json:"one_time_keyboard,omitempty"`

	// Optional. Use this parameter if you want to show the keyboard to specific users only. Targets: 1) users that
	// are @mentioned in the text of the Message object; 2) if the bot's message is a reply (has reply_to_message_id),
	// sender of the original message.
	//
	// Example: A user requests to change the bot‘s language, bot replies to the request with a keyboard to select
	//          the new language. Other users in the group don’t see the keyboard.
	Selective bool `json:"selective,omitempty"`
}

func NewKeyboardMarkupReply(keyboard [][]KeyboardButton, resize, oneTime bool, selective ...bool) ReplyKeyboardMarkup {
	isSelective := false
	if len(selective) > 0 {
		isSelective = selective[0]
	}

	reply := ReplyKeyboardMarkup{
		ReplyMarkup: ReplyMarkup{
			t: ReplyMarkupTypeKeyboardMarkup,
		},
		Keyboard:        keyboard,
		ResizeKeyboard:  resize,
		OneTimeKeyboard: oneTime,
		Selective:       isSelective,
	}

	return reply
}

// This object represents an inline keyboard that appears right next to the message it belongs to.
//
// Note: This will only work in Telegram versions released after 9 April, 2016. Older clients
//       will display unsupported message.
type InlineKeyboardMarkup struct {
	ReplyMarkup

	// Array of button rows, each represented by an Array of InlineKeyboardButton objects
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

func NewInlineKeyboardMarkupReply(keyboard [][]InlineKeyboardButton) InlineKeyboardMarkup {
	reply := InlineKeyboardMarkup{
		ReplyMarkup: ReplyMarkup{
			t: ReplyMarkupTypeInlineKeyboardMarkup,
		},
		InlineKeyboard: keyboard,
	}

	return reply
}
