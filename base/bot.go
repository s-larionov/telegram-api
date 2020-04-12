package base

import (
	"context"
	"sync"

	log "github.com/sirupsen/logrus"

	"github.com/s-larionov/telegram-api"
	"github.com/s-larionov/telegram-api/models"
)

type handler func(update models.Update) error

type Bot struct {
	Flow          *Flow
	API           *telegram.API
	subscriptions []<-chan models.Update
	wg            sync.WaitGroup
}

func NewBot(api *telegram.API, flow *Flow) *Bot {
	return &Bot{
		API:  api,
		Flow: flow,
	}
}

func (b *Bot) Run(ctx context.Context) error {
	b.subscribe(ctx, models.UpdateTypeMessage, b.Flow.OnMessage)
	b.subscribe(ctx, models.UpdateTypeEditedMessage, b.Flow.OnMessageEdit)
	b.subscribe(ctx, models.UpdateTypeChannelPost, b.Flow.OnChannelPost)
	b.subscribe(ctx, models.UpdateTypeEditedChannelPost, b.Flow.OnChannelPostEdit)
	b.subscribe(ctx, models.UpdateTypeInlineQuery, b.Flow.OnInlineQuery)
	b.subscribe(ctx, models.UpdateTypeChosenInlineResult, b.Flow.OnChosenInlineResult)
	b.subscribe(ctx, models.UpdateTypeCallbackQuery, b.Flow.OnCallbackQuery)
	b.subscribe(ctx, models.UpdateTypeShippingQuery, b.Flow.OnShippingQuery)
	b.subscribe(ctx, models.UpdateTypePreCheckoutQuery, b.Flow.OnPreCheckoutQuery)
	b.subscribe(ctx, models.UpdateTypePoll, b.Flow.OnPoll)
	b.subscribe(ctx, models.UpdateTypePollAnswer, b.Flow.OnPollAnswer)

	b.wg.Wait()

	return nil
}

func (b *Bot) subscribe(ctx context.Context, t models.UpdateType, handler handler) {
	ch := b.API.Subscribe(t)
	b.wg.Add(1)
	go func() {
		defer b.wg.Done()

		var u models.Update
		for {
			select {
			case <-ctx.Done():
				// TODO: ctx.Err()
				return
			case u = <-ch:
				err := handler(u)
				if err == ErrUnsupportedEvent {
					log.WithError(err).Info("unsupported event")
				} else if err != nil {
					log.WithError(err).Error("unable to process request")
				}
			}
		}
	}()
}
