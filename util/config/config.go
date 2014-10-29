/**
 * 作用:配置文件操作
 * 日期:2014-8-26
 * 支飞亚
 */

package config

import (
	"fmt"
	"github.com/astaxie/beego"
	html "html/template"
	con "strconv"
	"strings"
)

/**
 * 键值对组合结构
 */
type Kv struct {
	Key   string
	Value string
}

/**
 * 作用: 比如 配置文件中有如下信息:classType=提交:0;付款:1;发货:2;結束:3;退单:4;作废:5
 * 调用该方法   GetKvArray(classType)
 * 返回一个Kv结构的数组    该数组中,每个元素的Key为冒号前部分的,Value为冒号后部分的值
 */
func GetKvArray(name string) []Kv {
	fmt.Println("GetKvArray")
	var rs []Kv
	arr := beego.AppConfig.Strings(name)
	if len(arr) > 1 {
		for _, s := range arr {
			var temp Kv
			temp.Key = strings.Split(s, ":")[0]
			temp.Value = strings.Split(s, ":")[1]
			rs = append(rs, temp)
		}
	}
	return rs
}

/*
获取 默认时间 格式
*/
func GetDefaultTimeFomart() string {
	return beego.AppConfig.String("DefaultTimeFormat")
}

/// <summary>
/// 获取操作权限
/// </summary>
/// <returns>Dictionary</returns>
func Tb_Navigation_ActionType() map[string]string {
	dic := map[string]string{}
	dic["Show"] = "显示"
	dic["View"] = "查看"
	dic["Add"] = "添加"
	dic["Edit"] = "修改"
	dic["Delete"] = "删除"
	return dic
}

/**
 * [多选]获取推荐状态的html,并判断他们的选中状态,参数recom为在数据库中存放的推荐状态  形如:  TOP|HOT|NEWEST
 * 参数1:配置文件的数组字符串KEY  参数2:在数据库中存放的推荐状态  形如:  TOP|HOT|NEWEST  参数3:生成的input标签的name
 */
func GetRecommendStatus(StatusString, recom string, tagName string) html.HTML {
	var rs string
	arr := GetKvArray(StatusString)
	for _, v := range arr {
		var checked string
		if strings.Contains(strings.ToLower(recom), strings.ToLower(v.Value)) {
			checked = "checked=\"checked\""
		} else {
		}
		rs += "<input  type=\"checkbox\" value=\"" + v.Value + "\" name=\"" + tagName + "\" " + checked + " /><label>" + v.Key + "</label>"
	}
	return html.HTML(rs)
}

/**
 * [单选]获取推荐状态的html,并判断他们的选中状态,参数recom为在数据库中存放的推荐状态  形如:  TOP|HOT|NEWEST
 * 参数1:配置文件的数组字符串KEY  参数2:在数据库中存放的推荐状态  形如:  TOP|HOT|NEWEST  参数3:生成的input标签的name
 */
func GetRadioRecommendStatus(StatusString, recom string, tagName string) html.HTML {
	var rs string
	arr := GetKvArray(StatusString)
	for _, v := range arr {
		var checked string
		if strings.Contains(strings.ToLower(recom), strings.ToLower(v.Value)) {
			checked = "checked=\"checked\""
		} else {
		}
		rs += "<input  type=\"radio\" value=\"" + v.Value + "\" name=\"" + tagName + "\" " + checked + " /><label>" + v.Key + "</label>"
	}
	return html.HTML(rs)
}

/**
 * 同上
 */
func GetRadioRecommendStatus2(StatusString string, recom int64, tagName string) html.HTML {
	fmt.Println("GetRadioRecommendStatus2")
	var rs string
	arr := GetKvArray(StatusString)
	if arr != nil {
		for _, v := range arr {
			var checked string
			if strings.Contains(strings.ToLower(con.Itoa(int(recom))), strings.ToLower(v.Value)) {
				checked = "checked=\"checked\""
			} else {
			}
			rs += "<input  type=\"radio\" value=\"" + v.Value + "\" name=\"" + tagName + "\" " + checked + " /><label>" + v.Key + "</label>"
		}
	}
	return html.HTML(rs)
}

/**
 * 通过name,Key 用来获取配置文件中的"配置文件数组"字符串的值
 * 如:   classType=10:增加;20:修改;30:删除;40:登陆
 * 		 GetVByKInConfig("classType","20")
 * 返回结果为:"登陆"
 */
func GetVByKInConfig(name string, key string) string {

	arr := GetKvArray(name)
	var rs string = ""
	for _, v := range arr {
		if v.Key == key {
			rs = v.Value
		}
	}
	return rs
}
func GetKByVInConfig(name string, v int64) string {
	value := con.Itoa(int(v))
	arr := GetKvArray(name)
	var rs string = ""
	for _, v := range arr {
		if v.Value == value {
			rs = v.Key
		}
	}
	return rs
}

/**
 * 初始化函数
 */
func init() {
	beego.AddFuncMap("GetRecommendStatus", GetRecommendStatus)
	beego.AddFuncMap("GetRadioRecommendStatus", GetRadioRecommendStatus)
	beego.AddFuncMap("GetRadioRecommendStatus2", GetRadioRecommendStatus2)
	beego.AddFuncMap("GetVByKInConfig", GetVByKInConfig)
	beego.AddFuncMap("GetKByVInConfig", GetKByVInConfig)
}
