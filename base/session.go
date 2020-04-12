package base

import (
	"sync"
)

type Session interface {
	GetUserID() int64
	GetState() State
	UpdateState(state State)
}

type session struct {
	state  State
	lock   sync.RWMutex
	userID int64
}

func NewSession(userID int64) Session {
	return &session{
		state:  NewState(),
		userID: userID,
	}
}

func (s *session) GetUserID() int64 {
	return s.userID
}

func (s *session) GetState() State {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.state
}

func (s *session) UpdateState(state State) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.state = state
}
