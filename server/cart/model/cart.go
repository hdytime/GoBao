package model

import "time"

type CartProduct struct {
	Id         int64     `json:"id"`         //购物车id
	UserId     int64     `json:"userId"`     //用户id
	ProductId  int64     `json:"productId"`  //商品id
	Price      float64   `json:"price"`      //商品价格
	Quantity   int64     `json:"quantity"`   //商品数量
	CreateTime time.Time `json:"createTime"` //创建时间
	UpdateTime time.Time `json:"updateTime"` //更新时间
}
