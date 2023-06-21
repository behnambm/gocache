// This driver provides memory storage for cache

package gocache

import (
	"errors"
	"sync"
)

type MemoryDriver struct {
	dataTable map[string][]byte
	mu        sync.RWMutex
}

func NewMemory() MemoryDriver {
	return MemoryDriver{
		dataTable: make(map[string][]byte),
	}
}

func (m *MemoryDriver) Set(key string, value []byte) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.dataTable[key] = value

	return nil
}

func (m *MemoryDriver) Get(key string) ([]byte, error) {
	m.mu.RLocker()
	defer m.mu.RUnlock()
	d, ok := m.dataTable[key]
	if !ok {
		return []byte{}, errors.New("key does not exist")
	}
	return d, nil
}
