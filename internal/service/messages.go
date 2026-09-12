package service

import (
	"entrytest/internal/model"
	"time"
)

func (s *MessagesService) Messages(m model.Message) (uint64, time.Time) {
	idx := s.Count.Add(1) - 1
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	s.Storage[idx] = msg{
		Message:   m.Message,
		CreatedAt: time.Now(),
	}
	return idx, s.Storage[idx].CreatedAt
}
