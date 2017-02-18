package main

import (
	"sync"
)

// MemoryBackend implements Backend interface. It stores the records in memory
type MemoryBackend struct {
	mx   *sync.RWMutex
	keys map[string]*KVRecord
}

// NewMemoryBackend creates new instance of MemoryBackend
func NewMemoryBackend() *MemoryBackend {
	return &MemoryBackend{
		mx:   &sync.RWMutex{},
		keys: make(map[string]*KVRecord),
	}
}

// Set puts the record into storage
func (mb *MemoryBackend) Set(record *KVRecord) {
	mb.mx.Lock()
	defer mb.mx.Unlock()

	mb.keys[record.Key] = record
}

// Get gets the record by key. Second return parameter will be `false` if record not found
func (mb *MemoryBackend) Get(key string) (*KVRecord, bool) {
	mb.mx.RLock()
	defer mb.mx.RUnlock()

	if value, ok := mb.keys[key]; ok {
		return value, true
	}

	return nil, false
}

// Keys returns all the records from the storage
func (mb *MemoryBackend) Keys() []*KVRecord {
	mb.mx.RLock()
	defer mb.mx.RUnlock()

	records := make([]*KVRecord, len(mb.keys))
	indexCounter := 0
	for _, value := range mb.keys {
		records[indexCounter] = value
		indexCounter++
	}

	return records
}
