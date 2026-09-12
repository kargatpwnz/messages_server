package service

import "entrytest/internal/model"

func (s *MessagesService) GetMessages() []model.Message {
	messages := make([]model.Message, 0, s.Count.Load())
	s.Mutex.Lock()
	defer s.Mutex.Unlock()

	for i := s.Count.Load(); i != 0; i-- {
		messages = append(messages, model.Message{
			ID:      i - 1,
			Message: s.Storage[i-1]})
	}
	return messages
}
