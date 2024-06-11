package lib

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"imaotai_helper/constant"
	"imaotai_helper/dto/token"
	"reflect"
)

type DBError struct {
	error
}

func NewDBError(err error) *DBError {
	e := &DBError{
		error: err,
	}
	return e
}

type ParamError struct {
	error
}

func NewParamError(err error) *ParamError {
	e := &ParamError{
		error: err,
	}
	return e
}

type ServiceError struct {
	error
}

func NewServiceError(err error) *ServiceError {
	e := &ServiceError{
		error: err,
	}
	return e
}

type RpcError struct {
	error
}

func NewRpcError(err error) *RpcError {
	e := &RpcError{
		error: err,
	}
	return e
}

var ErrorFunc []func(*gin.Context, *Response, error) bool

type Response struct {
	ErrNo  int         `json:"err_no"`         // 错误码
	ErrMsg string      `json:"err_msg"`        // 错误信息
	Data   interface{} `json:"data,omitempty"` // 返回结果
}

func SetErrMsg(msg string) func(response *Response) {
	return func(r *Response) {
		r.ErrMsg = msg
	}
}

func SetErrNo(errNo int) func(*Response) {
	return func(r *Response) {
		r.ErrNo = errNo
	}
}

func SetData(data interface{}) func(*Response) {
	return func(r *Response) {
		r.Data = data
	}
}

func NewResponse(ops ...func(response *Response)) *Response {
	//默认的resp
	resp := &Response{}
	for _, op := range ops {
		op(resp)
	}
	return resp
}

func SetResponse(resp *Response, ops ...func(response *Response)) {
	for _, op := range ops {
		op(resp)
	}
}

func SetReqErrorResponse(c *gin.Context, err error) {
	resp := NewResponse(SetErrNo(constant.ParamErr), SetErrMsg("参数错误"))
	c.JSON(200, resp)
	return
}

func SetContextErrorResponse(c *gin.Context, resp *Response, err error) {
	ErrorFunc = []func(*gin.Context, *Response, error) bool{SetParamErrorResponse, SetServiceErrorResponse, SetDBErrorResponse, SetRPCErrorResponse}
	for _, v := range ErrorFunc {
		if v(c, resp, err) {
			break
		}
	}

}

func UpdateToken(c *gin.Context) {
	// 更新token
	tokenStr := c.GetHeader(constant.HEADER_JWT)
	if len(tokenStr) == 0 {
		return
	}
	parsedToken, err := jwt.ParseWithClaims(tokenStr, &token.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名的密钥
		return []byte(constant.JWT_SALT), nil
	})
	if err != nil {
		return
	}
	if claim, ok := parsedToken.Claims.(*token.CustomClaims); ok && parsedToken.Valid {
		mobile := claim.Mobile
		adminUid := claim.AdminUid
		token, err := GenerateJwt(mobile, adminUid)
		if err != nil {
			return
		}
		c.Header(constant.HEADER_JWT, token)
	}
}

func SetContextSuccessResponse(c *gin.Context, resp *Response) {
	c.JSON(200, resp)
	// 更新token
	UpdateToken(c)
	return
}

func SetParamErrorResponse(c *gin.Context, resp *Response, err error) bool {
	paramError := &ParamError{}
	if reflect.TypeOf(err) != reflect.TypeOf(paramError) {
		return false
	}
	SetResponse(resp, SetErrNo(constant.ParamErr), SetErrMsg(err.Error()))
	c.JSON(200, resp)
	return true
}
func SetServiceErrorResponse(c *gin.Context, resp *Response, err error) bool {
	serviceError := &ServiceError{}
	if reflect.TypeOf(err) != reflect.TypeOf(serviceError) {
		return false
	}
	SetResponse(resp, SetErrNo(constant.ServerErr), SetErrMsg(err.Error()))
	c.JSON(200, resp)
	return true
}
func SetDBErrorResponse(c *gin.Context, resp *Response, err error) bool {
	dbError := &DBError{}
	if reflect.TypeOf(err) != reflect.TypeOf(dbError) {
		return false
	}
	SetResponse(resp, SetErrNo(constant.DBErr), SetErrMsg(err.Error()))
	c.JSON(200, resp)
	return true
}
func SetRPCErrorResponse(c *gin.Context, resp *Response, err error) bool {
	rpcError := &RpcError{}
	if reflect.TypeOf(err) != reflect.TypeOf(rpcError) {
		return false
	}
	SetResponse(resp, SetErrNo(constant.RPCErr), SetErrMsg(err.Error()))
	c.JSON(200, resp)
	return true
}
