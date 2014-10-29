package fieldsValue

import (
	"github.com/astaxie/beego/orm"
	"mcms/models"
)

/**
 * 添加一个扩展字段值
 * @param fieVal 字段值对象
 * @return id 字段值id，error 错误信息
 * @author tongfei
 * @date 2014-08-29
 * @version v1.0
 */
func Add(fieVal models.Tb_fields_value) (int64, error) {
	//实例化Orm对象
	o := orm.NewOrm()
	id, err := o.Insert(&fieVal)

	if err != nil {
		return 0, err
	}

	return id, nil
}

/**
 * 添加多条扩展字段值
 * @param fieVals 字段值对象切片
 * @return num 影响行数，error 错误信息
 * @author tongfei
 * @date 2014-08-29
 * @version v1.0
 */
func AddMult(fieVals []models.Tb_fields_value) (int64, error) {

	//实例化Orm对象
	o := orm.NewOrm()

	if len(fieVals) > 0 {
		num, err := o.InsertMulti(len(fieVals), &fieVals)

		if err != nil {
			return 0, err
		}

		return num, nil
	}

	return 0, nil
}

/**
 * 获取一条字段值记录
 * @param catid 分类id，itemid 信息(产品)id, fieldid 字段id
 * @return fie_val 字段值
 * @author tongfei
 * @date 2014-08-29
 * @version v1.0
 */
func GetFieldValue(catid, itemid, fieldid int64) string {
	o := orm.NewOrm()
	var FieVal models.Tb_fields_value
	o.Raw("select * from tb_fields_value where catid = ? and fieldid = ? and itemid = ?", catid, fieldid, itemid).QueryRow(&FieVal)

	return FieVal.FieldValue
}

/**
 * 获取字段值记录切片
 * @param catid 分类id，itemid 信息(产品)id
 * @return fie_val 字段值
 * @author tongfei
 * @date 2014-08-29
 * @version v1.0
 */
func GetFieldValArr(catid, itemid int64) []models.Tb_fields_value {
	o := orm.NewOrm()
	FieVal := make([]models.Tb_fields_value, 0)
	o.Raw("select * from tb_fields_value where catid = ? and itemid = ?", catid, itemid).QueryRow(&FieVal)

	return FieVal
}

/**
 * 删除多条字段值记录
 * @param catid 分类id，itemid 信息(产品)id
 * @return error 错误信息
 * @author tongfei
 * @date 2014-08-29
 * @version v1.0
 */

func Delete(catid, itemid int64) error {
	o := orm.NewOrm()
	_, err := o.Raw("delete from tb_fields_value where catid = ? and itemid = ?", catid, itemid).Exec()

	return err
}
