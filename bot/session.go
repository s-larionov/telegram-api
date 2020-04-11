package bot

import (
	"sync"
)

type Session struct {
	state  *State
	lock   sync.RWMutex
	UserID int64
}

func NewSession(userID int64) *Session {
	return &Session{
		state:  NewState(),
		UserID: userID,
	}
}

func (s *Session) GetState() *State {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.state
}

func (s *Session) UpdateState(state *State) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.state = state
}
