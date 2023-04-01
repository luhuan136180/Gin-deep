package routers

import (
	"bluebell/controller"
	"bluebell/logger"
	"bluebell/middleWares"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetUpRouter(mode string) *gin.Engine {

	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode) //设置成发布模式
	} //不然就是调试模式
	r := gin.New()
	//使用自己编写的logger
	//r.Use(logger.GinLogger(), logger.GinRecovery(true), middleWares.RateLimitMiddleware(2*time.Second, 1))
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	//路由信息
	v1 := r.Group("/api/v1")
	//注册
	v1.POST("/signup", controller.SignUpHandler)
	//登录
	v1.POST("/login", controller.LoginHandler)
	//

	//继续访问该项目的内容需要先登录，所以需要认证处理，添加中间件
	v1.Use(middleWares.JWTAuthMiddleware()) //应用JWT认证中间件

	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDatailHandler)

		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		v1.GET("/posts", controller.GetPostListHandler)
		//
		v1.GET("/posts2", controller.GetPostListHandler2)

		v1.POST("/vote", controller.PostVoteController)
	}
	pprof.Register(r)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "404"})
	})

	return r
}
