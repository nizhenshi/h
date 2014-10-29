package basetype

import (
	"fmt"
	"github.com/astaxie/beego"
	html "html/template"
	dalbt "mcms/dal/basetype"
	m "mcms/models"
	"strings"
	// con "strconv"
)

func GetDeepPx(id int64, l []m.Tb_basetype) int {
	s := (dalbt.GetDeep(id, l) - 1) * 20
	fmt.Println("---------------s:", s)
	return s
}

/**
 *判断两个int数des1与des2是否相等，如果相等，返回指定的值str1,不相等，返回str2
 */
func JudgeEqualOrNo(des1 int, des2 int, str1 string, str2 string) html.HTML {
	str := ""
	if des2 == des1 {
		str = str1
	} else {
		str = str2
	}
	return html.HTML(str)
}

/**
 *判断两个int数des1与des2是否相等，如果相等，返回指定的值str1,不相等，返回str2
 */
func JudgeEqualOrNo2(des1 int64, des2 int64, str1 string, str2 string) html.HTML {
	str := ""
	if des2 == des1 {
		str = str1
	} else {
		str = str2
	}
	return html.HTML(strings.Trim(str, " "))
}

/**
 *判断两个int数des1与des2是否相等，如果相等，返回指定的值str1,不相等，返回str2
 */
func JudgeEqualOrNo_STRING(des1 string, des2 string, str1 string, str2 string) html.HTML {
	str := ""
	if des2 == des1 {
		str = str1
	} else {
		str = str2
	}
	return html.HTML(str)
}
func JudgeEqualOrNo_STRING2(des1 string, des2 string, str1 string, str2 string) string {
	str := ""
	if des2 == des1 {
		str = str1
	} else {
		str = str2
	}
	return str
}

func JudgeChild(id int64) html.HTML {
	_, n := dalbt.JudgeHasChild(id)
	rs := "false"
	if n > 0 {
		rs = "true"
	} else {
		rs = "false"
	}
	fmt.Println("-----------------rs:", rs)
	return html.HTML(rs)
}

func GetBaseTitle(Id int64) string {
	return dalbt.GetById(Id).Title

}
func Tb_baseType_Class_Name(Id int64) string {
	model := dalbt.GetById(Id)
	return model.Title
}
func init() {
	beego.AddFuncMap("GetDeepPx", GetDeepPx)
	beego.AddFuncMap("JudgeEqualOrNo_STRING2", JudgeEqualOrNo_STRING2)
	beego.AddFuncMap("JudgeEqualOrNo_STRING", JudgeEqualOrNo_STRING)
	beego.AddFuncMap("JudgeEqualOrNo2", JudgeEqualOrNo2)
	beego.AddFuncMap("JudgeEqualOrNo", JudgeEqualOrNo)
	beego.AddFuncMap("JudgeChild", JudgeChild)
	beego.AddFuncMap("GetBaseTitle", GetBaseTitle)
	beego.AddFuncMap("Tb_baseType_Class_Name", Tb_baseType_Class_Name)
}
