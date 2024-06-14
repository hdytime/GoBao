package model

import "time"

type Order struct {
	Id          int64     `json:"id"`          //订单id
	UserId      int64     `json:"userId"`      //用户id
	ProductId   int64     `json:"productId"`   //商品id
	ProductName string    `json:"productName"` //商品名称
	OrderSn     string    `json:"orderSn"`     //商品标识号
	UnitPrice   float64   `json:"unitPrice"`   //单价
	Quantity    int64     `json:"quantity"`    //数量
	TotalPrice  float64   `json:"totalPrice"`  //实付金额
	Status      int64     `json:"status"`      //状态
	PayTime     time.Time `json:"payTime"`     //支付时间
	CreateTime  time.Time `json:"createTime"`  //订单创建时间
	UpdateTime  time.Time `json:"updateTime"`  //订单更新时间
	Remark      string    `json:"remark"`      //备注
}
