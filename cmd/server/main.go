// Заготовка сервера. Закрывайте этапы по одному, пока не позеленеет go test ./tests/ -v
//
// Запуск: go run ./cmd/server — порт берётся из PORT, по умолчанию 8080.
// Панель уже раздаётся: откройте http://localhost:8080/ и смотрите, как этапы
// зеленеют по ходу работы.
package main

import (
	"entrytest/internal/app/http/http_server"
	"entrytest/internal/config"
	"entrytest/internal/service"
	"log"
)

func main() {
	cfg, err := config.CreateConfig()
	if err != nil {
		log.Fatal("Can't create config: ", err.Error())
	}

	// TODO Этап 1: GET /health           -> 200, тело "ok"
	// TODO Этап 2: POST /echo            -> тело запроса без изменений
	// TODO Этап 3: POST /echo            -> на application/json разобрать {"message": "..."} и вернуть JSON
	// TODO Этап 4: POST /messages        -> сохранить в памяти, 201
	// TODO Этап 5: GET /messages         -> все сообщения, новые сверху
	// TODO Этап 6: DELETE /messages/{id} -> 204, либо 404 если такого нет

	service := service.NewMessagesService()
	server := http_server.NewMessagesServer(service, cfg.Server.Port)

	log.Printf("сервер слушает http://localhost:%s", cfg.Server.Port)
	log.Fatal(server.ListenAndServe())
}
