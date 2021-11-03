package startup

import (
	"context"
	"fmt"
	"gitee.com/cristiane/micro-mall-sku/model/args"
	"gitee.com/cristiane/micro-mall-sku/vars"
	"gitee.com/kelvins-io/g2cache"
	"gitee.com/kelvins-io/kelvins"
	"gitee.com/kelvins-io/kelvins/setup"
	"gitee.com/kelvins-io/kelvins/util/queue_helper"
	"github.com/qiniu/qmgo"
	"log"
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
	// 3 g2cache
	if vars.G2CacheSetting != nil && vars.G2CacheSetting.RedisConfDSN != "" {
		vars.G2CacheEngine, err = newG2Cache(vars.G2CacheSetting, nil, nil)
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
	if vars.G2CacheEngine != nil {
		vars.G2CacheEngine.Close()
	}
	return nil
}

func newG2Cache(g2cacheSetting *vars.G2CacheSettingS, out g2cache.OutCache, local g2cache.LocalCache) (*g2cache.G2Cache, error) {
	if g2cacheSetting == nil {
		return nil, fmt.Errorf("g2cacheSetting is nil")
	}
	if g2cacheSetting.CacheMonitor {
		g2cache.CacheMonitor = true
		if g2cacheSetting.CacheMonitorSecond > 0 {
			g2cache.CacheMonitorSecond = g2cacheSetting.CacheMonitorSecond
		}
	}
	if g2cacheSetting.CacheDebug {
		g2cache.CacheDebug = true
	}
	if g2cacheSetting.OutCachePubSub {
		g2cache.OutCachePubSub = true
	}

	if g2cacheSetting.EntryLazyFactor > 0 {
		g2cache.EntryLazyFactor = g2cacheSetting.EntryLazyFactor
	}
	if g2cacheSetting.GPoolWorkerNum > 0 {
		g2cache.DefaultGPoolWorkerNum = g2cacheSetting.GPoolWorkerNum
	}
	if g2cacheSetting.GPoolJobQueueChanLen > 0 {
		g2cache.DefaultGPoolJobQueueChanLen = g2cacheSetting.GPoolJobQueueChanLen
	}
	if g2cacheSetting.FreeCacheSize > 0 {
		g2cache.DefaultFreeCacheSize = g2cacheSetting.FreeCacheSize
	}
	if len(g2cacheSetting.PubSubRedisChannel) != 0 {
		g2cache.DefaultPubSubRedisChannel = g2cacheSetting.PubSubRedisChannel
	}
	if len(g2cacheSetting.RedisConfDSN) <= 0 {
		return nil, fmt.Errorf("g2cacheSetting.RedisConfDSN is empty")
	} else {
		g2cache.DefaultRedisConf.DSN = g2cacheSetting.RedisConfDSN
	}
	if g2cacheSetting.RedisConfDB >= 0 {
		g2cache.DefaultRedisConf.DB = g2cacheSetting.RedisConfDB
	}
	if len(g2cacheSetting.RedisConfPwd) > 0 {
		g2cache.DefaultRedisConf.Pwd = g2cacheSetting.RedisConfPwd
	}
	if g2cacheSetting.RedisConfMaxConn > 0 {
		g2cache.DefaultRedisConf.MaxConn = g2cacheSetting.RedisConfMaxConn
	}
	if g2cacheSetting.PubSubRedisConfDSN != "" {
		g2cache.DefaultPubSubRedisConf.DSN = g2cacheSetting.PubSubRedisConfDSN
	} else {
		g2cache.DefaultPubSubRedisConf.DSN = g2cacheSetting.RedisConfDSN
	}
	if g2cacheSetting.PubSubRedisConfDB >= 0 {
		g2cache.DefaultPubSubRedisConf.DB = g2cacheSetting.PubSubRedisConfDB
	} else {
		g2cache.DefaultPubSubRedisConf.DB = g2cacheSetting.RedisConfDB
	}
	if g2cacheSetting.PubSubRedisConfPwd != "" {
		g2cache.DefaultPubSubRedisConf.Pwd = g2cacheSetting.PubSubRedisConfPwd
	} else {
		g2cache.DefaultPubSubRedisConf.Pwd = g2cacheSetting.RedisConfPwd
	}
	if g2cacheSetting.PubSubRedisConfMaxConn > 0 {
		g2cache.DefaultPubSubRedisConf.MaxConn = g2cacheSetting.PubSubRedisConfMaxConn
	} else {
		g2cache.DefaultPubSubRedisConf.MaxConn = g2cacheSetting.RedisConfMaxConn
	}

	return g2cache.New(out, local)
}
