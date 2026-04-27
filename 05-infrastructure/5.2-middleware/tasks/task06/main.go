// Задание: Интеграционная мини-задача — публичный и защищённый endpoint
//
// Собери API с двумя маршрутами:
//   GET /health       — публичный, без авторизации
//   GET /api/v1/me    — только с валидным Bearer-токеном
//
// Требования:
//   - цепочка middleware только на защищённом endpoint
//   - структурированное логирование запросов (slog)
//   - единый формат 401 ошибки
//   - замена верификатора без изменения handler'ов
//
// Ожидаемый результат:
//   $ curl http://localhost:8080/health
//   {"status":"ok"}
//
//   $ curl http://localhost:8080/api/v1/me
//   {"error":"authorization header is empty"}
//
//   $ curl -H "Authorization: Bearer valid-token" http://localhost:8080/api/v1/me
//   {"user_id":"user-123","role":"admin"}

package main

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"strings"
	"time"
)

type Claims struct {
	UserID string
	Role   string
}

type contextKey string

const claimsKey contextKey = "claims"

type TokenVerifier interface {
	Verify(token string) (Claims, error)
}

type mockVerifier struct{}

func (v *mockVerifier) Verify(token string) (Claims, error) {
	if token == "valid-token" {
		return Claims{UserID: "user-123", Role: "admin"}, nil
	}
	return Claims{}, errors.New("invalid token")
}

type Middleware func(http.Handler) http.Handler

func Chain(h http.Handler, mws ...Middleware) http.Handler {
	for i := len(mws) - 1; i >= 0; i-- {
		h = mws[i](h)
	}
	return h
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

// writeError отправляет JSON-ошибку
func writeError(w http.ResponseWriter, status int, msg string) {
	writeJSON(w, status, map[string]string{"error": msg})
}

// TODO: реализуй LoggingMiddleware(logger *slog.Logger) Middleware
// Логируй: method, path, status, duration_ms
func LoggingMiddleware(logger *slog.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Оборачиваем ResponseWriter, чтобы перехватить статус ответа
			type statusRecorder struct {
				http.ResponseWriter
				status int
			}
			rec := &statusRecorder{
				ResponseWriter: w,
				status:         http.StatusOK,
			}

			next.ServeHTTP(rec, r)

			duration := time.Since(start).Milliseconds()
			// Логируем после выполнения
			logger.Info("http request",
				"method", r.Method,
				"path", r.URL.Path,
				"status", rec.status,
				"duration_ms", duration,
			)
		})
	}
}

// TODO: реализуй AuthMiddleware(verifier TokenVerifier) Middleware
// Извлекай Bearer-токен, верифицируй, клади claims в context
// На ошибке — 401 {"error":"<msg>"}
func AuthMiddleware(verifier TokenVerifier) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			//Получаем заголовок Authorization
			authHeder := r.Header.Get("Authorization")
			if authHeder == "" {
				writeError(w, http.StatusUnauthorized, "authorization header is empty")
				return
			}
			//Проверяем формат "Bearer <token>"
			token, found := strings.CutPrefix(authHeder, "Bearer ")
			if !found {
				writeError(w, http.StatusUnauthorized, "invalid authorization header format")
				return
			}

			if token == "" {
				writeError(w, http.StatusUnauthorized, "authorization token is empty")
				return
			}
			// Верифицируем токен
			claims, err := verifier.Verify(token)
			if err != nil {
				writeError(w, http.StatusUnauthorized, err.Error())
				return
			}
			//Кладём claims в контекст
			ctx := context.WithValue(r.Context(), claimsKey, claims)
			//Передаём управление дальше
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func meHandler(w http.ResponseWriter, r *http.Request) {
	// TODO: достань claims из r.Context() и верни {"user_id":"...","role":"..."}
	claims, ok := r.Context().Value(claimsKey).(Claims)

	if !ok {
		writeError(w, http.StatusUnauthorized, "user not found in context")
		return
	} else {
		writeJSON(w, http.StatusOK, map[string]string{
			"user_id": claims.UserID,
			"role":    claims.Role,
		})
	}
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	verifier := &mockVerifier{}

	mux := http.NewServeMux()

	// TODO: зарегистрируй маршруты:
	// /health — только с LoggingMiddleware
	mux.Handle("GET /health", LoggingMiddleware(logger)(http.HandlerFunc(healthHandler)))
	// /api/v1/me — с Chain(LoggingMiddleware + AuthMiddleware)
	protectedChain := Chain(
		http.HandlerFunc(meHandler),
		AuthMiddleware(verifier),  // сначала проверяем токен
		LoggingMiddleware(logger), // потом логируем (снаружи)
	)
	mux.Handle("GET /api/v1/me", protectedChain)

	srv := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	logger.Info("server started", "addr", ":8080")
	if err := srv.ListenAndServe(); err != nil {
		logger.Error("server failed", "error", err)
	}
}
