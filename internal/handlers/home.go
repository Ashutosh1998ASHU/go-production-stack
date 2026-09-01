package handlers

import (
	"fmt"
	"go-production-stack/internal/redis"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// if r.URL.Query().Get("panic") == "true" {
	// 	panic("intentional panic")
	// }

	count, err := redis.Client.Incr(redis.Ctx, "visits").Result()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Hello from Dockerized GO server + Redis!\n\nVisitor Count: %d\n", count)
}
