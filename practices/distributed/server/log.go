package server

import (
	"fmt"
	"sync"
	"time"
)

type Log struct {
	mu      sync.Mutex
	records []Record
}

func NewLog() *Log {
	return &Log{}
}

// Store is used to save a record to the log.
func (l *Log) Store(record Record) (uint64, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	record.RequestTime = time.Now().Format("2006-01-02 15:04:05")
	record.Offset = uint64(len(l.records))
	l.records = append(l.records, record)

	return record.Offset, nil
}

// Each time we read a record given an index, we use that index to look up the record in the slice.
// If the offset given by the client doesn’t exist, we return an error saying that the offset doesn’t exist.
func (l *Log) Read(offset uint64) (Record, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	if offset >= uint64(len(l.records)) {
		return Record{}, ErrOffsetNotFound
	}

	return l.records[offset], nil
}

type Record struct {
	Value       []byte `json:"value"`
	ClientIP    string `json:"client_ip"`
	RequestTime string
	Offset      uint64 `json:"offset"`
}

var ErrOffsetNotFound = fmt.Errorf("offset not found")
