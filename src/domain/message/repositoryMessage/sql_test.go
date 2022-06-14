package repositoryMessage

import (
	"chatGo/src/domain/message"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestDatabase_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	r := NewMockRepository(ctrl)

	messageModel := MessageModel{
		ID:        "uuid-id",
		CreatedAt: time.Time{},
		From:      "fromUser",
		Kind:      "typeOfMessage",
		Message:   "this is a message",
	}

	entity := message.Message{
		ID:      "uuid-id",
		From:    "fromUser",
		Kind:    "typeOfMessage",
		Message: "this is a message",
	}
	r.EXPECT().Create(&messageModel)

	repositoryMessage := NewRepository(r)
	repositoryMessage.Create(entity)
}

func TestDatabase_FindWithLimit(t *testing.T) {
	assertTestify := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	r := NewMockRepository(ctrl)

	var models []MessageModel

	r.EXPECT().GetWithLimit(&models, 50).Return(&[]MessageModel{
		{
			ID:        "uuid-id1",
			CreatedAt: time.Time{},
			From:      "fromUser",
			Kind:      "typeOfMessage",
			Message:   "this is a message",
		},
		{
			ID:        "uuid-id2",
			CreatedAt: time.Time{},
			From:      "fromUser",
			Kind:      "typeOfMessage",
			Message:   "this is a message",
		},
		{
			ID:        "uuid-id3",
			CreatedAt: time.Time{},
			From:      "fromUser",
			Kind:      "typeOfMessage",
			Message:   "this is a message",
		},
	})

	expectedResult := []message.Message{
		{
			ID:           "uuid-id1",
			RegisterDate: time.Time{},
			From:         "fromUser",
			Kind:         "typeOfMessage",
			Message:      "this is a message",
		},
		{
			ID:           "uuid-id2",
			RegisterDate: time.Time{},
			From:         "fromUser",
			Kind:         "typeOfMessage",
			Message:      "this is a message",
		},
		{
			ID:           "uuid-id3",
			RegisterDate: time.Time{},
			From:         "fromUser",
			Kind:         "typeOfMessage",
			Message:      "this is a message",
		},
	}

	repositoryMessage := NewRepository(r)
	result := repositoryMessage.FindWithLimit(50)

	assertTestify.Equal(expectedResult, result)
}
