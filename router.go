package main

import (
	"github.com/forseason/env"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

var store redis.Store

// 这个函数用于注册路由
func (this *app) initRouter() {
	r := this.r
	r.Use(this.middleware.CORS())

	// 初始化session组件
	var err error
	store, err = redis.NewStore(30, "tcp", env.Get("REDIS_HOST", ""), env.Get("REDIS_PASSWORD", ""), []byte("__100steps__100steps__100steps__"))
	if err != nil {
		panic(err)
	}
	r.Use(sessions.Sessions(env.Get("SESSION_NAME", "gin-layout"), store))

	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "ok")
	})

	api := r.Group("/api")
	api.GET("/user", this.userController.Index)
	api.POST("/user", this.userController.Register)
	api.POST("/session", this.userController.Login)

	authedApi := api.Group("", this.middleware.Authentication())
	authedApi.GET("/me", this.userController.ShowMe)
}
