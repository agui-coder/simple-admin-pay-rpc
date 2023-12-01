package refund

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"

	"github.com/agui-coder/simple-admin-pay-rpc/internal/svc"
	"github.com/agui-coder/simple-admin-pay-rpc/utils/errorhandler"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRefundLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRefundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRefundLogic {
	return &UpdateRefundLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateRefundLogic) UpdateRefund(in *pay.RefundInfo) (*pay.BaseResp, error) {
	query := l.svcCtx.DB.Refund.UpdateOneID(*in.Id).
		SetNotNilNo(in.No).
		SetNotNilAppID(in.AppId).
		SetNotNilChannelID(in.ChannelId).
		SetNotNilChannelCode(in.ChannelCode).
		SetNotNilOrderID(in.OrderId).
		SetNotNilOrderNo(in.OrderNo).
		SetNotNilMerchantOrderID(in.MerchantOrderId).
		SetNotNilMerchantRefundID(in.MerchantRefundId).
		SetNotNilNotifyURL(in.NotifyUrl).
		SetNotNilReason(in.Reason).
		SetNotNilUserIP(in.UserIp).
		SetNotNilChannelOrderNo(in.ChannelOrderNo).
		SetNotNilChannelRefundNo(in.ChannelRefundNo).
		SetNotNilSuccessTime(pointy.GetTimePointer(in.SuccessTime, 0)).
		SetNotNilChannelErrorCode(in.ChannelErrorCode).
		SetNotNilChannelErrorMsg(in.ChannelErrorMsg).
		SetNotNilChannelNotifyData(in.ChannelNotifyData)

	if in.Status != nil {
		query.SetNotNilStatus(pointy.GetPointer(uint8(*in.Status)))
	}
	if in.PayPrice != nil {
		query.SetNotNilPayPrice(in.PayPrice)
	}
	if in.RefundPrice != nil {
		query.SetNotNilRefundPrice(in.RefundPrice)
	}

	err := query.Exec(l.ctx)

	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &pay.BaseResp{Msg: errormsg.UpdateSuccess}, nil
}
