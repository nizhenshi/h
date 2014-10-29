package menu

import (
	"github.com/astaxie/beego/context"
	"github.com/astaxie/beego/orm"
	"html/template"
	"mcms/dal/menu"
	"mcms/models"
	"mcms/util"
	"strconv"
	"strings"
)

func Tb_Menu_CheckTop(Id int64) bool {
	var model models.Tb_menu
	model.Id = Id
	o := orm.NewOrm()
	o.Read(&model)
	if model.ParentId == 0 {
		return true
	}
	return false
}

/*获取 导航列表中 菜单名 缩略 空格*/
func Tb_Menu_GetWidth(Id int64) int {
	return (menu.Select_Nav_Deep(Id) - 2) * 24
}

/*获取 菜单 深度*/
func Select_Nav_Deep(Id int64) int {
	return menu.Select_Nav_Deep(Id)
}

/*根据 父ID 获取 所有 子菜单*/
func Tb_Menu_Select_List_Parent(parentId int) []*models.Tb_menu {
	return menu.Tb_Menu_Select_All(parentId)
}

/*获取 添加导航菜单页面中 title 缩略 空格*/
func Tb_Menu_Add_SelectTitle(Id int64, title string) template.HTML {
	deep := menu.Select_Nav_Deep(Id)
	strTemp := ""
	for i := 0; i < deep; i++ {
		strTemp += "&nbsp;&nbsp;"
	}
	strTemp += "|--" + title
	return template.HTML(strTemp)
}

/*获取 父菜单名称*/
func Tb_Menu_GetParentName(parentId int64) string {
	var strTemp = "--无父级菜单--"
	if parentId != 0 {
		model := menu.GetModel(parentId)
		strTemp = model.Title
	}
	return strTemp
}

/*添加 导航菜单 验证 是否 有此权限*/
func Tb_Menu_Check_Role(navaction, current string) bool {
	if strings.Contains(navaction, current) {
		return true
	}
	return false
}

/*获取 添加 菜单下拉列表菜单*/
func Tb_Menu_GetAddNav_List(parentId int) string {
	list := menu.Tb_Menu_Select_All(parentId)
	var strTemp string = ""
	for _, m := range list {
		strTemp += "{id:" + strconv.FormatInt(m.Id, 10) + ", pId:" + strconv.Itoa(m.ParentId) + ", name:'" + m.Title + "'},"
	}
	if strTemp != "" {
		strTemp = util.Substr(strTemp, util.Len(strTemp)-1)
	}
	return strTemp
}
func Tb_Menu_CheckDeep(Id int64, currDeep int) bool {
	var deep int = menu.Select_Nav_Deep(Id)
	if deep == currDeep {
		return true
	}
	return false
}

/*
前台 页面 调用 方法
*/
func Tb_Menu_Select_System_Menu(ctx *context.Context) template.HTML {
	var strTemp template.HTML = ""
	strTemp = Tb_Menu_Menu_System()
	return strTemp
}

/*获取 导航菜单  admin/index  超级管理员页面*/
func Tb_Menu_Menu_System() template.HTML {
	var str string = ""
	o := orm.NewOrm()
	list := []*models.Tb_menu{}
	model := &models.Tb_menu{}
	o.QueryTable(model.TableName()).All(&list)
	var i int = 0
	for _, l := range menu.Select_List(list, 5) {
		str += "<div class='list-group' name='" + l.Title + "'>"
		str += "<ul>"
		for _, t := range menu.Select_List(list, int(l.Id)) {
			str += "<li>"
			str += "<a navid='" + t.Name + "' class='item'><span>" + t.Title + "</span></a>"
			str += "<ul>"
			i = 0
			for _, f := range menu.Select_List(list, int(t.Id)) {
				str += "<li><a navid='" + f.Name + "' href='" + f.LinkUrl + "' target='" + f.Target + "' class='item"
				if i == 0 {
					str += " selected "
				}
				str += "'><span>" + f.Title + "</span> </a></li>"
				i++
			}
			str += "</ul></li>"
		}
		str += "</ul>"
		str += "</div>"
	}
	return template.HTML(str)
}

/*
/*获取 导航菜单  admin/index  系统管理员页面
func Tb_Menu_Menu_Admin(groupId string) template.HTML {
	var str string = ""
	o := orm.NewOrm()
	list := []*models.Tb_menu{}
	model := &models.Tb_menu{}
	o.QueryTable(model.TableName()).All(&list)
	for _, l := range menu.Select_List(list, 56) {
		if Tb_Role_Check_Menu_Role(groupId, l.Id, "Show") {
			str += "<div class='list-group' name='" + l.Title + "'>"
			str += "<ul>"
			for _, t := range menu.Select_List(list, int(l.Id)) {
				if Tb_Role_Check_Menu_Role(groupId, t.Id, "Show") {
					str += "<li>"
					str += "<a navid='" + t.Name + "' class='item'><span>" + t.Title + "</span></a>"
					str += "<ul>"
					for _, f := range menu.Select_List(list, int(t.Id)) {
						if Tb_Role_Check_Menu_Role(groupId, f.Id, "Show") {
							str += "<li><a navid='" + f.Name + "' href='" + f.LinkUrl + "' target='" + f.Target + "' class='item'><span>" + f.Title + "</span> </a></li>"
						}
					}
					str += "</ul></li>"
				}
			}
			str += "</ul></div>"
		}
	}
	return template.HTML(str)
}*/

/*=================会员前台菜单===================*/

func Tb_Menu_PersonMenu(menuId int64) template.HTML {
	var strTemp = ""
	strTemp += "<div class=\"help-sub-l lf\">"
	strTemp += "<div class=\"left_top\">"
	strTemp += "<span style=\"background:none; padding-left:6px;\">会员中心</span>"
	strTemp += "</div>"
	strTemp += "<div class=\"memu\">"
	strTemp += "<ul>"

	list := menu.Tb_Menu_Select_All(int(menuId))
	for _, m := range menu.Select_List(list, int(menuId)) {
		strTemp += "<li>"
		strTemp += "<h2>" + m.Title + "</h2>"
		for _, l := range menu.Select_List(list, int(m.Id)) {
			strTemp += "<p>"
			strTemp += "<a href=\"" + l.LinkUrl + "\">" + l.Title + "</a>"
			strTemp += "</p>"
		}
		strTemp += "</li>"

	}
	strTemp += "</ul></div></div>"
	return template.HTML(strTemp)
}
