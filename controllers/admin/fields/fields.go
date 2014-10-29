package fields

import (
	"github.com/astaxie/beego"
	dalbt "mcms/dal/basetype"
	"mcms/dal/fields"
	"mcms/models"
	dlg "mcms/util/lhdialog"
	"mcms/util/pager"
	"strings"
	"time"
)

type FieldsController struct {
	beego.Controller
}

/**
 * 返回多条字段记录
 * @param
 * @return
 * @author tongfei
 * @date 2014-08-29
 * @version v1.0
 */
func (this *FieldsController) List() {
	typeid := this.GetString("typeid")

	cur_page, _ := this.GetInt("pno")
	if cur_page == 0 {
		cur_page = 1
	}
	var po pager.PageOptions
	po.TableName = "tb_fields"
	po.Currentpage = int(cur_page)
	po.EnableFirstLastLink = true
	po.EnablePreNexLink = true
	po.Conditions = " and typeid = " + typeid + " order by resort, id desc, add_time "
	totalItem, _, rs, pagerhtml := pager.GetPagerLinks(&po, this.Ctx)
	var fields []models.Tb_fields
	rs.QueryRows(&fields)
	this.Data["Fields"] = fields
	this.Data["pagerhtml"] = pagerhtml
	this.Data["PageSize"] = po.PageSize
	this.Data["totalItem"] = totalItem

	// 产品typeid=1 文章typeid=2
	this.Data["Typeid"] = typeid

	this.TplNames = "admin/fields/list.html"
}

/**
 * 添加一条字段记录
 * @param
 * @return
 * @author tongfei
 * @date 2014-08-29
 * @version v1.0
 */
func (this *FieldsController) Add() {
	var alert_num int
	var alert_msg string

	if this.Ctx.Input.Method() == "POST" {
		var Fields models.Tb_fields

		if err := this.ParseForm(&Fields); err != nil {
			alert_msg = "对不起,数据格式化失败！"
			alert_num = 5
		} else {
			Fields.AddTime, _ = time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
			_, err = fields.Add(Fields)

			if err == nil {
				alert_msg = "恭喜你,添加成功！"
				alert_num = 3
			} else {
				alert_msg = "对不起,添加失败！"
				alert_num = 5
			}
		}
	}

	if alert_num == 3 {
		this.Data["alert"] = dlg.Tips(alert_msg, 3, "succ.png")
	} else if alert_num == 5 {
		this.Data["alert"] = dlg.Tips(alert_msg, 5, "fail.png")
	}

	typeid := this.GetString("typeid")
	var topid int64

	if typeid == "1" {
		topid, _ = beego.AppConfig.Int64("ProTopCategoryId")
	} else if typeid == "2" {
		topid, _ = beego.AppConfig.Int64("ArticleTopCatid")
	}
	// 顶级分类id
	this.Data["Topid"] = topid
	// 顶级分类名称
	this.Data["CatName"] = dalbt.GetById(topid).Title
	this.Data["Typeid"] = typeid
	this.Data["List"] = dalbt.GetZtreeJsonByFid(topid)
	this.TplNames = "admin/fields/add.html"
}

/**
 * 更新一条字段记录
 * @param
 * @return
 * @author tongfei
 * @date 2014-08-29
 * @version v1.0
 */
func (this *FieldsController) Update() {
	var Fields models.Tb_fields

	if this.Ctx.Input.Method() == "POST" {

		if err := this.ParseForm(&Fields); err != nil {
			this.Data["alert"] = dlg.Tips("对不起，数据格式化失败", 5, "fail.png")
		}

		_, err := fields.Update(Fields)

		if err != nil {
			this.Data["alert"] = dlg.Tips("对不起，更新失败", 5, "fail.png")
		} else {
			this.Data["alert"] = dlg.Tips("更新成功！", 3, "succ.png")
		}

		this.Data["Fields"] = Fields

	} else {
		id, _ := this.GetInt("id")

		var err error
		Fields, err = fields.GetById(id)

		if err != nil {
			this.Data["alert"] = dlg.Tips("对不起，读取数据失败", 5, "fail.png")
		}

		this.Data["CatName"] = dalbt.GetById(Fields.Catid).Title

		this.Data["Fields"] = Fields
	}

	var topid int64

	if Fields.Typeid == 1 {
		topid, _ = beego.AppConfig.Int64("ProTopCategoryId")
	} else if Fields.Typeid == 2 {
		topid, _ = beego.AppConfig.Int64("ArticleTopCatid")
	}

	this.Data["List"] = dalbt.GetZtreeJsonByFid(topid)
	this.TplNames = "admin/fields/update.html"
}

/**
 * 更新一条字段排序
 * @param
 * @return
 * @author tongfei
 * @date 2014-08-29
 * @version v1.0
 */
func (this *FieldsController) Resort() {
	id, _ := this.GetInt("fid")
	resort, _ := this.GetInt("fresort")

	var Fields models.Tb_fields

	Fields, _ = fields.GetById(id)
	Fields.Resort = resort
	_, err := fields.Update(Fields)

	if err == nil {
		this.Ctx.WriteString("y")
	} else {
		this.Ctx.WriteString("n")
	}

}

/**
 * ajax删除多条字段记录
 * @param
 * @return
 * @author tongfei
 * @date 2014-08-29
 * @version v1.0
 */
func (this *FieldsController) JsonDelMulti() {
	mult_ids := strings.Split(this.GetString("ids"), ",")
	err := fields.DeleteMult(mult_ids)

	if err != nil {
		this.Ctx.WriteString("n")
	} else {
		this.Ctx.WriteString("y")
	}
}
