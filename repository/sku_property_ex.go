package repository

import (
	"context"
	"sync"

	"gitee.com/cristiane/micro-mall-sku/model/args"
	"gitee.com/cristiane/micro-mall-sku/model/mongodb"
	"gitee.com/cristiane/micro-mall-sku/vars"
)

var oneSkuPropertyEx sync.Once

func createIndexesOfSkuPropertyEx() {
	oneSkuPropertyEx.Do(func() {
		var uniques = []string{"sku_code,shop_id"}
		var indexes = []string{"shape", "shop_id", "sku_code"}
		vars.MongoDBDatabase.Collection(mongodb.TableSkuPropertyEx).EnsureIndexes(context.Background(), uniques, indexes)
	})
}

func CreateSkuPropertyEx(ctx context.Context, req interface{}) (err error) {
	if vars.MongoDBDatabase == nil {
		return nil
	}
	// 创建索引
	createIndexesOfSkuPropertyEx()
	// 插入记录
	_, err = vars.MongoDBDatabase.Collection(mongodb.TableSkuPropertyEx).InsertOne(ctx, req)
	return
}

func GetSkuPropertyExList(ctx context.Context, query map[string]interface{}) ([]args.SkuPropertyEx, error) {
	var skuExList []args.SkuPropertyEx
	err := vars.MongoDBDatabase.Collection(mongodb.TableSkuPropertyEx).Find(ctx, query).Sort("sku_code").Limit(100).All(&skuExList)
	return skuExList, err
}
