package web

import (
	"github.com/astaxie/beego"
	h "mcms/controllers/web/h"
)

func init() {
	beego.Router("/h", &h.HController{}, "*:Index")
}
