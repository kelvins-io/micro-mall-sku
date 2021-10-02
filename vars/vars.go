package vars

import (
	"gitee.com/kelvins-io/common/queue"
	"gitee.com/kelvins-io/kelvins/config/setting"
	"gitee.com/kelvins-io/kelvins/util/queue_helper"
	"github.com/qiniu/qmgo"
)

var (
	EmailConfigSetting              *EmailConfigSettingS
	MongoDBSetting                  *MongoDBSettingS
	MongoDBDatabase                 *qmgo.Database
	SkuInventorySearchNoticeSetting *setting.QueueAMQPSettingS
	SkuInventorySearchNoticeServer  *queue.MachineryQueue
	SkuInventorySearchNoticePusher  *queue_helper.PublishService
)
