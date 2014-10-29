package article

import (
	"github.com/astaxie/beego/orm"
	"mcms/models"
)

/**
 * 添加一条信息
 * @param article 信息对象
 * @return id 字段id，error 错误信息
 * @author tongfei
 * @date 2014-08-29
 * @version v1.0
 */
func Add(article models.Tb_article) (int64, error) {

	//实例化Orm对象
	o := orm.NewOrm()
	id, err := o.Insert(&article)

	if err != nil {
		return 0, err
	}

	return id, nil
}

/**
 * 更新一条信息
 * @param article 信息对象
 * @return num 影响的行数，error 错误信息
 * @author tongfei
 * @date 2014-08-29
 * @version v1.0
 */
func Update(article models.Tb_article) (int64, error) {
	o := orm.NewOrm()
	num, err := o.Update(&article)

	if err != nil {
		return 0, err
	}

	return num, nil
}

/**
 * 获取一个信息对象
 * @param id 信息id
 * @return Article 信息对象，error 错误信息
 * @author tongfei
 * @date 2014-08-29
 * @version v1.0
 */
func GetById(id int64) (models.Tb_article, error) {
	o := orm.NewOrm()
	var Article, ArtNil models.Tb_article
	Article.Id = id
	err := o.Read(&Article)

	if err != nil {
		return ArtNil, err
	}

	return Article, nil
}

/**
 * 获取信息对象的切片
 * @param catid 信息所属分类的id, 状态status, 调取条数limitNum， 跳过的记录数offsetNum
 * @return arts 信息对象切片，error 错误信息
 * @author tongfei
 * @date 2014-08-29
 * @version v1.0
 */
func List(catid, status, limitNum, offsetNum int64) ([]models.Tb_article, error) {
	o := orm.NewOrm()
	Fields := new(models.Tb_article)
	tb_article := Fields.TableName()
	orm_art := o.QueryTable(tb_article)

	if catid > 0 {
		orm_art = orm_art.Filter("catid", catid)
	}

	orm_art = orm_art.Filter("status", status)
	orm_art = orm_art.Limit(limitNum, offsetNum)
	arts := make([]models.Tb_article, 0)
	_, err := orm_art.OrderBy("-add_time").All(&arts)

	if err != nil {
		return nil, err
	}

	return arts, nil
}

/**
 * 删除一条信息
 * @param cid 信息id
 * @return error 错误信息
 * @author tongfei
 * @date 2014-08-29
 * @version v1.0
 */
func Delete(id int64) error {
	o := orm.NewOrm()
	var Article models.Tb_article
	Article.Id = id

	if _, err := o.Delete(&Article); err != nil {
		return err
	}

	return nil
}

/**
 * 删除多条条信息
 * @param ids 字符串切片
 * @return error 错误信息
 * @author tongfei
 * @date 2014-08-29
 * @version v1.0
 */
func DeleteMult(ids []string) error {
	o := orm.NewOrm()
	Article := new(models.Tb_article)
	tb_article := Article.TableName()
	orm_art := o.QueryTable(tb_article)

	if _, err := orm_art.Filter("id__in", ids).Delete(); err != nil {
		return err
	}

	return nil
}

/**
 * 通过sql语句  获取文章
 */
func GetArticleBySQL(sql string) []models.Tb_article {
	o := orm.NewOrm()
	var all []models.Tb_article
	o.Raw(sql).QueryRows(&all)
	return all
}
