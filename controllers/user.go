package controllers

import (
	"CMP1066/models"
)

type UserController struct {
	MainController
}

func (c *UserController) Get() {
	_, users := models.GetAllUsers()

	c.Data["Form"] = &models.User{Status: true}
	c.Data["Users"] = users
}

func (c *UserController) Post() {
	user := models.User{}
	c.ParseForm(&user)
	
    userid := models.AddOne(user)
    c.Data["json"] = map[string]interface{}{"User": userid }
    c.ServeJSON()
}