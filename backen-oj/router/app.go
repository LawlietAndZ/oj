package router

import (
	_ "backen-oj/docs"
	"backen-oj/middlewares"
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



	//公有方法
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
	//注册
	r.POST("/register",service.Register)
	//用户提交排行榜
	r.GET("/rank-list",service.GetRankList)

	// 分类列表
	r.GET("/category-list", service.GetCategoryList)



	//管理员私有方法
	authAdmin := r.Group("/admin", middlewares.AuthAdminCheck())
	//问题创建
	authAdmin.POST("/problem-create",service.ProblemCreate)
	//问题修改
	authAdmin.GET("/problem-modify",service.ProblemModify)

	// 分类创建
	authAdmin.POST("/category-create", service.CategoryCreate)
	// 分类修改
	authAdmin.PUT("/category-modify", service.CategoryModify)
	// 分类删除
	authAdmin.DELETE("/category-delete", service.CategoryDelete)


	// 用户私有方法
	authUser := r.Group("/user", middlewares.AuthUserCheck())
	// 代码提交
	authUser.POST("/submit", service.Submit)



	return r
}
