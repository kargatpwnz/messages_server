package http_server

import (
	"encoding/json"
	"entrytest/internal/service"
	"net/http"
	"time"
)

type response struct {
	Id        uint64    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Message   string    `json:"message"`
}

func (s *MessagesServer) Messages(w http.ResponseWriter, r *http.Request) {
	var err error
	m := service.Message{}
	if err = json.NewDecoder(r.Body).Decode(&m); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	idx, createdAt := s.messagesService.Messages(m)

	rsp := response{
		Id:        idx,
		CreatedAt: createdAt,
		Message:   m.Message,
	}

	bytes, err := json.Marshal(rsp)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(bytes)
}
