package token

type Token struct {
	Iss      string `json:"iss"`
	Exp      int64  `json:"exp"`
	UserId   int64  `json:"userId"`
	DeviceId string `json:"deviceId"`
	Iat      int64  `json:"iat"`
}
