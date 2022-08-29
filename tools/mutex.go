package tools

import (
	"sync"
	"time"
)

type Mutex struct {
	*sync.Mutex
	cleanTimer *time.Timer
	V          map[string]interface{}
}

func NewMutex() *Mutex {
	m := new(Mutex)
	m.Mutex = new(sync.Mutex)
	m.V = make(map[string]interface{})
	return m
}

func (m *Mutex) Write(key string, v interface{}) {
	m.Lock()
	defer m.Unlock()
	m.V[key] = v
}

func (m *Mutex) WriteMany(v map[string]interface{}) {
	m.Lock()
	defer m.Unlock()
	for k, v := range v {
		m.V[k] = v
	}
}

func (m *Mutex) WriteManyWithClean(v map[string]interface{}) {
	m.Lock()
	defer m.Unlock()
	for k := range m.V {
		delete(m.V, k)
	}
	for k, v := range v {
		m.V[k] = v
	}
}

func (m *Mutex) Read(key string) interface{} {
	m.Lock()
	defer m.Unlock()
	return m.V[key]
}

func (m *Mutex) Delete(key string) interface{} {
	m.Lock()
	defer m.Unlock()
	v := m.V[key]
	delete(m.V, key)
	return v
}

func (m *Mutex) WriteWithTTL(key string, v interface{}, ttl time.Duration) {
	m.Write(key, v)
	if m.cleanTimer != nil {
		m.cleanTimer.Stop()
	}
	m.cleanTimer = time.AfterFunc(ttl, func() {
		m.Delete(key)
	})
}

func (m *Mutex) Clean() {
	m.Lock()
	defer m.Unlock()
	for k := range m.V {
		delete(m.V, k)
	}
}
