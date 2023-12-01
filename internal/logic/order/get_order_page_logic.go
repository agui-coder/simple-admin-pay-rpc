package order

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"

	"github.com/agui-coder/simple-admin-pay-rpc/utils/errorhandler"

	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/agui-coder/simple-admin-pay-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderPageLogic {
	return &GetOrderPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderPageLogic) GetOrderPage(in *pay.OrderPageReq) (*pay.OrderListResp, error) {
	page, err := l.svcCtx.Model.Order.QueryPage(l.ctx, in)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	orders := page.List
	orderInfos := make([]*pay.OrderInfo, len(orders))
	for i, order := range orders {
		orderInfos[i] = &pay.OrderInfo{
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
		}
	}
	return &pay.OrderListResp{Data: orderInfos, Total: page.PageDetails.Total}, nil
}
