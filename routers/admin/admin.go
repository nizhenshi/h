package admin

import (
	"github.com/astaxie/beego"
	m "mcms/controllers"
	menu "mcms/controllers/admin/menu"
)

func init() {
	/*登陆 相关路由*/
	beego.Router("/admin", &m.LoginController{}, "*:AdminLogin")
	beego.Router("/login", &m.LoginController{})
	beego.Router("/admin/center", &m.LoginController{}, "get:Center")
	/*菜单*/
	beego.Router("/admin/menu/list", &menu.MenuController{}, "get:MenuList")
	beego.Router("/admin/menu/delete", &menu.MenuController{}, "post:DeleteInfo")
	beego.Router("/admin/menu/add", &menu.MenuController{}, "get:AddMenu;post:AddMenuInfo")
	beego.Router("/admin/menu/uorder", &menu.MenuController{}, "post:UpdateOrder")
}
