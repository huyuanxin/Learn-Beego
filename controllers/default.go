package controllers

import (
	_ "beego-learning/models"
	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/orm"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "Index.html"
}
