package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Tb_fields struct {
	Id           int64
	Typeid       int64
	Catid        int64
	Name         string `orm:"size(20);unique"`
	Title        string `orm:"size(30)"`
	HtmlType     string `orm:"size(10)"`
	DefaultValue string `orm:"type(text)"`
	OptionValue  string `orm:"type(text)"`
	InputLimit   string `orm:"size(50)"`
	Width        int64  `orm:"size(5)"`
	Height       int64  `orm:"size(5)"`
	Status       int64  `orm:"size(1)"`
	Note         string `orm:"size(30)"`
	Resort       int64  `orm:"size(11)"`
	AddTime      time.Time
}

func (this *Tb_fields) TableName() string {
	return "tb_fields"
}

func init() {
	orm.Debug = true
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Tb_fields))
}
