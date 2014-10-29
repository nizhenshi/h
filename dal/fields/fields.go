package fields

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	html "html/template"
	fv "mcms/dal/fieldsValue"
	"mcms/models"
	tpl "mcms/util/htmlFields"
	"strconv"
	"strings"
)

/**
 * 添加一个扩展字段
 * @param fields 字段对象
 * @return id 字段id，error 错误信息
 * @author tongfei
 * @date 2014-08-29
 * @version v1.0
 */
func Add(fields models.Tb_fields) (int64, error) {
	//实例化Orm对象
	o := orm.NewOrm()
	id, err := o.Insert(&fields)

	if err != nil {
		return 0, err
	}

	return id, nil
}

/**
 * 更新一个扩展字段
 * @param fields 字段对象
 * @return num 影响的行数，error 错误信息
 * @author tongfei
 * @date 2014-08-29
 * @version v1.0
 */
func Update(fields models.Tb_fields) (int64, error) {
	o := orm.NewOrm()
	num, err := o.Update(&fields)

	if err != nil {
		return 0, err
	}

	return num, nil
}

/**
 * 获取一个扩展字段对象
 * @param id 字段id
 * @return Fields 字段对象，error 错误信息
 * @author tongfei
 * @date 2014-08-29
 * @version v1.0
 */
func GetById(id int64) (models.Tb_fields, error) {
	o := orm.NewOrm()
	var Fields, FieNil models.Tb_fields
	Fields.Id = id
	err := o.Read(&Fields)

	if err != nil {
		return FieNil, err
	}

	return Fields, nil
}

/**
 * 获取扩展字段值的map
 * @param catid 分类, itemid 所属元素id
 * @return Fields 字段对象，error 错误信息
 * @author tongfei
 * @date 2014-08-29
 * @version v1.0
 */
func GetFValByIdItemid(catid, itemid int64) (map[string]string, error) {
	o := orm.NewOrm()
	Fields := make([]models.Tb_fields, 0)
	orm_art := o.QueryTable("tb_fields")
	orm_art.Filter("catid", catid).All(&Fields)
	fieVals := make(map[string]string)

	for _, v := range Fields {
		var fiename string
		fiename = v.Name
		fieVals[fiename] = fv.GetFieldValue(catid, itemid, v.Id)
	}

	return fieVals, nil
}

/**
 * 获取一个扩展字段值
 * @param catid 分类, itemid 所属元素id
 * @return Fields 字段对象，error 错误信息
 * @author tongfei
 * @date 2014-08-29
 * @version v1.0
 */
func GetOneFVal(catid, itemid int64, name string) string {
	fieMap, _ := GetFValByIdItemid(catid, itemid)

	return fieMap[name]
}

/**
 * 获取字段对象切片
 * @param catid 字段所属分类的id
 * @return fies 字段对象，error 错误信息
 * @author tongfei
 * @date 2014-08-29
 * @version v1.0
 */
func List(catid int64) ([]models.Tb_fields, error) {
	o := orm.NewOrm()
	Fields := new(models.Tb_fields)
	tb_fields := Fields.TableName()
	orm_fie := o.QueryTable(tb_fields)

	if catid > 0 {
		orm_fie = orm_fie.Filter("catid", catid)
	}

	fies := make([]models.Tb_fields, 0)
	_, err := orm_fie.OrderBy("-resort", "-addtime").All(&fies)

	if err != nil {
		return nil, err
	}

	return fies, nil
}

/**
 * 删除一个字段记录
 * @param id 字段id
 * @return error 错误信息
 * @author tongfei
 * @date 2014-08-29
 * @version v1.0
 */
func Delete(id int64) error {
	o := orm.NewOrm()
	var Fields models.Tb_fields
	Fields.Id = id

	if _, err := o.Delete(&Fields); err != nil {
		return err
	}

	return nil
}

/**
 * 删除多个字段记录
 * @param ids 字符串切片
 * @return error 错误信息
 * @author tongfei
 * @date 2014-08-29
 * @version v1.0
 */
func DeleteMult(ids []string) error {
	o := orm.NewOrm()
	Fields := new(models.Tb_fields)
	tb_fields := Fields.TableName()
	orm_fie := o.QueryTable(tb_fields)

	if _, err := orm_fie.Filter("id__in", ids).Delete(); err != nil {
		return err
	}

	return nil
}

/**
 * 获取特定分类下的字段的html代码
 * @param catid 字段所属分类的id，itemid 所属的信息(商品)id
 * @return html_fields html模板代码
 * @author tongfei
 * @date 2014-08-29
 * @version v1.0
 */
func GetFieldsHtml(catid int64, itemid int64) html.HTML {
	var Fields []models.Tb_fields
	var err error
	var html_fields string = ""

	Fields, err = List(catid)

	if err != nil {

	}

	for _, v := range Fields {
		var html_code, type_value, extAttr string

		if itemid == int64(0) {
			type_value = v.DefaultValue
		} else {
			type_value = fv.GetFieldValue(catid, itemid, v.Id)
		}

		if v.InputLimit != "" {
			extAttr += " datatype=\"" + v.InputLimit + "\" "
		}
		if v.Note != "" {
			extAttr += " errormsg=\"" + v.Note + "\" "
		}

		extAttr += " fid = \"" + strconv.Itoa(int(v.Id)) + "\" "

		switch v.HtmlType {

		case "text":
			html_code = tpl.HtmlText(v.Title, v.Name, type_value, extAttr)

		case "thumb":
			html_code = tpl.HtmlThumb(v.Title, v.Name, extAttr)

		case "textarea":
			html_code = tpl.HtmlTextarea(v.Title, v.Name, type_value, extAttr)

		case "select":
			arr := strings.Split(v.OptionValue, "#")
			int_value, _ := strconv.Atoi(type_value)
			html_code = tpl.HtmlSelect(v.Title, v.Name, arr, int64(int_value), extAttr, true)

		case "checkbox":
			arr := strings.Split(v.OptionValue, "#")
			arr_level := strings.Split(type_value, "#")
			html_code = tpl.HtmlCheckBox(v.Title, v.Name, arr, arr_level, extAttr)
		case "radio":
			arr := strings.Split(v.OptionValue, "#")
			int_value, _ := strconv.Atoi(type_value)
			html_code = tpl.HtmlRadio(v.Title, v.Name, arr, int64(int_value), extAttr)

		case "editor":
			html_code = tpl.HtmlEditor(v.Title, v.Name, extAttr)

		default:

		}

		html_fields += html_code
	}

	return html.HTML(html_fields)
}

/**
 * 查看name字段相同的个数   用来ajax验证其唯一性
 */
func AjaxNameUnique(name string) (int64, error) {
	o := orm.NewOrm()
	num, err := o.QueryTable("Tb_fields").Filter("Name", name).Count()
	return num, err
}

/**
 * 判断是发威产品的通用属性
 */
func JudgeIsProCommon(name string) (int64, error) {
	o := orm.NewOrm()
	topid, _ := beego.AppConfig.Int64("ProTopExtendCategoryId")
	num, err := o.QueryTable("Tb_fields").Filter("Name", name).Filter("Catid", topid).Count()
	return num, err
}

func init() {
	beego.AddFuncMap("GetOneFVal", GetOneFVal)
}
