package pkg

import (
	"telegram/pkg/models"
)

const urlTemplate = "https://api.telegram.org/bot%s/%s"

type Bot struct {
	token string
}

func NewBot(token string) *Bot {
	return &Bot{
		token: token,
	}
}

// Use this method to specify a url and receive incoming updates via an outgoing webhook. Whenever there is an update
// for the bot, we will send an HTTPS POST request to the specified url, containing a JSON-serialized Update.
// In case of an unsuccessful request, we will give up after a reasonable amount of attempts. Returns True on success.
//
// If you'd like to make sure that the Webhook request comes from Telegram, we recommend using a secret path in the URL,
// e.g. https://www.example.com/<token>. Since nobody else knows your bot‘s token, you can be pretty sure it’s us.
//
// Notes
// 1. You will not be able to receive updates using getUpdates for as long as an outgoing webhook is set up.
// 2. To use a self-signed certificate, you need to upload your public key certificate using certificate parameter. Please upload as InputFile, sending a String will not work.
// 3. Ports currently supported for Webhooks: 443, 80, 88, 8443.
//
// NEW! If you're having any trouble setting up webhooks, please check out this
// [amazing guide to Webhooks](https://core.telegram.org/bots/webhooks).
func (b *Bot) SetWebhook(request models.WebhookRequest) {
}

// Use this method to remove webhook integration if you decide to switch back to getUpdates. Returns True on success. Requires no parameters.
func (b *Bot) DeleteWebhook() {
}

// Use this method to receive incoming updates using long polling [wiki](https://en.wikipedia.org/wiki/Push_technology#Long_polling).
// An Array of Update objects is returned.
//
// Notes
// 1. This method will not work if an outgoing webhook is set up.
// 2. In order to avoid getting duplicate updates, recalculate offset after each server response.
func (b *Bot) GetUpdates(request models.UpdateRequest) []models.Update {
	return nil
}

func (b *Bot) GetMe() *models.User {
	return nil
}

func (b *Bot) SendMessage() *models.User {
	return nil
}

func (b *Bot) makeRequest() {
}

func (b *Bot) makeJSONRequest() {
}

func (b *Bot) makeMultipartRequest() {
}
