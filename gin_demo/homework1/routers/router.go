package routers

import (
	"gin_demo/homework1/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	//创建引擎
	r := gin.Default()

	//注册:添加
	r.POST("/register", controller.RegisterHandler)
	//登录
	r.POST("/login", controller.LoginHandler)
	//查询
	r.GET("/get", controller.FindName)
	//更新
	r.PUT("/update", controller.UpdateHandler)
	//注销
	r.DELETE("/delete", controller.DeleteHandler)

	return r
}
