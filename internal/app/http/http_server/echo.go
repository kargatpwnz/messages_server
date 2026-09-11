package http_server

import (
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
		echo, err = s.messagesService.EchoJSON(body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	default:
		log.Println("got unknown contentType:", contentType)
		return
	}

	w.Write(echo)
}
