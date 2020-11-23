package repository

import (
	"gitee.com/cristiane/micro-mall-sku/model/mysql"
	"gitee.com/kelvins-io/kelvins"
	"xorm.io/xorm"
)

// sku库存
func CreateSkuInventory(tx *xorm.Session, model *mysql.SkuInventory) (err error) {
	_, err = tx.Table(mysql.TableSkuInventory).Insert(model)
	return
}

func GetSkuInventory(sqlSelect string, shopId int64, skuCode string) (*mysql.SkuInventory, error) {
	var model mysql.SkuInventory
	_, err := kelvins.XORM_DBEngine.Table(mysql.TableSkuInventory).
		Select(sqlSelect).
		Where("sku_code = ?", skuCode).
		Where("shop_id = ?", shopId).
		Get(&model)
	return &model, err
}

func CheckSkuInventoryExist(shopId int64, skuCode string) (exist bool, err error) {
	var model mysql.SkuInventory
	_, err = kelvins.XORM_DBEngine.Table(mysql.TableSkuInventory).
		Select("id").
		Where("sku_code = ?", skuCode).
		Where("shop_id = ?", shopId).
		Get(&model)
	if err != nil {
		return false, err
	}
	if model.Id > 0 {
		return true, nil
	}
	return false, nil
}

func GetSkuInventoryListByShopId(sqlSelect string, shopId int64, pageSize, pageNum int) ([]mysql.SkuInventory, error) {
	var result = make([]mysql.SkuInventory, 0)
	session := kelvins.XORM_DBEngine.Table(mysql.TableSkuInventory)
	session = session.Select(sqlSelect).Where("amount >= 0")
	if shopId > 0 {
		session = session.Where("shop_id = ?", shopId)
	}
	if pageSize > 0 && pageNum >= 1 {
		session = session.Limit(pageSize, (pageNum-1)*pageSize)
	}
	err := session.Desc("create_time").Find(&result)
	return result, err
}

func GetSkuInventoryList(sqlSelect string, shopIdList []int64, skuCodeList []string) ([]*mysql.SkuInventory, error) {
	var result = make([]*mysql.SkuInventory, 0)
	session := kelvins.XORM_DBEngine.Table(mysql.TableSkuInventory).
		Select(sqlSelect)
	if shopIdList != nil && len(shopIdList) > 0 {
		session = session.In("shop_id", shopIdList)
	}
	if skuCodeList != nil && len(skuCodeList) > 0 {
		session = session.In("sku_code", skuCodeList)
	}
	err := session.Find(&result)
	return result, err
}

func UpdateInventory(tx *xorm.Session, where, maps map[string]interface{}) (int64, error) {
	return tx.Table(mysql.TableSkuInventory).Where(where).Update(maps)
}
