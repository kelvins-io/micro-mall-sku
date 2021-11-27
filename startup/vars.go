package startup

import (
	"context"
	"log"

	"gitee.com/cristiane/micro-mall-sku/model/args"
	"gitee.com/cristiane/micro-mall-sku/vars"
	"gitee.com/kelvins-io/kelvins"
	"gitee.com/kelvins-io/kelvins/setup"
	"gitee.com/kelvins-io/kelvins/util/queue_helper"
	"github.com/qiniu/qmgo"
)

// SetupVars 加载变量
func SetupVars() error {
	var err error
	// 1 mongodb
	err = setupMongodb()
	if err != nil {
		return err
	}
	// 2 sku queue
	err = setupQueueSkuInventorySearchNotice()
	if err != nil {
		return err
	}

	return err
}

func setupQueueSkuInventorySearchNotice() error {
	var err error
	if vars.SkuInventorySearchNoticeSetting != nil {
		vars.SkuInventorySearchNoticeServer, err = setup.NewAMQPQueue(vars.SkuInventorySearchNoticeSetting, nil)
		if err != nil {
			return err
		}
		vars.SkuInventorySearchNoticePusher, err = queue_helper.NewPublishService(
			vars.SkuInventorySearchNoticeServer, &queue_helper.PushMsgTag{
				DeliveryTag:    args.SkuInventorySearchNoticeTag,
				DeliveryErrTag: args.SkuInventorySearchNoticeTagErr,
				RetryCount:     vars.SkuInventorySearchNoticeSetting.TaskRetryCount,
				RetryTimeout:   vars.SkuInventorySearchNoticeSetting.TaskRetryTimeout,
			}, kelvins.BusinessLogger)
		if err != nil {
			return err
		}
	}

	return err
}

func setupMongodb() error {
	if vars.MongoDBSetting == nil || vars.MongoDBSetting.Uri == "" {
		return nil
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
	vars.MongoDBDatabase = client.Database(vars.MongoDBSetting.Database)

	return nil
}

func StopFunc() error {

	return nil
}
