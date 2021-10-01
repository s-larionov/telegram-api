package models

const (
	// UpdateTypeMessage New incoming message of any kind — text, photo, sticker, etc.
	UpdateTypeMessage UpdateType = "message"

	// UpdateTypeEditedMessage New version of a message that is known to the bot and was edited
	UpdateTypeEditedMessage UpdateType = "edited_message"

	// UpdateTypeChannelPost New incoming channel post of any kind — text, photo, sticker, etc.
	UpdateTypeChannelPost UpdateType = "channel_post"

	// UpdateTypeEditedChannelPost New version of a channel post that is known to the bot and was edited
	UpdateTypeEditedChannelPost UpdateType = "edited_channel_post"

	// UpdateTypeInlineQuery New incoming inline query
	UpdateTypeInlineQuery UpdateType = "inline_query"

	// UpdateTypeChosenInlineResult The result of an inline query that was chosen by a user and sent to their chat partner. Please see our
	// documentation on the [feedback collecting](https://core.telegram.org/bots/inline#collecting-feedback) for details
	// on how to enable these updates for your bot.
	UpdateTypeChosenInlineResult UpdateType = "chosen_inline_result"

	// UpdateTypeCallbackQuery New incoming callback query
	UpdateTypeCallbackQuery UpdateType = "callback_query"

	// UpdateTypeShippingQuery New incoming shipping query. Only for invoices with flexible price
	UpdateTypeShippingQuery UpdateType = "shipping_query"

	// UpdateTypePreCheckoutQuery New incoming pre-checkout query. Contains full information about checkout
	UpdateTypePreCheckoutQuery UpdateType = "pre_checkout_query"

	// UpdateTypePoll New poll state. Bots receive only updates about stopped polls and polls, which are sent by the bot
	UpdateTypePoll UpdateType = "poll"

	// UpdateTypePollAnswer A user changed their answer in a non-anonymous poll. Bots receive new votes only in polls that were sent by the bot itself.
	UpdateTypePollAnswer UpdateType = "poll_answer"
)

type UpdateType string

type UpdateRequest struct {
	// Identifier of the first update to be returned. Must be greater by one than the highest among the identifiers
	// of previously received updates. By default, updates starting with the earliest unconfirmed update are returned.
	// An update is considered confirmed as soon as getUpdates is called with an offset higher than its update_id.
	// The negative offset can be specified to retrieve updates starting from -offset update from the end of the updates
	// queue. All previous updates will forgotten.
	Offset int64 `json:"offset,omitempty"`

	// Limits the number of updates to be retrieved. Values between 1—100 are accepted. Defaults to 100.
	Limit int64 `json:"limit,omitempty"`

	// Timeout in seconds for long polling. Defaults to 0, i.e. usual short polling. Should be positive, short polling
	// should be used for testing purposes only.
	TimeoutInSeconds int64 `json:"timeout,omitempty"`

	// A JSON-serialized list of the update types you want your bot to receive. For example, specify [“message”,
	// “edited_channel_post”, “callback_query”] to only receive updates of these types. See Update for a complete list
	// of available update types. Specify an empty list to receive all updates regardless of type (default).
	// If not specified, the previous setting will be used.
	// Please note that this parameter doesn't affect updates created before the call to the getUpdates, so unwanted
	// updates may be received for a short period of time.
	AllowedUpdates []UpdateType `json:"allowed_updates,omitempty"`
}

// Update This object represents an incoming update.
// At most one of the optional parameters can be present in any given update.
type Update struct {
	// The update‘s unique identifier. Update identifiers start from a certain positive number and increase sequentially.
	// This ID becomes especially handy if you’re using Webhooks, since it allows you to ignore repeated updates or
	// to restore the correct update sequence, should they get out of order. If there are no new updates for
	// at least a week, then identifier of the next update will be chosen randomly instead of sequentially.
	ID int64 `json:"update_id"`

	// Optional. New incoming message of any kind — text, photo, sticker, etc.
	Message *Message `json:"message,omitempty"`

	// Optional. New version of a message that is known to the bot and was edited
	EditedMessage *Message `json:"edited_message,omitempty"`

	// Optional. New incoming channel post of any kind — text, photo, sticker, etc.
	ChannelPost *Message `json:"channel_post,omitempty"`

	// Optional. New version of a channel post that is known to the bot and was edited
	EditedChannelPost *Message `json:"edited_channel_post,omitempty"`

	// Optional. New incoming inline query
	InlineQuery *InlineQuery `json:"inline_query,omitempty"`

	// Optional. The result of an inline query that was chosen by a user and sent to their chat partner. Please see our
	// documentation on the [feedback collecting](https://core.telegram.org/bots/inline#collecting-feedback) for details
	// on how to enable these updates for your bot.
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`

	// Optional. New incoming callback query
	CallbackQuery *CallbackQuery `json:"callback_query,omitempty"`

	// Optional. New incoming shipping query. Only for invoices with flexible price
	ShippingQuery *ShippingQuery `json:"shipping_query,omitempty"`

	// Optional. New incoming pre-checkout query. Contains full information about checkout
	PreCheckoutQuery *PreCheckoutQuery `json:"pre_checkout_query,omitempty"`

	// Optional. New poll state. Bots receive only updates about stopped polls and polls, which are sent by the bot
	Poll *Poll `json:"poll,omitempty"`

	// Optional. A user changed their answer in a non-anonymous poll. Bots receive new votes only in polls that were
	// sent by the bot itself.
	PollAnswer *PollAnswer `json:"poll_answer,omitempty"`
}

func (u Update) GetType() UpdateType {
	switch {
	case u.Message != nil:
		return UpdateTypeMessage
	case u.EditedMessage != nil:
		return UpdateTypeEditedMessage
	case u.ChannelPost != nil:
		return UpdateTypeChannelPost
	case u.EditedChannelPost != nil:
		return UpdateTypeEditedChannelPost
	case u.InlineQuery != nil:
		return UpdateTypeInlineQuery
	case u.ChosenInlineResult != nil:
		return UpdateTypeChosenInlineResult
	case u.CallbackQuery != nil:
		return UpdateTypeCallbackQuery
	case u.ShippingQuery != nil:
		return UpdateTypeShippingQuery
	case u.PreCheckoutQuery != nil:
		return UpdateTypePreCheckoutQuery
	case u.Poll != nil:
		return UpdateTypePoll
	case u.PollAnswer != nil:
		return UpdateTypePollAnswer
	default:
	}

	return ""
}
