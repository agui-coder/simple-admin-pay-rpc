package refund

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/predicate"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/refund"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"
	"github.com/agui-coder/simple-admin-pay-rpc/utils/errorhandler"
	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/agui-coder/simple-admin-pay-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRefundPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRefundPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRefundPageLogic {
	return &GetRefundPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRefundPageLogic) GetRefundPage(in *pay.RefundPageReq) (*pay.RefundPageResp, error) {
	var predicates []predicate.Refund
	if in.ChannelCode != nil {
		predicates = append(predicates, refund.ChannelCodeEQ(*in.ChannelCode))
	}
	if in.MerchantOrderId != nil {
		predicates = append(predicates, refund.MerchantOrderIDContains(*in.MerchantOrderId))
	}
	if in.MerchantRefundId != nil {
		predicates = append(predicates, refund.MerchantRefundIDContains(*in.MerchantRefundId))

	}
	if in.ChannelOrderNo != nil {
		predicates = append(predicates, refund.ChannelOrderNoContains(*in.ChannelOrderNo))
	}
	if in.ChannelRefundNo != nil {
		predicates = append(predicates, refund.ChannelRefundNoContains(*in.ChannelRefundNo))
	}
	if in.Status != nil {
		predicates = append(predicates, refund.StatusEQ(*pointy.GetStatusPointer(in.Status)))
	}
	if in.CreateTime != nil {
		predicates = append(predicates, refund.CreatedAtGT(*pointy.GetTimePointer(&in.CreateTime[0], 0)))
		predicates = append(predicates, refund.CreatedAtLT(*pointy.GetTimePointer(&in.CreateTime[1], 0)))
	}
	result, err := l.svcCtx.DB.Refund.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &pay.RefundPageResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &pay.RefundInfo{
			Id:                &v.ID,
			CreatedAt:         pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt:         pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			Status:            pointy.GetPointer(uint32(v.Status)),
			No:                &v.No,
			ChannelCode:       &v.ChannelCode,
			OrderId:           &v.OrderID,
			OrderNo:           &v.OrderNo,
			MerchantOrderId:   &v.MerchantOrderID,
			MerchantRefundId:  &v.MerchantRefundID,
			PayPrice:          &v.PayPrice,
			RefundPrice:       &v.RefundPrice,
			Reason:            &v.Reason,
			UserIp:            &v.UserIP,
			ChannelOrderNo:    &v.ChannelOrderNo,
			ChannelRefundNo:   &v.ChannelRefundNo,
			SuccessTime:       pointy.GetPointer(v.SuccessTime.UnixMilli()),
			ChannelErrorCode:  &v.ChannelErrorCode,
			ChannelErrorMsg:   &v.ChannelErrorMsg,
			ChannelNotifyData: &v.ChannelNotifyData,
		})
	}

	return resp, nil
}
