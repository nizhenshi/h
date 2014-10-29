package admin

import (
	"github.com/astaxie/beego"
	"mcms/controllers/admin/fields"
)

func init() {
	beego.Router("/admin/fields/add", &fields.FieldsController{}, "*:Add")
	beego.Router("/admin/fields/list", &fields.FieldsController{}, "*:List")
	beego.Router("/admin/fields/update", &fields.FieldsController{}, "*:Update")
	beego.Router("/admin/fields/delete", &fields.FieldsController{}, "Post:JsonDelMulti")
	beego.Router("/admin/fields/resort", &fields.FieldsController{}, "Get:Resort")
}
