package model

import "time"

type Pay struct {
	Id             int64     `json:"id"`             //支付id
	UserId         int64     `json:"userId"`         //用户id
	PaySn          string    `json:"paySn"`          //支付Sn
	OrderSn        string    `json:"orderId"`        //订单Sn
	TradeState     int64     `json:"status"`         //支付状态
	Paytotal       float64   `json:"paytotal"`       //支付总额
	TransactionID  int64     `json:"transactionID"`  //交易id
	TradeStateDesc string    `json:"tradeStateDesc"` //交易状态描述
	PayTime        time.Time `json:"payTime"`        //支付时间
	CreateTime     time.Time `json:"createTime"`     //创建时间
	UpdateTime     time.Time `json:"updateTime"`     //更新时间
}
