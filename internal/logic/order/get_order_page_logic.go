package order

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/ent"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/order"
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
	page, err := l.QueryPage(in)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	orders := page.List
	orderInfos := make([]*pay.OrderInfo, len(orders))
	for i, orderInfo := range orders {
		orderInfos[i] = &pay.OrderInfo{
			Id:              &orderInfo.ID,
			CreatedAt:       pointy.GetPointer(orderInfo.CreatedAt.UnixMilli()),
			UpdatedAt:       pointy.GetPointer(orderInfo.UpdatedAt.UnixMilli()),
			Status:          pointy.GetPointer(uint32(orderInfo.Status)),
			ChannelCode:     &orderInfo.ChannelCode,
			MerchantOrderId: &orderInfo.MerchantOrderID,
			Subject:         &orderInfo.Subject,
			Body:            &orderInfo.Body,
			Price:           &orderInfo.Price,
			ChannelFeeRate:  &orderInfo.ChannelFeeRate,
			ChannelFeePrice: &orderInfo.ChannelFeePrice,
			UserIp:          &orderInfo.UserIP,
			ExpireTime:      pointy.GetPointer(orderInfo.ExpireTime.UnixMilli()),
			SuccessTime:     pointy.GetPointer(orderInfo.SuccessTime.UnixMilli()),
			NotifyTime:      pointy.GetPointer(orderInfo.NotifyTime.UnixMilli()),
			ExtensionId:     &orderInfo.ExtensionID,
			No:              &orderInfo.No,
			RefundPrice:     &orderInfo.RefundPrice,
			ChannelUserId:   &orderInfo.ChannelUserID,
			ChannelOrderNo:  &orderInfo.ChannelOrderNo,
		}
	}
	return &pay.OrderListResp{Data: orderInfos, Total: page.PageDetails.Total}, nil
}

func (l *GetOrderPageLogic) QueryPage(in *pay.OrderPageReq) (*ent.OrderPageList, error) {
	query := l.svcCtx.DB.Order.Query().Where()
	if in.ChannelCode != nil {
		query.Where(order.ChannelCodeEQ(*in.ChannelCode))
	}
	if in.MerchantOrderId != nil {
		query.Where(order.MerchantOrderIDContains(*in.ChannelCode))
	}
	if in.ChannelOrderNo != nil {
		query.Where(order.ChannelOrderNoContains(*in.ChannelOrderNo))
	}
	if in.No != nil {
		query.Where(order.NoContains(*in.No))
	}
	if in.Status != nil {
		query.Where(order.StatusEQ(*pointy.GetStatusPointer(in.Status)))
	}
	startTime := in.CreateTime[0]
	endTime := in.CreateTime[1]
	if startTime > 0 {
		query.Where(order.CreatedAtGTE(*pointy.GetTimePointer(&startTime, 0)))
	}
	if endTime > 0 {
		query.Where(order.CreatedAtLTE(*pointy.GetTimePointer(&endTime, 0)))
	}
	query.Order(ent.Desc(order.FieldID))
	page, err := query.Page(l.ctx, in.Page, in.PageSize)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	return page, nil
}
