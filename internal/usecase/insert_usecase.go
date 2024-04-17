package usecase

import (
	"time"

	"github.com/rafaelsouzaribeiro/logs/internal/entity"
)

func NewOrderUseCase(repository entity.LogsInterfaceRepository) *LogsUseCase {
	return &LogsUseCase{
		repository: repository,
	}
}

func (o *LogsUseCase) Save(Topic string, Message string, CreatedAt time.Time) error {
	rs := entity.NewLogKafka(Topic, Message, CreatedAt)

	_, err := o.repository.Insert(*rs)

	if err != nil {
		return err
	}

	return nil
}
