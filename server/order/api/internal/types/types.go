// Code generated by goctl. DO NOT EDIT.
package types

type CreateSeckillOrderReq struct {
	ProductID    int64  `json:"productId"`
	ProductCount int64  `json:"productCount"`
	Remark       string `json:"remark"`
}

type CreateSeckillOrderResp struct {
	OrderSn string `json:"orderSn"`
}

type DeleteOrderReq struct {
	OrderSn string `json:"orderSn"`
}

type GetOrderDetailReq struct {
	UserID  int64  `json:"userId"`
	OrderSn string `json:"orderSn"`
}

type GetOrderDetailResp struct {
	ID           int64   `json:"id"`
	CreateTime   string  `json:"createTime"`
	UpdateTime   string  `json:"updateTime"`
	OrderSn      string  `json:"orderSn"`
	UserID       int64   `json:"userId"`
	ProductID    int64   `json:"productId"`
	Name         string  `json:"name"`
	ProductCount int64   `json:"productCount"`
	UnitPrice    float64 `json:"unitPrice"`
	TotalPrice   float64 `json:"totalPrice"`
	Status       int64   `json:"status"`
	Remark       string  `json:"remark"`
	PayTime      string  `json:"payTime"`
}

type GetOrderListReq struct {
	Page int64 `json:"page"`
	Size int64 `json:"size"`
}

type GetOrderListResp struct {
	OrderList []SmallOrder `json:"orderList"`
}

type Order struct {
	Id          int64   `json:"id"`
	UserId      int64   `json:"userId"`
	ProductId   int64   `json:"productId"`
	ProductName string  `json:"productName"`
	UnitPrice   float64 `json:"unitPrice"`
	Quantity    int64   `json:"quantity"`
	TotalPrice  float64 `json:"totalPrice"`
	Status      int64   `json:"status"`
	PayTime     int64   `json:"payTime"`
	CreateTime  int64   `json:"createTime"`
	UpdateTime  int64   `json:"updateTime"`
	Remark      string  `json:"remark"`
}

type SmallOrder struct {
	OrderSn    string  `json:"orderSn"`
	Name       string  `json:"name"`
	ProductID  int64   `json:"productId"`
	TotalPrice float64 `json:"totalPrice"`
	Status     string  `json:"status"`
}