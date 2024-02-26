package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/clwg/eve-analyzer/internal/consumer"
	"github.com/clwg/eve-analyzer/pkg/database"
	"github.com/clwg/eve-analyzer/pkg/filemonitor"
	"github.com/clwg/eve-analyzer/pkg/model"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	// Check if the command line argument for the file path is provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: <program> <path to eve.json>")
		return
	}

	// Get the file path from the command line arguments
	filePath := os.Args[1]

	database.SetDatasourceName(fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSLMODE"),
	))

	database.PostgresEventLogger()

	dataChan := make(chan model.Event)

	// Start the file monitor with the provided file path in a goroutine
	go filemonitor.OpenFile(filePath, dataChan)

	fmt.Println("Monitoring started on", filePath)

	// Setup a channel to listen for interrupt signals to ensure graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// Start the consumer in its own goroutine and wait for it to finish processing
	doneChan := make(chan bool)
	go func() {
		consumer.ConsumeData(dataChan)
		doneChan <- true
	}()

	// Wait for either a shutdown signal or for processing to complete
	select {
	case <-sigChan:
		fmt.Println("Shutdown signal received, exiting...")
	case <-doneChan:
		fmt.Println("Processing completed, exiting...")
	}

}
