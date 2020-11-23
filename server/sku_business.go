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
	result.List = list
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
			Code: sku_business.RetCode_SUCCESS,
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

func (s *SkuBusinessServer) FiltrateSkuPriceVersion(ctx context.Context, req *sku_business.FiltrateSkuPriceVersionRequest) (*sku_business.FiltrateSkuPriceVersionResponse, error) {
	result := &sku_business.FiltrateSkuPriceVersionResponse{Common: &sku_business.CommonResponse{
		Code: sku_business.RetCode_SUCCESS,
	}}
	list, retCode := service.FiltrateSkuPriceVersion(ctx, req)
	if retCode != code.Success {
		switch retCode {
		case code.SkuCodeNotExist:
			result.Common.Code = sku_business.RetCode_SKU_NOT_EXIST
		case code.SkuPriceVersionNotExist:
			result.Common.Code = sku_business.RetCode_SKU_PRICE_VERSION_NOT_EXIST
		case code.SkuPriceVersionPolicyNotSupport:
			result.Common.Code = sku_business.RetCode_SKU_PRICE_VERSION_POLICY_TYPE_NOT_SUPPORT
		case code.SkuPriceVersionPolicyDataFormatErr:
			result.Common.Code = sku_business.RetCode_SKU_PRICE_VERSION_POLICY_DATA_FORMAT_ERR
		default:
			result.Common.Code = sku_business.RetCode_ERROR
		}
		return result, nil
	}
	result.Result = list
	return result, nil
}

func (s *SkuBusinessServer) SearchSyncSkuInventory(ctx context.Context, req *sku_business.SearchSyncSkuInventoryRequest) (*sku_business.SearchSyncSkuInventoryResponse, error) {
	result := &sku_business.SearchSyncSkuInventoryResponse{
		Common: &sku_business.CommonResponse{
			Code: sku_business.RetCode_SUCCESS,
		},
		Info: make([]*sku_business.SkuInventoryInfo, 0),
	}
	list, retCode := service.SyncSkuInventory(ctx, req)
	if retCode != code.Success {
		result.Common.Code = sku_business.RetCode_ERROR
		return result, nil
	}
	result.Info = list
	return result, nil
}

func (s *SkuBusinessServer) SearchSkuInventory(ctx context.Context, req *sku_business.SearchSkuInventoryRequest) (*sku_business.SearchSkuInventoryResponse, error) {
	result := &sku_business.SearchSkuInventoryResponse{
		Common: &sku_business.CommonResponse{
			Code: sku_business.RetCode_SUCCESS,
		},
		List: nil,
	}
	list, retCode := service.SearchSkuInventory(ctx, req)
	if retCode != code.Success {
		result.Common.Code = sku_business.RetCode_ERROR
		return result, nil
	}
	result.List = list
	return result, nil
}
