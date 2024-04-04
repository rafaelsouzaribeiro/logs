package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Log_KafkaOuput struct {
	ID        primitive.ObjectID `json:"_id"`
	Topic     string             `json:"topic"`
	Message   string             `json:"message"`
	CreatedAt time.Time          `json:"createdAt"`
}
