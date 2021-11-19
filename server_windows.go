package main

import (
	"github.com/forseason/env"
)

func (this *app) serve() error {
	return this.r.Run(env.Get("SERVER_PORT", ":8080"))
}
