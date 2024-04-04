package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Log_Kafka struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Topic     string             `bson:"topic,omitempty"`
	Message   string             `bson:"message,omitempty"`
	CreatedAt time.Time          `bson:"createdAt,omitempty"`
}

func NewLogKafka(topic, message string, createdAt time.Time) *Log_Kafka {

	return &Log_Kafka{
		Topic:     topic,
		Message:   message,
		CreatedAt: createdAt,
	}
}
