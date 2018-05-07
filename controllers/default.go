package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Github"] = "Yugovtr"
	c.Data["Email"] = "vitormirandawork@gmail.com"
	c.TplName = "index.html"
}
