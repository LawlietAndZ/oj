package middlewares

import (
	"backen-oj/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthAdminCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		userClaim, err := helper.AnalyseToken(auth)
		if err != nil {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "你没有权限访问此功能",
			})
			return
		}
		if userClaim == nil || userClaim.IsAdmin != 1 {
			c.Abort()
			c.JSON(http.StatusOK, gin.H{
				"code": http.StatusUnauthorized,
				"msg":  "你没有权限访问此功能",
			})
			return
		}
		c.Next()
	}
}

