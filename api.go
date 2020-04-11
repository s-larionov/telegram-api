package telegram

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/s-larionov/telegram-api/events"
	"github.com/s-larionov/telegram-api/models"
	"github.com/s-larionov/telegram-api/request"
)

type API struct {
	subscribers *events.Container
	requester   *request.Requester
}

func NewAPI(token string) *API {
	return NewAPIWithClient(token, http.DefaultClient)
}

func NewAPIWithClient(token string, client *http.Client) *API {
	return &API{
		subscribers: events.NewContainer(),
		requester:   request.NewRequesterWithClient(token, client),
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
func (b *API) SetWebhook(request models.WebhookRequest) error {
	var err error
	if request.Certificate != "" {
		_, err = b.requester.MultipartRequest("setWebhook", request)
	} else {
		_, err = b.requester.JSONRequest("setWebhook", request)
	}

	return err
}

func (b *API) WebhookHandler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	log.WithField("body", string(body)).Trace("incoming request")

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
func (b *API) GetUpdates(request models.UpdateRequest) ([]models.Update, error) {
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
func (b *API) GetWebhookInfo() (*models.WebhookInfo, error) {
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
func (b *API) DeleteWebhook() error {
	_, err := b.requester.JSONRequest("deleteWebhook", []byte(""))

	return err
}

// A simple method for testing your bot's auth token. Requires no parameters. Returns basic information about
// the bot in form of a User object.
func (b *API) GetMe() (*models.User, error) {
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

// Use this method to forward messages of any kind.
func (b *API) ForwardMessage(request models.ForwardMessageRequest) (*models.Message, error) {
	return b.sendMessage("forwardMessage", request)
}

// Use this method to send text messages. On success, the sent Message is returned.
func (b *API) SendMessage(request models.MessageRequest) (*models.Message, error) {
	return b.sendMessage("sendMessage", request)
}

// Use this method to send photos.
func (b *API) SendPhoto(request models.PhotoMessageRequest) (*models.Message, error) {
	return b.sendMessage("sendPhoto", request)
}

// Use this method to send audio files, if you want Telegram clients to display them in the music player.
// Your audio must be in the .MP3 or .M4A format. On success, the sent Message is returned. Bots can currently
// send audio files of up to 50 MB in size, this limit may be changed in the future.
func (b *API) SendAudio(request models.AudioMessageRequest) (*models.Message, error) {
	return b.sendMessage("sendAudio", request)
}

// Use this method to send general files. On success, the sent Message is returned. Bots can currently send files
// of any type of up to 50 MB in size, this limit may be changed in the future.
func (b *API) SendDocument(request models.DocumentMessageRequest) (*models.Message, error) {
	return b.sendMessage("sendDocument", request)
}

// Use this method to send video files, Telegram clients support mp4 videos (other formats may be sent as Document).
// On success, the sent Message is returned. Bots can currently send video files of up to 50 MB in size,
// this limit may be changed in the future.
func (b *API) SendVideo(request models.VideoMessageRequest) (*models.Message, error) {
	return b.sendMessage("sendVideo", request)
}

// Use this method to send animation files (GIF or H.264/MPEG-4 AVC video without sound). On success, the sent
// Message is returned. Bots can currently send animation files of up to 50 MB in size, this limit may be changed
// in the future.
func (b *API) SendAnimation(request models.AnimationMessageRequest) (*models.Message, error) {
	return b.sendMessage("sendAnimation", request)
}

// Use this method to send audio files, if you want Telegram clients to display the file as a playable voice message.
// For this to work, your audio must be in an .OGG file encoded with OPUS (other formats may be sent as Audio
// or Document). On success, the sent Message is returned. Bots can currently send voice messages of up to 50 MB
// 	in size, this limit may be changed in the future.
func (b *API) SendVoice(request models.VoiceMessageRequest) (*models.Message, error) {
	return b.sendMessage("sendVoice", request)
}

// As of v.4.0, Telegram clients support rounded square mp4 videos of up to 1 minute long.
// Use this method to send video messages. On success, the sent Message is returned.
func (b *API) SendVideoNote(request models.VideoNoteMessageRequest) (*models.Message, error) {
	return b.sendMessage("sendVideoNote", request)
}

// Use this method to send a group of photos or videos as an album. On success,
// an array of the sent Messages is returned.
func (b *API) SendMediaGroup(request models.MediaGroupMessageRequest) ([]models.Message, error) {
	data, err := b.requester.JSONRequest("sendMediaGroup", request)
	if err != nil {
		return nil, err
	}

	var messages []models.Message
	err = json.Unmarshal(data, &messages)
	if err != nil {
		return nil, err
	}

	return messages, nil
}

// Use this method to send point on the map. On success, the sent Message is returned.
func (b *API) SendLocation(request models.LocationMessageRequest) (*models.Message, error) {
	return b.sendMessage("sendLocation", request)
}

// Use this method to edit live location messages. A location can be edited until its live_period expires or
// editing is explicitly disabled by a call to stopMessageLiveLocation. On success, if the edited message was sent
// by the bot, the edited Message is returned, otherwise True is returned.
func (b *API) EditMessageLiveLocation(request models.EditMessageLiveLocation) (bool, error) {
	_, err := b.requester.JSONRequest("editMessageLiveLocation", request)

	return err == nil, err
}

// Use this method to stop updating a live location message before live_period expires. On success, if the message
// was sent by the bot, the sent Message is returned, otherwise True is returned.
func (b *API) StopMessageLiveLocation(request models.StopMessageLiveLocation) (bool, error) {
	_, err := b.requester.JSONRequest("stopMessageLiveLocation", request)

	return err == nil, err
}

// Use this method to send information about a venue. On success, the sent Message is returned.
func (b *API) SendVenue(request models.VenueMessageRequest) (*models.Message, error) {
	return b.sendMessage("sendVenue", request)
}

// Use this method to send phone contacts. On success, the sent Message is returned.
func (b *API) SendContact(request models.ContactMessageRequest) (*models.Message, error) {
	return b.sendMessage("sendContact", request)
}

// Use this method to send a native poll. On success, the sent Message is returned.
func (b *API) SendPoll(request models.PollMessageRequest) (*models.Message, error) {
	return b.sendMessage("sendPoll", request)
}

// Use this method to stop a poll which was sent by the bot. On success, the stopped Poll with the final results
// is returned.
func (b *API) StopPoll(request models.PollMessageRequest) (*models.Poll, error) {
	data, err := b.requester.JSONRequest("stopPoll", request)
	if err != nil {
		return nil, err
	}

	var result models.Poll
	err = json.Unmarshal(data, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// Use this method to send a dice, which will have a random value from 1 to 6. On success, the sent Message is returned.
// (Yes, we're aware of the “proper” singular of die. But it's awkward, and we decided to help it change.
// One dice at a time!)
func (b *API) SendDice(request models.DiceMessageRequest) (*models.Message, error) {
	return b.sendMessage("sendDice", request)
}

// Use this method to delete a message, including service messages, with the following limitations:
// - A message can only be deleted if it was sent less than 48 hours ago.
// - A dice message in a private chat can only be deleted if it was sent more than 24 hours ago.
// - Bots can delete outgoing messages in private chats, groups, and supergroups.
// - Bots can delete incoming messages in private chats.
// - Bots granted can_post_messages permissions can delete outgoing messages in channels.
// - If the bot is an administrator of a group, it can delete any message there.
// - If the bot has can_delete_messages permission in a supergroup or a channel, it can delete any message there.
// Returns True on success.
//
// chatID    - Unique identifier for the target chat or username of the target channel (in the format @channelusername)
// messageID - Identifier of the message to delete
func (b *API) DeleteMessage(chatID string, messageID int64) (bool, error) {
	_, err := b.requester.JSONRequest("deleteMessage", map[string]interface{}{
		"chat_id":    chatID,
		"message_id": messageID,
	})

	return err == nil, err
}

// Use this method when you need to tell the user that something is happening on the bot's side.
// The status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients clear its
// typing status). Returns True on success.
// Example: The ImageBot needs some time to process a request and upload the image. Instead of sending a text message
//          along the lines of “Retrieving image, please wait…”, the bot may use sendChatAction with
//          action = upload_photo. The user will see a “sending photo” status for the bot.
// We only recommend using this method when a response from the bot will take a noticeable amount of time to arrive.
func (b *API) SendChatAction(chatID string, action models.ChatAction) (bool, error) {
	_, err := b.requester.JSONRequest("sendChatAction", map[string]interface{}{
		"chat_id": chatID,
		"action":  action,
	})

	return err == nil, err
}

// Use this method to get a list of profile pictures for a user. Returns a UserProfilePhotos object.
func (b *API) GetUserProfilePhotos(request models.UserProfilePhotosRequest) (*models.UserProfilePhotos, error) {
	data, err := b.requester.JSONRequest("getUserProfilePhotos", request)
	if err != nil {
		return nil, err
	}

	var response models.UserProfilePhotos
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// Use this method to get basic info about a file and prepare it for downloading. For the moment, bots can download
// files of up to 20MB in size. On success, a File object is returned. The file can then be downloaded via the
// link https://api.telegram.org/file/bot<token>/<file_path>, where <file_path> is taken from the response.
// It is guaranteed that the link will be valid for at least 1 hour. When the link expires, a new one can be
// requested by calling getFile again.
// Note: This function may not preserve the original file name and MIME type. You should save the file's MIME type
//       and name (if available) when the File object is received.
func (b *API) GetFile(fileID string) (string, error) {
	data, err := b.requester.JSONRequest("getFile", map[string]interface{}{
		"file_id": fileID,
	})
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// Use this method to kick a user from a group, a supergroup or a channel. In the case of supergroups and channels,
// the user will not be able to return to the group on their own using invite links, etc., unless unbanned first.
// The bot must be an administrator in the chat for this to work and must have the appropriate admin rights.
// Returns True on success.
//
// chatID         - Unique identifier for the target group or username of the target supergroup
//                  or channel (in the format @channelusername)
// userID         - Unique identifier of the target user
// untilTimestamp - Date when the user will be unbanned, unix time. If user is banned for more than 366 days or less
//                  than 30 seconds from the current time they are considered to be banned forever
func (b *API) KickChatMember(chatID string, userID int64, untilTimestamp ...int64) (bool, error) {
	r := map[string]interface{}{
		"chat_id": chatID,
		"user_id": userID,
	}

	if len(untilTimestamp) > 0 {
		r["until_date"] = untilTimestamp[0]
	}

	_, err := b.requester.JSONRequest("kickChatMember", r)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Use this method to unban a previously kicked user in a supergroup or channel. The user will not return
// to the group or channel automatically, but will be able to join via link, etc. The bot must be an administrator
// for this to work. Returns True on success.
//
// chatID         - Unique identifier for the target group or username of the target supergroup
//                  or channel (in the format @channelusername)
// userID         - Unique identifier of the target user
func (b *API) UnbanChatMember(chatID string, userID int64) (bool, error) {
	r := map[string]interface{}{
		"chat_id": chatID,
		"user_id": userID,
	}

	_, err := b.requester.JSONRequest("unbanChatMember", r)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Use this method to restrict a user in a supergroup. The bot must be an administrator in the supergroup for this
// to work and must have the appropriate admin rights. Pass True for all permissions to lift restrictions from a user.
// Returns True on success.
func (b *API) RestrictChatMember(request models.ChatMemberRestrictionsRequest) (bool, error) {
	_, err := b.requester.JSONRequest("restrictChatMember", request)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Use this method to promote or demote a user in a supergroup or a channel. The bot must be an administrator
// in the chat for this to work and must have the appropriate admin rights. Pass False for all boolean parameters
// to demote a user. Returns True on success.
func (b *API) PromoteChatMember(request models.ChatMemberPromotionRequest) (bool, error) {
	_, err := b.requester.JSONRequest("promoteChatMember", request)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Use this method to set a custom title for an administrator in a supergroup promoted by the bot.
// Returns True on success.
//
// chatID - Unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername)
// userID - Unique identifier of the target user
// title  - New custom title for the administrator; 0-16 characters, emoji are not allowed
func (b *API) SetChatAdministratorCustomTitle(chatID string, userID int64, title string) (bool, error) {
	r := map[string]interface{}{
		"chat_id":      chatID,
		"user_id":      userID,
		"custom_title": title,
	}

	_, err := b.requester.JSONRequest("setChatAdministratorCustomTitle", r)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Use this method to set default chat permissions for all members. The bot must be an administrator in the group
// or a supergroup for this to work and must have the can_restrict_members admin rights. Returns True on success.
//
// chatID      - Unique identifier for the target chat or username of the target supergroup
//               (in the format @supergroupusername)
// permissions - New default chat permissions
func (b *API) SetChatPermissions(chatID string, permissions models.ChatPermissions) (bool, error) {
	r := map[string]interface{}{
		"chat_id":     chatID,
		"permissions": permissions,
	}

	_, err := b.requester.JSONRequest("setChatPermissions", r)
	if err != nil {
		return false, err
	}

	return true, nil
}

// Use this method to generate a new invite link for a chat; any previously generated link is revoked. The bot
// must be an administrator in the chat for this to work and must have the appropriate admin rights.
// Returns the new invite link as String on success.
//
// Note: Each administrator in a chat generates their own invite links. Bots can't use invite links generated
//       by other administrators. If you want your bot to work with invite links, it will need to generate
//       its own link using exportChatInviteLink – after this the link will become available to the bot
//       via the getChat method. If your bot needs to generate a new invite link replacing its previous one,
//       use exportChatInviteLink again.
func (b *API) ExportChatInviteLink(chatID string) (string, error) {
	r := map[string]interface{}{
		"chat_id": chatID,
	}

	data, err := b.requester.JSONRequest("exportChatInviteLink", r)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

// Use this method to set a new profile photo for the chat. Photos can't be changed for private chats.
// The bot must be an administrator in the chat for this to work and must have the appropriate admin rights.
// Returns True on success.
func (b *API) SetChatPhoto(request models.ChatSetPhotoRequest) (bool, error) {
	_, err := b.requester.JSONRequest("setChatPhoto", request)

	return err == nil, err
}

// Use this method to delete a chat photo. Photos can't be changed for private chats. The bot must be an administrator
// in the chat for this to work and must have the appropriate admin rights. Returns True on success.
func (b *API) DeleteChatPhoto(chatID string) (bool, error) {
	r := map[string]interface{}{
		"chat_id": chatID,
	}

	_, err := b.requester.JSONRequest("deleteChatPhoto", r)

	return err == nil, err
}

// Use this method to change the title of a chat. Titles can't be changed for private chats. The bot must be
// an administrator in the chat for this to work and must have the appropriate admin rights.
// Returns True on success.
//
// chatID - Unique identifier for the target chat or username of the target channel (in the format @channelusername)
// title  - New chat title, 1-255 characters
func (b *API) SetChatTitle(chatID, title string) (bool, error) {
	r := map[string]interface{}{
		"chat_id": chatID,
		"title":   title,
	}

	_, err := b.requester.JSONRequest("setChatTitle", r)

	return err == nil, err
}

// Use this method to change the description of a group, a supergroup or a channel. The bot must be an administrator
// in the chat for this to work and must have the appropriate admin rights. Returns True on success.
//
// chatID       - Unique identifier for the target chat or username of the target channel
//                (in the format @channelusername)
// description  - New chat description, 1-255 characters
func (b *API) SetChatDescription(chatID, description string) (bool, error) {
	r := map[string]interface{}{
		"chat_id":     chatID,
		"description": description,
	}

	_, err := b.requester.JSONRequest("setChatDescription", r)

	return err == nil, err
}

// Use this method to pin a message in a group, a supergroup, or a channel. The bot must be an administrator in the chat
// for this to work and must have the ‘can_pin_messages’ admin right in the supergroup or ‘can_edit_messages’ admin
// right in the channel. Returns True on success.
//
// chatID              - Unique identifier for the target chat or username of the target channel
//                       (in the format @channelusername)
// messageID           - Identifier of a message to pin
// disableNotification - Pass True, if it is not necessary to send a notification to all chat members about the new pinned message. Notifications are always disabled in channels.
func (b *API) PinChatMessage(chatID string, messageID int64, disableNotification ...bool) (bool, error) {
	r := map[string]interface{}{
		"chat_id":    chatID,
		"message_id": messageID,
	}

	if len(disableNotification) > 0 {
		r["disable_notification"] = disableNotification[0]
	}

	_, err := b.requester.JSONRequest("pinChatMessage", r)

	return err == nil, err
}

// Use this method to unpin a message in a group, a supergroup, or a channel. The bot must be an administrator
// in the chat for this to work and must have the ‘can_pin_messages’ admin right in the supergroup
// or ‘can_edit_messages’ admin right in the channel. Returns True on success.
//
// chatID - Unique identifier for the target chat or username of the target channel (in the format @channelusername)
func (b *API) UnpinChatMessage(chatID string) (bool, error) {
	r := map[string]interface{}{
		"chat_id": chatID,
	}

	_, err := b.requester.JSONRequest("unpinChatMessage", r)

	return err == nil, err
}

// Use this method for your bot to leave a group, supergroup or channel. Returns True on success.
//
// chatID - Unique identifier for the target chat or username of the target channel (in the format @channelusername)
func (b *API) LeaveChat(chatID string) (bool, error) {
	r := map[string]interface{}{
		"chat_id": chatID,
	}

	_, err := b.requester.JSONRequest("leaveChat", r)

	return err == nil, err
}

// Use this method to get up to date information about the chat (current name of the user for one-on-one conversations,
// current username of a user, group or channel, etc.). Returns a Chat object on success.
//
// chatID - Unique identifier for the target chat or username of the target channel (in the format @channelusername)
func (b *API) GetChat(chatID string) (*models.Chat, error) {
	r := map[string]interface{}{
		"chat_id": chatID,
	}

	data, err := b.requester.JSONRequest("getChat", r)
	if err != nil {
		return nil, err
	}

	var chat models.Chat
	err = json.Unmarshal(data, &chat)
	if err != nil {
		return nil, err
	}

	return &chat, err
}

// Use this method to get a list of administrators in a chat. On success, returns an Array of ChatMember objects that
// contains information about all chat administrators except other bots. If the chat is a group or a supergroup and
// no administrators were appointed, only the creator will be returned.
//
// chatID - Unique identifier for the target chat or username of the target channel (in the format @channelusername)
func (b *API) GetChatAdministrators(chatID string) ([]string, error) {
	r := map[string]interface{}{
		"chat_id": chatID,
	}

	data, err := b.requester.JSONRequest("getChat", r)
	if err != nil {
		return nil, err
	}

	var response []string
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	return response, err
}

// Use this method to get the number of members in a chat. Returns Int on success.
//
// chatID - Unique identifier for the target chat or username of the target channel (in the format @channelusername)
func (b *API) GetChatMembersCount(chatID string) (int, error) {
	r := map[string]interface{}{
		"chat_id": chatID,
	}

	data, err := b.requester.JSONRequest("getChat", r)
	if err != nil {
		return 0, err
	}

	var response int
	err = json.Unmarshal(data, &response)
	if err != nil {
		return 0, err
	}

	return response, err
}

// Use this method to get information about a member of a chat. Returns a ChatMember object on success.
//
// chatID - Unique identifier for the target chat or username of the target channel (in the format @channelusername)
// userID - Unique identifier of the target user
func (b *API) GetChatMember(chatID string, userID int64) (*models.ChatMember, error) {
	r := map[string]interface{}{
		"chat_id": chatID,
		"user_id": userID,
	}

	data, err := b.requester.JSONRequest("getChatMember", r)
	if err != nil {
		return nil, err
	}

	var response models.ChatMember
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	return &response, err
}

// Use this method to set a new group sticker set for a supergroup. The bot must be an administrator in the chat
// for this to work and must have the appropriate admin rights. Use the field can_set_sticker_set optionally returned
// in getChat requests to check if the bot can use this method. Returns True on success.
//
// chatID - Unique identifier for the target chat or username of the target channel (in the format @channelusername)
// userID - Unique identifier of the target user
func (b *API) SetChatStickerSet(chatID, name string) (bool, error) {
	r := map[string]interface{}{
		"chat_id":          chatID,
		"sticker_set_name": name,
	}

	_, err := b.requester.JSONRequest("setChatStickerSet", r)

	return err == nil, err
}

// Use this method to delete a group sticker set from a supergroup. The bot must be an administrator in the chat
// for this to work and must have the appropriate admin rights. Use the field can_set_sticker_set optionally returned
// in getChat requests to check if the bot can use this method. Returns True on success.
//
// chatID - Unique identifier for the target chat or username of the target channel (in the format @channelusername)
func (b *API) DeleteChatStickerSet(chatID string) (bool, error) {
	r := map[string]interface{}{
		"chat_id": chatID,
	}

	_, err := b.requester.JSONRequest("deleteChatStickerSet", r)

	return err == nil, err
}

// Use this method to send answers to callback queries sent from inline keyboards. The answer will be displayed
// to the user as a notification at the top of the chat screen or as an alert. On success, True is returned.
//
// Alternatively, the user can be redirected to the specified Game URL. For this option to work, you must first create
// a game for your bot via @Botfather and accept the terms. Otherwise, you may use links like t.me/your_bot?start=XXXX
// that open your bot with a parameter.
func (b *API) AnswerCallbackQuery(request models.AnswerCallbackQuery) (bool, error) {
	_, err := b.requester.JSONRequest("answerCallbackQuery", request)

	return err == nil, err
}

// Use this method to change the list of the bot's commands. Returns True on success.
// commands - A list of bot commands to be set as the list of the bot's commands.
//            At most 100 commands can be specified.
func (b *API) SetMyCommands(request []models.BotCommand) (bool, error) {
	_, err := b.requester.JSONRequest("setMyCommands", request)

	return err == nil, err
}

// Use this method to get the current list of the bot's commands. Requires no parameters.
// Returns Array of BotCommand on success.
func (b *API) GetMyCommands() ([]models.BotCommand, error) {
	data, err := b.requester.JSONRequest("setMyCommands", []byte(""))
	if err != nil {
		return nil, err
	}

	var response []models.BotCommand
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Use this method to edit text and game messages. On success, if edited message is sent by the bot, the edited Message
// is returned, otherwise True is returned.
func (b *API) EditMessageText(request models.EditMessageTextRequest) (*models.Message, error) {
	return b.sendMessage("editMessageText", request)
}

// Use this method to edit captions of messages. On success, if edited message is sent by the bot,
// the edited Message is returned, otherwise True is returned.
func (b *API) EditMessageCaption(request models.EditMessageCaptionRequest) (*models.Message, error) {
	return b.sendMessage("editMessageCaption", request)
}

// Use this method to edit animation, audio, document, photo, or video messages. If a message is a part of a message
// album, then it can be edited only to a photo or a video. Otherwise, message type can be changed arbitrarily.
// When inline message is edited, new file can't be uploaded. Use previously uploaded file via its file_id or specify
// a URL. On success, if the edited message was sent by the bot, the edited Message is returned,
// otherwise True is returned.
func (b *API) EditMessageMedia(request models.EditMessageMediaRequest) (*models.Message, error) {
	return b.sendMessage("editMessageMedia", request)
}

// Use this method to edit only the reply markup of messages. On success, if edited message is sent by the bot,
// the edited Message is returned, otherwise True is returned.
func (b *API) EditMessageReplyMarkup(request models.EditMessageReplyMarkupRequest) (*models.Message, error) {
	return b.sendMessage("editMessageReplyMarkup", request)
}

// Use this method to send static .WEBP or animated .TGS stickers. On success, the sent Message is returned.
func (b *API) SendSticker(request models.SendStickerRequest) (*models.Message, error) {
	return b.sendMessage("sendSticker", request)
}

// Use this method to get a sticker set. On success, a StickerSet object is returned.
func (b *API) GetStickerSet(name string) (*models.StickerSet, error) {
	data, err := b.requester.JSONRequest("setMyCommands", map[string]interface{}{
		"name": name,
	})
	if err != nil {
		return nil, err
	}

	var response models.StickerSet
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// Use this method to upload a .PNG file with a sticker for later use in createNewStickerSet and addStickerToSet
// methods (can be used multiple times). Returns the uploaded File on success.
//
// userID  - User identifier of sticker file owner
// sticker - Png image with the sticker, must be up to 512 kilobytes in size, dimensions must not exceed 512px,
//           and either width or height must be exactly 512px. More info on Sending Files »
func (b *API) UploadStickerFile(userID int64, sticker models.InputFile) (*models.File, error) {
	data, err := b.requester.JSONRequest("uploadStickerFile", map[string]interface{}{
		"user_id":     userID,
		"png_sticker": sticker,
	})
	if err != nil {
		return nil, err
	}

	var response models.File
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// Use this method to create a new sticker set owned by a user. The bot will be able to edit the sticker
// set thus created. You must use exactly one of the fields png_sticker or tgs_sticker.
// Returns True on success.
func (b *API) CreateNewStickerSet(request models.NewStickerSetRequest) (bool, error) {
	_, err := b.requester.JSONRequest("createNewStickerSet", request)

	return err == nil, err
}

// Use this method to add a new sticker to a set created by the bot. You must use exactly one of the fields
// png_sticker or tgs_sticker. Animated stickers can be added to animated sticker sets and only to them.
// Animated sticker sets can have up to 50 stickers. Static sticker sets can have up to 120 stickers.
// Returns True on success.
func (b *API) AddStickerToSet(request models.AddStickerToSetSetRequest) (bool, error) {
	_, err := b.requester.JSONRequest("addStickerToSet", request)

	return err == nil, err
}

// Use this method to move a sticker in a set created by the bot to a specific position. Returns True on success.
//
// sticker  - File identifier of the sticker
// position - New sticker position in the set, zero-based
func (b *API) SetStickerPositionInSet(sticker string, position int) (bool, error) {
	_, err := b.requester.JSONRequest("setStickerPositionInSet", map[string]interface{}{
		"sticker":  sticker,
		"position": position,
	})

	return err == nil, err
}

// Use this method to delete a sticker from a set created by the bot. Returns True on success.
//
// sticker - File identifier of the sticker
func (b *API) DeleteStickerFromSet(sticker string) (bool, error) {
	_, err := b.requester.JSONRequest("deleteStickerFromSet", map[string]interface{}{
		"sticker": sticker,
	})

	return err == nil, err
}

// Use this method to set the thumbnail of a sticker set. Animated thumbnails can be set for animated sticker sets only.
// Returns True on success.
func (b *API) SetStickerSetThumb(request models.StickerSetThumbRequest) (bool, error) {
	_, err := b.requester.JSONRequest("setStickerSetThumb", request)

	return err == nil, err
}

func (b *API) Subscribe(t models.UpdateType) <-chan models.Update {
	return b.subscribers.Subscribe(t)
}

func (b *API) Unsubscribe(t models.UpdateType) {
	b.subscribers.Unsubscribe(t)
}

func (b *API) sendMessage(method string, request interface{}) (*models.Message, error) {
	data, err := b.requester.JSONRequest(method, request)
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
