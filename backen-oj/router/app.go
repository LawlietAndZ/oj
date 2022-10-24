package router

import (
	_ "backen-oj/docs"
	"backen-oj/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine{
	r := gin.Default()


	//swagger 配置
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.GET("/ping", service.Ping)

	//问题
	r.GET("/problem-list",service.GetProblemList)
	r.GET("/problem-detail",service.GetProblemDetail)

	//用户
	r.GET("user-detail",service.GetUserDetail)


	//提交记录
	r.GET("/submit-list",service.GetSubmitList)

	//登录
	r.POST("/login",service.Login)

	//邮箱发送验证码
	r.POST("/send-code",service.SendCode)
	return r
}
