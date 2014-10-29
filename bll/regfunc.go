package bll

import (
	"github.com/astaxie/beego"
	_ "mcms/bll/basetype"
	"mcms/bll/menu"
)

func init() {

	/*管理员 注册 函数*/
	beego.AddFuncMap("Manager_Get_Manager_Status", menu.Manager_Get_Manager_Status)
	beego.AddFuncMap("Tb_Manager_Status_List", menu.Tb_Manager_Status_List)
	/*菜单注册函数*/
	beego.AddFuncMap("Tb_Menu_Select_System_Menu", menu.Tb_Menu_Select_System_Menu)
	beego.AddFuncMap("Tb_Menu_GetParentName", menu.Tb_Menu_GetParentName)
	beego.AddFuncMap("Tb_Menu_GetAddNav_List", menu.Tb_Menu_GetAddNav_List)
	beego.AddFuncMap("Tb_Menu_CheckTop", menu.Tb_Menu_CheckTop)
	beego.AddFuncMap("Tb_Menu_GetWidth", menu.Tb_Menu_GetWidth)
}
