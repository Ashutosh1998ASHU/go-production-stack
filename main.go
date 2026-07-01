package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

var rdb = redis.NewClient(&redis.Options{
	Addr: "redis:6379",
})

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("Request received:", r.Method, r.URL.Path)
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	count, err := rdb.Incr(ctx, "visits").Result()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Hello from Dockerized GO server + Redis!\n\nVisitor Count: %d\n", count)
}

func main() {
	http.HandleFunc("/", homeHandler)
	fmt.Println("Starting server on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
