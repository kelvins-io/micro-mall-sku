package repository

import (
	"gitee.com/cristiane/micro-mall-sku/model/mysql"
	"xorm.io/xorm"
)

func CreateSkuPriceHistory(tx *xorm.Session, model *mysql.SkuPriceHistory) (err error) {
	_, err = tx.Table(mysql.TableSkuPriceHistory).Insert(model)
	return
}
