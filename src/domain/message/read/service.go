package read

import (
	"chatGo/src/domain/message"
	"chatGo/src/domain/message/repositoryMessage"
)

type (
	Finder interface {
		FindWithLimit(limit int) []message.Message
	}

	Service struct {
		Db *repositoryMessage.Database
	}
)

func NewService(db *repositoryMessage.Database) Service {
	return Service{
		Db: db,
	}
}

func (s Service) Execute(limit int) []message.Message {
	return s.Db.FindWithLimit(limit)
}
