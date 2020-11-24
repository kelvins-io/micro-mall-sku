package startup

import (
	"context"
	"gitee.com/cristiane/micro-mall-sku/vars"
	"gitee.com/kelvins-io/kelvins/setup"
	"github.com/qiniu/qmgo"
	"log"
)

// SetupVars 加载变量
func SetupVars() error {
	var err error
	if vars.QueueAMQPSettingUserRegisterNotice != nil && vars.QueueAMQPSettingUserRegisterNotice.Broker != "" {
		vars.QueueServerUserRegisterNotice, err = setup.NewAMQPQueue(vars.QueueAMQPSettingUserRegisterNotice, nil)
		if err != nil {
			return err
		}
	}

	if vars.QueueAMQPSettingUserStateNotice != nil && vars.QueueAMQPSettingUserStateNotice.Broker != "" {
		vars.QueueServerUserStateNotice, err = setup.NewAMQPQueue(vars.QueueAMQPSettingUserStateNotice, nil)
		if err != nil {
			return err
		}
	}

	// 初始化mongodb
	ctx := context.Background()
	var maxPoolSize = uint64(vars.MongoDBSetting.MaxPoolSize)
	var minPoolSize = uint64(vars.MongoDBSetting.MinPoolSize)

	mgoCfg := &qmgo.Config{
		Uri:         vars.MongoDBSetting.Uri,
		Database:    vars.MongoDBSetting.Database,
		MaxPoolSize: &maxPoolSize,
		MinPoolSize: &minPoolSize,
		Auth: &qmgo.Credential{
			AuthMechanism: "",
			AuthSource:    vars.MongoDBSetting.AuthSource,
			Username:      vars.MongoDBSetting.Username,
			Password:      vars.MongoDBSetting.Password,
			PasswordSet:   false,
		},
	}
	client, err := qmgo.NewClient(ctx, mgoCfg)
	if err != nil {
		log.Printf("mongodb connection err: %v", err)
		return err
	}
	err = client.Ping(30) // 30s
	if err != nil {
		log.Printf("mongodb ping timeout err: %v", err)
		return err
	}
	db := client.Database(vars.MongoDBSetting.Database)

	vars.MongoDBDatabase = db
	return err
}
