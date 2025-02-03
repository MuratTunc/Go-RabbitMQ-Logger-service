package main

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LogEntry struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Service   string             `bson:"service_name"`
	Message   string             `bson:"log_message"`
	Timestamp time.Time          `bson:"timestamp"`
	CreatedAt time.Time          `bson:"created_at"`
}
