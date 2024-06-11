package lib

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"imaotai_helper/constant"
	"imaotai_helper/dto/token"
	"strings"
	"time"
)

// 解析imaotai的token
func ParseImaoTaiToken(tokenStr string) (*token.Token, error) {
	chunks := strings.Split(tokenStr, ".")
	if len(chunks) < 2 {
		return nil, errors.New("token格式错误")
	}
	decoder := base64.URLEncoding
	// 解码payload
	payloadBytes, _ := decoder.DecodeString(chunks[1])
	// 因为没有密钥,只能部分解码
	payload := string(payloadBytes) + "}"
	token := &token.Token{}
	err := json.Unmarshal([]byte(payload), &token)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func NewClaims(mobile string, adminUid string) *token.CustomClaims {
	// 创建一个自定义的Claims
	claims := &token.CustomClaims{
		Random:   time.Now().String(),
		Mobile:   mobile,
		AdminUid: adminUid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // 设置过期时间为24小时后
			Issuer:    "moyu_helper",                                      // 签发者
		},
	}
	return claims
}

func GenerateJwt(mobile string, adminUid string) (string, error) {
	// 创建一个自定义的Claims
	claims := NewClaims(mobile, adminUid)
	// 创建一个新的Token，指定签名方法和Claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用密钥对Token进行签名
	signedToken, err := token.SignedString([]byte(constant.JWT_SALT))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
