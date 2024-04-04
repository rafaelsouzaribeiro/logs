package entity

type LogsInterfaceRepository interface {
	Insert(data Log_Kafka) (interface{}, error)
}
