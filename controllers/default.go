package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Layout = "Layout/Default/layout.html"
	c.TplName = "index.html"
	c.LayoutSections = make(map[string]string)
    c.LayoutSections["HtmlHead"] = "Layout/Default/html_head.html"
    c.LayoutSections["Scripts"] = "Layout/Default/scripts.html"
	c.LayoutSections["Sidebar"] = ""
	
	c.Data["Github"] = "Yugovtr"
	c.Data["Email"] = "vitormirandawork@gmail.com"
}
