package controllers

import (
	"CMP1066/models"
	"strconv"
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

func (c *UserController) Delete() {
	input   := c.Ctx.Input.Param(":id")
	idUser,_:= strconv.ParseInt(input,10,64)
	savedId := models.Delete(idUser)

	valid := false

	if savedId != 0 {
		valid = true
	} 

    c.Data["json"] = map[string]interface{}{"Success": valid }
    c.ServeJSON()
}