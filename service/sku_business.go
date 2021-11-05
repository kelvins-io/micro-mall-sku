package service

import (
	"context"
	"fmt"
	"gitee.com/cristiane/micro-mall-sku/model/args"
	"gitee.com/cristiane/micro-mall-sku/model/mysql"
	"gitee.com/cristiane/micro-mall-sku/pkg/code"
	"gitee.com/cristiane/micro-mall-sku/pkg/util"
	"gitee.com/cristiane/micro-mall-sku/proto/micro_mall_search_proto/search_business"
	"gitee.com/cristiane/micro-mall-sku/proto/micro_mall_sku_proto/sku_business"
	"gitee.com/cristiane/micro-mall-sku/repository"
	"gitee.com/kelvins-io/common/json"
	"gitee.com/kelvins-io/kelvins"
	"github.com/google/uuid"
	"strconv"
	"strings"
	"time"
)

func PutAwaySku(ctx context.Context, req *sku_business.PutAwaySkuRequest) (retCode int) {
	retCode = code.Success
	if req.OperationType == sku_business.OperationType_CREATE {
		exist, err := repository.CheckSkuInventoryExist(req.Sku.ShopId, req.Sku.SkuCode)
		if err != nil {
			kelvins.ErrLogger.Errorf(ctx, "CheckSkuInventoryExist %v,err: %v,ShopId: %v, SkuCode: %v", err, req.Sku.ShopId, req.Sku.SkuCode)
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
		defer func() {
			if retCode != code.Success {
				err := tx.Rollback()
				if err != nil {
					kelvins.ErrLogger.Errorf(ctx, "CreateSkuProperty Rollback err: %v", err)
					return
				}
			}
		}()
		opTxId := uuid.New().String()
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
			kelvins.ErrLogger.Errorf(ctx, "CreateSkuProperty err: %v, skuProperty: %v", err, json.MarshalToStringNoError(skuProperty))
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
			OpTxId:       opTxId,
			State:        0,
			Verify:       1,
			CreateTime:   time.Now(),
			UpdateTime:   time.Now(),
		}
		err = repository.CreateSkuInventoryRecordByTx(tx, skuPropertyRecord)
		if err != nil {
			kelvins.ErrLogger.Errorf(ctx, "CreateSkuInventoryRecordByTx err: %v, skuPropertyRecord: %v", err, json.MarshalToStringNoError(skuPropertyRecord))
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
			kelvins.ErrLogger.Errorf(ctx, "CreateSkuPriceHistory err: %v, skuPriceHistory: %v", err, json.MarshalToStringNoError(skuPriceHistory))
			retCode = code.ErrorServer
			return
		}
		// 增加库存记录
		skuInventory := mysql.SkuInventory{
			SkuCode:    req.Sku.SkuCode,
			Amount:     req.Sku.Amount,
			Price:      req.Sku.Price,
			ShopId:     req.Sku.ShopId,
			LastTxId:   opTxId,
			Version:    1, // 入库版本为1
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		}
		err = repository.CreateSkuInventory(tx, &skuInventory)
		if err != nil {
			kelvins.ErrLogger.Errorf(ctx, "CreateSkuInventory err: %v, skuInventory: %v", err, json.MarshalToStringNoError(skuInventory))
			retCode = code.ErrorServer
			return
		}
		err = tx.Commit()
		if err != nil {
			kelvins.ErrLogger.Errorf(ctx, "CreateSkuProperty Commit err: %v", err)
			retCode = code.TransactionFailed
			return
		}
		// 增加扩展属性
		kelvins.GPool.SendJob(func() {
			err = repository.CreateSkuPropertyMongoDB(ctx, &skuProperty)
			if err != nil {
				kelvins.ErrLogger.Errorf(ctx, "CreateSkuPropertyEx err: %v, skuExInfo: %v", err, json.MarshalToStringNoError(skuProperty))
			}
		})
		body := &args.SkuInventoryInfo{
			ShopId:        req.GetSku().GetShopId(),
			SkuCode:       req.GetSku().GetSkuCode(),
			Name:          req.GetSku().GetName(),
			Price:         req.GetSku().GetPrice(),
			Title:         req.GetSku().GetTitle(),
			SubTitle:      req.GetSku().GetSubTitle(),
			Desc:          req.GetSku().GetDesc(),
			Production:    req.GetSku().GetProduction(),
			Supplier:      req.GetSku().GetSupplier(),
			Category:      req.GetSku().GetCategory(),
			Color:         req.GetSku().GetColor(),
			ColorCode:     req.GetSku().GetColorCode(),
			Specification: req.GetSku().GetSpecification(),
			DescLink:      req.GetSku().GetDescLink(),
		}
		_ = skuInventorySearch(body)
		return
	} else if req.OperationType == sku_business.OperationType_PUT_AWAY {
		record, err := repository.GetSkuInventory("id,amount,last_tx_id", req.Sku.ShopId, req.Sku.SkuCode)
		if err != nil {
			kelvins.ErrLogger.Errorf(ctx, "CheckSkuInventoryExist err: %v, ShopId: %v, SkuCode: %v", err, req.Sku.ShopId, req.Sku.SkuCode)
			retCode = code.ErrorServer
			return
		}
		if record.Id <= 0 {
			retCode = code.SkuCodeNotExist
			return
		}
		tx := kelvins.XORM_DBEngine.NewSession()
		err = tx.Begin()
		if err != nil {
			kelvins.ErrLogger.Errorf(ctx, "CreateSkuInventoryRecord Begin err: %v", err)
			retCode = code.ErrorServer
			return
		}
		defer func() {
			if retCode != code.Success {
				err := tx.Rollback()
				if err != nil {
					kelvins.ErrLogger.Errorf(ctx, "CreateSkuInventoryRecord Rollback err: %v", err)
					return
				}
			}
		}()
		opTxId := uuid.New().String()
		skuInventoryRecord := mysql.SkuInventoryRecord{
			ShopId:       req.Sku.ShopId,
			SkuCode:      req.Sku.SkuCode,
			OpType:       0, // 入库
			OpUid:        req.OperationMeta.OpUid,
			OpIp:         req.OperationMeta.OpIp,
			AmountBefore: record.Amount,
			Amount:       req.Sku.Amount,
			OpTxId:       opTxId,
			State:        0,
			Verify:       1,
			CreateTime:   time.Now(),
			UpdateTime:   time.Now(),
		}
		err = repository.CreateSkuInventoryRecordByTx(tx, &skuInventoryRecord)
		if err != nil {
			kelvins.ErrLogger.Errorf(ctx, "CreateSkuInventoryRecordByTx  err: %v, record: %v", err, json.MarshalToStringNoError(skuInventoryRecord))
			retCode = code.ErrorServer
			return
		}
		updateSkuInventoryWhere := map[string]interface{}{
			"shop_id":    req.Sku.ShopId,
			"sku_code":   req.Sku.SkuCode,
			"amount":     record.Amount,
			"last_tx_id": record.LastTxId,
		}
		updateSkuInventoryMaps := map[string]interface{}{
			"amount":      record.Amount + req.Sku.Amount,
			"last_tx_id":  opTxId,
			"update_time": time.Now(),
		}
		rowAffected, err := repository.UpdateInventory(tx, updateSkuInventoryWhere, updateSkuInventoryMaps)
		if err != nil {
			kelvins.ErrLogger.Errorf(ctx, "UpdateInventory  err: %v, where: %v, maps: %v", err, json.MarshalToStringNoError(updateSkuInventoryWhere), json.MarshalToStringNoError(updateSkuInventoryMaps))
			retCode = code.ErrorServer
			return
		}
		if rowAffected != 1 {
			retCode = code.TransactionFailed
			return
		}
		err = tx.Commit()
		if err != nil {
			kelvins.ErrLogger.Errorf(ctx, "UpdateInventory Commit err: %v", err)
			retCode = code.TransactionFailed
			return
		}
	}

	return
}

func getSkuList(ctx context.Context, shopId int64, pageSize, pageNum int) (result []*sku_business.SkuInventoryInfo, retCode int) {
	retCode = code.Success
	result = make([]*sku_business.SkuInventoryInfo, 0)
	skuInventoryList, err := repository.GetSkuInventoryListByShopId(sqlSelectSkuInventory, shopId, pageSize, pageNum)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "GetSkuInventoryListByShopId err: %v, ShopId: %v", err, shopId)
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
		kelvins.ErrLogger.Errorf(ctx, "GetSkuPropertyList err: %v, skuCodeList: %v", err, json.MarshalToStringNoError(skuCodeList))
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

func SearchSkuInventory(ctx context.Context, req *sku_business.SearchSkuInventoryRequest) (result []*sku_business.SearchSkuInventoryEntry, retCode int) {
	result = make([]*sku_business.SearchSkuInventoryEntry, 0)
	retCode = code.Success
	searchKey := "micro-mall-sku:search-sku:" + req.GetKeyword()
	err := kelvins.G2CacheEngine.Get(searchKey, 120, &result, func() (interface{}, error) {
		ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
		defer cancel()
		list, ret := searchSkuInventory(ctx, req)
		if ret != code.Success {
			return &list, fmt.Errorf("searchSkuInventory ret %v", ret)
		}
		return &list, nil
	})
	if err != nil {
		retCode = code.ErrorServer
		return
	}
	return
}

const sqlSelectSkuInventory = "shop_id,sku_code,amount,version,last_tx_id"

func searchSkuInventory(ctx context.Context, req *sku_business.SearchSkuInventoryRequest) ([]*sku_business.SearchSkuInventoryEntry, int) {
	result := make([]*sku_business.SearchSkuInventoryEntry, 0)
	serverName := args.RpcServiceMicroMallSearch
	conn, err := util.GetGrpcClient(ctx, serverName)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
		return result, code.ErrorServer
	}
	//defer conn.Close()
	client := search_business.NewSearchBusinessServiceClient(conn)
	r := search_business.SkuInventorySearchRequest{
		SkuKey: req.Keyword,
	}
	rsp, err := client.SkuInventorySearch(ctx, &r)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "SkuInventorySearch %v,err: %v, req: %v", serverName, err, json.MarshalToStringNoError(r))
		return result, code.ErrorServer
	}
	if rsp.Common.Code != search_business.RetCode_SUCCESS {
		kelvins.ErrLogger.Errorf(ctx, "SkuInventorySearch req: %v, rsp: %v", json.MarshalToStringNoError(r), json.MarshalToStringNoError(rsp))
		return result, code.ErrorServer
	}
	if len(rsp.List) == 0 {
		return result, code.Success
	}
	shopIds := make([]int64, 0)
	skuCodes := make([]string, len(rsp.List))
	for i := range rsp.List {
		if rsp.List[i].ShopId == "" {
			continue
		}
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
	skuInventoryList, err := repository.GetSkuInventoryList(sqlSelectSkuInventory, shopIds, skuCodes)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "GetSkuInventoryList  err: %v, shopIds: %+v,skuCodes: %v", err, shopIds, json.MarshalToStringNoError(skuCodes))
		return result, code.ErrorServer
	}
	if len(skuInventoryList) == 0 {
		return result, code.Success
	}
	skuCodeToSkuInventory := make(map[string]mysql.SkuInventory)
	for i := 0; i < len(skuInventoryList); i++ {
		skuCodeToSkuInventory[skuInventoryList[i].SkuCode] = *skuInventoryList[i]
	}
	if len(skuCodes) == 0 {
		return result, code.Success
	}
	skuPropertyList, err := repository.GetSkuPropertyList(skuCodes)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "GetSkuPropertyList  err: %v, skuCodes: %v", err, json.MarshalToStringNoError(skuCodes))
		return result, code.ErrorServer
	}
	skuCodeToSkuProperty := make(map[string]mysql.SkuProperty)
	for i := 0; i < len(skuPropertyList); i++ {
		skuCodeToSkuProperty[skuPropertyList[i].Code] = skuPropertyList[i]
	}
	result = make([]*sku_business.SearchSkuInventoryEntry, 0, len(rsp.List))
	for i := 0; i < len(rsp.List); i++ {
		row := rsp.List[i]
		if rsp.List[i].ShopId == "" {
			continue
		}
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
		result = append(result, entry)
	}
	return result, code.Success
}

func GetSkuList(ctx context.Context, req *sku_business.GetSkuListRequest) (result []*sku_business.SkuInventoryInfo, retCode int) {
	retCode = code.Success
	result = make([]*sku_business.SkuInventoryInfo, 0)

	shopIdList := make([]int64, 0)
	if req.GetShopId() > 0 {
		shopIdList = append(shopIdList, req.GetShopId())
	}
	skuInventoryList, err := repository.GetSkuInventoryList("*", shopIdList, req.GetSkuCodeList())
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "GetSkuInventoryList err: %v,ShopId: %v, SkuCode: %v", err, req.GetShopId(), req.GetSkuCodeList())
		retCode = code.ErrorServer
		return
	}
	if len(skuInventoryList) == 0 {
		return
	}
	skuCodeToShop := make(map[string]mysql.SkuInventory, len(skuInventoryList))
	skuCodeList := make([]string, 0, len(skuInventoryList))
	for _, v := range skuInventoryList {
		skuCodeToShop[v.SkuCode] = *v
		skuCodeList = append(skuCodeList, v.SkuCode)
	}
	skuCodePropertyList, err := repository.GetSkuPropertyList(skuCodeList)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "GetSkuPropertyList err: %v, SkuCode: %v", err, json.MarshalToStringNoError(skuCodeList))
		retCode = code.ErrorServer
		return
	}
	if len(skuCodePropertyList) == 0 {
		return
	}

	result = make([]*sku_business.SkuInventoryInfo, len(skuCodePropertyList))
	for i := 0; i < len(skuCodePropertyList); i++ {
		skuInventoryInfo := &sku_business.SkuInventoryInfo{
			SkuCode:       skuCodePropertyList[i].Code,
			Name:          skuCodePropertyList[i].Name,
			Price:         skuCodePropertyList[i].Price,
			Title:         skuCodePropertyList[i].Title,
			SubTitle:      skuCodePropertyList[i].SubTitle,
			Desc:          skuCodePropertyList[i].Desc,
			Production:    skuCodePropertyList[i].Production,
			Supplier:      skuCodePropertyList[i].Supplier,
			Category:      int32(skuCodePropertyList[i].Category),
			Color:         skuCodePropertyList[i].Color,
			ColorCode:     int32(skuCodePropertyList[i].ColorCode),
			Specification: skuCodePropertyList[i].Specification,
			DescLink:      skuCodePropertyList[i].DescLink,
			State:         int32(skuCodePropertyList[i].State),
			Amount:        skuCodeToShop[skuCodePropertyList[i].Code].Amount,
			ShopId:        skuCodeToShop[skuCodePropertyList[i].Code].ShopId,
			Version:       int64(skuCodeToShop[skuCodePropertyList[i].Code].Version),
		}
		result[i] = skuInventoryInfo
	}
	return
}

func SyncSkuInventory(ctx context.Context, req *sku_business.SearchSyncSkuInventoryRequest) (result []*sku_business.SkuInventoryInfo, retCode int) {
	return getSkuList(ctx, req.ShopId, int(req.PageSize), int(req.PageNum))
}

func SupplementSkuProperty(ctx context.Context, req *sku_business.SupplementSkuPropertyRequest) int {
	// 检查商品是否存在
	exist, err := repository.CheckSkuInventoryExist(req.ShopId, req.SkuCode)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "CheckSkuInventoryExist err: %v,ShopId: %v, SkuCode: %+v", err, req.ShopId, req.SkuCode)
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
			kelvins.ErrLogger.Errorf(ctx, "CreateSkuPropertyEx err: %v, skuExInfo: %v", err, json.MarshalToStringNoError(skuExInfo))
			if strings.HasPrefix(err.Error(), "multiple write errors") {
				return code.SkuCodeExist
			}
			return code.ErrorServer
		}
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
	allOutTradeList := make([]string, len(req.List))
	for i := 0; i < len(req.List); i++ {
		allShopIdList[i] = req.List[i].ShopId
		allOutTradeList[i] = req.List[i].OutTradeNo
		if len(req.List[i].Detail) == 0 {
			continue
		}
		skuCodeList := make([]string, len(req.List[i].Detail))
		for j := 0; j < len(req.List[i].Detail); j++ {
			skuCodeList[j] = req.List[i].Detail[j].SkuCode
		}
		allSkuCodeList = append(allSkuCodeList, skuCodeList...)
	}
	// 取出外部订单号
	checkRecordWhere := map[string]interface{}{
		"out_trade_no": allOutTradeList,
	}
	checkRecordList, err := repository.FindSkuInventoryRecord("id", checkRecordWhere)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "FindSkuInventoryRecord err: %v, checkRecordWhere: %v", err, json.MarshalToStringNoError(checkRecordWhere))
		retCode = code.ErrorServer
		return
	}
	// 判断当前订单号是否已有库存记录--防止重复扣减
	if len(checkRecordList) > 0 {
		retCode = code.DeductInventoryRecordExist
		return
	}
	// 从DB里面取出这些商品
	inventoryList, err := repository.GetSkuInventoryList(sqlSelectSkuInventory, allShopIdList, allSkuCodeList)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "GetSkuInventoryList err: %v, allShopIdList: %v, skuCodeList: %v",
			err, json.MarshalToStringNoError(allShopIdList), json.MarshalToStringNoError(allSkuCodeList))
		retCode = code.ErrorServer
		return
	}

	// 收集数据库中商品剩余数量
	allShopIdSkuCodeAmount := make(map[string]int64)
	allShopIdSkuCodeLastTxId := make(map[string]string)
	for i := 0; i < len(inventoryList); i++ {
		key := fmt.Sprintf("%d-%s", inventoryList[i].ShopId, inventoryList[i].SkuCode)
		allShopIdSkuCodeAmount[key] = inventoryList[i].Amount // 如果shop_id + sku_code 是union key可以直接赋值
		allShopIdSkuCodeLastTxId[key] = inventoryList[i].LastTxId
	}

	// 统计哪些商品不够数量
	inventoryState := make(map[int64][]string)
	for i := 0; i < len(req.List); i++ {
		//allShopIdList[i] = req.List[i].ShopId
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
	defer func() {
		if retCode != code.Success {
			err := tx.Rollback()
			if err != nil {
				kelvins.ErrLogger.Errorf(ctx, "DeductInventory Rollback err: %v", err)
				return
			}
		}
	}()
	for i := 0; i < len(req.List); i++ {
		//allShopIdList[i] = req.List[i].ShopId
		if len(req.List[i].Detail) == 0 {
			continue
		}
		inventoryState = make(map[int64][]string)
		for j := 0; j < len(req.List[i].Detail); j++ {
			opTxId := uuid.New().String()
			amountKey := fmt.Sprintf("%d-%s", req.List[i].ShopId, req.List[i].Detail[j].SkuCode)
			lastTxIdKey := amountKey
			v, ok := allShopIdSkuCodeAmount[amountKey]
			if ok {
				// 记录库存扣减
				skuInventoryRecord := &mysql.SkuInventoryRecord{
					ShopId:       req.List[i].ShopId,
					SkuCode:      req.List[i].Detail[j].SkuCode,
					OutTradeNo:   req.List[i].OutTradeNo,
					OpType:       1, // 出库
					OpUid:        req.OperationMeta.OpUid,
					OpIp:         req.OperationMeta.OpIp,
					AmountBefore: v,
					Amount:       req.List[i].Detail[j].Amount,
					OpTxId:       opTxId,
					State:        0,
					Verify:       0,
					CreateTime:   time.Now(),
					UpdateTime:   time.Now(),
				}
				err = repository.CreateSkuInventoryRecordByTx(tx, skuInventoryRecord)
				if err != nil {
					kelvins.ErrLogger.Errorf(ctx, "CreateSkuInventoryRecordByTx err: %v, skuInventoryRecord: %v",
						err, json.MarshalToStringNoError(skuInventoryRecord))
					retCode = code.ErrorServer
					return
				}
				// 使用乐观锁扣减库存
				where := map[string]interface{}{
					"shop_id":    req.List[i].ShopId,
					"sku_code":   req.List[i].Detail[j].SkuCode,
					"amount":     v,
					"last_tx_id": allShopIdSkuCodeLastTxId[lastTxIdKey],
				}
				maps := map[string]interface{}{
					"amount":      v - req.List[i].Detail[j].Amount,
					"last_tx_id":  opTxId,
					"update_time": time.Now(),
				}
				rows, err := repository.UpdateInventory(tx, where, maps)
				if err != nil {
					kelvins.ErrLogger.Errorf(ctx, "DeductInventory err: %v, where: %v, maps: %v",
						err, json.MarshalToStringNoError(where), json.MarshalToStringNoError(maps))
					retCode = code.ErrorServer
					return
				}
				if rows != 1 {
					retCode = code.TransactionFailed
					inventoryState[req.List[i].ShopId] = append(inventoryState[req.List[i].ShopId], req.List[i].Detail[j].SkuCode)
					return
				}
				// 通过map key直接修改vale是允许的，但是value不是基本类型或非指针类型则不允许
				allShopIdSkuCodeAmount[amountKey] -= req.List[i].Detail[j].Amount
				allShopIdSkuCodeLastTxId[lastTxIdKey] = opTxId
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
		retCode = code.TransactionFailed
		return
	}
	retCode = code.Success
	return
}

func RestoreInventory(ctx context.Context, req *sku_business.RestoreInventoryRequest) (retCode int) {
	retCode = code.Success
	// 汇总商品
	allShopIdList := make([]int64, 0)
	allShopIdListSet := map[int64]struct{}{}
	allSkuCodeList := make([]string, 0)
	allSkuCodeListSet := map[string]struct{}{}
	allOutTradeList := make([]string, len(req.List))
	for i := 0; i < len(req.List); i++ {
		if _, ok := allShopIdListSet[req.List[i].ShopId]; !ok {
			allShopIdListSet[req.List[i].ShopId] = struct{}{}
			allShopIdList = append(allShopIdList, req.List[i].ShopId)
		}
		allOutTradeList[i] = req.List[i].OutTradeNo
		if len(req.List[i].Detail) == 0 {
			continue
		}
		skuCodeList := make([]string, 0)
		for j := 0; j < len(req.List[i].Detail); j++ {
			if _, ok := allSkuCodeListSet[req.List[i].Detail[j].SkuCode]; !ok {
				skuCodeList = append(skuCodeList, req.List[i].Detail[j].SkuCode)
				allSkuCodeListSet[req.List[i].Detail[j].SkuCode] = struct{}{}
			}
		}
		allSkuCodeList = append(allSkuCodeList, skuCodeList...)
	}
	// 取出外部订单号
	checkRecordWhere := map[string]interface{}{
		"out_trade_no": allOutTradeList,
		"op_type":      3, // 恢复
		"verify":       1, // 未核实的才能恢复
	}
	checkRecordList, err := repository.FindSkuInventoryRecord("id", checkRecordWhere)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "FindSkuInventoryRecord err: %v, checkRecordWhere: %v", err, json.MarshalToStringNoError(checkRecordWhere))
		retCode = code.ErrorServer
		return
	}
	// 当前订单号是否已经恢复过库存
	if len(checkRecordList) != 0 {
		retCode = code.RestoreInventoryRecordExist
		return
	}
	// 从DB里面取出这些商品
	inventoryList, err := repository.GetSkuInventoryList(sqlSelectSkuInventory, allShopIdList, allSkuCodeList)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "GetSkuInventoryList err: %v, allShopIdList: %v, skuCodeList: %v",
			err, json.MarshalToStringNoError(allShopIdList), json.MarshalToStringNoError(allSkuCodeList))
		retCode = code.ErrorServer
		return
	}
	// 收集数据库中商品剩余数量
	allShopIdSkuCodeAmount := make(map[string]int64)
	allShopIdSkuCodeLastTxId := make(map[string]string)
	for i := 0; i < len(inventoryList); i++ {
		key := fmt.Sprintf("%d-%s", inventoryList[i].ShopId, inventoryList[i].SkuCode)
		allShopIdSkuCodeAmount[key] = inventoryList[i].Amount // 如果shop_id + sku_code 是union key可以直接赋值
		allShopIdSkuCodeLastTxId[key] = inventoryList[i].LastTxId
	}
	// 开始扣减库存
	tx := kelvins.XORM_DBEngine.NewSession()
	err = tx.Begin()
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "RestoreInventory Begin err: %v", err)
		retCode = code.ErrorServer
		return
	}
	defer func() {
		if retCode != code.Success {
			err := tx.Rollback()
			if err != nil {
				kelvins.ErrLogger.Errorf(ctx, "RestoreInventory Rollback err: %v", err)
				return
			}
		}
	}()
	for i := 0; i < len(req.List); i++ {
		if len(req.List[i].Detail) == 0 {
			continue
		}
		for j := 0; j < len(req.List[i].Detail); j++ {
			opTxId := uuid.New().String()
			amountKey := fmt.Sprintf("%d-%s", req.List[i].ShopId, req.List[i].Detail[j].SkuCode)
			lastTxIdKey := amountKey
			v, ok := allShopIdSkuCodeAmount[amountKey]
			if ok {
				// 记录恢复记录
				skuInventoryRecord := &mysql.SkuInventoryRecord{
					ShopId:       req.List[i].ShopId,
					SkuCode:      req.List[i].Detail[j].SkuCode,
					OutTradeNo:   req.List[i].OutTradeNo,
					OpType:       3, // 恢复库存
					OpUid:        req.OperationMeta.OpUid,
					OpIp:         req.OperationMeta.OpIp,
					AmountBefore: v,
					Amount:       req.List[i].Detail[j].Amount,
					OpTxId:       opTxId,
					State:        0,
					Verify:       1,
					CreateTime:   time.Now(),
					UpdateTime:   time.Now(),
				}
				err = repository.CreateSkuInventoryRecordByTx(tx, skuInventoryRecord)
				if err != nil {
					kelvins.ErrLogger.Errorf(ctx, "CreateSkuInventoryRecordByTx err: %v, skuInventoryRecord: %v",
						err, json.MarshalToStringNoError(skuInventoryRecord))
					retCode = code.ErrorServer
					return
				}
				// 更新扣减记录
				updateRecordWhere := map[string]interface{}{
					"out_trade_no": req.List[i].OutTradeNo,
					"op_type":      1, // 出库
					"verify":       0, // 未核实的才能恢复
				}
				updateRecordMaps := map[string]interface{}{
					"verify":      1, // 已核实
					"op_tx_id":    opTxId,
					"update_time": time.Now(),
				}
				rows, err := repository.UpdateSkuInventoryRecordByTx(tx, updateRecordWhere, updateRecordMaps)
				if err != nil {
					kelvins.ErrLogger.Errorf(ctx, "UpdateSkuInventoryRecordByTx err: %v, where: %v, maps: %v",
						err, json.MarshalToStringNoError(updateRecordWhere), json.MarshalToStringNoError(updateRecordMaps))
					retCode = code.ErrorServer
					return
				}
				//if rows != 1 {
				//	errRollback := tx.Rollback()
				//	if errRollback != nil {
				//		kelvins.ErrLogger.Errorf(ctx, "UpdateSkuInventoryRecordByTx Rollback err: %v", errRollback)
				//	}
				//	retCode = code.TransactionFailed
				//	return
				//}
				// 使用乐观锁扣减库存
				where := map[string]interface{}{
					"shop_id":    req.List[i].ShopId,
					"sku_code":   req.List[i].Detail[j].SkuCode,
					"amount":     v,
					"last_tx_id": allShopIdSkuCodeLastTxId[lastTxIdKey],
				}
				maps := map[string]interface{}{
					"amount":      v + req.List[i].Detail[j].Amount,
					"last_tx_id":  opTxId,
					"update_time": time.Now(),
				}
				rows, err = repository.UpdateInventory(tx, where, maps)
				if err != nil {
					kelvins.ErrLogger.Errorf(ctx, "RestoreInventory err: %v, where: %v, maps: %v",
						err, json.MarshalToStringNoError(where), json.MarshalToStringNoError(maps))
					retCode = code.ErrorServer
					return
				}
				if rows != 1 {
					retCode = code.TransactionFailed
					return
				}
				allShopIdSkuCodeAmount[amountKey] += req.List[i].Detail[j].Amount
				allShopIdSkuCodeLastTxId[lastTxIdKey] = opTxId
			}
		}
	}
	err = tx.Commit()
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "RestoreInventory Commit err: %v", err)
		retCode = code.TransactionFailed
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
				kelvins.ErrLogger.Errorf(ctx, "GetSkuPriceHistory err: %v, where: %v", err, where)
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

func ConfirmSkuInventory(ctx context.Context, req *sku_business.ConfirmSkuInventoryRequest) (retCode int) {
	retCode = code.Success
	where := map[string]interface{}{
		"out_trade_no": req.OutTradeNo,
		"verify":       0, // 未确认的
	}
	opTxId := uuid.New().String()
	maps := map[string]interface{}{
		"op_tx_id":    opTxId,
		"verify":      1, // 确认的
		"op_ip":       req.OpMeta.OpIp,
		"op_uid":      req.OpMeta.OpUid,
		"update_time": time.Now(),
	}
	_, err := repository.UpdateSkuInventoryRecord(where, maps)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "UpdateSkuInventoryRecord err: %v, where :%v, maps: %v",
			err, json.MarshalToStringNoError(where), json.MarshalToStringNoError(maps))
		retCode = code.ErrorServer
		return
	}
	return
}
