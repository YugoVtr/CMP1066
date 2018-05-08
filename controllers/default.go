package controllers

import (
	"github.com/astaxie/beego"
	"CMP1066/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Github"] = "Yugovtr"
	c.Data["Email"] = "vitormirandawork@gmail.com"
	c.Data["User"] =  models.GetUser(1)
	c.Layout = "Layout/Default/layout.html"
	c.TplName = "index.html"

	c.LayoutSections = make(map[string]string)
    c.LayoutSections["HtmlHead"] = "Layout/Default/html_head.html"
    c.LayoutSections["Scripts"] = "Layout/Default/scripts.html"
    c.LayoutSections["Sidebar"] = ""
}
