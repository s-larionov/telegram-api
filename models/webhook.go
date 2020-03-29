package models

// Use this method to specify a url and receive incoming updates via an outgoing webhook. Whenever there is an update
// for the bot, we will send an HTTPS POST request to the specified url, containing a JSON-serialized Update.
// In case of an unsuccessful request, we will give up after a reasonable amount of attempts. Returns True on success.
//
// If you'd like to make sure that the Webhook request comes from Telegram, we recommend using a secret path in the URL,
// e.g. https://www.example.com/<token>. Since nobody else knows your bot‘s token, you can be pretty sure it’s us.
type WebhookRequest struct {
	// HTTPS url to send updates to. Use an empty string to remove webhook integration
	URL string `json:"url"`

	// Upload your public key certificate so that the root certificate in use can be checked
	// See our [self-signed guide](https://core.telegram.org/bots/self-signed) for details.
	Certificate InputFile `json:"certificate,omitempty"`

	// Maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery, 1-100.
	// Defaults to 40. Use lower values to limit the load on your bot‘s server, and higher values
	// to increase your bot’s throughput.
	MaxConnections int `json:"max_connections,omitempty"`

	// A JSON-serialized list of the update types you want your bot to receive. For example, specify [“message”,
	// “edited_channel_post”, “callback_query”] to only receive updates of these types. See Update for a complete
	// list of available update types. Specify an empty list to receive all updates regardless of type (default).
	// If not specified, the previous setting will be used.
	//
	// Please note that this parameter doesn't affect updates created before the call to the setWebhook, so unwanted
	// updates may be received for a short period of time.
	AllowedUpdates []UpdateType `json:"allowed_updates,omitempty"`
}

// Contains information about the current status of a webhook.
type WebhookInfo struct {
	// Webhook URL, may be empty if webhook is not set up
	URL string `json:"url,omitempty"`

	// True, if a custom certificate was provided for webhook certificate checks
	HasCustomCertificate bool `json:"has_custom_certificate,omitempty"`

	// Number of updates awaiting delivery
	PendingUpdateCount int64 `json:"pending_update_count"`

	// Optional. Unix time for the most recent error that happened when trying to deliver an update via webhook
	LastErrorTimestamp int64 `json:"last_error_date,omitempty"`

	// Optional. Error message in human-readable format for the most recent error that happened when trying
	// to deliver an update via webhook
	LastErrorMessage string `json:"last_error_message,omitempty"`

	// Optional. Maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery
	MaxConnections int `json:"max_connections,omitempty"`

	// Optional. A list of update types the bot is subscribed to. Defaults to all update types
	AllowedUpdates []UpdateType `json:"allowed_updates,omitempty"`
}

func (h WebhookInfo) IsEnabled() bool {
	return h.URL != ""
}
