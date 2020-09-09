package service

import (
	"context"
	"gitee.com/cristiane/micro-mall-sku/model/args"
	"gitee.com/cristiane/micro-mall-sku/model/mysql"
	"gitee.com/cristiane/micro-mall-sku/pkg/code"
	"gitee.com/cristiane/micro-mall-sku/pkg/util"
	"gitee.com/cristiane/micro-mall-sku/proto/micro_mall_shop_proto/shop_business"
	"gitee.com/cristiane/micro-mall-sku/proto/micro_mall_sku_proto/sku_business"
	"gitee.com/cristiane/micro-mall-sku/repository"
	"gitee.com/kelvins-io/kelvins"
	"time"
)

func PutAwaySku(ctx context.Context, req *sku_business.PutAwaySkuRequest) (retCode int) {
	retCode = code.Success
	if req.Sku.ShopId > 0 {
		serverName := args.RpcServiceMicroMallShop
		conn, err := util.GetGrpcClient(serverName)
		if err != nil {
			kelvins.ErrLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
			retCode = code.ErrorServer
			return
		}
		defer conn.Close()

		client := shop_business.NewShopBusinessServiceClient(conn)
		r := shop_business.GetShopMaterialRequest{
			ShopId: req.Sku.ShopId,
		}
		rsp, err := client.GetShopMaterial(ctx, &r)
		if err != nil {
			kelvins.ErrLogger.Errorf(ctx, "GetShopMaterial %v,err: %v, req: %+v", serverName, err, r)
			retCode = code.ErrorServer
			return
		}
		if rsp == nil || rsp.Material == nil || rsp.Material.ShopId <= 0 {
			retCode = code.ShopBusinessNotExist
			return
		}
	}

	if req.OperationType == sku_business.OperationType_CREATE {
		exist, err := repository.CheckSkuInventoryExist(req.Sku.SkuCode)
		if err != nil {
			kelvins.ErrLogger.Errorf(ctx, "CheckSkuInventoryExist %v,err: %v, SkuCode: %+v", err, req.Sku.SkuCode)
			retCode = code.ErrorServer
			return
		}
		if exist {
			retCode = code.SkuCodeExist
			return
		}
		skuProperty := mysql.SkuProperty{
			Code:          req.Sku.SkuCode,
			Price:         req.Sku.Price,
			Name:          req.Sku.Name,
			Desc:          req.Sku.Desc,
			Production:    req.Sku.Production,
			Supplier:      req.Sku.Supplier,
			Category:      int(req.Sku.Category),
			Title:         req.Sku.Title,
			SubTitle:      req.Sku.SubTitle,
			Color:         req.Sku.Color,
			ColorCode:     int(req.Sku.ColorCode),
			Specification: req.Sku.Specification,
			DescLink:      req.Sku.DescLink,
			State:         int(req.Sku.State),
			CreateTime:    time.Now(),
			UpdateTime:    time.Now(),
		}
		tx := kelvins.XORM_DBEngine.NewSession()
		err = repository.CreateSkuProperty(tx, &skuProperty)
		if err != nil {
			tx.Rollback()
			kelvins.ErrLogger.Errorf(ctx, "CreateSkuProperty err: %v, skuProperty: %+v", err, skuProperty)
			retCode = code.ErrorServer
			return
		}
		skuInventory := mysql.SkuInventory{
			SkuCode:    req.Sku.SkuCode,
			Amount:     req.Sku.Amount,
			Price:      req.Sku.Price,
			ShopId:     req.Sku.ShopId,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		}
		err = repository.CreateSkuInventory(tx, &skuInventory)
		if err != nil {
			tx.Rollback()
			kelvins.ErrLogger.Errorf(ctx, "CreateSkuInventory err: %v, skuInventory: %+v", err, skuInventory)
			retCode = code.ErrorServer
			return
		}
		tx.Commit()
	}

	return
}

func GetSkuList(ctx context.Context, req *sku_business.GetSkuListRequest) (result []args.SkuInventoryInfo, retCode int) {
	retCode = code.Success
	result = make([]args.SkuInventoryInfo, 0)
	skuInventoryList, err := repository.GetSkuInventoryListByShopId(req.ShopId)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "GetSkuInventoryListByShopId err: %v, ShopId: %+v", err, req.ShopId)
		retCode = code.ErrorServer
		return
	}
	skuCodeList := make([]string, len(skuInventoryList))
	skuCodeToInventory := make(map[string]mysql.SkuInventory)
	for i := 0; i < len(skuInventoryList); i++ {
		skuCodeList[i] = skuInventoryList[i].SkuCode
		skuCodeToInventory[skuInventoryList[i].SkuCode] = skuInventoryList[i]
	}

	skuPropertyList, err := repository.GetSkuPropertyList(skuCodeList)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "GetSkuPropertyList err: %v, skuCodeList: %+v", err, skuCodeList)
		retCode = code.ErrorServer
		return
	}
	result = make([]args.SkuInventoryInfo, len(skuPropertyList))
	for i := 0; i < len(skuPropertyList); i++ {
		skuInventoryInfo := args.SkuInventoryInfo{
			SkuCode:       skuPropertyList[i].Code,
			Name:          skuPropertyList[i].Name,
			Price:         skuPropertyList[i].Price,
			Title:         skuPropertyList[i].Title,
			SubTitle:      skuPropertyList[i].SubTitle,
			Desc:          skuPropertyList[i].Desc,
			Production:    skuPropertyList[i].Production,
			Supplier:      skuPropertyList[i].Supplier,
			Category:      int32(skuPropertyList[i].Category),
			Color:         skuPropertyList[i].Color,
			ColorCode:     int32(skuPropertyList[i].ColorCode),
			Specification: skuPropertyList[i].Specification,
			DescLink:      skuPropertyList[i].DescLink,
			State:         int32(skuPropertyList[i].State),
			Amount:        skuCodeToInventory[skuPropertyList[i].Code].Amount,
			ShopId:        skuCodeToInventory[skuPropertyList[i].Code].ShopId,
		}
		result[i] = skuInventoryInfo
	}

	return
}

func SupplementSkuProperty(ctx context.Context, req *sku_business.SupplementSkuPropertyRequest) int {
	if req.ShopId > 0 {
		serverName := args.RpcServiceMicroMallShop
		conn, err := util.GetGrpcClient(serverName)
		if err != nil {
			kelvins.ErrLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
			return code.ErrorServer
		}
		defer conn.Close()
		client := shop_business.NewShopBusinessServiceClient(conn)
		r := shop_business.GetShopMaterialRequest{
			ShopId: req.ShopId,
		}
		rsp, err := client.GetShopMaterial(ctx, &r)
		if err != nil {
			kelvins.ErrLogger.Errorf(ctx, "GetShopMaterial %v,err: %v, req: %+v", serverName, err, r)
			return code.ErrorServer
		}
		if rsp == nil || rsp.Material == nil || rsp.Material.ShopId <= 0 {
			return code.ShopBusinessNotExist
		}
	}

	if req.OperationType == sku_business.OperationType_CREATE {
		skuExInfo := args.SkuPropertyEx{
			OpUid:             req.OpUid,
			OpIp:              req.OpIp,
			ShopId:            req.ShopId,
			SkuCode:           req.SkuCode,
			Name:              req.Name,
			Size:              req.Size,
			Shape:             req.Shape,
			ProductionCountry: req.ProductionCountry,
			ProductionDate:    req.ProductionDate,
			ShelfLife:         req.ShelfLife,
		}
		err := repository.CreateSkuPropertyEx(ctx, &skuExInfo)
		if err != nil {
			kelvins.ErrLogger.Errorf(ctx, "CreateSkuPropertyEx err: %v, skuExInfo: %+v", err, skuExInfo)
			return code.ErrorServer
		}
		return code.Success
	} else if req.OperationType == sku_business.OperationType_UPDATE {

	}

	return code.Success
}
