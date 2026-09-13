package service

func (s *MessagesService) DeleteMessage(id uint64) bool {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	_, hasID := s.storage[id]
	if hasID {
		delete(s.storage, id)
	}
	return hasID
}
