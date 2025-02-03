package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// Import your models package if necessary (example: "yourapp/models")
)

// LogHandler receives logs from HTTP requests and stores them in MongoDB
func (app *Config) LogHandler(w http.ResponseWriter, r *http.Request) {
	var logEntry LogEntry // Use the LogEntry struct from models.go
	err := json.NewDecoder(r.Body).Decode(&logEntry)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Add timestamp
	logEntry.ID = primitive.NewObjectID()
	logEntry.CreatedAt = time.Now()

	// Insert log into MongoDB
	collection := app.DB.Database("logDB").Collection("logs")
	_, err = collection.InsertOne(context.TODO(), logEntry)
	if err != nil {
		http.Error(w, "Failed to insert log into MongoDB", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Log saved successfully")
}

// GetLogsHandler reads logs from MongoDB for a given service name
func (app *Config) GetLogsHandler(w http.ResponseWriter, r *http.Request) {
	// Extract service name from the URL query parameters
	serviceName := r.URL.Query().Get("service_name")
	if serviceName == "" {
		http.Error(w, "Missing service name", http.StatusBadRequest)
		return
	}

	// Find logs from MongoDB based on the service name
	collection := app.DB.Database("logDB").Collection("logs")
	filter := bson.M{"service_name": serviceName}

	cursor, err := collection.Find(context.TODO(), filter)
	if err != nil {
		http.Error(w, "Failed to fetch logs", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	var logs []LogEntry
	for cursor.Next(context.TODO()) {
		var log LogEntry
		if err := cursor.Decode(&log); err != nil {
			http.Error(w, "Error decoding log entry", http.StatusInternalServerError)
			return
		}
		logs = append(logs, log)
	}

	// Check for cursor iteration errors
	if err := cursor.Err(); err != nil {
		http.Error(w, "Cursor error", http.StatusInternalServerError)
		return
	}

	// Send logs as JSON response
	w.Header().Set("Content-Type", "application/json")
	if len(logs) == 0 {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "No logs found for the service")
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(logs)
	if err != nil {
		http.Error(w, "Error encoding logs", http.StatusInternalServerError)
	}
}
