package vars

import (
	"gitee.com/kelvins-io/kelvins/config/setting"
	"github.com/qiniu/qmgo"
)

var (
	EmailConfigSetting                 *EmailConfigSettingS
	MongoDBSetting                     *MongoDBSettingS
	MongoDBDatabase                    *qmgo.Database
	QueueAMQPSettingUserRegisterNotice *setting.QueueAMQPSettingS
	QueueAMQPSettingUserStateNotice    *setting.QueueAMQPSettingS
)
