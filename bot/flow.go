package bot

import (
	"errors"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/s-larionov/telegram-api/models"
)

var (
	ErrStepAlreadyExist = errors.New("step already exists in the flow")
	ErrUnsupportedEvent = errors.New("unsupported event")
)

type Flow struct {
	storage   Storage
	steps     []Step
	stepsLock sync.RWMutex
}

func NewFlow(storage Storage) *Flow {
	return &Flow{
		storage: storage,
	}
}

func NewFlowWithSteps(storage Storage, steps []Step) (*Flow, error) {
	f := NewFlow(storage)

	for _, step := range steps {
		err := f.AddStep(step)
		if err != nil {
			return nil, err
		}
	}

	return f, nil
}

func (f *Flow) AddStep(step Step) error {
	f.stepsLock.Lock()
	defer f.stepsLock.Unlock()

	for _, s := range f.steps {
		if s.GetName() == step.GetName() {
			return ErrStepAlreadyExist
		}
	}

	f.steps = append(f.steps, step)

	return nil
}

func (f *Flow) onMessage(u models.Update) error {
	log.WithFields(log.Fields{
		"from_id":    u.Message.From.ID,
		"message_id": u.Message.ID,
		"text":       u.Message.Text,
	}).Trace("incoming message")

	session := f.storage.Load(u.Message.From.ID)

	return f.process(u.GetType(), session, u)
}

func (f *Flow) onMessageEdit(u models.Update) error {
	log.WithFields(log.Fields{
		"from_id":    u.EditedMessage.From.ID,
		"message_id": u.EditedMessage.ID,
		"text":       u.EditedMessage.Text,
	}).Trace("message was edited")

	session := f.storage.Load(u.EditedMessage.From.ID)

	return f.process(u.GetType(), session, u)
}

func (f *Flow) onChannelPost(u models.Update) error {
	log.WithFields(log.Fields{
		"from_id": u.ChannelPost.From.ID,
		"chat_id": u.ChannelPost.Chat.ID,
		"post_id": u.ChannelPost.ID,
		"text":    u.ChannelPost.Text,
	}).Trace("incoming post to the channel")

	session := f.storage.Load(u.ChannelPost.From.ID)

	return f.process(u.GetType(), session, u)
}

func (f *Flow) onChannelPostEdit(u models.Update) error {
	log.WithFields(log.Fields{
		"from_id": u.EditedChannelPost.From.ID,
		"chat_id": u.EditedChannelPost.Chat.ID,
		"post_id": u.EditedChannelPost.ID,
		"text":    u.EditedChannelPost.Text,
	}).Trace("channel post was updated")

	session := f.storage.Load(u.EditedChannelPost.From.ID)

	return f.process(u.GetType(), session, u)
}

func (f *Flow) onInlineQuery(u models.Update) error {
	log.WithFields(log.Fields{
		"from_id":  u.InlineQuery.From.ID,
		"query_id": u.InlineQuery.ID,
		"query":    u.InlineQuery.Query,
	}).Trace("incoming inline query")

	session := f.storage.Load(u.InlineQuery.From.ID)

	return f.process(u.GetType(), session, u)
}

func (f *Flow) onChosenInlineResult(u models.Update) error {
	log.WithFields(log.Fields{
		"from_id":           u.ChosenInlineResult.From.ID,
		"query":             u.ChosenInlineResult.Query,
		"inline_message_id": u.ChosenInlineResult.InlineMessageID,
		"result_id":         u.ChosenInlineResult.ID,
	}).Trace("inline result was chosen")

	session := f.storage.Load(u.ChosenInlineResult.From.ID)

	return f.process(u.GetType(), session, u)
}

func (f *Flow) onCallbackQuery(u models.Update) error {
	log.WithFields(log.Fields{
		"from_id":      u.CallbackQuery.From.ID,
		"message_id":   u.CallbackQuery.Message.ID,
		"message_text": u.CallbackQuery.Message.Text,
		"chat":         u.CallbackQuery.ChatInstance,
		"query_data":   u.CallbackQuery.Data,
	}).Trace("incoming callback query")

	session := f.storage.Load(u.CallbackQuery.From.ID)

	return f.process(u.GetType(), session, u)
}

func (f *Flow) onShippingQuery(u models.Update) error {
	log.WithFields(log.Fields{
		"from_id": u.ShippingQuery.From.ID,
		"invoice": u.ShippingQuery.InvoicePayload,
		"id":      u.ShippingQuery.ID,
		"address": u.ShippingQuery.ShippingAddress.String(),
	}).Trace("incoming shipping query")

	session := f.storage.Load(u.ShippingQuery.From.ID)

	return f.process(u.GetType(), session, u)
}

func (f *Flow) onPreCheckoutQuery(u models.Update) error {
	log.WithFields(log.Fields{
		"from_id":  u.PreCheckoutQuery.From.ID,
		"invoice":  u.PreCheckoutQuery.InvoicePayload,
		"id":       u.PreCheckoutQuery.ID,
		"amount":   u.PreCheckoutQuery.TotalAmount,
		"currency": u.PreCheckoutQuery.Currency,
	}).Trace("incoming pre checkout query")

	session := f.storage.Load(u.PreCheckoutQuery.From.ID)

	return f.process(u.GetType(), session, u)
}

func (f *Flow) onPoll(u models.Update) error {
	log.WithFields(log.Fields{
		"poll_id":   u.Poll.ID,
		"poll_type": u.Poll.Type,
	}).Trace("incoming poll")

	session := f.storage.Load(u.PollAnswer.User.ID)

	return f.process(u.GetType(), session, u)
}

func (f *Flow) onPollAnswer(u models.Update) error {
	log.WithFields(log.Fields{
		"user_id":    u.PollAnswer.User.ID,
		"poll_id":    u.PollAnswer.PollID,
		"option_ids": u.PollAnswer.OptionIDs,
	}).Trace("incoming poll answer")

	session := f.storage.Load(u.PollAnswer.User.ID)

	return f.process(u.GetType(), session, u)
}

func (f *Flow) process(t models.UpdateType, session *Session, u models.Update) error {
	state := session.GetState()

	step, err := f.findStep(state, u)
	if err != nil {
		return err
	}

	if state.step != StepNone {
		finishedStep, err := f.findStepByName(state.GetCurrentStep())
		if err != nil {
			return err
		}

		err = finishedStep.OnLeave(session, u)
		if err != nil {
			return err
		}
	}

	state.SetCurrentStep(step.GetName())
	session.UpdateState(state)

	err = step.Process(session, u)
	if err != nil {
		return err
	}

	return nil
}

func (f *Flow) findStep(state *State, u models.Update) (Step, error) {
	f.stepsLock.RLock()
	defer f.stepsLock.RUnlock()

	for _, step := range f.steps {
		if !step.Supports(u) {
			continue
		}

		if !step.IsAllowedFrom(state.step) {
			continue
		}

		return step, nil
	}

	return nil, ErrUnsupportedEvent
}

func (f *Flow) findStepByName(name StepName) (Step, error) {
	f.stepsLock.RLock()
	defer f.stepsLock.RUnlock()

	for _, step := range f.steps {
		if step.GetName() != name {
			continue
		}

		return step, nil
	}

	return nil, ErrUnsupportedEvent
}
