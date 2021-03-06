package base

import (
	"errors"
	"reflect"
	"sync"

	"github.com/s-larionov/telegram-api/models"
)

var (
	ErrFieldNotFound        = errors.New("field doesn't exist")
	ErrElementMustBePointer = errors.New("element must be pointer")
	ErrUnableToSetPointer   = errors.New("unable to set pointer")
)

type State interface {
	SetLastStep(step StepName, u models.Update)
	GetLastStep() (StepName, models.Update)
	Set(field string, value interface{})
	Get(field string) (value interface{}, ok bool)
	Load(field string, element interface{}) error
}

type state struct {
	data   map[string]interface{}
	lock   sync.RWMutex
	step   StepName
	update models.Update
}

func NewState() State {
	return &state{
		data: make(map[string]interface{}),
		step: StepNone,
	}
}

func (s *state) SetLastStep(step StepName, u models.Update) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.step = step
	s.update = u
}

func (s *state) GetLastStep() (StepName, models.Update) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.step, s.update
}

func (s *state) Set(field string, value interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.data[field] = value
}

func (s *state) Get(field string) (value interface{}, ok bool) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	value, ok = s.data[field]

	return
}

func (s *state) Load(field string, element interface{}) error {
	s.lock.RLock()
	value, ok := s.data[field]
	s.lock.RUnlock()
	if !ok {
		return ErrFieldNotFound
	}

	iv := reflect.ValueOf(value)
	rv := reflect.ValueOf(element)

	if rv.Type().Kind() != reflect.Ptr {
		return ErrElementMustBePointer
	}

	rv = rv.Elem()
	rt := rv.Type()

	if iv.Type().Kind() != reflect.Ptr {
		rv.Set(iv.Elem())
		return nil
	}

	if rt.Kind() != reflect.Ptr {
		rv.Set(iv.Elem())

		return nil
	}

	// pointer inside pointer case
	iv = iv.Elem()
	if !iv.CanAddr() {
		return ErrUnableToSetPointer
	}

	rv.Set(iv.Addr())

	return nil
}
