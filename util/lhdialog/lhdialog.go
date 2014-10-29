package lhdialog

import (
	html "html/template"
	con "strconv"
)

/**/
func LhDialog_Tips(title string) string {
	return "<script>$.dialog.tips('" + title + "',3,'loading.gif');</script>"
}
func LhDialog_alert(content string) string {
	return "<script>$.dialog.alert('" + content + "');</script>"
}
func LhDialog_Diy(title, content, id, ok, cancel, icons string, islock bool) string {
	var strTemp = ""
	strTemp += "<script>"
	strTemp += "$.dialog({"
	strTemp += "title:'" + title + "',"
	if icons != "" {
		strTemp += "icon:'" + icons + "',"
	}
	strTemp += "id:'" + id + "',"
	if islock {
		strTemp += "lock:true,"
	} else {
		strTemp += "lock:false,"
	}
	strTemp += "content: '" + content + "',"
	strTemp += "ok: function(){" + ok + "},"
	strTemp += "cancel: function(){" + cancel + "}"
	strTemp += "});"
	strTemp += "</script>"
	return strTemp
}

/**
 *显示一个tip对话框,str:显示文本,time:时间,icon:图标   返回类型为普通HtmlString
 */
func Tips(str string, time int, icon string) html.HTML {
	return html.HTML("<script type='text/javascript'>$(function() {$.dialog.tips('" + str + "'," + con.Itoa(time) + ",'" + icon + "');});</script>")
}

/*
   弹出一个提示框，str提示文本，time显示时间，icon显示图标，href提示后跳转的路径

*/
func Alert(str string, time int, icon string, href string) html.HTML {
	return html.HTML("<script type='text/javascript'>$(function() {$.dialog.tips('" + str + "'," + con.Itoa(time) + ",'" + icon + "',function(){window.location.href='" + href + "';});});</script>")
}
