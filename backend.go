package main

type KVRecord struct {
	Key   string
	Value string
}

type Backend interface {
	Set(record *KVRecord)
	Get(key string) (*KVRecord, bool)
	Keys() []*KVRecord
}
