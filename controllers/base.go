package controllers

import (
	"GoLive/config"
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Get() {
	defer c.Head()
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *BaseController) Head() {
	if config.RunLevel != config.RUN_LEVEL_RELEASE {
		c.Ctx.Output.Header("Access-Control-Allow-Origin", "*")
	}
}
