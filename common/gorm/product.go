package gorm

import (
	"GoBao/server/product/model"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

const ProductDSN = "root:123456@tcp(127.0.0.1:3306)/gobao_product?charset=utf8mb4&parseTime=True&loc=Local"

var ProductDB *gorm.DB

func init() {
	newLogger := logger.New(log.New(os.Stdout, "gobao   ", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			Colorful:                  true,
			IgnoreRecordNotFoundError: true,
			LogLevel:                  logger.Error,
		})
	db, err := gorm.Open(mysql.Open(ProductDSN),
		&gorm.Config{
			Logger: newLogger,
		})
	if err != nil {
		logx.WithContext(context.Background()).Errorf("GORM connect ProductDB Error: %+v", err)
	}

	err = db.AutoMigrate(&model.Product{}, &model.ProductRecommend{}, &model.SeckillProduct{})
	if err != nil {
		logx.WithContext(context.Background()).Errorf("GORM AutoMigrate product ERROR:%+v", err)
	}
	ProductDB = db
}
