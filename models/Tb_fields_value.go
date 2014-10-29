package models

import (
	"github.com/astaxie/beego/orm"
)

type Tb_fields_value struct {
	Id         int64
	Catid      int64
	Fieldid    int64
	Itemid     int64
	FieldValue string `orm:"type(text)"`
}

func (this *Tb_fields_value) TableName() string {
	return "tb_fields_value"
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Tb_fields_value))
}
