package http_server

import (
	"entrytest/internal/model"
	"entrytest/internal/service"
	"net/http"
	"time"
)

type MessagesService interface {
	IsAlive() bool
	EchoRaw([]byte) []byte
	EchoJSON(service.Message) string
	Messages(service.Message) (uint64, time.Time)
	GetMessages() []model.Message
	DeleteMessage(uint64) bool
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
	mux.Handle("GET /", http.FileServer(http.Dir("frontend")))
	mux.HandleFunc("GET /health", server.Healthcheck)
	mux.HandleFunc("POST /echo", server.Echo)
	mux.HandleFunc("POST /messages", server.Messages)
	mux.HandleFunc("GET /messages", server.GetMessages)
	mux.HandleFunc("DELETE /messages/{id}", server.DeleteMessage)

	httpServer := http.Server{
		Addr:    ":" + port,
		Handler: mux}
	server.server = &httpServer

	return server
}
