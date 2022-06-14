package repositoryMessage

import (
	"chatGo/src/domain/message"
	"time"
)

type (
	Creator interface {
		Create(message *MessageModel)
	}

	Getter interface {
		GetWithLimit(models *[]MessageModel, limit int)
	}

	Repository interface {
		Creator
		Getter
	}

	Database struct {
		Repository
	}

	MessageModel struct {
		ID        string
		CreatedAt time.Time
		From      string
		Kind      string
		Message   string
	}
)

func NewRepository(db Repository) Database {
	return Database{
		Repository: db,
	}
}

func (d Database) Create(message message.Message) {
	payload := MessageModel{
		ID:      message.ID,
		From:    message.From,
		Kind:    message.Kind,
		Message: message.Message,
	}
	d.Repository.Create(&payload)
}

func (d Database) FindWithLimit(limit int) []message.Message {
	var models []MessageModel
	var entities []message.Message

	//d.db.Limit(limit).Find(&models).Order("created_at DESC")
	d.GetWithLimit(&models, limit)
	for _, model := range models {
		entity := message.Message{
			ID:           model.ID,
			From:         model.From,
			Kind:         model.Kind,
			Message:      model.Message,
			RegisterDate: model.CreatedAt,
		}
		entities = append(entities, entity)
	}
	return entities
}
