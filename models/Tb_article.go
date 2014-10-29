package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Tb_article struct {
	Id          int64
	Catid       int64
	Level       string
	Title       string `orm:"size(60)"`
	Thumb       string `orm:"size(60)"`
	Introduce   string `orm:"size(255)"`
	Author      string `orm:"size(30)"`
	Userid      int64
	Typeid      int64  `orm:"size(1)"`
	Areaid      string `orm:"size(10)"`
	Hits        int64
	CopyFrom    string `orm:"size(50)"`
	IsComment   int64
	Status      int64
	Content     string `orm:"type(text)"`
	SeoTitle    string `orm:"size(100)"`
	SeoKeyword  string `orm:"size(255)"`
	SeoDescribe string `orm:"size(255)"`
	EditTime    time.Time
	AddTime     time.Time
}

func (this *Tb_article) TableName() string {
	return "tb_article"
}

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(Tb_article))
}
