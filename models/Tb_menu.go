package models

/*
菜单实体类
*/

import (
	"github.com/astaxie/beego/orm"
)

type Tb_menu struct {
	Id       int64  `orm:"auto;pk;column(Id);" form:"id"`
	ParentId int    `orm:"column(ParentId)" form:"parentid"`
	Title    string `orm:"column(Title)" form:"title"`
	LinkUrl  string `orm:"column(LinkUrl)" form:"linkurl"`
	OrderTop int    `orm:"column(OrderTop)" form:"ordertop"`
	Remark   string `orm:"column(Remark)" form:"remark"`
	Target   string `orm:"column(Target)" form:"target"`
	Action   string `orm:"column(Action)" form:"action"`
	Name     string `orm:"column(Name)" form:"name"`
	ViewFlag int    `orm:"column(ViewFlag)" form:"viewFlag"`
}

func (this *Tb_menu) TableName() string {
	return "tb_menu"
}
func init() {
	orm.RegisterModel(new(Tb_menu))
}
