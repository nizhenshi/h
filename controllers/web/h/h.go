package h

import (
	"fmt"
	"github.com/astaxie/beego"
)

type HController struct {
	beego.Controller
}

func (this *HController) Index() []string {
	// signature := this.GetStrings("signature")
	// timestamp := this.GetStrings("timestamp")
	// nonce := this.GetStrings("nonce")
	echostr := this.GetStrings("echostr")
	fmt.Println(echostr)
	return echostr
}
