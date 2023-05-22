package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
	"server/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	//Get all mdts from the database
	c.Data["Mdts"], _ = models.GetMdt()
	c.TplName = "index.tpl"
}
