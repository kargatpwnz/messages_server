package service

func (s *MessagesService) DeleteMessage(id uint64) bool {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	_, hasID := s.Storage[id]
	if hasID {
		delete(s.Storage, id)
	}
	return hasID
}
