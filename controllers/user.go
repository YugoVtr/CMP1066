package controllers

import (
	"CMP1066/models"
)

type UserController struct {
	MainController
}

func (c *UserController) Get() {
	_, users := models.GetAllUsers()

	c.Data["Users"] = users
}
