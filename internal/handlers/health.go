package handlers

import (
	"fmt"
	"go-production-stack/internal/redis"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
}

func LiveHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Alive")
}

func ReadyHandler(w http.ResponseWriter, r *http.Request) {
	err := redis.Client.Ping(redis.Ctx).Err()
	if err != nil {
		http.Error(w, "Redis Not Ready", http.StatusServiceUnavailable)
		return
	}
	fmt.Fprintf(w, "Ready")
}
