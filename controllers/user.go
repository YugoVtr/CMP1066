package controllers

import (
	"github.com/astaxie/beego"
	"CMP1066/models"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	_, users := models.GetAllUsers()

	c.Data["Users"] = users
	c.Layout = "Layout/Default/layout.html"
	c.TplName = "users.html"

	c.LayoutSections = make(map[string]string)
    c.LayoutSections["HtmlHead"] = "Layout/Default/html_head.html"
    c.LayoutSections["Scripts"] = "Layout/Default/scripts.html"
    c.LayoutSections["Sidebar"] = ""
}
