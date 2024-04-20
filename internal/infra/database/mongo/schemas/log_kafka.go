package schemas

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
