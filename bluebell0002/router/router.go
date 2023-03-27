package router

import (
	"bluebell0002/controller"
	"bluebell0002/logger"
	"bluebell0002/middlewares"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetUpRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	//r.POST("/aaa", func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{
	//		"aaaaa": "1231",
	//	})
	//})

	v1 := r.Group("/api/v1")
	v1.POST("/signup", controller.SignUpHandler)
	//登录
	v1.POST("/login", controller.LoginHandler)

	//r.NoRoute(func(c *gin.Context) {
	//	c.JSON(http.StatusOK, gin.H{"msg": "404"})
	//})
	v1.Use(middlewares.JWTAuthMiddleware())
	{
		v1.GET("/community", controller.CommunityHandler)
		v1.GET("/community/:id", controller.CommunityDatailHandler)
		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		//v1.GET("/posts", controller.GetPostListHandler)
		////
		//v1.GET("/posts2", controller.GetPostListHandler2)
		//
		//v1.POST("/vote", controller.PostVoteController)
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "404"})
	})

	return r
}
