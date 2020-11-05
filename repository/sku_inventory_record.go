package repository

import (
	"gitee.com/cristiane/micro-mall-sku/model/mysql"
	"xorm.io/xorm"
)

func CreateSkuInventoryRecordByTx(tx *xorm.Session, model *mysql.SkuInventoryRecord) (err error) {
	_, err = tx.Table(mysql.TableSkuInventoryRecord).Insert(model)
	return
}
