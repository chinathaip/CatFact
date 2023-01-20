package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/chinathaip/catfact/service"
)

func main() {
	catFactService := service.NewCatFactService("https://catfact.ninja/fact")
	loggingService := service.NewLoggingService(catFactService)

	router := NewRouter(loggingService)

	e := router.RegisterRouter()

	srv := &http.Server{
		Addr:    ":8080",
		Handler: e,
	}
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	<-signals
	log.Println("Shutting down server...")
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("Could not gracefully shutdown the server: %v", err)
	}
	log.Println("Server stopped")
}
