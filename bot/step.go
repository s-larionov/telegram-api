package bot

import (
	"sync"

	"github.com/s-larionov/telegram-api"
	"github.com/s-larionov/telegram-api/models"
)

const StepNone StepName = ""

type StepName string

type Step interface {
	GetName() StepName
	IsAllowedFrom(StepName) bool
	AllowFrom(StepName)
	DenyFrom(StepName)
	Process(Session, models.Update) error
	OnLeave(Session, models.Update) error
	Supports(models.Update) bool
}

type StepBase struct {
	lock    sync.RWMutex
	allowed []StepName
	denied  []StepName
	API     *telegram.API
	Name    StepName
}

func NewStepBase(name StepName, api *telegram.API) StepBase {
	return StepBase{
		API:  api,
		Name: name,
	}
}

func (s *StepBase) GetName() StepName {
	return s.Name
}

func (s *StepBase) IsAllowedFrom(step StepName) bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	for _, st := range s.denied {
		if st == step {
			return false
		}
	}

	if len(s.allowed) == 0 {
		return true
	}

	for _, st := range s.allowed {
		if st == step {
			return true
		}
	}

	return false
}

func (s *StepBase) AllowFrom(step StepName) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.allowed = append(s.allowed, step)
}

func (s *StepBase) DenyFrom(step StepName) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.denied = append(s.denied, step)
}

func (s *StepBase) Process(_ Session, _ models.Update) error {
	return nil
}

func (s *StepBase) OnLeave(_ Session, _ models.Update) error {
	return nil
}

func (s *StepBase) Supports(_ models.Update) bool {
	return true
}
