package cmd

import (
	"fmt"
	"time"

	"github.com/rafaelsouzaribeiro/logs/internal/infra/database/mongo/connection"
	"github.com/rafaelsouzaribeiro/logs/internal/infra/di"
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
	insert.Save(topic, message, createdAt)
}
