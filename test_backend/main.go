package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

// тестовый сервер обработки запросов
func main() {
	port := flag.String("port", "9001", "порт для запуска сервера")
	delay := flag.Duration("sleep", 0, "задержка времени выполнения, можно симулировать раницу в быстродействии серверов")
	flag.Parse()

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("statusOK"))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/health" {
			if *delay > 0 {
				time.Sleep(*delay)
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("это простой http сервер"))
		}
	})

	addr := ":" + *port
	fmt.Printf("Starting server on %s with sleep %v\n", addr, *delay)
	if err := http.ListenAndServe(addr, nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
