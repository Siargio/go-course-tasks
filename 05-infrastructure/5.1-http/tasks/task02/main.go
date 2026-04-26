// Задание: Path-параметры
//
// Добавь маршрут GET /api/v1/users/{id}, который возвращает JSON с user_id.
//
// Ожидаемый результат:
//   $ go run main.go
//   server started on :8080
//
//   $ curl http://localhost:8080/api/v1/users/u-42
//   {"user_id":"u-42"}

package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type userResponse struct {
	UserID string `json:"user_id"`
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func handlerUser(w http.ResponseWriter, r *http.Request) {
	// Достаём id из пути
	id := r.PathValue("id")

	if id == "" {
		writeJSON(w, http.StatusBadRequest, map[string]string{"error": "id is required"})
		return
	}

	writeJSON(w, http.StatusOK, userResponse{UserID: id})
}

func main() {
	mux := http.NewServeMux()

	// TODO: зарегистрируй обработчик для "GET /api/v1/users/{id}"
	// Обработчик должен:
	//   1. Достать id через r.PathValue("id")
	//   2. Вернуть JSON {"user_id":"<id>"} со статусом 200
	// Подсказка: mux.HandleFunc("GET /api/v1/users/{id}", func(...) { ... })
	mux.HandleFunc("GET /api/v1/users/{id}", handlerUser)

	log.Println("server started on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
