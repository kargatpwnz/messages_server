package service

import (
	"entrytest/internal/model"
	"time"
)

func (s *MessagesService) PostMessage(m model.Message) (uint64, time.Time) {
	idx := s.count.Add(1) - 1
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.storage[idx] = msg{
		message:   m.Message,
		createdAt: time.Now(),
	}
	return idx, s.storage[idx].createdAt
}
