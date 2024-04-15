package cmd

import (
	"fmt"
	"time"

	"github.com/rafaelsouzaribeiro/logs/internal/infra/database/connection"
	"github.com/rafaelsouzaribeiro/logs/internal/infra/di"
	"github.com/rafaelsouzaribeiro/logs/internal/usecase/dto"
	"github.com/spf13/viper"
)

func Insert(topic, message string, createdAt time.Time) {
	//DB_HOST := os.Getenv("DB_HOST")
	DB_HOST := viper.GetString("DB_HOST")

	if DB_HOST == "" {
		fmt.Println("Error: DB_HOST environment variable is not set.")
		return
	}

	db, _err := connection.GetMongoDataBase("logs", DB_HOST)

	if _err != nil {
		panic(_err)
	}

	insert := di.NewLogUseCase(db)
	insert.Save(dto.Log_KafkaInputDTO{
		Topic:     topic,
		Message:   message,
		CreatedAt: createdAt,
	})
}
