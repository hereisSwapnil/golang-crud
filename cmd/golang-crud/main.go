package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/hereisSwapnil/golang-crud/internal/config"
	"github.com/hereisSwapnil/golang-crud/internal/http/handlers/student"
	"github.com/hereisSwapnil/golang-crud/internal/storage/sqlite"
)

func main() {
	// load config
	config := config.LoadConfig()
	fmt.Println("Config loaded successfully ✅")

	// database setup
	storage, err := sqlite.New(config)
	if err != nil {
		log.Fatalf("failed to create storage: %v", err)
	}
	fmt.Println("Storage created successfully ✅")
	
	// server setup
	router := http.NewServeMux()
	
	router.Handle("POST /api/v1/student", student.New(storage))

	server := &http.Server{
		Addr:    config.HttpServer.Address,
		Handler: router,
	}

	// making a done channel
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)



	// start server
	// running this in a separate go routine
	fmt.Println("Server is starting... ✅")
	fmt.Printf("Server is running on %s\n", config.HttpServer.Address)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Fatalf("failed to start server: %v", err)
		}
	}()

	<-done
	fmt.Println("Server is shutting down... ✅")
	
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("failed to shutdown server: %v", err)
	}

	fmt.Println("Server is shutdown successfully ✅")
}