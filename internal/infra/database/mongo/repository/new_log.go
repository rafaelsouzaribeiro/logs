package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type LogsInterfaceRepository struct {
	Colletion *mongo.Collection
}

func NewLogRepository(c *mongo.Collection) *LogsInterfaceRepository {
	return &LogsInterfaceRepository{
		Colletion: c,
	}
}
