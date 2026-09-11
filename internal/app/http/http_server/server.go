package http_server

import (
	"net/http"
)

type MessagesService interface {
	IsAlive() bool
	EchoRaw([]byte) []byte
	EchoJSON([]byte) ([]byte, error)
}

type MessagesServer struct {
	messagesService MessagesService
	server          *http.Server
}

func (s *MessagesServer) ListenAndServe() error {
	return s.server.ListenAndServe()
}

func NewMessagesServer(service MessagesService, port string) *MessagesServer {
	server := &MessagesServer{messagesService: service}

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("frontend")))
	mux.HandleFunc("/health", server.Healthcheck)
	mux.HandleFunc("/echo", server.Echo)

	httpServer := http.Server{
		Addr:    ":" + port,
		Handler: mux}
	server.server = &httpServer

	return server
}
