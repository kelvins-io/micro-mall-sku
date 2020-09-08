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
		Code: 0,
		Msg:  "",
	}
	if req.Sku.SkuCode == "" {
		result.Common.Code = sku_business.RetCode_SKU_NOT_EXIST
		result.Common.Msg = errcode.GetErrMsg(code.SkuCodeNotExist)
		return &result, nil
	}
	retCode := service.PutAwaySku(ctx, req)
	if retCode != code.Success {
		result.Common.Code = sku_business.RetCode_ERROR
		result.Common.Msg = errcode.GetErrMsg(code.ErrorServer)
		return &result, nil
	}
	result.Common.Code = sku_business.RetCode_SUCCESS
	result.Common.Msg = errcode.GetErrMsg(code.Success)
	return &result, nil

}

func (s *SkuBusinessServer) GetSkuList(ctx context.Context, req *sku_business.GetSkuListRequest) (*sku_business.GetSkuListResponse, error) {
	return &sku_business.GetSkuListResponse{}, nil
}
