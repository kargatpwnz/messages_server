package http_server

import (
	"encoding/json"
	"entrytest/internal/model"
	"io"
	"log"
	"net/http"
)

func (s *MessagesServer) Echo(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Echo ReadAll err", err.Error())
		return
	}

	var echo []byte
	contentType := r.Header.Get("Content-Type")
	switch contentType {
	case "text/plain":
		w.Header().Set("Content-Type", "text/plain")
		echo = s.messagesService.EchoRaw(body)
	case "application/json":
		w.Header().Set("Content-Type", "application/json")

		m := model.Message{}
		err := json.Unmarshal(body, &m)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		text := s.messagesService.EchoJSON(m)
		echo, err = json.Marshal(model.Message{
			Message: text,
		})
		if err != nil {
			log.Println("cant Marshall response", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	default:
		log.Println("got unknown contentType:", contentType)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Write(echo)
}
