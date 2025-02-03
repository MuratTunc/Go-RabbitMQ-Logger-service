package main

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LogEntry struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Service   string             `json:"service_name" bson:"service_name"`
	Message   string             `json:"log_message" bson:"log_message"`
	Timestamp time.Time          `json:"timestamp,omitempty" bson:"timestamp,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
}
