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
	c.Layout = "layout/default.html"
	c.TplName = "view/user.html"

	c.LayoutSections = make(map[string]string)
    c.LayoutSections["HtmlHead"] = ""
    c.LayoutSections["Scripts"] = ""
    c.LayoutSections["Sidebar"] = ""
}
