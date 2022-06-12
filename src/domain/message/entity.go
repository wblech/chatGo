package message

import (
	"github.com/google/uuid"
	"time"
)

type Message struct {
	ID           string
	From         string
	Kind         string
	Message      string
	RegisterDate time.Time
}

func NewMessage(from string, kind string, message string) Message {
	id := uuid.New()
	return Message{
		ID:      id.String(),
		From:    from,
		Kind:    kind,
		Message: message,
	}
}
