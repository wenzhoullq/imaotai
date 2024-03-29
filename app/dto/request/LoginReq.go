package request

type LoginRequest struct {
	Mobile string `json:"mobile" binding:"required"`
	Code   string `json:"code" binding:"required"`
}

type GetVerifyCodeReq struct {
	Mobile string `json:"mobile" binding:"required"`
}

type UpdateAddressRequest struct {
	Mobile  string `json:"mobile" binding:"required"`
	Address string `json:"address" binding:"required"`
}

type UpdateTokenRequest struct {
	Mobile string `json:"mobile" binding:"required"`
	Code   string `json:"code" binding:"required"`
}
