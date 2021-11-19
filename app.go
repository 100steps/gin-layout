package main

import (
	"github.com/100steps/gin-layout/controller"
	"github.com/100steps/gin-layout/cron"
	"github.com/100steps/gin-layout/middleware"
	"github.com/forseason/env"
	"github.com/gin-gonic/gin"
)

// 生成app容器并加载依赖的函数，这一层应当注入控制器依赖
func newApp(
	middleware *middleware.Middleware,
	crontab *cron.Cron,
	userController *controller.UserController,
) *app {
	switch env.Get("SERVER_MODE", "debug") {
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		gin.SetMode(gin.DebugMode)
	}
	app := &app{
		r:              gin.New(),
		middleware:     middleware,
		crontab:        crontab,
		userController: userController,
	}
	app.initRouter()
	return app
}

// 这里也要记得添加上对应的属性
type app struct {
	r              *gin.Engine
	middleware     *middleware.Middleware
	crontab        *cron.Cron
	userController *controller.UserController
}
