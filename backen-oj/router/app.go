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
	r.GET("/problem-list",service.GetProblemList)
	r.GET("/problem-detail",service.GetProblemDetail)
	return r
}
