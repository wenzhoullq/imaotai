package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"imaotai_helper/constant"
	"imaotai_helper/dao"
	"imaotai_helper/dto/token"
	"net/http"
	"strings"
)

func CheckJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenStr := strings.TrimSpace(c.Request.Header.Get(constant.HEADER_JWT))
		if len(tokenStr) == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		parsedToken, err := jwt.ParseWithClaims(tokenStr, &token.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			// 验证签名的密钥
			return []byte(constant.JWT_SALT), nil
		})
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// 检查Token是否有效
		claims, ok := parsedToken.Claims.(*token.CustomClaims)
		if !ok || !parsedToken.Valid {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		adminUid := claims.AdminUid
		//检查Uid是否有效
		AdminDao := dao.NewAdminDao()
		_, err = AdminDao.GetAdminByUid(adminUid)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		//配置UID
		c.Set("admin_uid", adminUid)
		c.Set("mobile", claims.Mobile)
		c.Next()
	}
}
