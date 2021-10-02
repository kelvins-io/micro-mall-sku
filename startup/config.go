package startup

import (
	"gitee.com/cristiane/micro-mall-sku/vars"
	"gitee.com/kelvins-io/kelvins/config"
	"gitee.com/kelvins-io/kelvins/config/setting"
)

const (
	SectionEmailConfig              = "email-config"
	SectionSkuInventorySearchNotice = "sku-inventory-search-notice"
	SectionMongoDB                  = "mongodb-config"
)

// LoadConfig 加载配置对象映射
func LoadConfig() error {
	// 加载email数据源
	vars.EmailConfigSetting = new(vars.EmailConfigSettingS)
	config.MapConfig(SectionEmailConfig, vars.EmailConfigSetting)
	// 商品库存搜素通知
	vars.SkuInventorySearchNoticeSetting = new(setting.QueueAMQPSettingS)
	config.MapConfig(SectionSkuInventorySearchNotice, vars.SkuInventorySearchNoticeSetting)
	// 加载mongodb 配置
	vars.MongoDBSetting = new(vars.MongoDBSettingS)
	config.MapConfig(SectionMongoDB, vars.MongoDBSetting)
	return nil
}
