package http_server

import (
	"net/http"
	"strconv"
)

func (s *MessagesServer) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(r.PathValue("id"), 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	deleted := s.messagesService.DeleteMessage(id)
	if deleted {
		w.WriteHeader(http.StatusNoContent)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
