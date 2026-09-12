package service

import (
	"sync"
	"sync/atomic"
	"time"
)

type msg struct {
	CreatedAt time.Time
	Message   string
}
type (
	storage map[uint64]msg // probably we can store messages in slice to keep insertion order without extra Sort on GET /messages
)

type MessagesService struct {
	Count   atomic.Uint64
	Storage storage // this should be repository interface
	Mutex   sync.RWMutex
}

func NewMessagesService() *MessagesService {
	return &MessagesService{
		Count:   atomic.Uint64{},
		Storage: make(storage),
	}
}

func (s *MessagesService) IsAlive() bool {
	return true
}
