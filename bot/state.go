package bot

import (
	"errors"
	"reflect"
	"sync"
)

var (
	ErrFieldNotFound        = errors.New("field doesn't exist")
	ErrElementMustBePointer = errors.New("element must be pointer")
	ErrUnableToSetPointer   = errors.New("unable to set pointer")
)

type State struct {
	data map[string]interface{}
	lock sync.RWMutex
	step StepName
}

func NewState() *State {
	return &State{
		data: make(map[string]interface{}),
		step: StepNone,
	}
}

func (s *State) SetCurrentStep(step StepName) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.step = step
}

func (s *State) GetCurrentStep() StepName {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.step
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

func (s *State) Load(field string, element interface{}) error {
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
