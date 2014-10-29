/**
 * 本文件用于注册数据库驱动,注册数据库
 */

package dal

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", beego.AppConfig.String("connectString"))
	orm.RunSyncdb("default", false, true)
}
