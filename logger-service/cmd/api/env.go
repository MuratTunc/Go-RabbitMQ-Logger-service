package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from a .env file
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌ Error loading .env file")
	}
}

// Environment variables
var (
	// Logger Service Config
	ServicePort = os.Getenv("LOGGER_SERVICE_PORT")
	ServiceName = os.Getenv("LOGGER_SERVICE_NAME")

	// MongoDB Config
	MongoUser   = os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	MongoPass   = os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	MongoHost   = os.Getenv("MONGO_HOST")
	MongoPort   = os.Getenv("MONGO_PORT")
	MongoDBName = os.Getenv("MONGO_DB_NAME")

	// RabbitMQ Config
	RabbitUser  = os.Getenv("RABBITMQ_USER")
	RabbitPass  = os.Getenv("RABBITMQ_PASS")
	RabbitHost  = os.Getenv("RABBITMQ_HOST")
	RabbitPort  = os.Getenv("RABBITMQ_PORT")
	RabbitQueue = os.Getenv("RABBITMQ_QUEUE")
)

// PrintEnvVariables prints all environment variables for debugging
func PrintEnvVariables() {
	fmt.Println("🔧 Loaded Environment Variables:")
	fmt.Printf("Service Name: %s\n", ServiceName)
	fmt.Printf("Service Port: %s\n", ServicePort)
	fmt.Println("📦 MongoDB Config:")
	fmt.Printf("  - Host: %s\n", MongoHost)
	fmt.Printf("  - Port: %s\n", MongoPort)
	fmt.Printf("  - Database: %s\n", MongoDBName)
	fmt.Printf("  - Username: %s\n", MongoUser)
	fmt.Printf("  - Password: %s\n", MongoPass)
	fmt.Println("📡 RabbitMQ Config:")
	fmt.Printf("  - Host: %s\n", RabbitHost)
	fmt.Printf("  - Port: %s\n", RabbitPort)
	fmt.Printf("  - Queue: %s\n", RabbitQueue)
	fmt.Printf("  - Username: %s\n", RabbitUser)
	fmt.Printf("  - Password: %s\n", RabbitPass)
}
