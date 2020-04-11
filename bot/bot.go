package bot

import (
	"context"
	"sync"

	"telegram"
	"telegram/models"
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
	b.subscribe(ctx, models.UpdateTypeMessage, b.Flow.onMessage)
	b.subscribe(ctx, models.UpdateTypeEditedMessage, b.Flow.onMessageEdit)
	b.subscribe(ctx, models.UpdateTypeChannelPost, b.Flow.onChannelPost)
	b.subscribe(ctx, models.UpdateTypeEditedChannelPost, b.Flow.onChannelPostEdit)
	b.subscribe(ctx, models.UpdateTypeInlineQuery, b.Flow.onInlineQuery)
	b.subscribe(ctx, models.UpdateTypeChosenInlineResult, b.Flow.onChosenInlineResult)
	b.subscribe(ctx, models.UpdateTypeCallbackQuery, b.Flow.onCallbackQuery)
	b.subscribe(ctx, models.UpdateTypeShippingQuery, b.Flow.onShippingQuery)
	b.subscribe(ctx, models.UpdateTypePreCheckoutQuery, b.Flow.onPreCheckoutQuery)
	b.subscribe(ctx, models.UpdateTypePoll, b.Flow.onPoll)
	b.subscribe(ctx, models.UpdateTypePollAnswer, b.Flow.onPollAnswer)

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
				// TODO: errors
				_ = handler(u)
			}
		}
	}()
}
