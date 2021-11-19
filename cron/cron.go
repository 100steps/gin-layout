package cron

import (
	// "log"

	"github.com/google/wire"
	"github.com/robfig/cron/v3"
)

var ProviderSet = wire.NewSet(NewCron)

func NewCron() *Cron {
	crontab := cron.New()
	// 每分钟打印一次hello
	// crontab.AddFunc("* * * * *", func() {
	// 	log.Println("hello")
	// })
	return &Cron{c: crontab}
}

type Cron struct {
	c *cron.Cron
}

func (this *Cron) Excute() {
	go this.run()
}

func (this *Cron) run() {
	this.c.Start()
	select {}
}
