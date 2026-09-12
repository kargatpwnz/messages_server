package service

import (
	"entrytest/internal/model"
	"sort"
)

func (s *MessagesService) GetMessages() []model.Message {
	s.Mutex.Lock()

	messages := make([]model.Message, 0, len(s.Storage))
	for id, message := range s.Storage {
		messages = append(messages, model.Message{
			ID:      id,
			Message: message})
	}
	s.Mutex.Unlock()

	sort.Slice(messages, func(i, j int) bool {
		return messages[i].ID > messages[j].ID
	})
	return messages
}
