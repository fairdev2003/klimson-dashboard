package helpers

import "sync"

type StateHub struct {
	mu     sync.RWMutex
	values map[string]interface{}
}

var instance = &StateHub{
	values: make(map[string]interface{}),
}

func GetState() *StateHub {
	return instance
}

func (s *StateHub) Set(key string, value interface{}) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.values[key] = value
}

func (s *StateHub) Get(key string) (interface{}, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.values[key]
	return val, ok
}
