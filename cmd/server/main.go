// Заготовка сервера. Закрывайте этапы по одному, пока не позеленеет go test ./tests/ -v
//
// Запуск: go run ./cmd/server — порт берётся из PORT, по умолчанию 8080.
// Панель уже раздаётся: откройте http://localhost:8080/ и смотрите, как этапы
// зеленеют по ходу работы.
package main

import (
	"entrytest/internal/config"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.CreateConfig()
	if err != nil {
		log.Fatal("Can't create config: ", err.Error())
	}
	mux := http.NewServeMux()

	// Панель из frontend/. Каталог берётся относительно рабочего, поэтому
	// запускайте из корня модуля: go run ./cmd/server
	mux.Handle("/", http.FileServer(http.Dir("frontend")))

	// TODO Этап 1: GET /health           -> 200, тело "ok"
	// TODO Этап 2: POST /echo            -> тело запроса без изменений
	// TODO Этап 3: POST /echo            -> на application/json разобрать {"message": "..."} и вернуть JSON
	// TODO Этап 4: POST /messages        -> сохранить в памяти, 201
	// TODO Этап 5: GET /messages         -> все сообщения, новые сверху
	// TODO Этап 6: DELETE /messages/{id} -> 204, либо 404 если такого нет

	log.Printf("сервер слушает http://localhost:%s", cfg.Server.Port)
	log.Fatal(http.ListenAndServe(":"+cfg.Server.Port, mux))
}
