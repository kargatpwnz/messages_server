package service

import (
	"sync"
	"sync/atomic"
)

type (
	storage map[uint64]string
)

type MessagesService struct {
	Count   atomic.Uint64
	Storage storage
	Mutex   sync.Mutex
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
