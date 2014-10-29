package routers

/*
CMS 过滤器
*/
import (
	"github.com/astaxie/beego/context"
	"strings"
)

/*
过滤器
*/
var FilterUser = func(ctx *context.Context) {
	userId := 1
	var lowerUrl string = strings.ToLower(ctx.Request.RequestURI)
	if userId == 0 && strings.Contains(lowerUrl, "/admin") {
		if !noCheckUrl(lowerUrl) {
			ctx.Redirect(301, "/admin/loginout")
		}
	}

}

/*排除的URL */
func noCheckUrl(lowerUrl string) bool {
	var flag bool = false
	flag = flag || strings.Contains(lowerUrl, "/admin/center")
	flag = flag || strings.Contains(lowerUrl, "/static")
	flag = flag || strings.Contains(lowerUrl, "/admin/loginout")
	flag = flag || lowerUrl == "/admin"
	return flag
}
