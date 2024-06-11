package request

type RegisterReq struct {
	Mobile   string `json:"mobile" binding:"required"`
	PassWord string `json:"pass_word" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Email    string `json:"email" binding:"required"`
}
type LoginReq struct {
	Mobile   string `json:"mobile" binding:"required"`
	PassWord string `json:"pass_word" binding:"required"`
}

type GetAdminInfoReq struct {
	Mobile string `json:"mobile" binding:"required"`
}
