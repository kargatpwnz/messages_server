package http_server

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *MessagesServer) GetMessages(w http.ResponseWriter, r *http.Request) {
	msgs := s.messagesService.GetMessages()
	if err := json.NewEncoder(w).Encode(msgs); err != nil {
		log.Println("GetMessages: Encode err ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}
}
