package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"


	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Starting Peril server...")

	// 1. Declare the RabbitMQ connection string
	const rabbitConnString = "amqp://guest:guest@localhost:5672/"

	// 2. Establish a connection to the RabbitMQ server
	conn, err := amqp.Dial(rabbitConnString)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)

	}

	// 3. Defer closing the connection
	defer func() {
		fmt.Println("Closing RabbitMQ connection...")
		conn.Close()
	 	}()

	// 4. Print success message
	fmt.Println("Successfully connected to RabbitMQ!")

	// 5. Block and wait for a termination signal (Ctrl+C)
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	fmt.Println("Server running. Press Ctrl+C to stop...")
	<-signalChan

	// 6. Print shutdown message
	fmt.Println("Shutting down gracefully...")
}
