/**
 * tb_basetype的dal文件,并包含简单的业务逻辑操作
 * 2014-8-26
 * 支飞亚
 * 功能描述:
 * 1:插入,删除,查询,修改
 * 2:判断是否有子项
 * 3:获取ztree的json数据
 * 4:获取子项(在指定的数据源内/在数据库中)
 * 5:根据Id获取所有后代
 * 6:判断一个id是否在指定的切片内
 * 7:通过Id  判断该Id对应的类别是否有子项,并返回bool和数量
 * 8:获取类别表中，id对应的记录的深度
 */

package basetype

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	html "html/template"
	m "mcms/models"
	con "strconv"
	"strings"
)

/**
* 插入数据库一条记录，返回值为插入的Id
 */
func Add(tm *m.Tb_basetype) (int64, error) {
	o := orm.NewOrm()
	id, err := o.Insert(tm)
	return id, err
}

/**
 * 通过id,获取一个Tb_basetype的实例(数据源为数据库)
 */
func GetById(id int64) m.Tb_basetype {
	o := orm.NewOrm()
	var e m.Tb_basetype
	e.Id = id
	o.Read(&e)
	return e
}

/**
 * 通过id,在指定的数据源内获取一个Tb_basetype的实例
 */
func GetById2(id int64, l []m.Tb_basetype) m.Tb_basetype {
	fmt.Println("GetById2")
	var e m.Tb_basetype
	for _, v := range l {
		if v.Id == id {
			e = v
		}
	}
	return e
}

/**
 * 删除操作,通过id  参数形如：“1” 或  “1,2,4,23”,返回影响的行数
 */
func Delete(ids string) int64 {
	o := orm.NewOrm()
	sql := "delete from Tb_basetype where id in(" + ids + ")"
	res, err := o.Raw(sql).Exec()
	var num int64 = 0
	if err == nil {
		num, _ = res.RowsAffected()
	}
	return num
}

/**
 * 用一个Tb_basetype实体(实体必须含有Id)来更新数据库，返回影响的行数
 */
func Update(tm m.Tb_basetype) (int64, error) {
	o := orm.NewOrm()
	num := int64(0)
	var err error
	tm2 := m.Tb_basetype{Id: tm.Id}
	if o.Read(&tm2) == nil {
		tm2.Date = tm.Date
		tm2.Describe = tm.Describe
		tm2.IsShow = tm.IsShow
		tm2.ImageUrl = tm.ImageUrl
		tm2.OrderTop = tm.OrderTop
		tm2.ParentId = tm.ParentId
		tm2.Remark = tm.Remark
		tm2.Title = tm.Title
		num, err = o.Update(&tm2)
	}
	return num, err
}

/**
*根据pId获取"类别实体"的子类(不包含孙子辈的)的对象集合切片slice
 */
func GetOnlyChildSlices(pId int64) []m.Tb_basetype {

	o := orm.NewOrm()
	var l []m.Tb_basetype
	//切记查询的时候,QueryTable("Tb_basetype")参数注意大小写,一定要和实体结构类一模一样
	o.QueryTable("Tb_basetype").Filter("ParentId", pId).OrderBy("Ordertop").Limit(-1).All(&l)
	return l
}

/**
*根据pId在指定的[数据源tm]内获取"类别实体"的子类(不包含孙子辈的)的对象集合切片slice
 */
func GetOnlyChildSlices2(pId int64, tm []m.Tb_basetype) []m.Tb_basetype {
	fmt.Println("GetOnlyChildSlices2")
	var l []m.Tb_basetype
	for _, v := range tm {
		if v.ParentId == pId {
			l = append(l, v)
		}
	}
	return l
}

/**
*根据pId获取菜单子类(包含孙子辈的)的对象集合切片slice
 */
func GetAllChildSlices(pId int64) []m.Tb_basetype {
	fmt.Println("GetAllChildSlices")
	//var all []m.Tb_basetype //所有的
	var rs []m.Tb_basetype //符合条件的
	o := orm.NewOrm()
	ids := ""
	o.Raw("select  getTreeCategory(" + con.Itoa(int(pId)) + ")").QueryRow(&ids)
	ids = DealIds(ids, pId)
	o.Raw("select *  from tb_basetype where id in (" + ids + ")").QueryRows(&rs)
	return rs
}

func DealIds(ids string, pId int64) string {
	fmt.Println("DealIds")
	arr := strings.Split(ids, ",")
	var rs string = ""
	if len(arr) > 1 {
		for _, v := range arr {
			if v != con.Itoa(int(pId)) {
				rs += v + ","
			}
		}
	}
	rs = strings.Trim(rs, ",")
	return rs
}

/**
*判断一个id是否在指定的切片内
 */
func JudgeIdInSlices(id int64, sm []m.Tb_basetype) bool {
	fmt.Println("JudgeIdInSlices")
	var rs bool
	for _, v := range sm {
		if v.Id == id {
			rs = true
			break
		}
	}
	return rs
}

/**
*获取所有长辈的id,不能被外部代码调用,只能在本包内调用
 */
func getAllEldershipId(tm m.Tb_basetype, sm []m.Tb_basetype) []m.Tb_basetype {
	fmt.Println("getAllEldershipId")
	var rs []m.Tb_basetype
	for _, v := range sm {
		if tm.ParentId == int64(0) {
			return rs

		} else {

			if tm.ParentId == v.Id {
				rs = append(rs, v)
				var tempSlice = getAllEldershipId(v, sm)
				//切片追加切片
				for _, v2 := range tempSlice {
					rs = append(rs, v2)
				}
			}
		}
	}
	return rs
}

/**
*通过Id  判断该Id对应的类别是否有子项,并返回bool和数量
 */
func JudgeHasChild(id int64) (bool, int) {
	fmt.Println("JudgeHasChild")
	o := orm.NewOrm()
	var l []m.Tb_basetype
	var b bool //是否有子
	var c int  //数量
	_, err := o.QueryTable("Tb_basetype").OrderBy("Ordertop").Limit(-1).All(&l)
	if err == nil {
		for _, v := range l {
			if v.ParentId == id {
				b = true
				c++
			}
		}
	}
	return b, c
}

/**
*通过Id  判断该Id对应的类别是否在[数据源l]中有子项,并返回bool和数量
 */
func JudgeHasChild2(id int64, l []m.Tb_basetype) (bool, int) {
	fmt.Println("JudgeHasChild2")
	var b bool //是否有子
	var c int  //数量
	for _, v := range l {
		if v.ParentId == id {
			b = true
			c++
		}
	}

	return b, c
}

/**
*获取类别表中，id对应的记录的深度
 */
func GetDeep(id int64, l []m.Tb_basetype) int {

	var deep int = 1
	var temp m.Tb_basetype = GetById2(id, l)
	if temp.ParentId == 0 {
		return deep
	} else {
		deep += GetDeep(temp.ParentId, l)
	}
	return deep
}

/**
*获取 zTree [类别（Tb_basetype）]下拉列表json数据源
 */
func GetZtreeJsonByFid(fid int64) string {
	fmt.Println("GetZtreeJsonByFid")
	var z string = ""
	if fid == 0 {
		z = "[{id:0,pId:0,name:\"--无父级类别--\",open:true},"
	} else {
		tm := GetById(fid)
		z = "[{id:" + con.Itoa(int(tm.Id)) + ",pId:" + con.Itoa(int(tm.ParentId))
		z += ",name:\"" + tm.Title + "\",nocheck:true,open:true},"
	}
	tmenu := GetAllChildSlices(fid)
	if tmenu != nil {
		for i, v := range tmenu {
			if i == 0 {
				z += "{id:" + con.Itoa(int(v.Id)) + ", pId:" + con.Itoa(int(v.ParentId)) + ", name:\"" + v.Title + "\"}"
			} else {
				z += ",{id:" + con.Itoa(int(v.Id)) + ", pId:" + con.Itoa(int(v.ParentId)) + ", name:\"" + v.Title + "\"}"
			}
		}
	}
	z += "]"
	return z
}

/**
* 通过Id,获取cms的Menu列表页tr(用于TreeTable的ajax的动态请求)
 */
func GetTrById(id int64, basetypePid int64, l []m.Tb_basetype) string {
	e := GetById2(id, l)
	deep := GetDeep(id, l)
	haschild, _ := JudgeHasChild2(id, l)
	hc := "false"
	if haschild {
		hc = "true"
	}
	var str string = "<tr id='" + con.Itoa(int(e.Id)) + "' pId='" + con.Itoa(int(e.ParentId)) + "' hasChild='" + hc + "'>" +
		//<div class=\"checker\" >
		"<td align=\"center\">" + "<span>" + "<input  type=\"checkbox\" class='checkbox' value=\"" + con.Itoa(int(e.Id)) + "\" />" + "</span></td>" +
		"<td>" + con.Itoa(int(e.Id)) + "</td>" +
		"<td  controller=\"true\"  >" +
		"<span style=\"display:inline-block;width:" + con.Itoa((deep-1)*24) +
		"px;\"></span><span class=\"folder-line\"></span><span class=\"folder-open\"></span>" +
		e.Title +
		"</td>" +
		"<td >" + string(JudgeEqualOrNo(1, int(e.IsShow), "<font color='green'>显示</font>", "<font color='red'>隐藏</font>")) + "</td>" +
		"<td ><input type=\"text\" class=\"sort\" value=\"" + con.Itoa(int(e.OrderTop)) + "\" /></td>" +
		"<td style=\"text-align:center;\">" +
		"<a href=\"/admin/basetype/edit?pid=" + con.Itoa(int(e.Id)) + "&menuRootId=" + con.Itoa(int(basetypePid)) + "\">添加子类</a>&nbsp;&nbsp;<a href=\"/admin/basetype/edit?id=" + con.Itoa(int(e.Id)) + "&menuRootId=" + con.Itoa(int(basetypePid)) + "\">修改</a> </td></tr>"
	return str
}

/**
 *判断两个int数des1与des2是否相等，如果相等，返回指定的值str1,不相等，返回str2
 */
func JudgeEqualOrNo(des1 int, des2 int, str1 string, str2 string) html.HTML {
	str := ""
	if des2 == des1 {
		str = str1
	} else {
		str = str2
	}
	return html.HTML(str)
}

/**
 * 通过类别id  类别的根id  获取导航
 */
func GetDaoHang(cid, rootcid int64) html.HTML {
	fmt.Println("GetDaoHang")
	var rs string
	source := GetAllChildSlices(rootcid)
	t := GetById(cid)
	t2 := GetById(rootcid)
	farr := getAllEldershipId(t, source)
	farr = append(farr, t2)
	farr = append(farr, t)
	for i := 0; i < len(farr); i++ {
		for _, v2 := range farr {
			if rootcid == v2.ParentId {
				if i != 0 {
					rs += ">  <a href='/news/list2?cid=" + con.Itoa(int(v2.Id)) + "'>" + v2.Title + "</a>"
				} else {
					rs += ">  <a>" + v2.Title + "</a>"
				}
				rootcid = v2.Id
			}
			continue
		}

	}
	rs = strings.Trim(strings.Trim(strings.Trim(rs, ">"), " "), "")
	return html.HTML(rs)
}

/**
 * 通过类别id  类别的根id  获取导航
 */
func GetDaoHang2(cid, rootcid int64) html.HTML {
	fmt.Println("GetDaoHang2")
	var rs string
	source := GetAllChildSlices(rootcid)
	t := GetById(cid)
	t2 := GetById(rootcid)
	farr := getAllEldershipId(t, source)
	farr = append(farr, t2)
	farr = append(farr, t)
	for i := 0; i < len(farr); i++ {
		for _, v2 := range farr {
			if rootcid == v2.ParentId {
				if i != 0 {
					rs += ">  <a href='/news/list3?cid=" + con.Itoa(int(v2.Id)) + "'>" + v2.Title + "</a>"
				} else {
					rs += ">  <a>" + v2.Title + "</a>"
				}
				rootcid = v2.Id
			}
			continue
		}

	}
	rs = strings.Trim(strings.Trim(strings.Trim(rs, ">"), " "), "")
	return html.HTML(rs)
}

func GetRightMenu(cid int64) html.HTML {
	fmt.Println("GetRightMenu")
	all := GetAllChildSlices(0)
	dp := GetDeep(cid, all)
	t := GetById(cid)
	var rs string
	if dp > 1 && dp <= 3 {
		b, num := JudgeHasChild(cid)
		if b && num > 0 {
			p1 := GetById(t.ParentId)
			all1 := GetOnlyChildSlices(p1.Id)
			for _, v2 := range all1 {
				rs += "<dl><dt><span><a class=\"sel\" href=\"#\">" + v2.Title + "</a></span></dt>"
				allBrothers2 := GetOnlyChildSlices(v2.Id)
				rs += "<dd class=\"hide\" id=\"sm1\">"
				for _, v1 := range allBrothers2 {
					rs += "<a href='/news/list2?cid=" + con.Itoa(int(v1.Id)) + "' target='_blank'>" + v1.Title + "</a><br/>"
				}
				rs += "</dd></dl>"
			}
		} else {
			rs = "<dl>"
			allBrothers := GetOnlyChildSlices(t.ParentId)
			rs += "<dd class=\"hide\" id=\"sm1\">"
			for _, v1 := range allBrothers {
				rs += "<a href='/news/list2?cid=" + con.Itoa(int(v1.Id)) + "' target='_blank'>" + v1.Title + "</a><br/>"
			}
			rs += "</dd>"
			rs += "</dl>"
		}
	} else {
		p2 := GetById(t.ParentId)
		p1 := GetById(p2.ParentId)
		all1 := GetOnlyChildSlices(p1.Id)
		for _, v2 := range all1 {
			rs += "<dl><dt><span><a class=\"sel\" href=\"#\">" + v2.Title + "</a></span></dt>"
			allBrothers2 := GetOnlyChildSlices(v2.Id)
			rs += "<dd class=\"hide\" id=\"sm1\">"
			for _, v1 := range allBrothers2 {
				rs += "<a href='/news/list2?cid=" + con.Itoa(int(v1.Id)) + "' target='_blank'>" + v1.Title + "</a><br/>"
			}
			rs += "</dd></dl>"
		}
	}

	return html.HTML(rs)
}

func GetRightMenu3(cid int64) html.HTML {
	fmt.Println("GetRightMenu3")
	all := GetAllChildSlices(0)
	dp := GetDeep(cid, all)
	t := GetById(cid)
	var rs string
	if dp > 1 && dp <= 3 {
		b, num := JudgeHasChild(cid)
		if b && num > 0 {
			p1 := GetById(t.ParentId)
			all1 := GetOnlyChildSlices(p1.Id)
			for _, v2 := range all1 {
				rs += "<dl><dt><span><a class=\"sel\" href=\"#\">" + v2.Title + "</a></span></dt>"
				allBrothers2 := GetOnlyChildSlices(v2.Id)
				rs += "<dd class=\"hide\" id=\"sm1\">"
				for _, v1 := range allBrothers2 {
					rs += "<a href='/news/list3?cid=" + con.Itoa(int(v1.Id)) + "' target='_blank'>" + v1.Title + "</a><br/>"
				}
				rs += "</dd></dl>"
			}
		} else {
			rs = "<dl>"
			allBrothers := GetOnlyChildSlices(t.ParentId)
			rs += "<dd class=\"hide\" id=\"sm1\">"
			for _, v1 := range allBrothers {
				rs += "<a href='/news/list3?cid=" + con.Itoa(int(v1.Id)) + "' target='_blank'>" + v1.Title + "</a><br/>"
			}
			rs += "</dd>"
			rs += "</dl>"
		}
	} else {
		p2 := GetById(t.ParentId)
		p1 := GetById(p2.ParentId)
		all1 := GetOnlyChildSlices(p1.Id)
		for _, v2 := range all1 {
			rs += "<dl><dt><span><a class=\"sel\" href=\"#\">" + v2.Title + "</a></span></dt>"
			allBrothers2 := GetOnlyChildSlices(v2.Id)
			rs += "<dd class=\"hide\" id=\"sm1\">"
			for _, v1 := range allBrothers2 {
				rs += "<a href='/news/list3?cid=" + con.Itoa(int(v1.Id)) + "' target='_blank'>" + v1.Title + "</a><br/>"
			}
			rs += "</dd></dl>"
		}
	}

	return html.HTML(rs)
}

func GetRightMenu2(cid int64) html.HTML {
	fmt.Println("GetRightMenu2")
	var rs string
	t := GetById(cid)
	_, num := JudgeHasChild(cid)
	if num != 0 {
		pt := GetById(t.ParentId)
		all := GetOnlyChildSlices(t.Id)
		rs += "<div class=\"tit\">" + pt.Title + "</div>"
		rs += "<div class=\"left2_menu\" id=\"helpnav\">"
		rs += "<ul>"
		rs += "<dl><dt><span><a class=\"sel\" href='/jy/list?cid=" + con.Itoa(int(t.Id)) + "'>" + t.Title + "</a></span></dt>"
		rs += "<dd class=\"hide\" id=\"sm1\">"
		for _, v := range all {
			if v.Id == 123 {
				rs += "<a href='/jy/message?cid=" + con.Itoa(int(v.Id)) + "' target='blank'>" + v.Title + "</a><br/>"
			} else if v.Id == 128 {
				rs += "<a href='/jy/apply?cid=" + con.Itoa(int(v.Id)) + "' target='blank'>" + v.Title + "</a><br/>"
			} else {
				rs += "<a href='/jy/list?cid=" + con.Itoa(int(v.Id)) + "' target='blank'>" + v.Title + "</a><br/>"
			}
		}
		rs += "</dd></dl>"
		rs += "</ul>"
		rs += "</div>"
	} else {
		pt := GetById(t.ParentId)
		gt := GetById(pt.ParentId)
		all := GetOnlyChildSlices(pt.Id)
		rs += "<div class=\"tit\">" + gt.Title + "</div>"
		rs += "<div class=\"left2_menu\" id=\"helpnav\">"
		rs += "<ul>"
		rs += "<dl><dt><span><a class=\"sel\" href='/jy/list?cid=" + con.Itoa(int(pt.Id)) + "'>" + pt.Title + "</a></span></dt>"
		rs += "<dd class=\"hide\" id=\"sm1\">"
		for _, v := range all {
			if v.Id == 123 {
				rs += "<a href='/jy/message?cid=" + con.Itoa(int(v.Id)) + "' target='blank'>" + v.Title + "</a><br/>"
			} else if v.Id == 128 {
				rs += "<a href='/jy/apply?cid=" + con.Itoa(int(v.Id)) + "' target='blank'>" + v.Title + "</a><br/>"
			} else {
				rs += "<a href='/jy/list?cid=" + con.Itoa(int(v.Id)) + "' target='blank'>" + v.Title + "</a><br/>"
			}
		}
		rs += "</dd></dl>"
		rs += "</ul>"
		rs += "</div>"
	}
	return html.HTML(rs)
}
