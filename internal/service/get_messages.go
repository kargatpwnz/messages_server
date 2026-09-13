package service

import (
	"entrytest/internal/model"
	"sort"
)

func (s *MessagesService) GetMessages() []model.Message {
	s.mutex.RLock()

	messages := make([]model.Message, 0, len(s.storage))
	for id, message := range s.storage {
		messages = append(messages, model.Message{
			ID:        id,
			Message:   message.message,
			CreatedAt: message.createdAt})
	}
	s.mutex.RUnlock()

	sort.Slice(messages, func(i, j int) bool {
		return messages[i].ID > messages[j].ID
	})
	return messages
}
