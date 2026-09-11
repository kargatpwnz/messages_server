package service

type MessagesService struct {
}

func NewMessagesService() *MessagesService {
	return &MessagesService{}
}

func (s *MessagesService) IsAlive() bool {
	return true
}
