package controllers

import (
	"github.com/astaxie/beego"
	"strings"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Prepare() {
	c.Layout = "layout/default.html"

	//Custom layout
	uri := c.Ctx.Input.Data()["RouterPattern"].(string)
	view := strings.Split(uri,"/")[1]

	c.TplName = "view/" + view + ".html"
	c.LayoutSections = make(map[string]string)
    c.LayoutSections["HtmlHead"] = "html_head/" + view + ".html"
    c.LayoutSections["Scripts"] = "scripts/" + view + ".html"
	c.LayoutSections["Sidebar"] = "sidebar/" + view + ".html"
}
