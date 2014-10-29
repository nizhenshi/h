package article

import (
	"github.com/astaxie/beego"
	"math/rand"
	"mcms/dal/article"
	dalbt "mcms/dal/basetype"
	"mcms/dal/fields"
	fv "mcms/dal/fieldsValue"
	"mcms/models"
	"mcms/util/pager"
	"strconv"
	"strings"
	"time"
	dlg "ywl/util/lhdialog"
)

type ArticleController struct {
	beego.Controller
}

func (this *ArticleController) Add() {
	var postCat, cat_id int64
	if this.Ctx.Input.Method() == "POST" {
		var Article models.Tb_article
		if err := this.ParseForm(&Article); err != nil {
			this.Data["alert"] = dlg.Tips("对不起，数据格式化失败！", 5, "fail.png")
			return
		}
		//获取分类，用于获取该分类下的扩展字段
		catid := Article.Catid
		postCat = catid
		Article.Typeid = 1
		Article.Userid = 1
		Article.EditTime, _ = time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
		Article.AddTime, _ = time.Parse("2006-01-02 15:04:05", strings.Trim(this.GetString("addtime"), " "))
		arr_level := this.GetStrings("Level")
		Article.Level = strings.Join(arr_level, ",")
		itemid, err := article.Add(Article)
		if err != nil {
			this.Data["alert"] = dlg.Tips("对不起，添加失败", 5, "fail.png")
			return
		}
		var Fields []models.Tb_fields // 扩展字段实体
		var FieldsVal []models.Tb_fields_value

		Fields, _ = fields.List(catid)

		for _, v := range Fields {
			var Fie models.Tb_fields_value
			var tmp_field string
			tmp_field = v.Name

			Fie.Catid = catid
			Fie.Fieldid = v.Id
			Fie.Itemid = itemid
			if v.HtmlType == "checkbox" {
				Fie.FieldValue = strings.Join(this.GetStrings(tmp_field), "#")
			} else if v.HtmlType == "radio" {
				Fie.FieldValue = this.GetString(tmp_field)
			} else {
				Fie.FieldValue = this.GetString(tmp_field)
			}

			FieldsVal = append(FieldsVal, Fie)
		}

		// 多条插入扩展字段值
		_, err = fv.AddMult(FieldsVal)

		if err != nil {
			this.Data["alert"] = dlg.Tips("对不起，添加失败", 5, "fail.png")
		} else {
			this.Data["alert"] = dlg.Tips("添加成功！", 3, "succ.png")
		}
	}

	this.Data["LevelArr"] = this.GetLevelArr()
	this.Data["StatusArr"] = this.GetStatusArr()
	this.Data["AddTime"] = time.Now().Format("2006-01-02 15:04:05")
	if postCat == 0 {
		cat_id, _ = this.GetInt("catid")
	} else {
		cat_id = postCat
	}
	this.Data["Catid"] = cat_id

	this.Data["Hits"] = 50 + rand.Intn(150)
	this.Data["CatName"] = dalbt.GetById(cat_id).Title
	this.Data["List"] = dalbt.GetZtreeJsonByFid(1)

	this.TplNames = "admin/article/add.html"
}

func (this *ArticleController) List() {
	catid, _ := this.GetInt("catid")

	allCatidSons := dalbt.GetAllChildSlices(catid)

	var allSons string
	for _, v := range allCatidSons {
		allSons += strconv.Itoa(int(v.Id)) + ","
	}
	allSons += strconv.Itoa(int(catid))
	cur_page, _ := this.GetInt("pno")

	if cur_page == 0 {
		cur_page = 1
	}
	var po pager.PageOptions
	po.TableName = "tb_article"
	po.Currentpage = int(cur_page)
	po.EnableFirstLastLink = true
	po.EnablePreNexLink = true
	po.Conditions = " and catid in(" + allSons + ") order by id desc, add_time "
	totalItem, _, rs, pagerhtml := pager.GetPagerLinks(&po, this.Ctx)
	var Articles []models.Tb_article
	rs.QueryRows(&Articles)
	this.Data["Articles"] = Articles
	this.Data["pagerhtml"] = pagerhtml
	this.Data["PageSize"] = po.PageSize
	this.Data["totalItem"] = totalItem

	this.Data["StatusArr"] = this.GetLevelArr()
	// 信息分类的顶级id
	this.Data["Catid"] = catid
	this.Data["CatName"] = dalbt.GetById(catid).Title
	this.Data["List"] = dalbt.GetZtreeJsonByFid(catid)

	this.TplNames = "admin/article/list.html"
}

func (this *ArticleController) Update() {
	this.Data["LevelArr"] = this.GetLevelArr()
	this.Data["StatusArr"] = this.GetStatusArr()

	if this.Ctx.Input.Method() == "POST" {
		// 实例化对象
		var Article models.Tb_article
		// 赋值
		if err := this.ParseForm(&Article); err != nil {
			this.Data["alert"] = dlg.Tips("对不起，数据序列化失败", 5, "fail.png")
		}

		// 获取分类，用于查询该分类的扩展字段
		catid := Article.Catid
		itemid := Article.Id

		Article.EditTime, _ = time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"))
		Article.AddTime, _ = time.Parse("2006-01-02 15:04:05", strings.Trim(this.GetString("AddTime"), " "))
		_, err := article.Update(Article)

		// 查询分类下的扩展字段
		var Fields []models.Tb_fields // 扩展字段实体
		var FieldsVal []models.Tb_fields_value

		Fields, _ = fields.List(catid)

		for _, v := range Fields {
			var Fie models.Tb_fields_value
			var tmp_field string
			tmp_field = v.Name

			Fie.Catid = catid
			Fie.Fieldid = v.Id
			Fie.Itemid = itemid
			if v.HtmlType == "checkbox" {
				Fie.FieldValue = strings.Join(this.GetStrings(tmp_field), "#")
			} else if v.HtmlType == "radio" {
				Fie.FieldValue = this.GetString(tmp_field)
			} else {
				Fie.FieldValue = this.GetString(tmp_field)
			}

			FieldsVal = append(FieldsVal, Fie)
		}

		// 删除旧信息
		err = fv.Delete(catid, itemid)

		if err != nil {
			this.Data["alert"] = dlg.Tips("对不起，更新失败", 5, "fail.png")
			return
		}
		// 多条插入扩展字段值
		_, err = fv.AddMult(FieldsVal)

		if err != nil {
			this.Data["alert"] = dlg.Tips("对不起，更新失败", 5, "fail.png")
		} else {
			this.Data["alert"] = dlg.Tips("更新成功！", 3, "succ.png")
		}

		this.Data["Status"] = strconv.Itoa(int(Article.Status))
		//this.Data["HtmlFields"] = fields.GetFieldsHtml(Article.Catid, Article.Id)
		this.Data["CatName"] = dalbt.GetById(Article.Catid).Title

		if dalbt.JudgeIdInSlices(Article.Catid, dalbt.GetAllChildSlices(14)) {
			this.Data["List"] = dalbt.GetZtreeJsonByFid(14)
		} else if dalbt.JudgeIdInSlices(Article.Catid, dalbt.GetAllChildSlices(332)) {
			this.Data["List"] = dalbt.GetZtreeJsonByFid(332)
		}

		this.Data["Article"] = Article

	} else {
		id, _ := this.GetInt("id")
		var Article models.Tb_article

		Article, _ = article.GetById(id)
		// 获取分类catid
		catid := Article.Catid
		this.Data["Status"] = strconv.Itoa(int(Article.Status))
		this.Data["CheckedLevel"] = strings.Split(Article.Level, ",")
		this.Data["HtmlFields"] = fields.GetFieldsHtml(catid, id)
		this.Data["CatName"] = dalbt.GetById(catid).Title

		if dalbt.JudgeIdInSlices(Article.Catid, dalbt.GetAllChildSlices(14)) {
			this.Data["List"] = dalbt.GetZtreeJsonByFid(14)
		} else if dalbt.JudgeIdInSlices(Article.Catid, dalbt.GetAllChildSlices(332)) {
			this.Data["List"] = dalbt.GetZtreeJsonByFid(332)
		}

		this.Data["Article"] = Article
	}
	this.TplNames = "admin/article/update.html"
}

/**
 * ajax删除多条信息记录
 * @param
 * @return
 * @author tongfei
 * @date 2014-08-29
 * @version v1.0
 */
func (this *ArticleController) JsonDelMulti() {
	mult_ids := strings.Split(this.GetString("ids"), ",")
	err := article.DeleteMult(mult_ids)

	if err != nil {
		this.Ctx.WriteString("n")
	} else {
		this.Ctx.WriteString("y")
	}
}

/**
 * ajax动态加载字段
 * @param
 * @return
 * @author tongfei
 * @date 2014-08-29
 * @version v1.0
 */
func (this *ArticleController) JsonFields() {
	id, _ := this.GetInt("fid")
	cat_id, _ := this.GetInt("catid")
	html_fields := fields.GetFieldsHtml(cat_id, id)
	this.Ctx.WriteString(string(html_fields))
}

/**
 * 返回信息状态的切片
 * @param
 * @return LevelArr 级别切片
 * @author tongfei
 * @date 2014-08-29
 * @version v1.0
 */
func (this *ArticleController) GetLevelArr() []map[string]string {
	level_arr := beego.AppConfig.Strings("ArticleLevelStatus")
	var arrLevel []map[string]string

	for _, v := range level_arr {
		var temp_map map[string]string
		temp_map = make(map[string]string)

		item := strings.Split(v, ":")
		temp_map["Key"] = item[0]
		temp_map["Value"] = item[1]

		arrLevel = append(arrLevel, temp_map)
	}

	return arrLevel
}

/**
 * 返回信息 审核的切片
 * @param
 * @return StatusArr 级别切片
 * @author tongfei
 * @date 2014-08-29
 * @version v1.0
 */
func (this *ArticleController) GetStatusArr() []map[string]string {
	status_arr := beego.AppConfig.Strings("ArtcileCheckStatus")
	var sattusLevel []map[string]string

	for _, v := range status_arr {
		var temp_map map[string]string
		temp_map = make(map[string]string)

		item := strings.Split(v, ":")
		temp_map["Key"] = item[0]
		temp_map["Value"] = item[1]

		sattusLevel = append(sattusLevel, temp_map)
	}

	return sattusLevel
}
