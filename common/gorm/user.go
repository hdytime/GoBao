package gorm

import (
	"GoBao/server/user/model"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

const UserDSN = "root:123456@tcp(127.0.0.1:3306)/gobao_user?charset=utf8mb4&parseTime=True&loc=Local"

var UserDB *gorm.DB

func init() {
	newLogger := logger.New(log.New(os.Stdout, "gobao   ", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			Colorful:                  true,
			IgnoreRecordNotFoundError: true,
			LogLevel:                  logger.Error,
		})
	db, err := gorm.Open(mysql.Open(UserDSN),
		&gorm.Config{
			Logger: newLogger,
		})
	if err != nil {
		logx.WithContext(context.Background()).Errorf("GORM connect UserDB Error: %+v", err)
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		logx.WithContext(context.Background()).Errorf("GORM AutoMigrate user ERROR:%+v", err)
	}
	UserDB = db
}
