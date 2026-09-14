package http_server

import (
	"encoding/json"
	"entrytest/internal/model"
	"log"
	"net/http"
)

func (s *MessagesServer) PostMessage(w http.ResponseWriter, r *http.Request) {
	var err error
	req := message{}
	if err = json.NewDecoder(r.Body).Decode(&req); err != nil || len(req.Message) == 0 { // use validation pkg
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, createdAt := s.messagesService.PostMessage(req.Message)
	rsp := model.Message{
		ID:        id,
		CreatedAt: createdAt,
		Message:   req.Message,
	}

	bytes, err := json.Marshal(rsp)
	if err != nil {
		log.Println("Marshall message err:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if _, err = w.Write(bytes); err != nil {
		log.Println("PostMessage Write err:", err.Error())
	}
}
