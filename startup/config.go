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
	SectionG2Cache                  = "micro-mall-g2cache"
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
	//加载G2Cache二级缓存配置
	vars.G2CacheSetting = new(vars.G2CacheSettingS)
	config.MapConfig(SectionG2Cache, vars.G2CacheSetting)
	return nil
}
