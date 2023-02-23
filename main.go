package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	shutdown := make(chan os.Signal, 1)

	signal.Notify(shutdown, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		log.Println("server started at http://localhost:8080/metrics")
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()
	<-shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Printf("server shutdown with error: %s\n", err)
		return
	}
	log.Println("server gracefully stoped")
}
