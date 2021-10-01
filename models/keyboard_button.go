package models

// KeyboardButton This object represents one button of the reply keyboard. For simple text buttons String can be used instead
// of this object to specify text of the button. Optional fields request_contact, request_location,
// and request_poll are mutually exclusive.
//
// Note: request_contact and request_location options will only work in Telegram versions released after 9 April, 2016.
//       Older clients will display unsupported message.
// Note: request_poll option will only work in Telegram versions released after 23 January, 2020.
//       Older clients will display unsupported message.
type KeyboardButton struct {
	// Text of the button. If none of the optional fields are used, it will be sent as a message
	// when the button is pressed
	Text string `json:"text"`

	// Optional. If True, the user's phone number will be sent as a contact when the button is pressed.
	// Available in private chats only
	RequestContact bool `json:"request_contact,omitempty"`

	// Optional. If True, the user's current location will be sent when the button is pressed.
	// Available in private chats only
	RequestLocation bool `json:"request_location,omitempty"`

	// Optional. If specified, the user will be asked to create a poll and send it to the bot when the button is pressed.
	// Available in private chats only
	RequestPoll *KeyboardButtonPollType `json:"request_poll,omitempty"`
}

func NewKeyboardButton(text string, contact, location bool, poll *KeyboardButtonPollType) KeyboardButton {
	button := KeyboardButton{
		Text:            text,
		RequestContact:  contact,
		RequestLocation: location,
		RequestPoll:     poll,
	}

	return button
}

// KeyboardButtonPollType This object represents type of a poll, which is allowed to be created and sent when the corresponding
// button is pressed.
type KeyboardButtonPollType struct {
	// Optional. If quiz is passed, the user will be allowed to create only polls in the quiz mode. If regular is passed,
	// only regular polls will be allowed. Otherwise, the user will be allowed to create a poll of any type.
	Type PollType `json:"type"`
}
