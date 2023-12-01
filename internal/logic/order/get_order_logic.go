package order

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"

	"github.com/agui-coder/simple-admin-pay-rpc/utils/errorhandler"

	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/agui-coder/simple-admin-pay-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderLogic {
	return &GetOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderLogic) GetOrder(in *pay.IDReq) (*pay.OrderInfo, error) {
	order, err := l.svcCtx.DB.Order.Get(l.ctx, in.Id)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	return &pay.OrderInfo{
		Id:              &order.ID,
		CreatedAt:       pointy.GetPointer(order.CreatedAt.UnixMilli()),
		UpdatedAt:       pointy.GetPointer(order.UpdatedAt.UnixMilli()),
		Status:          pointy.GetPointer(uint32(order.Status)),
		AppId:           &order.AppID,
		ChannelId:       &order.ChannelID,
		ChannelCode:     &order.ChannelCode,
		MerchantOrderId: &order.MerchantOrderID,
		Subject:         &order.Subject,
		Body:            &order.Body,
		NotifyUrl:       &order.NotifyURL,
		Price:           &order.Price,
		ChannelFeeRate:  &order.ChannelFeeRate,
		ChannelFeePrice: &order.ChannelFeePrice,
		UserIp:          &order.UserIP,
		ExpireTime:      pointy.GetPointer(order.ExpireTime.UnixMilli()),
		SuccessTime:     pointy.GetPointer(order.SuccessTime.UnixMilli()),
		NotifyTime:      pointy.GetPointer(order.NotifyTime.UnixMilli()),
		ExtensionId:     &order.ExtensionID,
		No:              &order.No,
		RefundPrice:     &order.RefundPrice,
		ChannelUserId:   &order.ChannelUserID,
		ChannelOrderNo:  &order.ChannelOrderNo,
	}, nil
}
