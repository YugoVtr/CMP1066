package controllers

import (
	"CMP1066/models"
	"strconv" 			//Para converter string em int64
	"crypto/sha256"		//Para fazer o hash da senha 
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

	verify := models.GetUserByLogin(user.Login)
	
	if verify.Id != 0 { 
		c.Data["json"] = map[string]interface{}{"User": "Username already exists"}
	} else { 
		//SHA-256
		hashSenha := sha256.Sum256([]byte(user.Password))
		user.Password = string(hashSenha[:])
		
		userid := models.AddOne(user)
		c.Data["json"] = map[string]interface{}{"User": userid }
	}

    c.ServeJSON()
}

func (c *UserController) Delete() {
	input   := c.Ctx.Input.Param(":id")
	idUser,_:= strconv.ParseInt(input,10,64)
	var valid bool = (models.Delete(idUser) == 0)

    c.Data["json"] = map[string]interface{}{"Success": valid }
    c.ServeJSON()
}