package lockP

import (
	"sync"
)

type SafeLocks struct {
	m *sync.Map
}

func NewSafeLocks() *SafeLocks {
	return &SafeLocks{
		m: &sync.Map{},
	}
}

func (s *SafeLocks) Lock(key string) {
	val, _ := s.m.LoadOrStore(key, &sync.Mutex{})
	val.(*sync.Mutex).Lock()
}

func (s *SafeLocks) Unlock(key string) {
	val, _ := s.m.Load(key)
	val.(*sync.Mutex).Unlock()
}
