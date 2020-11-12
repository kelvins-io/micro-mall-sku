package server

import (
	"context"
	"gitee.com/cristiane/micro-mall-sku/pkg/code"
	"gitee.com/cristiane/micro-mall-sku/proto/micro_mall_sku_proto/sku_business"
	"gitee.com/cristiane/micro-mall-sku/service"
	"gitee.com/kelvins-io/common/errcode"
)

type SkuBusinessServer struct{}

func NewSkuBusinessServer() sku_business.SkuBusinessServiceServer {
	return new(SkuBusinessServer)
}

func (s *SkuBusinessServer) PutAwaySku(ctx context.Context, req *sku_business.PutAwaySkuRequest) (*sku_business.PutAwaySkuResponse, error) {
	var result sku_business.PutAwaySkuResponse
	result.Common = &sku_business.CommonResponse{
		Code: sku_business.RetCode_SUCCESS,
		Msg:  "",
	}
	if req.Sku.SkuCode == "" {
		result.Common.Code = sku_business.RetCode_SKU_NOT_EXIST
		result.Common.Msg = errcode.GetErrMsg(code.SkuCodeNotExist)
		return &result, nil
	}
	retCode := service.PutAwaySku(ctx, req)
	if retCode != code.Success {
		if retCode == code.SkuCodeExist {
			result.Common.Code = sku_business.RetCode_SKU_EXIST
		} else if retCode == code.ShopBusinessNotExist {
			result.Common.Code = sku_business.RetCode_SHOP_NOT_EXIST
		} else if retCode == code.TransactionFailed {
			result.Common.Code = sku_business.RetCode_TRANSACTION_FAILED
		} else {
			result.Common.Code = sku_business.RetCode_ERROR
		}
		result.Common.Msg = errcode.GetErrMsg(retCode)
		return &result, nil
	}
	return &result, nil
}

func (s *SkuBusinessServer) GetSkuList(ctx context.Context, req *sku_business.GetSkuListRequest) (*sku_business.GetSkuListResponse, error) {
	var result sku_business.GetSkuListResponse
	result.List = make([]*sku_business.SkuInventoryInfo, 0)
	list, retCode := service.GetSkuList(ctx, req)
	if retCode != code.Success {
		return &result, errcode.TogRPCError(retCode)
	}
	result.List = make([]*sku_business.SkuInventoryInfo, len(list))
	for i := 0; i < len(list); i++ {
		info := &sku_business.SkuInventoryInfo{
			SkuCode:       list[i].SkuCode,
			Name:          list[i].Name,
			Price:         list[i].Price,
			Title:         list[i].Title,
			SubTitle:      list[i].SubTitle,
			Desc:          list[i].Desc,
			Production:    list[i].Production,
			Supplier:      list[i].Supplier,
			Category:      list[i].Category,
			Color:         list[i].Color,
			ColorCode:     list[i].ColorCode,
			Specification: list[i].Specification,
			DescLink:      list[i].DescLink,
			State:         list[i].State,
			Amount:        list[i].Amount,
			ShopId:        list[i].ShopId,
			Version:       int64(list[i].Version),
		}
		result.List[i] = info
	}

	return &result, nil
}

func (s *SkuBusinessServer) SupplementSkuProperty(ctx context.Context, req *sku_business.SupplementSkuPropertyRequest) (*sku_business.SupplementSkuPropertyResponse, error) {
	var result sku_business.SupplementSkuPropertyResponse
	result.Common = &sku_business.CommonResponse{
		Code: sku_business.RetCode_SUCCESS,
		Msg:  "",
	}
	if req.SkuCode == "" {
		result.Common.Code = sku_business.RetCode_SKU_NOT_EXIST
		result.Common.Msg = errcode.GetErrMsg(code.SkuCodeNotExist)
		return &result, nil
	}
	retCode := service.SupplementSkuProperty(ctx, req)
	if retCode != code.Success {
		if retCode == code.SkuCodeExist {
			result.Common.Code = sku_business.RetCode_SKU_EXIST
		} else if retCode == code.ShopBusinessNotExist {
			result.Common.Code = sku_business.RetCode_SHOP_NOT_EXIST
		} else if retCode == code.TransactionFailed {
			result.Common.Code = sku_business.RetCode_TRANSACTION_FAILED
		} else {
			result.Common.Code = sku_business.RetCode_ERROR
		}
		result.Common.Msg = errcode.GetErrMsg(retCode)
		return &result, nil
	}
	return &result, nil
}

func (s *SkuBusinessServer) DeductInventory(ctx context.Context, req *sku_business.DeductInventoryRequest) (*sku_business.DeductInventoryResponse, error) {
	var result sku_business.DeductInventoryResponse
	result.Common = &sku_business.CommonResponse{
		Code: sku_business.RetCode_SUCCESS,
		Msg:  errcode.GetErrMsg(code.Success),
	}
	_, retCode := service.DeductInventory(ctx, req)
	if retCode != code.Success {
		if retCode == code.SkuAmountNotEnough {
			result.Common.Code = sku_business.RetCode_SKU_AMOUNT_NOT_ENOUGH
		} else if retCode == code.TransactionFailed {
			result.Common.Code = sku_business.RetCode_TRANSACTION_FAILED
		} else {
			result.Common.Code = sku_business.RetCode_ERROR
		}
		result.Common.Msg = errcode.GetErrMsg(retCode)
		return &result, nil
	}
	result.IsSuccess = true
	return &result, nil
}

func (s *SkuBusinessServer) RestoreInventory(ctx context.Context, req *sku_business.RestoreInventoryRequest) (*sku_business.RestoreInventoryResponse, error) {
	var result = sku_business.RestoreInventoryResponse{
		Common: &sku_business.CommonResponse{
			Code: sku_business.RetCode_ERROR,
			Msg:  "",
		},
		IsSuccess: false,
	}
	retCode := service.RestoreInventory(ctx, req)
	if retCode != code.Success {
		if retCode == code.SkuAmountNotEnough {
			result.Common.Code = sku_business.RetCode_SKU_AMOUNT_NOT_ENOUGH
		} else if retCode == code.TransactionFailed {
			result.Common.Code = sku_business.RetCode_TRANSACTION_FAILED
		} else {
			result.Common.Code = sku_business.RetCode_ERROR
		}
		result.Common.Msg = errcode.GetErrMsg(retCode)
		return &result, nil
	}
	result.IsSuccess = true
	return &result, nil
}
