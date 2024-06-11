package response

type GetUserListResp struct {
	Uid      string `json:"uid"`
	Mobile   string `json:"mobile"` // 手机号
	Status   string `json:"status"` // 状态
	Address  string `json:"address"`
	UserName string `json:"user_name"`
	ExpTime  string `json:"exp_time"`
}

type GetAdminInfoResp struct {
	Uid    string `json:"uid"`
	Mobile string `json:"mobile"`
	Role   int    `json:"role"`
}
