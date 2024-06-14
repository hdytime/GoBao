package model

type User struct {
	ID          int64   `json:"id"`          //用户id
	Username    string  `json:"username"`    //用户名
	Password    string  `json:"password"`    //密码
	Money       float64 `json:"money"`       //余额
	Sex         int64   `json:"sex"`         //性别
	PhoneNumber int64   `json:"phoneNumber"` //手机号
	Email       string  `json:"email"`       //邮箱
	Sign        string  `json:"sign"`        //签名
}
