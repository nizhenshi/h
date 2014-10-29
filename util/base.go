package util

import (
	"strconv"
	"strings"
)

/*通用 分页 代码*/
/// <summary>
/// 返回分页页码
/// </summary>
/// <param name="pageSize">页面大小</param>
/// <param name="pageIndex">当前页</param>
/// <param name="totalCount">总记录数</param>
/// <param name="linkUrl">链接地址，__id__代表页码</param>
/// <param name="centSize">中间页码数量</param>
/// <returns></returns>
func OutPageList(pageSize, pageIndex, totalCount int, linkUrl string, centSize int) string {
	//计算页数
	if totalCount < 1 || pageSize < 1 {
		return ""
	}
	var pageCount int = totalCount / pageSize
	if pageCount < 1 {
		return ""
	}
	if totalCount%pageSize > 0 {
		pageCount += 1
	}
	if pageCount <= 1 {
		return ""
	}
	var pageStr, pageId, firstBtn, lastBtn, firstStr, lastStr string
	pageStr = ""
	pageId = "__id__"
	firstBtn = "<a href=\"" + ReplaceStr(linkUrl, pageId, strconv.Itoa(pageIndex-1)) + "\">«上一页</a>"
	lastBtn = "<a href=\"" + ReplaceStr(linkUrl, pageId, strconv.Itoa(pageIndex+1)) + "\">下一页»</a>"
	firstStr = "<a href=\"" + ReplaceStr(linkUrl, pageId, "1") + "\">1</a>"
	lastStr = "<a href=\"" + ReplaceStr(linkUrl, pageId, strconv.Itoa(pageCount)) + "\">" + strconv.Itoa(pageCount) + "</a>"

	if pageIndex <= 1 {
		firstBtn = "<span class=\"disabled\">«上一页</span>"
	}
	if pageIndex >= pageCount {
		lastBtn = "<span class=\"disabled\">下一页»</span>"
	}
	if pageIndex == 1 {
		firstStr = "<span class=\"current\">1</span>"
	}
	if pageIndex == pageCount {
		lastStr = "<span class=\"current\">" + strconv.Itoa(pageCount) + "</span>"
	}
	firstNum := pageIndex - (centSize / 2) //中间开始的页码
	if pageIndex < centSize {
		firstNum = 2
	}
	lastNum := pageIndex + centSize - ((centSize / 2) + 1) //中间结束的页码
	if lastNum >= pageCount {
		lastNum = pageCount - 1
	}
	pageStr += "<span>共" + strconv.Itoa(totalCount) + "记录</span>"
	pageStr += firstBtn + firstStr
	if pageIndex >= centSize {
		pageStr += "<span>...</span>\n"
	}
	for i := firstNum; i <= lastNum; i++ {
		if i == pageIndex {
			pageStr += "<span class=\"current\">" + strconv.Itoa(i) + "</span>"
		} else {
			pageStr += "<a href=\"" + ReplaceStr(linkUrl, pageId, strconv.Itoa(i)) + "\">" + strconv.Itoa(i) + "</a>"
		}
	}
	if pageCount-pageIndex > centSize-(centSize/2) {
		pageStr += "<span>...</span>"
	}
	pageStr += lastStr + lastBtn + "<input name=\"pageIndex\"/><input type=\"submit\" value=\"提交\"/>"
	return pageStr
}

/*替换URL 字符 配合分页使用*/
func ReplaceStr(originalStr, oldStr, newStr string) string {
	if strings.Contains(originalStr, "?") {
		if strings.Contains(originalStr, "pageIndex") {
			urlString := strings.Split(originalStr, "pageIndex=")
			loca := strings.Index(urlString[1], "&")
			if loca > 0 {
				originalStr = string(urlString[0]) + "pageIndex=" + newStr + SubStrings(urlString[1], loca)
			} else {
				originalStr = string(urlString[0]) + "pageIndex=" + newStr
			}
		} else {
			originalStr += "&pageIndex=" + newStr
		}
	} else {
		originalStr += "?pageIndex=" + newStr
	}
	return originalStr
}

/*截取字符串,从 l 开始截取到 最后
///原符串
///截取位置
*/
func SubStrings(strContent string, l int) string {
	r := []rune(strContent) //将字符串转换成 数据
	var rtnString string
	if len(r) <= l {
		return strContent
	}
	for i := l; i < len(r); i++ {
		rtnString += string(r[i])
	}
	return rtnString
}
func Substr(strContent string, l int) string {
	r := []rune(strContent) //将字符串转换成 数据
	var rtnString string
	if len(r) <= l {
		return strContent
	}
	for i := 0; i < l; i++ {
		rtnString += string(r[i])
	}
	return rtnString
}
func Len(strContent string) int {
	r := []rune(strContent) //将字符串转换成 数据
	return len(r)
}
