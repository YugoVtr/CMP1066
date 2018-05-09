package controllers

type IndexController struct {
	MainController
}

func (c *IndexController) Get() {
	c.Data["Github"] = "Yugovtr"
	c.Data["Email"] = "vitormirandawork@gmail.com"
}
