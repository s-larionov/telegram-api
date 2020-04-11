package bot

import (
	"sync"
)

type Storage interface {
	Load(userID int64) *Session
	Store(session *Session)
}

func NewInMemoryStorage() Storage {
	return &inMemory{
		storage: make(map[int64]*Session),
		lock:    &sync.Mutex{},
	}
}

type inMemory struct {
	storage map[int64]*Session
	lock    sync.Locker
}

func (s *inMemory) Load(userID int64) *Session {
	s.lock.Lock()
	defer s.lock.Unlock()

	session, ok := s.storage[userID]
	if ok {
		return session
	}

	session = NewSession(userID)

	s.storage[userID] = session

	return session
}

func (s *inMemory) Store(session *Session) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.storage[session.UserID] = session
}
