package bot

import (
	"sync"
)

type State struct {
	data map[string]interface{}
	lock sync.RWMutex
	Step Step
}

func NewState() *State {
	return &State{
		data: make(map[string]interface{}),
	}
}

func (s *State) Set(field string, value interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.data[field] = value
}

func (s *State) Get(field string) (value interface{}, ok bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	value, ok = s.data[field]

	return
}
