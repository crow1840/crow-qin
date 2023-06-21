package router

import (
	_ "crow-qin/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func ServerRouter() {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//创建一个groupV1组，里边放一个检查存活的接口
	groupV1 := r.Group("/api/v1")
	{
		groupV1.GET("/ping", Ping)
		groupV1.GET("/version", Version)
		groupV1.GET("/liubei", LiuBei) //这是一个测试接口

		//登录
		groupV1.POST("/login", Login)           //登录服务
		groupV1.POST("/logout", Logout)         //用户登出
		groupV1.POST("/login/refresh", Refresh) //生成新token

		//admin接口
		groupV1.POST("/admin/user", admin(), CreateUser)                //创建用户
		groupV1.PUT("/admin/user_password", admin(), ResetUserPassword) //重置用户密码
		groupV1.GET("/admin/users", admin(), GetUsers)                  //获取用户信息，支持模糊查询
		groupV1.GET("/admin/users/:uuid", admin(), GetUser)             //删除指定用户
		groupV1.DELETE("/admin/users/:uuid", admin(), DeleteUser)       //删除指定用户
		groupV1.PUT("/admin/users", admin(), UpdateUsersInfo)           //修改用户信息

		//普通用户接口
		groupV1.PUT("/user/info", user(), UpdateUserInfo)         //用户修改自己的信息
		groupV1.PUT("/user/password", user(), UpdateUserPassword) //用户修改自己的密码
		groupV1.GET("/user/info", user(), GetUserInfo)            //用户查询自己信息
	}

	//前端相关路由
	webGroup := r.Group("")
	{
		webGroup.GET("/", indexHtml)
		webGroup.StaticFS("/js", http.Dir("./web/js"))
		//r.StaticFile("/deamon.js", "./web/deamon.js")
	}

	r.Run(":1840")
}
