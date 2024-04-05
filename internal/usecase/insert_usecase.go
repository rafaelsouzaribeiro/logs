package usecase

import (
	"github.com/rafaelsouzaribeiro/logs/internal/entity"
	"github.com/rafaelsouzaribeiro/logs/internal/usecase/dto"
)

func NewOrderUseCase(repository entity.LogsInterfaceRepository) *LogsUseCase {
	return &LogsUseCase{
		repository: repository,
	}
}

func (o *LogsUseCase) Save(data dto.Log_KafkaInputDTO) error {
	rs := entity.NewLogKafka(data.Topic, data.Message, data.CreatedAt)

	_, err := o.repository.Insert(*rs)

	if err != nil {
		return err
	}

	return nil
}
