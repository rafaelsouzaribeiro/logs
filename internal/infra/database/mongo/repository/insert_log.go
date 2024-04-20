package repository

import (
	"context"

	"github.com/rafaelsouzaribeiro/logs/internal/entity"
	schema "github.com/rafaelsouzaribeiro/logs/internal/infra/database/mongo/schemas"
)

func (s *LogsInterfaceRepository) Insert(data entity.Log_Kafka) (interface{}, error) {

	insert := schema.Log_Kafka{
		Topic:     data.Topic,
		Message:   data.Message,
		CreatedAt: data.CreatedAt,
	}

	res, err := s.Colletion.InsertOne(context.TODO(), insert)

	if err != nil {
		return nil, err
	}

	return res.InsertedID, nil

}
