package sql

import (
	"chatGo/src/domain/message/repository"
	"chatGo/src/infrastructure/settings"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Start(config *settings.GlobalConfig) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/chat?charset=utf8&parseTime=True&loc=Local", config.DbUsername, config.DbPassword, config.DbHost)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(repository.MessageModel{})
	if err != nil {
		panic("Couldn't create migration")
	}

	return db
}
