package http_server

import (
	"encoding/json"
	"net/http"
)

func (s *MessagesServer) GetMessages(w http.ResponseWriter, r *http.Request) {
	msgs := s.messagesService.GetMessages()

	json.NewEncoder(w).Encode(msgs)
}
