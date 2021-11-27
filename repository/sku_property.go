package repository

import (
	"context"
	"sync"

	"gitee.com/cristiane/micro-mall-sku/model/mongodb"
	"gitee.com/cristiane/micro-mall-sku/model/mysql"
	"gitee.com/cristiane/micro-mall-sku/vars"
	"gitee.com/kelvins-io/kelvins"
	"xorm.io/xorm"
)

var oneSkuProperty sync.Once

// sku 商品属性
func CreateSkuProperty(tx *xorm.Session, model *mysql.SkuProperty) (err error) {
	_, err = tx.Table(mysql.TableSkuProperty).Insert(model)
	return
}

func GetSkuPropertyList(skuCodeList []string) ([]mysql.SkuProperty, error) {
	var result = make([]mysql.SkuProperty, 0)
	err := kelvins.XORM_DBEngine.Table(mysql.TableSkuProperty).In("code", skuCodeList).Find(&result)
	return result, err
}

func createIndexesOfSkuProperty() {
	oneSkuProperty.Do(func() {
		var uniques = []string{"code"}
		var indexes = []string{"name", "title", "sub_title", "color"}
		vars.MongoDBDatabase.Collection(mongodb.TableSkuProperty).EnsureIndexes(context.Background(), uniques, indexes)
	})
}

func CreateSkuPropertyMongoDB(ctx context.Context, req interface{}) (err error) {
	if vars.MongoDBDatabase == nil {
		return nil
	}
	// 创建索引
	createIndexesOfSkuProperty()
	// 插入记录
	_, err = vars.MongoDBDatabase.Collection(mongodb.TableSkuProperty).InsertOne(ctx, req)
	return
}
