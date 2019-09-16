package main

import (
	"GoLive/config"
	_ "GoLive/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run(config.ListenIp + ":" + config.HttpPort)
	config.Gwg.Wait()
}
