package token

import "github.com/golang-jwt/jwt/v4"

// ImaotaiToken
type Token struct {
	Iss      string `json:"iss"`
	Exp      int64  `json:"exp"`
	UserId   int64  `json:"userId"`
	DeviceId string `json:"deviceId"`
	Iat      int64  `json:"iat"`
}

// 系统自定义的token
type CustomClaims struct {
	Random   string `json:"random"`
	Mobile   string `json:"mobile"`
	AdminUid string `json:"adminUid"`
	jwt.RegisteredClaims
}
