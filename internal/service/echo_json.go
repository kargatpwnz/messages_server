package service

import "encoding/json"

// TODO: move to proper folder
type Message struct {
	Message string
}

func (s *MessagesService) EchoJSON(bytes []byte) ([]byte, error) {
	m := Message{}
	err := json.Unmarshal(bytes, &m)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
