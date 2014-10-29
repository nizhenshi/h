/**
 * 实体模型,基础类别表
 */
package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Tb_basetype struct {
	Id       int64     `orm:"pk;column(Id)"`
	ParentId int64     `orm:"column(ParentId)"`
	Title    string    `orm:"column(Title)"`
	Date     time.Time `orm:"column(Date)"`
	OrderTop int64     `orm:"column(OrderTop)"`
	ImageUrl string    `orm:"column(ImageUrl)"`
	IsShow   int64     `orm:"column(IsShow)"`
	Describe string    `orm:"column(Describe)"`
	Remark   string    `orm:"column(Remark)"`
	Sort     int64     `orm:"column(Sort)"`
}

func (this *Tb_basetype) TableName() string {
	return "tb_basetype"
}

func init() {
	orm.RegisterModel(new(Tb_basetype))
}
