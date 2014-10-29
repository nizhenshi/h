package html

import (
	"fmt"
	html "html/template"
	con "strconv"
	m "ywl/models"
	"ywl/util/config"
)

/**
 * 生成一个下拉列表,根据Tb_basetype类别数组
 * @arr 数据源
 * @selectValue  默认选择的值
 * @extString 扩展内容 比如" name='xxx'  id='aa'  style='width:10px;' "
 * @boolDefault 是否添加默认项,即:添加一个[--请选择--]项,值默认为0
 * 返回一个html字符串
 */
func SelectInput(arr []m.Tb_basetype, selectValue int64, extString string, boolDefault bool) html.HTML {
	var rs string = "<select " + extString + ">"
	if boolDefault {
		rs += "<option value='0'>--请选择--</option>"
	}
	for _, v := range arr {
		if selectValue == v.Id {
			rs += "<option selected value=" + con.Itoa(int(v.Id)) + ">" + v.Title + "</option>"
		} else {
			rs += "<option value=" + con.Itoa(int(v.Id)) + ">" + v.Title + "</option>"
		}
	}
	rs += "</select>"
	return html.HTML(rs)
}

/**
 * 生成一个下拉列表,根据util/config包的GetKvArray方法返回的类别数组
 */
func ConfigSelectInput(arr []config.Kv, selectValue string, extString string, boolDefault bool) html.HTML {
	var rs string = "<select " + extString + ">"
	if boolDefault {
		rs += "<option value='0'>--请选择--</option>"
	}
	for _, v := range arr {
		if selectValue == v.Value {
			rs += "<option selected value=" + v.Value + ">" + v.Key + "</option>"
		} else {
			rs += "<option value=" + v.Value + ">" + v.Key + "</option>"
		}
	}
	rs += "</select>"
	return html.HTML(rs)
}

func DataToggleCheckbox() {

}

func MultipleSelectinput() {

}

func DataRadioButton() {

}

func init() {
	fmt.Println("----------------------------")
}
