package gorm

import (
	"GoBao/server/cart/model"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

const CartDSN = "root:123456@tcp(127.0.0.1:3306)/gobao_cart?charset=utf8mb4&parseTime=True&loc=Local"

var CartDB *gorm.DB

func init() {
	newLogger := logger.New(log.New(os.Stdout, "gobao   ", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			Colorful:                  true,
			IgnoreRecordNotFoundError: true,
			LogLevel:                  logger.Error,
		})
	db, err := gorm.Open(mysql.Open(CartDSN),
		&gorm.Config{
			Logger: newLogger,
		})
	if err != nil {
		logx.WithContext(context.Background()).Errorf("GORM connect CartDB Error: %+v", err)
	}
	err = db.AutoMigrate(model.CartProduct{})
	if err != nil {
		logx.WithContext(context.Background()).Errorf("GORM AutoMigrate Cart ERROR:%+v", err)
	}

	CartDB = db
}
