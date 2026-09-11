package service

// TODO: move to proper folder
type Message struct {
	Message string
}

func (s *MessagesService) EchoJSON(m Message) string {
	return m.Message
}
