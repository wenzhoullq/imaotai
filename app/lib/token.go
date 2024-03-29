package lib

import (
	"encoding/base64"
	"encoding/json"
	"strings"
	"zuoxingtao/dto/token"
)

func ParseToken(tokenStr string) (*token.Token, error) {
	chunks := strings.Split(tokenStr, ".")
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
