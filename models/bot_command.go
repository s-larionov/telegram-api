package models

// BotCommand This object represents a bot command.
type BotCommand struct {
	// Text of the command, 1-32 characters. Can contain only lowercase English letters, digits and underscores.
	Command string `json:"command"`

	// Description of the command, 3-256 characters.
	Description string `json:"description"`
}
