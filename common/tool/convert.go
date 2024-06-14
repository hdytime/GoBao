package tool

import (
	"strconv"
	"time"
)

func ConvertStringToInt64(s string) int64 {
	res, _ := strconv.ParseInt(s, 10, 64)
	return res
}

func ConvertStringToFloat64(s string) float64 {
	res, _ := strconv.ParseFloat(s, 64)
	return res
}

func ConvertStringToTimeUnix(s string) int64 {
	t, _ := time.Parse(time.RFC3339, s)
	return t.Unix()
}
