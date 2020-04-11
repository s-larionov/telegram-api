package events

import (
	"sync"

	"github.com/s-larionov/telegram-api/models"
)

const subscriberChannelBufferSize = 5

type Container struct {
	mutex       sync.RWMutex
	subscribers map[models.UpdateType][]chan models.Update
}

func NewContainer() *Container {
	return &Container{
		subscribers: make(map[models.UpdateType][]chan models.Update),
	}
}

func (c *Container) Subscribe(t models.UpdateType) <-chan models.Update {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	subscribers, ok := c.subscribers[t]
	if !ok {
		subscribers = make([]chan models.Update, 0)
	}

	ch := make(chan models.Update, subscriberChannelBufferSize)
	subscribers = append(subscribers, ch)

	c.subscribers[t] = subscribers

	return ch
}

func (c *Container) Unsubscribe(t models.UpdateType) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	subscribers, ok := c.subscribers[t]
	if !ok {
		return
	}

	for i, ch := range subscribers {
		close(ch)
		subscribers[i] = nil
	}

	delete(c.subscribers, t)
}

func (c *Container) Emit(update models.Update) {
	c.mutex.RLock()
	subscribers, ok := c.subscribers[update.GetType()]
	c.mutex.RUnlock()

	if !ok {
		return
	}

	for _, ch := range subscribers {
		ch <- update
	}
}
