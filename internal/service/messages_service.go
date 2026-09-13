package service

import (
	"sync"
	"sync/atomic"
	"time"
)

type msg struct {
	createdAt time.Time
	message   string
}
type (
	storage map[uint64]msg // probably we can store messages in slice to keep insertion order without extra Sort on GET /messages
)

type MessagesService struct {
	count   atomic.Uint64
	storage storage // this should be repository interface
	mutex   sync.RWMutex
}

func NewMessagesService() *MessagesService {
	return &MessagesService{
		count:   atomic.Uint64{},
		storage: make(storage),
	}
}

func (s *MessagesService) IsAlive() bool {
	return true
}
