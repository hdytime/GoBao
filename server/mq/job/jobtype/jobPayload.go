package jobtype

type DeferCloseSeckillOrderPayload struct {
	UserID       int64
	OrderSn      string
	RedisKey     string
	RedisField   string
	SeckillCount int64
}
