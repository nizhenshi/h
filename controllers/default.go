package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["uname"] = "admin"
	this.Data["IsLogin"] = checkAccount(this.Ctx)
	this.TplNames = "index.html"
}
