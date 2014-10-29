package basetype

import (
	"fmt"
	"github.com/astaxie/beego"
	dalbt "mcms/dal/basetype"
	m "mcms/models"
	dlg "mcms/util/lhdialog"
	con "strconv"
	"strings"
	"time"
)

type BasetypeController struct {
	beego.Controller
}

/**
 * 类别列表
 */
func (this *BasetypeController) List() {
	pid, _ := this.GetInt("pid")
	this.Data["pid"] = pid
	l := dalbt.GetOnlyChildSlices(pid)
	var list []m.Tb_basetype
	for _, v := range l {
		v.ParentId = 0 //把展示的parent修改为0 要不,就不能用TreeTable
		list = append(list, v)
		this.Data["list"] = list
	}
	this.Data["list"] = list
	this.TplNames = "admin/basetype/list.html"
}

/**
 * 编辑,相应get请求
 */
func (this *BasetypeController) Edit() {
	id, _ := this.GetInt("id")
	pid, _ := this.GetInt("pid")
	menuRootId, _ := this.GetInt("menuRootId")
	this.Data["menuRootId"] = menuRootId
	this.Data["Id"] = id
	this.Data["strParentMenuList"] = dalbt.GetZtreeJsonByFid(menuRootId)
	if id == 0 { //新加
		this.Data["IsShow"] = int64(1)
		this.Data["ParentId"] = pid
		this.Data["Date"] = time.Now().Format("2006-01-02 15:04:05")
		this.Data["OrderTop"] = 0
		if pid == 0 {
			this.Data["ParentName"] = "--无父级菜单--"
		} else {
			this.Data["ParentName"] = dalbt.GetById(pid).Title
		}
	} else { //编辑
		tm := dalbt.GetById(id)
		this.Data["ParentId"] = tm.ParentId
		this.Data["IsShow"] = tm.IsShow
		this.Data["Title"] = tm.Title
		this.Data["Date"] = tm.Date.Format("2006-01-02 15:04:05")
		this.Data["OrderTop"] = tm.OrderTop
		this.Data["ImageUrl"] = tm.ImageUrl
		this.Data["Describe"] = tm.Describe
		this.Data["Remark"] = tm.Remark
		if tm.ParentId == 0 {
			this.Data["ParentName"] = "--无父级菜单--"
		} else {
			this.Data["ParentName"] = dalbt.GetById(tm.ParentId).Title
		}
	}
	fmt.Println("edit.html")
	this.TplNames = "admin/basetype/edit.html"
}

/**
 * 编辑页面,响应post请求
 */
func (this *BasetypeController) EditPost() {
	menuRootId, _ := this.GetInt("menuRootId")
	this.Data["menuRootId"] = menuRootId
	this.Data["IsShow"] = 1
	tm := m.Tb_basetype{}
	tm.Id, _ = this.GetInt("Id")
	tm.Date, _ = time.Parse("2006-01-02 15:04:05", strings.Trim(this.GetString("Date"), " "))
	tm.Describe = this.GetString("Describe")
	tm.IsShow, _ = this.GetInt("IsShow")
	tm.ImageUrl = this.GetString("ImageUrl")
	tm.OrderTop, _ = this.GetInt("OrderTop")
	tm.ParentId, _ = this.GetInt("ParentId")
	tm.Remark = this.GetString("Remark")
	tm.Title = this.GetString("Title")
	if tm.Id == 0 { //添加
		tm.Date, _ = time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
		id, err := dalbt.Add(&tm)
		if err == nil {
			this.Data["Id"] = id
			this.Data["alert"] = dlg.Tips("恭喜你,添加成功！", 3, "succ.png")
			//lh.InsertLog(this.Ctx, lh.ADD, "Tb_Basetype添加 id:"+con.Itoa(int(id))) //日志记录
		} else {
			this.Data["alert"] = dlg.Tips("对不起,添加失败！", 5, "fail.png")
		}

	} else { //编辑
		if num, _ := dalbt.Update(tm); num > 0 {
			this.Data["Id"] = tm.Id
			this.Data["alert"] = dlg.Tips("恭喜你,编辑成功！", 3, "succ.png")
		} else {
			this.Data["Id"] = tm.Id
			this.Data["alert"] = dlg.Tips("对不起,编辑失败！", 5, "fail.png")
		}
	}
	this.Data["ParentId"] = tm.ParentId
	this.Data["IsShow"] = tm.IsShow
	this.Data["Title"] = tm.Title
	this.Data["Date"] = tm.Date.Format("2006-01-02 15:04:05")
	this.Data["OrderTop"] = tm.OrderTop
	this.Data["ImageUrl"] = tm.ImageUrl
	this.Data["Describe"] = tm.Describe
	this.Data["Remark"] = tm.Remark
	if tm.ParentId == 0 {
		this.Data["ParentName"] = "--无父级菜单--"
	} else {
		this.Data["ParentName"] = dalbt.GetById(tm.ParentId).Title
	}
	this.Data["strParentMenuList"] = dalbt.GetZtreeJsonByFid(menuRootId)
	this.TplNames = "admin/basetype/edit.html"
}

/**
 * ajax,动态加载treetable
 */
func (this *BasetypeController) GetTrList() {
	str := ""
	pid, _ := this.GetInt("pId")
	menuRootId, _ := this.GetInt("menuRootId")
	all := dalbt.GetAllChildSlices(0)
	fmt.Println("-------------------+++++++++++++++++++++++++l", all)
	var l []m.Tb_basetype
	l = dalbt.GetOnlyChildSlices(pid)
	for _, v := range l {
		str += dalbt.GetTrById(v.Id, menuRootId, all)
	}
	this.Ctx.WriteString(str)
}

/**
 * 用于ajax删除
 */
func (this *BasetypeController) Remove() {
	ids := this.GetString("ids")
	var num int64 = 0
	if len(ids) > 0 {
		num = dalbt.Delete(ids)
		//lh.InsertLog(this.Ctx, lh.DELETE, "Tb_Basetype删除,ids,"+ids) //日志记录
	}
	this.Ctx.WriteString(con.Itoa(int(num)))
}
