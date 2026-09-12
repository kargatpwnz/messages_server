package service

import (
	"entrytest/internal/model"
	"sort"
)

func (s *MessagesService) GetMessages() []model.Message {
	s.Mutex.RLock()

	messages := make([]model.Message, 0, len(s.Storage))
	for id, message := range s.Storage {
		messages = append(messages, model.Message{
			ID:        id,
			Message:   message.Message,
			CreatedAt: message.CreatedAt})
	}
	s.Mutex.RUnlock()

	sort.Slice(messages, func(i, j int) bool {
		return messages[i].ID > messages[j].ID
	})
	return messages
}
