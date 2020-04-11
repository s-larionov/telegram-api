package bot

import (
	"sync"

	"telegram/models"
)

type StepName string

type Step interface {
	GetName() StepName
	IsAllowedFrom(step Step) bool
	AllowFrom(step Step) error
	DenyFrom(step Step) error
	Process(*Session, interface{}) error
	OnFinish(*Session) error
	Supports(t models.UpdateType) bool
}

type StepBase struct {
	lock    sync.RWMutex
	allowed []Step
	denied  []Step
	Name    StepName
}

func (s *StepBase) GetName() StepName {
	return s.Name
}

func (s *StepBase) IsAllowedFrom(step Step) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	name := step.GetName()

	for _, st := range s.denied {
		if st.GetName() == name {
			return false
		}
	}

	if len(s.allowed) == 0 {
		return true
	}

	for _, st := range s.allowed {
		if st.GetName() == name {
			return true
		}
	}

	return false
}

func (s *StepBase) AllowFrom(step Step) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.allowed = append(s.allowed, step)

	return nil
}

func (s *StepBase) DenyFrom(step Step) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.denied = append(s.denied, step)

	return nil
}

func (s *StepBase) Process(_ *Session, _ interface{}) error {
	return nil
}

func (s *StepBase) OnFinish(_ *Session) error {
	return nil
}

func (s *StepBase) Supports(_ models.UpdateType) bool {
	return true
}
