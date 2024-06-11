package request

type AddFlowerUserRequest struct {
	//Uid     string `json:"uid" binding:"required"`
	Mobile  string `json:"mobile" binding:"required"`
	Address string `json:"address" binding:"required"`
}
type StartFollowerUserRequest struct {
	Uid string `json:"uid" binding:"required"`
}
type SuspendFollowerUserRequest struct {
	Uid string `json:"uid" binding:"required"`
}
type DeleteFollowerUserRequest struct {
	Uid string `json:"uid" binding:"required"`
}

type ActivationUserRequest struct {
	Uid  string `json:"uid" binding:"required"`
	Code string `json:"code" binding:"required"`
}

type GetVerifyCodeReq struct {
	Uid string `json:"uid" binding:"required"`
}

type UpdateAddressRequest struct {
	Uid     string `json:"uid" binding:"required"`
	Address string `json:"address" binding:"required"`
}

type UpdateTokenRequest struct {
	Mobile string `json:"mobile" binding:"required"`
	Code   string `json:"code" binding:"required"`
}

type GetFlowerUserListReq struct {
	//UID      string `json:"uid" binding:"required"`
	Page     int `json:"page" `
	PageSize int `json:"page_size" `
}
