package menu

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	html "html/template"
	"mcms/models"
	"strconv"
)

func Add(model models.Tb_menu) (number int64, flag bool) {
	o := orm.NewOrm()
	o.Begin()
	num, err := o.Insert(&model)
	if err != nil {
		o.Rollback()
		return 0, false
	}
	o.Commit()
	return num, true
}
func Update_Order(Id, orderTop int64) {
	o := orm.NewOrm()
	var model models.Tb_menu
	sql := "update " + model.TableName() + " set ordertop =" + strconv.FormatInt(orderTop, 10) + " where  id=" + strconv.FormatInt(Id, 10)
	res, err := o.Raw(sql).Exec()
	if err == nil {
		res.RowsAffected()
	}
}
func Update(model models.Tb_menu) (number int64, flag bool) {
	o := orm.NewOrm()
	o.Begin()
	var oldModel models.Tb_menu
	oldModel.Id = model.Id
	o.Read(&oldModel)
	oldModel.Action = model.Action
	oldModel.LinkUrl = model.LinkUrl
	oldModel.Name = model.Name
	oldModel.OrderTop = model.OrderTop
	oldModel.ParentId = model.ParentId
	oldModel.Remark = model.Remark
	oldModel.Target = model.Target
	oldModel.Title = model.Title
	oldModel.ViewFlag = model.ViewFlag

	num, err := o.Update(&oldModel)
	if err != nil {
		o.Rollback()
		return 0, false
	}
	o.Commit()
	return num, true
}
func GetModel(Id int64) models.Tb_menu {
	model := models.Tb_menu{}
	model.Id = Id
	o := orm.NewOrm()
	err := o.Read(&model)
	//o.Raw("select * from " + model.TableName() + " where id=" + strconv.FormatInt(Id, 10)).QueryRow(&model)
	if err != nil {
		model.Id = 0
		model.Action = ""
		model.LinkUrl = "/admin/center"
		model.Name = ""
		model.OrderTop = 0
		model.ParentId = 0
		model.Remark = ""
		model.Target = "mainframe"
		model.Title = ""
		model.ViewFlag = 0
	}
	return model
}
func GetModel_ChannelName(channelName string) models.Tb_menu {
	model := models.Tb_menu{}
	model.Name = channelName
	o := orm.NewOrm()
	o.Read(&model, "name")
	return model
}
func Delete(id int64) bool {
	o := orm.NewOrm()
	if id == 0 {
		return false
	}
	model := &models.Tb_menu{Id: id}
	o.Begin()
	_, err := o.Delete(model)
	if err != nil {
		o.Rollback()
		fmt.Println("菜单表：" + err.Error())
		return false
	}
	o.Commit()
	return true
}
func DeleteList(Id string) bool {
	o := orm.NewOrm()
	model := models.Tb_menu{}
	if Id == "" {
		return false
	}
	o.Begin()
	_, err := o.QueryTable(model.TableName()).Filter("id__in", Id).Delete()
	if err != nil {
		o.Rollback()
		return false
	}
	o.Commit()
	return true
}

/*根据 父ID 获取 所有 子菜单 */
func Select_List_Parent(parentId int) []*models.Tb_menu {
	o := orm.NewOrm()
	list := []*models.Tb_menu{}
	model := models.Tb_menu{}
	_, err := o.QueryTable(model.TableName()).Filter("ParentId", parentId).All(&list)
	if err != nil {
		fmt.Println("菜单表：" + err.Error())
	}
	return list
}

/*根据 父ID 获取 所有 子菜单 ,不操作数据 库*/
func Select_List(oldList []*models.Tb_menu, parentId int) []*models.Tb_menu {
	list := []*models.Tb_menu{}
	for _, l := range oldList {
		if l.ParentId == parentId {
			list = append(list, l)
		}
	}
	list = SortList(list)
	return list
}

/*获取 子菜单 项 -1 获取所有 菜单 */
func Select_ParentId(parentId int) []*models.Tb_menu {
	o := orm.NewOrm()
	m := models.Tb_menu{}
	list := []*models.Tb_menu{}
	qs := o.QueryTable(m.TableName())
	if parentId > -1 {
		qs = qs.Filter("Parentid", parentId)
	}
	qs.All(&list)
	return list
}

/*
根据 父ID 获取 所有 菜单列表
*/
func Tb_Menu_Select_All(parentId int) []*models.Tb_menu {
	item := []*models.Tb_menu{}
	classList := Select_ParentId(-1)
	item = select_Class_Slice(parentId, item, classList)
	return item
}
func select_Class_Slice(parentId int, item, classList []*models.Tb_menu) []*models.Tb_menu {
	if len(classList) > 0 {
		tempList := []*models.Tb_menu{}
		for _, v := range classList {
			if v.ParentId == parentId {
				tempList = append(tempList, v)
			}
		}
		tempList = SortList(tempList)
		for _, v := range tempList {
			item = append(item, v)
			item = select_Class_Slice(int(v.Id), item, classList)
		}
	}
	return item
}

/*
获取 菜单深度
*/
func Select_Nav_Deep(Id int64) int {
	deep := 1
	t := &models.Tb_menu{}
	t.Id = Id
	o := orm.NewOrm()
	o.Read(t)
	if t.ParentId == 0 {
		return deep
	} else {
		deep += Select_Nav_Deep(int64(t.ParentId))
	}
	return deep
}

/*
对 菜单列表 进行排序
*/
func SortList(list []*models.Tb_menu) []*models.Tb_menu {
	l := len(list)
	for i := 0; i < l; i++ {
		for j := i + 1; j < l; j++ {
			if list[i].OrderTop > list[j].OrderTop {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
	return list
}

/**
 * 设置当前页面的菜单的根菜单,用来生成2即菜单
 */
func SetRootMenuid(mid int64) html.HTML {
	o := orm.NewOrm()
	var l []models.Tb_menu
	rs1 := "<input type='hidden' id=\"mid\" name='mid' value='" + strconv.Itoa(int(mid)) + "'>"
	rs2 := ""
	var submenus string
	o.QueryTable("Tb_menu").Filter("ParentId", mid).All(&l)
	if len(l) > 0 {
		submenus += "<div class='nav_lm'><div class='nav_lm_nr'>"
		for _, v := range l {
			submenus += "<a href='" + v.LinkUrl + "'>" + v.Title + "</a>"
		}
		submenus += "</div></div>"
		rs2 += "<input type='hidden' id=\"submenus\" name='submenus' value=\"" + submenus + "\">"
	}
	return html.HTML(rs1 + rs2)
}
