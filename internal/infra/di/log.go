package di

import (
	database "github.com/rafaelsouzaribeiro/logs/internal/infra/database/mongo/repository"
	"github.com/rafaelsouzaribeiro/logs/internal/infra/enum"
	"github.com/rafaelsouzaribeiro/logs/internal/usecase"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewLogUseCase(db *mongo.Database) usecase.LogsUseCase {

	collection := db.Collection(string(enum.Log_Kafka))
	repo := database.NewLogRepository(collection)

	first := usecase.NewLogsUseCase(repo)

	return *first
}
