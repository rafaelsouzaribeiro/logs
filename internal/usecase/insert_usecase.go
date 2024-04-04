package usecase

import (
	"time"

	"github.com/rafaelsouzaribeiro/logs/internal/entity"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Log_KafkaInputDTO struct {
	ID        primitive.ObjectID `json:"_id"`
	Topic     string             `json:"topic"`
	Message   string             `json:"message"`
	CreatedAt time.Time          `json:"createdAt"`
}

func NewOrderUseCase(repository entity.LogsInterfaceRepository) *LogsUseCase {
	return &LogsUseCase{
		repository: repository,
	}
}

func (o *LogsUseCase) Save(data Log_KafkaInputDTO) error {
	rs := entity.NewLogKafka(data.Topic, data.Message, data.CreatedAt)

	_, err := o.repository.Insert(*rs)

	if err != nil {
		return err
	}

	return nil
}
