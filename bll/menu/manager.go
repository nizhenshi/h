package menu

/*
管理层 后台 业务层
秦鹏超
2014-08-27
*/
import (
	"html/template"
)

/*获取  管理员 状态  0 正常 1 待审核 */
func Manager_Get_Manager_Status(flag int) template.HTML {
	var strTemp = "<font color='red'>待审核</font>"
	if flag == 0 {
		strTemp = "正常"
	}
	return template.HTML(strTemp)
}
func Tb_Manager_Status_List(oldflag, flag int) bool {
	if oldflag == flag {
		return true
	}
	return false
}
