package model

import "time"

type Product struct {
	Id            int64     `json:"id"`            // 商品id
	Name          string    `json:"name"`          // 商品名称
	Price         float64   `json:"price"`         // 商品价格
	DiscountPrice float64   `json:"discountPrice"` // 秒杀价
	Stock         int64     `json:"stock"`         // 库存数量
	Status        int64     `json:"status"`        // 商品状态
	CreateTime    time.Time `json:"createTime"`    // 创建时间
	UpdateTime    time.Time `json:"updateTime"`    // 更新时间
}

type SeckillProduct struct {
	Id           int64     `json:"id"`           // 商品id
	Name         string    `json:"name"`         // 商品名称
	Price        float64   `json:"price"`        // 商品价格
	Stock        int64     `json:"stock"`        // 库存数量
	Status       int64     `json:"status"`       // 商品状态
	CreateTime   time.Time `json:"createTime"`   // 创建时间
	UpdateTime   time.Time `json:"updateTime"`   // 更新时间
	SeckillPrice float64   `json:"seckillPrice"` // 秒杀价
	StockCount   int64     `json:"stockCount"`   // 秒杀库存
	StartTime    string    `json:"startTime"`    // 开始时间
	Time         int64     `json:"time"`         // 持续时间
}

type ProductRecommend struct {
	ID         int64     `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime;not null;index:ix_update_time,priority:1;default:CURRENT_TIMESTAMP" json:"update_time"`
	DeleteTime time.Time `gorm:"column:delete_time;type:datetime;not null;default:CURRENT_TIMESTAMP" json:"delete_time"`
	DelState   int64     `gorm:"column:del_state;type:tinyint;not null" json:"del_state"`
	ProductID  int64     `gorm:"column:product_id;type:bigint unsigned;not null;comment:商品id" json:"product_id"`
}
