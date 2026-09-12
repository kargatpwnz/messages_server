package http_server

import (
	"encoding/json"
	"entrytest/internal/model"
	"log"
	"net/http"
)

func (s *MessagesServer) Messages(w http.ResponseWriter, r *http.Request) {
	var err error
	m := model.Message{}
	if err = json.NewDecoder(r.Body).Decode(&m); err != nil || len(m.Message) == 0 { // use validation pkg
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	idx, createdAt := s.messagesService.Messages(m)
	rsp := model.Message{
		ID:        idx,
		CreatedAt: createdAt,
		Message:   m.Message,
	}

	bytes, err := json.Marshal(rsp)
	if err != nil {
		log.Println("Marshall message err:", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(bytes)
}
