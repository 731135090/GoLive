package routers

import (
	"GoLive/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.BaseController{})

    beego.Router("/ws/cheat", &controllers.WebSocketController{})


    beego.Router("/admin/ws/cheat", &controllers.AdminWsController{})
}
