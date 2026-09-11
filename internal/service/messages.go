package service

import "time"

func (s *MessagesService) Messages(m Message) (uint64, time.Time) {
	idx := s.Count.Add(1)
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	s.Storage[idx] = m.Message
	return idx, time.Now()
}
