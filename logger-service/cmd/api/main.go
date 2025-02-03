/* Use environment variables from env.go
   Connect to MongoDB
   Connect to RabbitMQ
   Listen for log messages & store them in MongoDB
   Run an HTTP server
*/

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Config struct holds database and RabbitMQ connections
type Config struct {
	DB *mongo.Client
}

func main() {
	// Print environment variables from env.go
	PrintEnvVariables()

	// Connect to MongoDB
	db, err := connectToDB()
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	defer db.Disconnect(context.TODO())

	// Connect to RabbitMQ
	rabbitConn, rabbitChannel, err := connectToRabbitMQ()
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
	}
	defer rabbitConn.Close()
	defer rabbitChannel.Close()

	// Create Config instance
	app := Config{DB: db}

	// Start RabbitMQ consumer to listen for logs
	go app.consumeLogs(rabbitChannel)

	// Start HTTP server
	fmt.Printf("🚀 %s is running on port: %s\n", ServiceName, ServicePort)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", ServicePort),
		Handler: app.routes(),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Panic("❌ Server crashed:", err)
	}
}

// connectToDB initializes a connection to MongoDB
func connectToDB() (*mongo.Client, error) {
	// Build MongoDB URI using environment variables and specify the authentication database
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=admin",
		MongoUser, MongoPass, MongoHost, MongoPort, MongoDBName)

	// Set client options and apply the URI
	clientOptions := options.Client().ApplyURI(uri)

	// Connect to the MongoDB client
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	// Ping MongoDB to check connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	// Print a success message if connected
	fmt.Println("✅ Connected to MongoDB")
	return client, nil
}

// connectToRabbitMQ initializes a RabbitMQ connection
func connectToRabbitMQ() (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial(fmt.Sprintf(
		"amqp://%s:%s@%s:%s/",
		RabbitUser,
		RabbitPass,
		RabbitHost,
		RabbitPort,
	))
	if err != nil {
		return nil, nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, err
	}

	_, err = ch.QueueDeclare(
		"log_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, nil, err
	}

	fmt.Println("✅ Connected to RabbitMQ")
	return conn, ch, nil
}

// consumeLogs listens for messages from RabbitMQ and stores them in MongoDB
func (app *Config) consumeLogs(ch *amqp.Channel) {
	msgs, err := ch.Consume(
		"log_queue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("❌ Failed to consume messages from RabbitMQ:", err)
	}

	fmt.Println("🎧 Listening for logs on RabbitMQ...")

	for msg := range msgs {
		var logEntry LogEntry
		json.Unmarshal(msg.Body, &logEntry)

		logEntry.ID = primitive.NewObjectID()
		logEntry.CreatedAt = time.Now()

		collection := app.DB.Database("logDB").Collection("logs")
		_, err := collection.InsertOne(context.TODO(), logEntry)
		if err != nil {
			log.Println("❌ Failed to insert log into MongoDB:", err)
		}

		fmt.Println("✅ Log saved:", logEntry)
	}
}
