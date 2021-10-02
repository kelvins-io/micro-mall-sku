package service

import (
	"context"
	"gitee.com/cristiane/micro-mall-sku/model/args"
	"gitee.com/cristiane/micro-mall-sku/vars"
	"gitee.com/kelvins-io/common/json"
	"gitee.com/kelvins-io/kelvins"
	"github.com/google/uuid"
)

func skuInventorySearch(info *args.SkuInventoryInfo) error {
	kelvins.GPool.SendJob(func() {
		var ctx = context.TODO()
		msg := &args.CommonBusinessMsg{
			Type:    args.SkuInventorySearchNotice,
			Tag:     args.GetMsg(args.SkuInventorySearchNotice),
			UUID:    uuid.New().String(),
			Content: json.MarshalToStringNoError(info),
		}
		vars.SkuInventorySearchNoticePusher.PushMessage(ctx, msg)
	})
	return nil
}
