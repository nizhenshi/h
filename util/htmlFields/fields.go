// +----------------------------------------------------------------------
// | 用于生成扩展字段和调查问卷的html代码
// +----------------------------------------------------------------------
// | golang version: go1.3
// +----------------------------------------------------------------------
// | beego version: beego 1.3.2 (beego framework)
// +----------------------------------------------------------------------
// | Author: tongfei
// +----------------------------------------------------------------------

package htmlFields

import (
	"strconv"
)

/**
 * 生成一个文本框的html代码
 * @param title 文本框标题，name 文本框name属性，selectedValue 默认值，extAttr 扩展属性
 * @return html_code html代码
 * @author tongfei
 * @date 2014-08-27
 * @version v1.0
 */
func HtmlText(title string, name string, selectedValue string, extAttr string) string {
	var html_code string = ""

	start, end := HtmlBase(title)
	html_code += start
	html_code += "	<input type=\"text\" class=\"input normal extendFields\" name=\"" + name + "\" value=\"" + selectedValue + "\"" + extAttr + "/>"
	html_code += end

	return html_code
}

/**
 * 生成一个文本域的html代码
 * @param title 标题，name name属性，selectedValue 默认值，extAttr 扩展属性
 * @return html_code html代码
 * @author tongfei
 * @date 2014-08-27
 * @version v1.0
 */
func HtmlTextarea(title string, name string, selectedValue string, extAttr string) string {
	var html_code string = ""

	start, end := HtmlBase(title)
	html_code += start
	html_code += "<textarea name=\"" + name + "\" class=\"span6" + "\"" + extAttr + ">" + selectedValue + "</textarea>"
	html_code += end

	return html_code
}

/**
 * 生成一个上传图片的html代码
 * @param title 标题，name name属性，extAttr 扩展属性
 * @return html_code html代码
 * @author tongfei
 * @date 2014-08-27
 * @version v1.0
 */

func HtmlThumb(title string, name string, extAttr string) string {
	var html_code string = ""

	start, end := HtmlBase(title)
	html_code += start
	html_code += "<input type=\"file\" name=\"" + name + "\"" + extAttr + "/>"
	html_code += end

	return html_code
}

/**
 * 生成一个上传文件的html代码
 * @param title 标题，name name属性，extAttr 扩展属性
 * @return html_code html代码
 * @author tongfei
 * @date 2014-08-27
 * @version v1.0
 */
func HtmlFile(title string, name string, extAttr string) string {
	var html_code string = ""
	start, end := HtmlBase(title)
	html_code += start
	// 上传文件
	html_code += end

	return html_code
}

/**
 * 生成一个编辑器的html代码
 * @param title 标题，name name属性，extAttr 扩展属性
 * @return html_code html代码
 * @author tongfei
 * @date 2014-08-27
 * @version v1.0
 */
func HtmlEditor(title string, name string, extAttr string) string {
	var html_code string = ""
	start, end := HtmlBase(title)
	html_code += start
	html_code += "<script type=\"text/plain\"  id=\"" + name + "\" name=\"" + name + "\"></script>"
	html_code += "<script type=\"text/javascript\">"
	html_code += "var options = {\"fileUrl\":\"/web/ue/upload\",\"imageUrl\":\"/web/ue/upload\",\"filePath\":\"\",\"imagePath\":\"\",\"initialFrameWidth\":\"90%\",\"initialFrameHeight\":\"200\",\"initialContent\":\"\"};"
	html_code += "var ue = UE.getEditor(\"" + name + "\", options);"
	html_code += "</script>"
	html_code += end

	return html_code
}

/**
 * 生成一个日期的html代码
 * @param title 标题，name name属性，extAttr 扩展属性
 * @return html_code html代码
 * @author tongfei
 * @date 2014-08-27
 * @version v1.0
 */
func HtmlDate(title string, name string, extAttr string) string {
	var html_code string = ""

	return html_code
}

/**
 * 生成一个下拉列表的html代码
 * @param title 标题，name name属性，arr 选项值的切片，selectedValue 默认值，extAttr 扩展属性，isTitle 是否有选项提示标题
 * @return html_code html代码
 * @author tongfei
 * @date 2014-08-27
 * @version v1.0
 */

func HtmlSelect(title string, name string, arr []string, selectedValue int64, extAttr string, isTitle bool) string {
	var html_code string = ""

	start, end := HtmlBase(title)
	html_code += start
	var type_item string

	if isTitle {
		type_item += "<option value=\"0\">请选择</option>"
	}

	for k, v := range arr {
		k++

		if int64(k) == selectedValue {
			type_item += "<option selected=\"selected\" value=\"" + strconv.Itoa(k) + "\">" + v + "</option>"
		} else {
			type_item += "<option value=\"" + strconv.Itoa(k) + "\">" + v + "</option>"
		}
	}

	html_code += "</select>"

	html_code += "<div class=\"rule-single-select\">"
	html_code += "<select name=\"" + name + "\" id=\"ddl\"" + name + " style=\"width:180px;height:30px\">"
	html_code += type_item
	html_code += "</select>"
	html_code += "</div>"
	html_code += end
	return html_code
}

/**
 * 生成一个复选框的html代码
 * @param title 标题，name name属性，arr 选项值的切片， arrLevel 默认值切片，extAttr 扩展属性
 * @return html_code html代码
 * @author tongfei
 * @date 2014-08-27
 * @version v1.0
 */
func HtmlCheckBox(title string, name string, arr []string, arrLevel []string, extAttr string) string {
	var html_code string = ""

	start, end := HtmlBase(title)
	html_code += start
	var type_item string = ""
	var a_item string = ""
	var is_add bool

	for k, v := range arr {
		k++
		var checkedAttr string = ""
		var selectedAttr string = ""

		if len(arrLevel) > int(0) {
			is_add = true

			for _, vv := range arrLevel {
				vv_int, _ := strconv.Atoi(vv)

				if k == vv_int {
					checkedAttr = " checked=\"checked\" "
					selectedAttr = " selected "
				}
			}

		} else {
			is_add = true
		}

		if is_add {
			a_item += "<a class=\"extendAitem " + selectedAttr + "\" connType=\"checkbox\" connName=\"" + name + "_cb\" href=\"javascript:;\">" + v + "</a>"
		}

		type_item += "<input class=\"" + name + "_cb\"" + checkedAttr + "type=\"checkbox\" value=\"" + strconv.Itoa(k) + "\" name=\"" + name + "\"" + extAttr + "/>"
		type_item += "<label>" + v + "</label>"
	}

	html_code += "<div class=\"rule-multi-checkbox multi-checkbox\">"
	html_code += "<div class=\"boxwrap extendDiv \">"
	html_code += a_item
	html_code += "</div>"
	html_code += "<span id=\"cbl" + name + "\" style=\"display: none;\">"
	html_code += type_item
	html_code += "</span>"
	html_code += "</div>"
	html_code += end

	return html_code
}

/**
 * 生成一个单选按钮的html代码
 * @param title 标题，name name属性，arr 选项值的切片，selectedValue 默认值，extAttr 扩展属性
 * @return html_code html代码
 * @author tongfei
 * @date 2014-08-27
 * @version v1.0
 */
func HtmlRadio(title string, name string, arr []string, selectedValue int64, extAttr string) string {
	var html_code string = ""

	start, end := HtmlBase(title)
	html_code += start
	var type_item string = ""
	var a_item string = ""

	for k, v := range arr {
		k++

		if int64(k) == selectedValue {
			a_item += "<a class=\"extendAitem selected\" connType=\"radio\" connName=\"" + name + "_rb\" href=\"javascript:;\">" + v + "</a>"
			type_item += "<input class=\"" + name + "_rb\" checked=\"checked\" type=\"radio\" value=\"" + strconv.Itoa(k) + "\" name=\"" + name + "\"" + extAttr + " />"

		} else {
			a_item += "<a class=\"extendAitem\" connType=\"radio\" connName=\"" + name + "_rb\" href=\"javascript:;\">" + v + "</a>"
			type_item += "<input class=\"" + name + "_rb\" type=\"radio\" value=\"" + strconv.Itoa(k) + "\" name=\"" + name + "\"" + extAttr + " />"
		}

		type_item += "<label>" + v + "</label>"
	}

	html_code += "<div class=\"rule-multi-radio multi-radio\">"
	html_code += "<div class=\"boxwrap extendDiv \">"
	html_code += a_item
	html_code += "</div>"
	html_code += "<span id=\"rbl" + name + "\" style=\"display: none;\">"
	html_code += type_item
	html_code += "</span>"
	html_code += "</div>"
	html_code += end
	html_code += end

	return html_code
}

/**
 * 生成一个html代码公共代码
 * @param title 代码元素的标题
 * @return start 公共头部，end 公共尾部
 * @author tongfei
 * @date 2014-08-27
 * @version v1.0
 */
func HtmlBase(title string) (string, string) {
	var start string = ""
	var end string = ""

	start += "<dl>"
	start += " <dt>" + title + "</dt>"
	start += "  <dd>"

	end += " </dd>"
	end += "</dl>"

	return start, end
}
