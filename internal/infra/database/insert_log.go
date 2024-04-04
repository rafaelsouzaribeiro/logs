package database

import (
	"context"

	"github.com/rafaelsouzaribeiro/logs/internal/entity"
)

func (s *LogsInterfaceRepository) Insert(data entity.Log_Kafka) (interface{}, error) {

	res, err := s.Colletion.InsertOne(context.TODO(), data)

	if err != nil {
		return nil, err
	}

	return res.InsertedID, nil

}
