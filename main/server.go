package main

import (
	"GoLive/config"
	_ "GoLive/routers"
	"github.com/astaxie/beego"
)

func main() {
	config.Init()
	beego.Run(config.ListenIp + ":" + config.HttpPort)
}
