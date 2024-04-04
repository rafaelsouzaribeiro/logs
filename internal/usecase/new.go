package usecase

import "github.com/rafaelsouzaribeiro/logs/internal/entity"

type LogsUseCase struct {
	repository entity.LogsInterfaceRepository
}

func NewLogsUseCase(e entity.LogsInterfaceRepository) *LogsUseCase {
	return &LogsUseCase{
		repository: e,
	}
}
