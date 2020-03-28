package models

const (
	TypePrivate    ChatType = "private"
	TypeGroup      ChatType = "group"
	TypeSuperGroup ChatType = "supergroup"
	TypeChannel    ChatType = "channel"
)

// ChatType of chat, can be either “private”, “group”, “supergroup” or “channel”
type ChatType string
