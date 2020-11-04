package service

import (
	"context"
	"fmt"
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
		err = tx.Begin()
		if err != nil {
			kelvins.ErrLogger.Errorf(ctx, "CreateSkuProperty Begin err: %v", err)
			retCode = code.ErrorServer
			return
		}
		// 存储商品属性-基本属性
		err = repository.CreateSkuProperty(tx, &skuProperty)
		if err != nil {
			errRollback := tx.Rollback()
			if errRollback != nil {
				kelvins.ErrLogger.Errorf(ctx, "CreateSkuProperty Rollback err: %v", errRollback)
			}
			kelvins.ErrLogger.Errorf(ctx, "CreateSkuProperty err: %v, skuProperty: %+v", err, skuProperty)
			retCode = code.ErrorServer
			return
		}
		// 增加库存记录
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
			errRollback := tx.Rollback()
			if errRollback != nil {
				kelvins.ErrLogger.Errorf(ctx, "CreateSkuProperty Rollback err: %v", errRollback)
			}
			kelvins.ErrLogger.Errorf(ctx, "CreateSkuInventory err: %v, skuInventory: %+v", err, skuInventory)
			retCode = code.ErrorServer
			return
		}
		// 插入价格历史
		skuPriceHistory := &mysql.SkuPriceHistory{
			ShopId:     req.Sku.ShopId,
			SkuCode:    req.Sku.SkuCode,
			Price:      req.Sku.Price,
			Tsp:        int(time.Now().Unix()),
			Reason:     req.Sku.Name,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
			OpUid:      req.OpUid,
			OpIp:       req.OpIp,
		}
		err = repository.CreateSkuPriceHistory(tx, skuPriceHistory)
		if err != nil {
			errRollback := tx.Rollback()
			if errRollback != nil {
				kelvins.ErrLogger.Errorf(ctx, "CreateSkuProperty Rollback err: %v", errRollback)
			}
			kelvins.ErrLogger.Errorf(ctx, "CreateSkuPriceHistory err: %v, skuPriceHistory: %+v", err, skuPriceHistory)
			retCode = code.ErrorServer
			return
		}
		err = tx.Commit()
		if err != nil {
			kelvins.ErrLogger.Errorf(ctx, "CreateSkuProperty Commit err: %v", err)
			retCode = code.ErrorServer
			return
		}
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

func DeductInventory(ctx context.Context, req *sku_business.DeductInventoryRequest) (result *args.DeductInventoryRsp, retCode int) {
	result = &args.DeductInventoryRsp{List: make([]args.InventoryState, 0)}
	retCode = code.Success

	// 汇总商品
	allShopIdList := make([]int64, len(req.List))
	allSkuCodeList := make([]string, 0)
	allShopIdSkuCodeAmount := make(map[string]int64)
	for i := 0; i < len(req.List); i++ {
		allShopIdList[i] = req.List[i].ShopId
		if len(req.List[i].Detail) == 0 {
			continue
		}
		skuCodeList := make([]string, len(req.List[i].Detail))
		for j := 0; j < len(req.List[i].Detail); j++ {
			skuCodeList[j] = req.List[i].Detail[j].SkuCode
		}
		allSkuCodeList = append(allSkuCodeList, skuCodeList...)
	}

	// 从DB里面取出这些商品
	inventoryList, err := repository.GetSkuInventoryList(allShopIdList, allSkuCodeList)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "GetSkuInventoryList err: %v, allShopIdList: %v, skuCodeList: %v", err, allShopIdList, allSkuCodeList)
		retCode = code.ErrorServer
		return
	}

	// 收集数据库中商品剩余数量
	for i := 0; i < len(inventoryList); i++ {
		key := fmt.Sprintf("%d-%s", inventoryList[i].ShopId, inventoryList[i].SkuCode)
		allShopIdSkuCodeAmount[key] = inventoryList[i].Amount // 如果shop_id + sku_code 是union key可以直接赋值
	}

	// 统计哪些商品不够数量
	inventoryState := make(map[int64][]string)
	for i := 0; i < len(req.List); i++ {
		allShopIdList[i] = req.List[i].ShopId
		if len(req.List[i].Detail) == 0 {
			continue
		}
		// 依赖于请求数据正常排序，同一个店铺的商品聚合在一起
		for j := 0; j < len(req.List[i].Detail); j++ {
			// 判断请求中的商品是否达到购买条件
			amountKey := fmt.Sprintf("%d-%s", req.List[i].ShopId, req.List[i].Detail[j].SkuCode)
			v, ok := allShopIdSkuCodeAmount[amountKey]
			if !ok {
				inventoryState[req.List[i].ShopId] = append(inventoryState[req.List[i].ShopId], req.List[i].Detail[j].SkuCode)
			} else {
				// 如果数据库中sku数量小于要购买的数量
				if v < req.List[i].Detail[j].Amount {
					inventoryState[req.List[i].ShopId] = append(inventoryState[req.List[i].ShopId], req.List[i].Detail[j].SkuCode)
				}
			}
		}
	}
	// 检查是否有不满足购买条件的商品
	if len(inventoryState) > 0 {
		result.List = make([]args.InventoryState, 0)
		for k, v := range inventoryState {
			state := args.InventoryState{
				ShopId:   k,
				SkuCodes: v,
			}
			result.List = append(result.List, state)
		}
		return result, code.SkuAmountNotEnough
	}

	// 开始扣减库存
	result.List = make([]args.InventoryState, 0)
	tx := kelvins.XORM_DBEngine.NewSession()
	err = tx.Begin()
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "DeductInventory Begin err: %v", err)
		retCode = code.ErrorServer
		return
	}
	for i := 0; i < len(req.List); i++ {
		allShopIdList[i] = req.List[i].ShopId
		if len(req.List[i].Detail) == 0 {
			continue
		}
		inventoryState = make(map[int64][]string)
		for j := 0; j < len(req.List[i].Detail); j++ {
			amountKey := fmt.Sprintf("%d-%s", req.List[i].ShopId, req.List[i].Detail[j].SkuCode)
			v, ok := allShopIdSkuCodeAmount[amountKey]
			if ok {
				// 使用乐观锁扣减库存
				where := map[string]interface{}{
					"shop_id":  req.List[i].ShopId,
					"sku_code": req.List[i].Detail[j].SkuCode,
					"amount":   v,
				}
				maps := map[string]interface{}{
					"amount":      v - req.List[i].Detail[j].Amount,
					"update_time": time.Now(),
				}
				rows, err := repository.DeductInventory(tx, where, maps)
				if err != nil {
					errRollback := tx.Rollback()
					if errRollback != nil {
						kelvins.ErrLogger.Errorf(ctx, "DeductInventory Rollback err: %v", errRollback)
					}
					kelvins.ErrLogger.Errorf(ctx, "DeductInventory err: %v, where: %v, maps: %v", err, where, maps)
					retCode = code.ErrorServer
					return
				}
				if rows <= 0 {
					errRollback := tx.Rollback()
					if errRollback != nil {
						kelvins.ErrLogger.Errorf(ctx, "DeductInventory Rollback err: %v", errRollback)
					}
					retCode = code.SkuAmountNotEnough
					inventoryState[req.List[i].ShopId] = append(inventoryState[req.List[i].ShopId], req.List[i].Detail[j].SkuCode)
				}
			}
		}
		if len(inventoryState) > 0 {
			for k, v := range inventoryState {
				state := args.InventoryState{
					ShopId:   k,
					SkuCodes: v,
				}
				result.List = append(result.List, state)
			}
			return result, code.SkuAmountNotEnough
		}
	}
	err = tx.Commit()
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "DeductInventory Commit err: %v", err)
		retCode = code.ErrorServer
		return
	}
	retCode = code.Success
	return
}
