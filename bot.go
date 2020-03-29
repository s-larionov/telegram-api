package telegram

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"

	"telegram/events"
	"telegram/models"
	telegram "telegram/request"
)

type Bot struct {
	subscribers *events.Container
	requester   *telegram.Requester
}

func NewBot(token string) *Bot {
	return NewBotWithClient(token, http.DefaultClient)
}

func NewBotWithClient(token string, client *http.Client) *Bot {
	return &Bot{
		subscribers: events.NewContainer(),
		requester:   telegram.NewRequesterWithClient(token, client),
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
func (b *Bot) SetWebhook(request models.WebhookRequest) error {
	var err error
	if request.Certificate != "" {
		_, err = b.requester.MultipartRequest("setWebhook", request)
	} else {
		_, err = b.requester.JSONRequest("setWebhook", request)
	}

	return err
}

func (b *Bot) WebhookHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	log.WithField("body", string(body)).Debug("incoming request")

	var update models.Update
	err := json.Unmarshal(body, &update)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	b.subscribers.Emit(update)
}

// Use this method to receive incoming updates using long polling [wiki](https://en.wikipedia.org/wiki/Push_technology#Long_polling).
// An Array of Update objects is returned.
//
// Notes
// 1. This method will not work if an outgoing webhook is set up.
// 2. In order to avoid getting duplicate updates, recalculate offset after each server response.
func (b *Bot) GetUpdates(request models.UpdateRequest) ([]models.Update, error) {
	data, err := b.requester.JSONRequest("getUpdates", request)
	if err != nil {
		return nil, err
	}

	var updates []models.Update
	err = json.Unmarshal(data, &updates)
	if err != nil {
		return nil, err
	}

	return updates, nil
}

// Use this method to get current webhook status. Requires no parameters. On success, returns a WebhookInfo object.
// If the bot is using getUpdates, will return an object with the url field empty.
func (b *Bot) GetWebhookInfo() (*models.WebhookInfo, error) {
	data, err := b.requester.JSONRequest("getWebhookInfo", []byte(""))
	if err != nil {
		return nil, err
	}

	var webhook models.WebhookInfo
	err = json.Unmarshal(data, &webhook)
	if err != nil {
		return nil, err
	}

	return &webhook, nil
}

// Use this method to remove webhook integration if you decide to switch back to getUpdates. Returns True on success.
// Requires no parameters.
func (b *Bot) DeleteWebhook() error {
	_, err := b.requester.JSONRequest("deleteWebhook", []byte(""))

	return err
}

func (b *Bot) GetMe() (*models.User, error) {
	data, err := b.requester.JSONRequest("getMe", []byte(""))
	if err != nil {
		return nil, err
	}

	var user models.User
	err = json.Unmarshal(data, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (b *Bot) SendMessage(request models.MessageRequest) (*models.Message, error) {
	data, err := b.requester.JSONRequest("sendMessage", request)
	if err != nil {
		return nil, err
	}

	var msg models.Message
	err = json.Unmarshal(data, &msg)
	if err != nil {
		return nil, err
	}

	return &msg, nil
}

func (b *Bot) Subscribe(t models.UpdateType) <-chan models.Update {
	return b.subscribers.Subscribe(t)
}

func (b *Bot) Unsubscribe(t models.UpdateType) {
	b.subscribers.Unsubscribe(t)
}
