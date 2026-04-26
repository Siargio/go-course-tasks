// Задание: Health-check endpoint
//
// Подними HTTP-сервер на :8080 и реализуй маршрут GET /health,
// который возвращает статус 200 и тело "ok".
//
// Ожидаемый результат:
//   $ go run main.go
//   2025/01/01 12:00:00 server started on :8080
//
//   $ curl -i http://localhost:8080/health
//   HTTP/1.1 200 OK
//   Content-Length: 2
//
//   ok

package main

import (
	"fmt"
	"log"
	"net/http"
)

func handlerHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("OK"))
	if err != nil {
		fmt.Println("fail to write HTTP response:", err)
	}
}

func main() {
	// компонент, который выбирает нужный handler по маршруту.
	mux := http.NewServeMux()

	// TODO: зарегистрируй обработчик для "GET /health"
	// Обработчик должен:
	//   1. Установить статус 200 (http.StatusOK)
	//   2. Записать в тело ответа строку "ok"
	// Подсказка: mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) { ... })

	// TODO: запусти сервер на :8080 с помощью http.ListenAndServe(":8080", mux)
	mux.HandleFunc("GET /health", handlerHealth)

	log.Println("server started on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
