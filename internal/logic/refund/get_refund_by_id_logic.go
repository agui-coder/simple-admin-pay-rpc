package refund

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"

	"github.com/agui-coder/simple-admin-pay-rpc/internal/svc"
	"github.com/agui-coder/simple-admin-pay-rpc/utils/errorhandler"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetRefundByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRefundByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRefundByIdLogic {
	return &GetRefundByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRefundByIdLogic) GetRefundById(in *pay.IDReq) (*pay.RefundInfo, error) {
	result, err := l.svcCtx.DB.Refund.Get(l.ctx, in.Id)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &pay.RefundInfo{
		Id:                &result.ID,
		CreatedAt:         pointy.GetPointer(result.CreatedAt.Unix()),
		UpdatedAt:         pointy.GetPointer(result.UpdatedAt.Unix()),
		Status:            pointy.GetPointer(uint32(result.Status)),
		No:                &result.No,
		AppId:             &result.AppID,
		ChannelId:         &result.ChannelID,
		ChannelCode:       &result.ChannelCode,
		OrderId:           &result.OrderID,
		OrderNo:           &result.OrderNo,
		MerchantOrderId:   &result.MerchantOrderID,
		MerchantRefundId:  &result.MerchantRefundID,
		NotifyUrl:         &result.NotifyURL,
		PayPrice:          &result.PayPrice,
		RefundPrice:       &result.RefundPrice,
		Reason:            &result.Reason,
		UserIp:            &result.UserIP,
		ChannelOrderNo:    &result.ChannelOrderNo,
		ChannelRefundNo:   &result.ChannelRefundNo,
		SuccessTime:       pointy.GetPointer(result.SuccessTime.UnixMilli()),
		ChannelErrorCode:  &result.ChannelErrorCode,
		ChannelErrorMsg:   &result.ChannelErrorMsg,
		ChannelNotifyData: &result.ChannelNotifyData,
	}, nil
}
