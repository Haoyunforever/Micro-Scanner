package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/Haoyunforever/Micro-Scanner/pkg/utils"
)

// JWT token验证中间件
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code uint32

		code = http.StatusOK
		token := c.GetHeader("Authorization")
		fmt.Println(token)
		if token == "" {
			code = http.StatusNotFound
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  "鉴权失败",
			})
			return
		}
		fmt.Println(token)
		claims, err := utils.ParseToken(token)
		fmt.Println(err)
		if err != nil {
			code = http.StatusUnauthorized
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  "鉴权失败",
			})
			c.Abort()
			return
		}

		if time.Now().Unix() > claims.ExpiresAt {
			code = http.StatusUnauthorized
			c.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg":  "权限过期，请重新登陆",
			})
			c.Abort()
			return
		}

		//ctl.InitUserInfo(c.Request.Context(), &ctl.UserInfo{Id: claims.Id})
		c.Next()
	}
}
