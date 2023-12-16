package refund

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"

	"github.com/agui-coder/simple-admin-pay-rpc/ent/predicate"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/refund"
	"github.com/agui-coder/simple-admin-pay-rpc/internal/svc"
	"github.com/agui-coder/simple-admin-pay-rpc/utils/errorhandler"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetRefundListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRefundListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRefundListLogic {
	return &GetRefundListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRefundListLogic) GetRefundList(in *pay.RefundListReq) (*pay.RefundListResp, error) {
	var predicates []predicate.Refund
	if in.No != nil {
		predicates = append(predicates, refund.NoContains(*in.No))
	}
	if in.ChannelCode != nil {
		predicates = append(predicates, refund.ChannelCodeContains(*in.ChannelCode))
	}
	if in.OrderNo != nil {
		predicates = append(predicates, refund.OrderNoContains(*in.OrderNo))
	}
	result, err := l.svcCtx.DB.Refund.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &pay.RefundListResp{}
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
