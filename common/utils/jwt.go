package utils

import (
	"GoBao/common/ctxData"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var AccessSecret = []byte("gobaoAccessSecret")
var RefreshSecret = []byte("gobaoRefreshSecret")
var AccessExpireDuration = 30 * time.Minute
var RefreshExpireDuration = 24 * time.Hour

func GenerateJwtToken(secretKey string, iat, seconds, userID int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims[ctxData.JwtKeyUserID] = userID
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

//type Claims struct {
//	UserID int64 `json:"user_id"`
//	jwt.StandardClaims
//}
//
//func GenerateToken(userID int64, secret string, expireDuration time.Duration) (string, error) {
//	now := time.Now()
//	expireTime := now.Add(expireDuration)
//	claims := &Claims{
//		UserID: userID,
//		StandardClaims: jwt.StandardClaims{
//			ExpiresAt: expireTime.Unix(),
//			IssuedAt:  now.Unix(),
//		},
//	}
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//	return token.SignedString([]byte(secret))
//}
//
//func ParseAccessToken(tokenStr, AccessSecret string) (*Claims, bool, error) {
//	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
//		return []byte(AccessSecret), nil
//	})
//	if err != nil {
//		return nil, true, err
//	}
//	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
//		return claims, false, nil
//	}
//	return nil, true, nil
//}
//
//func ParseRefreshToken(tokenStr, RefreshSecret string) (*Claims, bool, error) {
//	token, err := jwt.ParseWithClaims(tokenStr, &Claims{}, func(token *jwt.Token) (interface{}, error) {
//		return []byte(RefreshSecret), nil
//	})
//	if err != nil {
//		return nil, true, err
//	}
//	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
//		return claims, false, nil
//	}
//	return nil, true, err
//}

//type MyClaims struct {
//	ID    int64
//	State string `json:"state"`
//	jwt.StandardClaims
//}
//
//// var accessSecret = []byte("liuxian123")
//// var refreshSecret = []byte("123456789")
//
//// GetToken 获取accessToken和refreshToken
//func GetToken(id int64, state, a, r string) (string, string) {
//	// accessToken 的数据
//	accessSecret := []byte(a)
//	refreshSecret := []byte(r)
//	aT := MyClaims{
//		id,
//		state,
//		jwt.StandardClaims{
//			Issuer:    "AR",
//			IssuedAt:  time.Now().Unix(),
//			ExpiresAt: time.Now().Add(3 * time.Minute).Unix(),
//		},
//	}
//	// refreshToken 的数据
//
//	rT := MyClaims{
//		id,
//		state,
//		jwt.StandardClaims{
//			Issuer:    "RT",
//			IssuedAt:  time.Now().Unix(),
//			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
//		},
//	}
//	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, aT)
//	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, rT)
//	accessTokenSigned, err := accessToken.SignedString(accessSecret)
//	if err != nil {
//		fmt.Println("获取Token失败，Secret错误")
//		return "", ""
//	}
//	refreshTokenSigned, err := refreshToken.SignedString(refreshSecret)
//	if err != nil {
//		fmt.Println("获取Token失败，Secret错误")
//		return "", ""
//	}
//	return accessTokenSigned, refreshTokenSigned
//}
//
//func ParseToken(accessTokenString, refreshTokenString, a, r string) (*MyClaims, bool, error) {
//	accessSecret := []byte(a)
//	refreshSecret := []byte(r)
//	accessToken, err := jwt.ParseWithClaims(accessTokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
//		return accessSecret, nil
//	})
//	if claims, ok := accessToken.Claims.(*MyClaims); ok && accessToken.Valid {
//		return claims, false, nil
//	}
//
//	refreshToken, err := jwt.ParseWithClaims(refreshTokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
//		return refreshSecret, nil
//	})
//	if err != nil {
//		return nil, false, err
//	}
//	if claims, ok := refreshToken.Claims.(*MyClaims); ok && refreshToken.Valid {
//		return claims, true, nil
//	}
//
//	return nil, false, errors.New("invalid token")
//}
