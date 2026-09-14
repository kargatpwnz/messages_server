package http_server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type message struct {
	Message string `json:"message"`
}

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

		m := message{}
		err = json.Unmarshal(body, &m)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		text := s.messagesService.EchoJSON(m.Message)
		echo, err = json.Marshal(message{
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

	if _, err = w.Write(echo); err != nil {
		log.Println("Echo Write err:", err.Error())
	}
}
