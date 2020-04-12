package base

import (
	"sync"
)

type Storage interface {
	Load(userID int64) (Session, error)
	Store(session Session) error
}

func NewInMemoryStorage() Storage {
	return &inMemory{
		storage: make(map[int64]Session),
		lock:    &sync.Mutex{},
	}
}

type inMemory struct {
	storage map[int64]Session
	lock    sync.Locker
}

func (s *inMemory) Load(userID int64) (Session, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	session, ok := s.storage[userID]
	if ok {
		return session, nil
	}

	session = NewSession(userID)

	s.storage[userID] = session

	return session, nil
}

func (s *inMemory) Store(session Session) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.storage[session.GetUserID()] = session

	return nil
}
