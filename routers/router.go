package routers

import (
	"github.com/astaxie/beego"
	"mcms/controllers"
	_ "mcms/routers/admin"
	_ "mcms/routers/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.InsertFilter("/*", beego.BeforeRouter, FilterUser) //注册过滤器
}
