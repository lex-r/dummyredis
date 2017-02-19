package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewMemoryBackend(t *testing.T) {
	backend := NewMemoryBackend()
	assert.IsType(t, &MemoryBackend{}, backend)
}

func TestSet(t *testing.T) {
	backend := NewMemoryBackend()
	kvRecord := &KVRecord{"key", "value"}
	backend.Set(kvRecord)
}

func TestGetOnEmptyStorage(t *testing.T) {
	backend := NewMemoryBackend()
	record, exists := backend.Get("key")

	assert.Nil(t, record)
	assert.False(t, exists)
}

func TestGet(t *testing.T) {
	backend := NewMemoryBackend()
	loadFixturesIntoBackend(backend)

	records := getFixtures()
	for _, record := range records {
		actual, exists := backend.Get(record.Key)
		assert.Equal(t, record, actual)
		assert.True(t, exists)
	}
}

func TestKeysOnEmptyStorage(t *testing.T) {
	backend := NewMemoryBackend()
	records := backend.Keys()

	assert.Empty(t, records)
}

func TestKeys(t *testing.T) {
	backend := NewMemoryBackend()
	loadFixturesIntoBackend(backend)

	keys := backend.Keys()
	records := getFixtures()

	assert.Equal(t, len(records), len(keys))
	for _, record := range records {
		assert.Contains(t, keys, record)
	}
}

// loadFixturesIntoBackend puts fixtures into backend
func loadFixturesIntoBackend(backend *MemoryBackend) {
	records := getFixtures()

	for _, record := range records {
		backend.Set(record)
	}
}

// getFixtures returns list of KVRecord objects for testing
func getFixtures() []*KVRecord {
	records := make([]*KVRecord, 0)
	records = append(records, &KVRecord{"key1", "value1"})
	records = append(records, &KVRecord{"key2", "value2"})
	records = append(records, &KVRecord{"key3", "value3"})

	return records
}
