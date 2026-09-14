package service

import "time"

func (s *MessagesService) PostMessage(message string) (uint64, time.Time) {
	idx := s.count.Add(1) - 1
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.storage[idx] = msg{
		message:   message,
		createdAt: time.Now(),
	}
	return idx, s.storage[idx].createdAt
}
