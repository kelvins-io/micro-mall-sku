package repository

import (
	"gitee.com/cristiane/micro-mall-sku/model/mysql"
	"gitee.com/kelvins-io/kelvins"
	"xorm.io/xorm"
)

func CreateSkuPriceHistory(tx *xorm.Session, model *mysql.SkuPriceHistory) (err error) {
	_, err = tx.Table(mysql.TableSkuPriceHistory).Insert(model)
	return
}

func GetSkuPriceHistory(sqlSelect string, where interface{}, orderByDesc []string, limit int) ([]mysql.SkuPriceHistory, error) {
	result := make([]mysql.SkuPriceHistory, 0)
	err := kelvins.XORM_DBEngine.Table(mysql.TableSkuPriceHistory).Select(sqlSelect).
		Where(where).
		Desc(orderByDesc...).
		Limit(limit).
		Find(&result)
	return result, err
}
