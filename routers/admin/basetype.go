package admin

import (
	"github.com/astaxie/beego"
	bt "mcms/controllers/admin/basetype"
)

func init() {
	beego.Router("/admin/basetype/list", &bt.BasetypeController{}, "get:List")
	beego.Router("/admin/basetype/edit", &bt.BasetypeController{}, "*:Edit")
	beego.Router("/admin/basetype/editpost", &bt.BasetypeController{}, "Post:EditPost")
	///cms/basetype/gettrlist
	beego.Router("/admin/basetype/gettrlist", &bt.BasetypeController{}, "Post:GetTrList")
	beego.Router("/admin/basetype/remove", &bt.BasetypeController{}, "Post:Remove")
}
