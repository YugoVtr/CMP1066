package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Layout = "layout/default.html"
	c.TplName = "view/index.html"
	c.LayoutSections = make(map[string]string)
    c.LayoutSections["HtmlHead"] = "html_head/index.html"
    c.LayoutSections["Scripts"] = ""
	c.LayoutSections["Sidebar"] = ""
	
	c.Data["Github"] = "Yugovtr"
	c.Data["Email"] = "vitormirandawork@gmail.com"
}
