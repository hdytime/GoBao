// Code generated by goctl. DO NOT EDIT.
package types

type AddProductToCartReq struct {
	ProductID int64 `json:"productId"`
	Count     int64 `json:"count"`
}

type CartProduct struct {
	Id            int64   `json:"id"`
	UserId        int64   `json:"userId"`
	ProductId     int64   `json:"productId"`
	Price         float64 `json:"price"`
	DiscountPrice float64 `json:"discountPrice"`
	Quantity      int64   `json:"quantity"`
	CreateTime    int64   `json:"createTime"`
	UpdateTime    int64   `json:"updateTime"`
}

type DeleteProductFromCartReq struct {
	CartID int64 `json:"productId"`
}

type UpdateCartProductDetailReq struct {
	CartID int64 `json:"cartId"`
	Count  int64 `json:"count"`
}

type GetCartListResp struct {
	CartProducts []CartProduct `json:"cartProducts"`
}
