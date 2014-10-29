package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"mcms/dal/menu"
	"mcms/models"
	"mcms/util"
	"mcms/util/config"
	"mcms/util/lhdialog"
	"strconv"
)

type MenuController struct {
	beego.Controller
}

func (this *MenuController) MenuList() {
	parentId, err := this.GetInt("parentId")
	if err != nil {
		parentId = 0
	}
	list := menu.Tb_Menu_Select_All(int(parentId))
	this.Data["list"] = list
	this.Data["Ctx"] = this.Ctx
	this.Data["msg"] = ""
	this.TplNames = "admin/menu/list.html"
}
func (this *MenuController) DeleteInfo() {
	var msg = "{\"status\":\"0\",\"msg\":\"数据删除失败\"}"
	Id := this.GetString("Id")
	if Id == "" {
		this.Ctx.WriteString(msg)
		return
	}
	var flag bool = menu.DeleteList(Id)
	if flag {
		msg = "{\"status\":\"1\",\"msg\":\"数据删除成功\"}"
	}
	this.Ctx.WriteString(msg)
}
func (this *MenuController) AddMenu() {
	Id, err := this.GetInt("Id")
	if err != nil {
		Id = 0
	}
	ParentId, err1 := this.GetInt("parentId")
	if err1 != nil {
		ParentId = 0
	}
	var model models.Tb_menu
	model = menu.GetModel(Id)
	this.Data["parentId"] = ParentId
	this.Data["model"] = model
	this.Data["rolelist"] = config.Tb_Navigation_ActionType()
	this.Data["msg"] = ""
	this.TplNames = "admin/menu/add.html"
}

func (this *MenuController) AddMenuInfo() {
	model := models.Tb_menu{}
	err := this.ParseForm(&model)

	list := config.Tb_Navigation_ActionType()
	var strTemp string = ""
	for k, _ := range list {
		if this.GetString("cblActionType_"+k) == "on" {
			strTemp += k + ","
		}
	}
	if strTemp != "" {
		strTemp = util.Substr(strTemp, len(strTemp)-1)
	}
	model.Action = strTemp
	if err != nil {
		fmt.Println(err.Error())
	}
	if model.Id == 0 {
		_, flag := menu.Add(model)
		if flag {
			this.Data["msg"] = lhdialog.LhDialog_Diy("提示", "数据操作成功,是否继续添加？", "div1", "window.open('/admin/menu/add?Id=0','_self');return false", "", "success.gif", true)
		} else {
			this.Data["msg"] = lhdialog.LhDialog_Tips("数据操作失败，请联系管理员")
		}
	} else {
		_, flag := menu.Update(model)
		if flag {
			this.Data["msg"] = lhdialog.LhDialog_Diy("提示", "数据操作成功,是否继续修改？", "div1", "window.open('/admin/menu/add?Id="+strconv.FormatInt(model.Id, 10)+"','_self');return false", "window.open('/admin/menu/list','_self');return false", "success.gif", true)
		} else {
			this.Data["msg"] = lhdialog.LhDialog_Tips("数据操作失败，请联系管理员")
		}
	}
	this.Data["rolelist"] = config.Tb_Navigation_ActionType()
	this.Data["parentId"] = int64(model.ParentId)
	this.Data["model"] = model
	this.TplNames = "admin/menu/add.html"
}
func (this *MenuController) UpdateOrder() {
	var msg = "{\"status\":\"1\",\"msg\":\"数据操作成功\"}"
	Id, _ := this.GetInt("id")
	orderTop, _ := this.GetInt("orderTop")
	menu.Update_Order(Id, orderTop)
	this.Ctx.WriteString(msg)
}
