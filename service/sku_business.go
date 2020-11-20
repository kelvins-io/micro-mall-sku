package service

import (
	"context"
	"fmt"
	"gitee.com/cristiane/micro-mall-sku/model/args"
	"gitee.com/cristiane/micro-mall-sku/model/mysql"
	"gitee.com/cristiane/micro-mall-sku/pkg/code"
	"gitee.com/cristiane/micro-mall-sku/pkg/util"
	"gitee.com/cristiane/micro-mall-sku/proto/micro_mall_search_proto/search_business"
	"gitee.com/cristiane/micro-mall-sku/proto/micro_mall_shop_proto/shop_business"
	"gitee.com/cristiane/micro-mall-sku/proto/micro_mall_sku_proto/sku_business"
	"gitee.com/cristiane/micro-mall-sku/repository"
	"gitee.com/kelvins-io/kelvins"
	"strconv"
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
		exist, err := repository.CheckSkuInventoryExist(req.Sku.ShopId, req.Sku.SkuCode)
		if err != nil {
			kelvins.ErrLogger.Errorf(ctx, "CheckSkuInventoryExist %v,err: %v,ShopId: %v, SkuCode: %+v", err, req.Sku.ShopId, req.Sku.SkuCode)
			retCode = code.ErrorServer
			return
		}
		if exist {
			retCode = code.SkuCodeExist
			return
		}
		tx := kelvins.XORM_DBEngine.NewSession()
		err = tx.Begin()
		if err != nil {
			kelvins.ErrLogger.Errorf(ctx, "CreateSkuProperty Begin err: %v", err)
			retCode = code.ErrorServer
			return
		}
		// 存储商品属性-基本属性
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
		skuPropertyRecord := &mysql.SkuInventoryRecord{
			ShopId:       req.Sku.ShopId,
			SkuCode:      req.Sku.SkuCode,
			OpType:       0, // 入库
			OpUid:        req.OperationMeta.OpUid,
			OpIp:         req.OperationMeta.OpIp,
			AmountBefore: 0,
			Amount:       req.Sku.Amount,
			OpTxId:       fmt.Sprintf("%d-%s", req.Sku.ShopId, req.Sku.SkuCode),
			State:        0,
			CreateTime:   time.Now(),
			UpdateTime:   time.Now(),
		}
		err = repository.CreateSkuInventoryRecordByTx(tx, skuPropertyRecord)
		if err != nil {
			errRollback := tx.Rollback()
			if errRollback != nil {
				kelvins.ErrLogger.Errorf(ctx, "CreateSkuProperty Rollback err: %v", errRollback)
			}
			kelvins.ErrLogger.Errorf(ctx, "CreateSkuInventoryRecordByTx err: %v, skuPropertyRecord: %+v", err, skuPropertyRecord)
			retCode = code.ErrorServer
			return
		}
		// 插入价格历史
		skuPriceHistory := &mysql.SkuPriceHistory{
			ShopId:     req.Sku.ShopId,
			SkuCode:    req.Sku.SkuCode,
			Price:      req.Sku.Price,
			Reason:     "入库",
			Version:    1, // 入库版本为1
			OpUid:      req.OperationMeta.OpUid,
			OpIp:       req.OperationMeta.OpIp,
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
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
		// 增加库存记录
		skuInventory := mysql.SkuInventory{
			SkuCode:    req.Sku.SkuCode,
			Amount:     req.Sku.Amount,
			Price:      req.Sku.Price,
			ShopId:     req.Sku.ShopId,
			Version:    1, // 入库版本为1
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
		err = tx.Commit()
		if err != nil {
			kelvins.ErrLogger.Errorf(ctx, "CreateSkuProperty Commit err: %v", err)
			retCode = code.ErrorServer
			return
		}

		// 增加扩展属性
		go func() {
			err = repository.CreateSkuPropertyMongoDB(ctx, &skuProperty)
			if err != nil {
				kelvins.ErrLogger.Errorf(ctx, "CreateSkuPropertyEx err: %v, skuExInfo: %+v", err, skuProperty)
			}
		}()
		return code.Success
	} else if req.OperationType == sku_business.OperationType_UPDATE {
		exist, err := repository.CheckSkuInventoryExist(req.Sku.ShopId, req.Sku.SkuCode)
		if err != nil {
			kelvins.ErrLogger.Errorf(ctx, "CheckSkuInventoryExist %v,err: %v,ShopId: %v, SkuCode: %+v", err, req.Sku.ShopId, req.Sku.SkuCode)
			retCode = code.ErrorServer
			return
		}
		if !exist {
			retCode = code.SkuCodeNotExist
			return
		}
		// 增加库存

	}

	return
}

func getSkuList(ctx context.Context, shopId int64, pageSize, pageNum int) (result []*sku_business.SkuInventoryInfo, retCode int) {
	retCode = code.Success
	result = make([]*sku_business.SkuInventoryInfo, 0)
	skuInventoryList, err := repository.GetSkuInventoryListByShopId(shopId, pageSize, pageNum)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "GetSkuInventoryListByShopId err: %v, ShopId: %+v", err, shopId)
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
	result = make([]*sku_business.SkuInventoryInfo, len(skuPropertyList))
	for i := 0; i < len(skuPropertyList); i++ {
		skuInventoryInfo := &sku_business.SkuInventoryInfo{
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
			Version:       int64(skuCodeToInventory[skuPropertyList[i].Code].Version),
		}
		result[i] = skuInventoryInfo
	}

	return
}

func SearchSkuInventory(ctx context.Context, req *sku_business.SearchSkuInventoryRequest) ([]*sku_business.SearchSkuInventoryEntry, int) {
	result := make([]*sku_business.SearchSkuInventoryEntry, 0)
	serverName := args.RpcServiceMicroMallSearch
	conn, err := util.GetGrpcClient(serverName)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
		return result, code.ErrorServer
	}
	defer conn.Close()
	client := search_business.NewSearchBusinessServiceClient(conn)
	r := search_business.SkuInventorySearchRequest{
		SkuKey: req.Keyword,
	}
	rsp, err := client.SkuInventorySearch(ctx, &r)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "SkuInventorySearch %v,err: %v, req: %+v", serverName, err, r)
		return result, code.ErrorServer
	}
	if rsp.Common.Code != search_business.RetCode_SUCCESS {
		kelvins.ErrLogger.Errorf(ctx, "SkuInventorySearch %v,err: %v, rsp: %+v", serverName, err, rsp)
		return result, code.ErrorServer
	}
	if len(rsp.List) == 0 {
		return result, code.Success
	}
	shopIds := make([]int64, 0)
	skuCodes := make([]string, len(rsp.List))
	for i := 0; i < len(rsp.List); i++ {
		shopId, err := strconv.ParseInt(rsp.List[i].ShopId, 10, 64)
		if err != nil {
			kelvins.ErrLogger.Errorf(ctx, "SearchSkuInventory ParseI shopId err: %v, shopId: %s", err, rsp.List[i].ShopId)
			return result, code.ErrorServer
		}
		if shopId > 0 {
			shopIds = append(shopIds, shopId)
		}
		skuCodes[i] = rsp.List[i].SkuCode
	}
	skuInventoryList, err := repository.GetSkuInventoryList(shopIds, skuCodes)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "GetSkuInventoryList  err: %v, shopIds: %+v,skuCodes: %+v", err, shopIds, skuCodes)
		return result, code.ErrorServer
	}
	if len(skuInventoryList) == 0 {
		return result, code.Success
	}
	skuCodeToSkuInventory := make(map[string]*mysql.SkuInventory)
	for i := 0; i < len(skuInventoryList); i++ {
		skuCodeToSkuInventory[skuInventoryList[i].SkuCode] = skuInventoryList[i]
	}
	skuPropertyList, err := repository.GetSkuPropertyList(skuCodes)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "GetSkuPropertyList  err: %v, skuCodes: %+v", err, skuCodes)
		return result, code.ErrorServer
	}
	skuCodeToSkuProperty := make(map[string]mysql.SkuProperty)
	for i := 0; i < len(skuPropertyList); i++ {
		skuCodeToSkuProperty[skuPropertyList[i].Code] = skuPropertyList[i]
	}
	result = make([]*sku_business.SearchSkuInventoryEntry, len(rsp.List))
	for i := 0; i < len(rsp.List); i++ {
		row := rsp.List[i]
		shopId, _ := strconv.ParseInt(rsp.List[i].ShopId, 10, 0)
		entry := &sku_business.SearchSkuInventoryEntry{
			Info: &sku_business.SkuInventoryInfo{
				SkuCode:       row.SkuCode,
				Name:          row.SkuName,
				Price:         skuCodeToSkuProperty[row.SkuCode].Price,
				Title:         skuCodeToSkuProperty[row.SkuCode].Title,
				SubTitle:      skuCodeToSkuProperty[row.SkuCode].SubTitle,
				Desc:          skuCodeToSkuProperty[row.SkuCode].Desc,
				Production:    skuCodeToSkuProperty[row.SkuCode].Production,
				Supplier:      skuCodeToSkuProperty[row.SkuCode].Supplier,
				Category:      int32(skuCodeToSkuProperty[row.SkuCode].Category),
				Color:         skuCodeToSkuProperty[row.SkuCode].Color,
				ColorCode:     int32(skuCodeToSkuProperty[row.SkuCode].ColorCode),
				Specification: skuCodeToSkuProperty[row.SkuCode].Specification,
				DescLink:      skuCodeToSkuProperty[row.SkuCode].DescLink,
				State:         int32(skuCodeToSkuProperty[row.SkuCode].State),
				Version:       int64(skuCodeToSkuInventory[row.SkuCode].Version),
				Amount:        skuCodeToSkuInventory[row.SkuCode].Amount,
				ShopId:        shopId,
			},
			Score: row.Score,
		}
		result[i] = entry
	}
	return result, code.Success
}

func GetSkuList(ctx context.Context, req *sku_business.GetSkuListRequest) (result []*sku_business.SkuInventoryInfo, retCode int) {
	return getSkuList(ctx, req.ShopId, 0, 0)
}

func SyncSkuInventory(ctx context.Context, req *sku_business.SearchSyncSkuInventoryRequest) (result []*sku_business.SkuInventoryInfo, retCode int) {
	return getSkuList(ctx, req.ShopId, int(req.PageSize), int(req.PageNum))
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
	// 检查商品是否存在
	exist, err := repository.CheckSkuInventoryExist(req.ShopId, req.SkuCode)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "CheckSkuInventoryExist %v,err: %v,ShopId: %v, SkuCode: %+v", err, req.ShopId, req.SkuCode)
		return code.ErrorServer
	}
	if !exist {
		return code.SkuCodeNotExist
	}

	if req.OperationType == sku_business.OperationType_CREATE {
		skuExInfo := args.SkuPropertyEx{
			OpUid:             req.OperationMeta.OpUid,
			OpIp:              req.OperationMeta.OpIp,
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

func DeductInventory(ctx context.Context, req *sku_business.DeductInventoryRequest) (result *args.OperationInventoryRsp, retCode int) {
	result = &args.OperationInventoryRsp{List: make([]args.InventoryState, 0)}
	retCode = code.Success
	// 汇总商品
	allShopIdList := make([]int64, len(req.List))
	allSkuCodeList := make([]string, 0)
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
	allShopIdSkuCodeAmount := make(map[string]int64)
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
				// 记录库存扣减
				skuInventoryRecord := &mysql.SkuInventoryRecord{
					ShopId:       req.List[i].ShopId,
					SkuCode:      req.List[i].Detail[j].SkuCode,
					OpType:       1,
					OpUid:        req.OperationMeta.OpUid,
					OpIp:         req.OperationMeta.OpIp,
					AmountBefore: v,
					Amount:       req.List[i].Detail[j].Amount,
					OpTxId:       req.List[i].OutTradeNo,
					State:        0,
					CreateTime:   time.Now(),
					UpdateTime:   time.Now(),
				}
				err = repository.CreateSkuInventoryRecordByTx(tx, skuInventoryRecord)
				if err != nil {
					errRollback := tx.Rollback()
					if errRollback != nil {
						kelvins.ErrLogger.Errorf(ctx, "DeductInventory Rollback err: %v", errRollback)
					}
					kelvins.ErrLogger.Errorf(ctx, "CreateSkuInventoryRecordByTx err: %v, skuInventoryRecord: %v", err, skuInventoryRecord)
					retCode = code.ErrorServer
					return
				}
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
				rows, err := repository.UpdateInventory(tx, where, maps)
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
					retCode = code.TransactionFailed
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

func RestoreInventory(ctx context.Context, req *sku_business.RestoreInventoryRequest) (retCode int) {
	retCode = code.Success
	// 汇总商品
	allShopIdList := make([]int64, len(req.List))
	allSkuCodeList := make([]string, 0)
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
	allShopIdSkuCodeAmount := make(map[string]int64)
	for i := 0; i < len(inventoryList); i++ {
		key := fmt.Sprintf("%d-%s", inventoryList[i].ShopId, inventoryList[i].SkuCode)
		allShopIdSkuCodeAmount[key] = inventoryList[i].Amount // 如果shop_id + sku_code 是union key可以直接赋值
	}

	// 开始扣减库存
	tx := kelvins.XORM_DBEngine.NewSession()
	err = tx.Begin()
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "RestoreInventory Begin err: %v", err)
		retCode = code.ErrorServer
		return
	}
	for i := 0; i < len(req.List); i++ {
		allShopIdList[i] = req.List[i].ShopId
		if len(req.List[i].Detail) == 0 {
			continue
		}
		for j := 0; j < len(req.List[i].Detail); j++ {
			amountKey := fmt.Sprintf("%d-%s", req.List[i].ShopId, req.List[i].Detail[j].SkuCode)
			v, ok := allShopIdSkuCodeAmount[amountKey]
			if ok {
				// 记录库存扣减
				skuInventoryRecord := &mysql.SkuInventoryRecord{
					ShopId:       req.List[i].ShopId,
					SkuCode:      req.List[i].Detail[j].SkuCode,
					OpType:       3, // 恢复库存
					OpUid:        req.OperationMeta.OpUid,
					OpIp:         req.OperationMeta.OpIp,
					AmountBefore: v,
					Amount:       req.List[i].Detail[j].Amount,
					OpTxId:       req.List[i].OutTradeNo,
					State:        0,
					CreateTime:   time.Now(),
					UpdateTime:   time.Now(),
				}
				err = repository.CreateSkuInventoryRecordByTx(tx, skuInventoryRecord)
				if err != nil {
					errRollback := tx.Rollback()
					if errRollback != nil {
						kelvins.ErrLogger.Errorf(ctx, "RestoreInventory Rollback err: %v", errRollback)
					}
					kelvins.ErrLogger.Errorf(ctx, "CreateSkuInventoryRecordByTx err: %v, skuInventoryRecord: %v", err, skuInventoryRecord)
					retCode = code.ErrorServer
					return
				}
				// 使用乐观锁扣减库存
				where := map[string]interface{}{
					"shop_id":  req.List[i].ShopId,
					"sku_code": req.List[i].Detail[j].SkuCode,
					"amount":   v,
				}
				maps := map[string]interface{}{
					"amount":      v + req.List[i].Detail[j].Amount,
					"update_time": time.Now(),
				}
				rows, err := repository.UpdateInventory(tx, where, maps)
				if err != nil {
					errRollback := tx.Rollback()
					if errRollback != nil {
						kelvins.ErrLogger.Errorf(ctx, "RestoreInventory Rollback err: %v", errRollback)
					}
					kelvins.ErrLogger.Errorf(ctx, "RestoreInventory err: %v, where: %v, maps: %v", err, where, maps)
					retCode = code.ErrorServer
					return
				}
				if rows <= 0 {
					errRollback := tx.Rollback()
					if errRollback != nil {
						kelvins.ErrLogger.Errorf(ctx, "RestoreInventory Rollback err: %v", errRollback)
					}
					retCode = code.TransactionFailed
				}
			}
		}
	}
	err = tx.Commit()
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "RestoreInventory Commit err: %v", err)
		retCode = code.ErrorServer
		return
	}
	retCode = code.Success
	return
}

func FiltrateSkuPriceVersion(ctx context.Context, req *sku_business.FiltrateSkuPriceVersionRequest) ([]*sku_business.FiltrateSkuPriceVersionResult, int) {
	result := make([]*sku_business.FiltrateSkuPriceVersionResult, 0)
	retCode := code.Success
	// 价格策略
	switch req.PolicyType {
	case sku_business.SkuPricePolicyFiltrateType_VERSION_SECTION:
	case sku_business.SkuPricePolicyFiltrateType_VERSION_LOWER:
	case sku_business.SkuPricePolicyFiltrateType_VERSION_UPPER:
	default:
		return result, code.SkuPriceVersionPolicyNotSupport
	}
	// 查询价格是否在预定版本范围
	for i := 0; i < len(req.SetList); i++ {
		resultOne := &sku_business.FiltrateSkuPriceVersionResult{
			ShopId:  req.SetList[i].ShopId,
			SkuCode: make([]string, 0),
		}
		for j := 0; j < len(req.SetList[i].EntryList); j++ {
			row := req.SetList[i].EntryList[j]
			where := map[string]interface{}{
				"shop_id":  req.SetList[i].ShopId,
				"sku_code": row.SkuCode,
			}
			orderByDesc := []string{"version"}
			limit := int(req.LimitUpper)
			skuPriceHistoryList, err := repository.GetSkuPriceHistory("price,version", where, orderByDesc, limit)
			if err != nil {
				kelvins.ErrLogger.Errorf(ctx, "GetSkuPriceHistory err: %v, where: %+v,shopIdList: %v,skuCodeList:%v", err, where, nil, nil)
				return result, code.ErrorServer
			}
			if len(skuPriceHistoryList) == 0 {
				resultOne.SkuCode = append(resultOne.SkuCode, row.SkuCode)
				retCode = code.SkuPriceVersionNotExist
				continue
			}
			upperVersion := skuPriceHistoryList[len(skuPriceHistoryList)-1]
			lowerVersion := skuPriceHistoryList[0]
			switch req.PolicyType {
			case sku_business.SkuPricePolicyFiltrateType_VERSION_SECTION:
				if row.Version > int64(upperVersion.Version) || row.Version < int64(upperVersion.Version) {
					resultOne.SkuCode = append(resultOne.SkuCode, row.SkuCode)
					retCode = code.SkuPriceVersionNotExist
					continue
				}
			case sku_business.SkuPricePolicyFiltrateType_VERSION_LOWER:
				if row.Version < int64(lowerVersion.Version) {
					resultOne.SkuCode = append(resultOne.SkuCode, row.SkuCode)
					retCode = code.SkuPriceVersionNotExist
					continue
				}
			case sku_business.SkuPricePolicyFiltrateType_VERSION_UPPER:
				if row.Version > int64(upperVersion.Version) {
					resultOne.SkuCode = append(resultOne.SkuCode, row.SkuCode)
					retCode = code.SkuPriceVersionNotExist
					continue
				}
			}
		}
		if len(resultOne.SkuCode) > 0 {
			result = append(result, resultOne)
		}
	}

	return result, retCode
}
