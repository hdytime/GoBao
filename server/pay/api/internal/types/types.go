// Code generated by goctl. DO NOT EDIT.
package types

type OrderPayReq struct {
	OrderSn string `json:"orderSn"`
}

type OrderPayResp struct {
	PayTotalPrice float64 `json:"payTotalPrice"`
	Paysn         string  `json:"paysn"`
}
