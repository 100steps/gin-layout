// +build wireinject

package main

import (
	"github.com/100steps/gin-layout/controller"
	"github.com/100steps/gin-layout/cron"
	"github.com/100steps/gin-layout/middleware"
	"github.com/100steps/gin-layout/service"
	"github.com/google/wire"
)

// DO NOT EDIT.
func initApp() (*app, error) {
	panic(wire.Build(middleware.ProviderSet, cron.ProviderSet, controller.ProviderSet, service.ProviderSet, newApp))
}
