package http_server

import (
	"log"
	"net/http"
)

func (s *MessagesServer) Healthcheck(w http.ResponseWriter, r *http.Request) {
	alive := s.messagesService.IsAlive()
	if alive {
		if _, err := w.Write([]byte("ok")); err != nil {
			log.Println("Healthcheck Write err:", err.Error())
		}
	}
}
