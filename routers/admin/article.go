package admin

import (
	"github.com/astaxie/beego"
	"mcms/controllers/admin/article"
)

func init() {
	beego.Router("/admin/article/add", &article.ArticleController{}, "*:Add")
	beego.Router("/admin/article/list", &article.ArticleController{}, "*:List")
	beego.Router("/admin/article/update", &article.ArticleController{}, "*:Update")
	beego.Router("/admin/article/delete", &article.ArticleController{}, "*:JsonDelMulti")
	beego.Router("/admin/article/jsonfields", &article.ArticleController{}, "Post:JsonFields")
}
