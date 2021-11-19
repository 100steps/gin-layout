package middleware

import (
	"github.com/100steps/gin-layout/service"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewMiddleware)

type Middleware struct {
	userService *service.UserService
}

func NewMiddleware(userService *service.UserService) *Middleware {
	return &Middleware{userService: userService}
}
