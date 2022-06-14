package sql

import (
	"chatGo/src/domain/message/repositoryMessage"
	"chatGo/src/infrastructure/settings"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GormImpl struct {
	Db *gorm.DB
}

func Start(config *settings.GlobalConfig) *GormImpl {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/chat?charset=utf8&parseTime=True&loc=Local", config.DbUsername, config.DbPassword, config.DbHost)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(repositoryMessage.MessageModel{})
	if err != nil {
		panic("Couldn't create migration")
	}

	return &GormImpl{
		Db: db,
	}
}

func (g *GormImpl) Create(message *repositoryMessage.MessageModel) {
	g.Db.Create(message)
}

func (g *GormImpl) GetWithLimit(message *[]repositoryMessage.MessageModel, limit int) *[]repositoryMessage.MessageModel {
	g.Db.Limit(limit).Find(&message).Order("created_at DESC")
	return message
}
