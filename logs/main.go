package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/rs/cors"
	"golang.org/x/net/context"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	corsOptions := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Origin", "Content-Type", "Accept"},
		Debug:          true,
	})

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
	})

	mux.HandleFunc("/events", func(w http.ResponseWriter, r *http.Request) {
		queueName := r.URL.Query().Get("queue")
		if queueName == "" {
			http.Error(w, "Missing queue parameter", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		fmt.Fprintf(w, "data: Logs in Action.\n\n")
		w.(http.Flusher).Flush()

		for {
			logEntry, err := rdb.RPop(ctx, queueName).Result()
			if logEntry != "" {
				fmt.Println(logEntry)
			}
			if err == redis.Nil {
				continue
			} else if err != nil {
				log.Printf("Error fetching log from Redis: %v", err)
				break
			}
			fmt.Fprintf(w, "data: %s\n\n", logEntry)
			w.(http.Flusher).Flush()
		}
	})

	handler := corsOptions.Handler(mux)

	fmt.Println("Listening on PORT: 1234")
	log.Fatal(http.ListenAndServe(":1234", handler))
}
