package repository

import (
	"chatGo/src/domain/message"
	"gorm.io/gorm"
	"time"
)

type MessageModel struct {
	ID        string
	CreatedAt time.Time
	From      string
	Kind      string
	Message   string
}

type GormDB struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) GormDB {
	return GormDB{db: db}
}

func (g GormDB) Create(message message.Message) {
	payload := MessageModel{
		ID:      message.ID,
		From:    message.From,
		Kind:    message.Kind,
		Message: message.Message,
	}
	g.db.Create(&payload)
}

func (g GormDB) GetByID(id string) message.Message {
	result := &MessageModel{}
	g.db.First(result, id)
	return message.Message{
		ID:      result.ID,
		From:    result.From,
		Kind:    result.Kind,
		Message: result.Message,
	}
}

func (g GormDB) FindWithLimit(limit int) []message.Message {
	var models []MessageModel
	var entities []message.Message

	g.db.Limit(limit).Find(&models).Order("created_at DESC")

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
