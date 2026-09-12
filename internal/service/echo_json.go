package service

import "entrytest/internal/model"

func (s *MessagesService) EchoJSON(m model.Message) string {
	return m.Message
}
