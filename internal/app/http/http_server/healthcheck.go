package http_server

import "net/http"

func (s *MessagesServer) Healthcheck(w http.ResponseWriter, r *http.Request) {
	alive := s.messagesService.IsAlive()
	if alive {
		w.Write([]byte("ok"))
	}
}
