package create

import (
	"chatGo/src/domain/message"
)

type (
	Creator interface {
		Create(message message.Message)
	}

	Service struct {
		Creator Creator
	}
)

func NewService(creator Creator) Service {
	return Service{
		Creator: creator,
	}
}

func (s Service) Execute(message message.Message) {
	s.Creator.Create(message)
}
