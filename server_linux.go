package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"syscall"

	_ "github.com/100steps/gin-layout/dao/migration"
	"github.com/forseason/env"
	"github.com/fvbock/endless"
)

func (this *app) serve() error {
	if env.Get("ENABLE_ENDLESS", "true") != "true" {
		return this.r.Run(env.Get("SERVER_PORT", ":8080"))
	}
	server := endless.NewServer(env.Get("SERVER_PORT", ":8080"), this.r)
	server.BeforeBegin = func(add string) {
		pid := syscall.Getpid()
		log.Printf("Actual pid is %d", pid)
		ioutil.WriteFile("pid", []byte(fmt.Sprintf("%d", pid)), 0777)
	}

	return server.ListenAndServe()
}
