package model

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-common/consts"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"
	"github.com/agui-coder/simple-admin-pay-rpc/utils/money"
	"time"

	"github.com/agui-coder/simple-admin-pay-rpc/ent"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/order"
	"github.com/agui-coder/simple-admin-pay-rpc/utils/errorhandler"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type OrderModel struct {
	*ent.OrderClient
}

func NewOrderModel(client *ent.OrderClient) *OrderModel {
	return &OrderModel{client}
}

func (m *OrderModel) QueryByAppIdAndMerchantOrderId(ctx context.Context, appId uint64, merchantOrderId string) (*ent.Order, error) {
	order, err := m.Query().Where(order.AppIDEQ(appId), order.MerchantOrderIDEQ(merchantOrderId)).Only(ctx)
	if ent.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, errorhandler.DefaultEntError(logx.WithContext(ctx), err, appId)
	}
	return order, nil
}

func (m *OrderModel) QueryPage(ctx context.Context, in *pay.OrderPageReq) (*ent.OrderPageList, error) {
	query := m.Query().Where()
	if in.AppId != nil {
		query.Where(order.AppIDEQ(*in.AppId))
	}
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
	page, err := query.Page(ctx, in.Page, in.PageSize)
	if err != nil {
		return nil, errorhandler.DefaultEntError(logx.WithContext(ctx), err, in)
	}
	return page, nil
}

func (m *OrderModel) ValidateOrderCanSubmit(ctx context.Context, id uint64) (*ent.Order, error) {
	order, err := m.Get(ctx, id)
	if err != nil {
		return nil, errorhandler.DefaultEntError(logx.WithContext(ctx), err, id)
	}
	if consts.SUCCESS == order.Status { // 校验状态，发现已支付
		return nil, errorx.NewInvalidArgumentError("pay order status is success")
	}
	if consts.WAITING != order.Status { // 校验状态，必须是待支付
		return nil, errorx.NewInvalidArgumentError("pay order status is not waiting")
	}
	if time.Now().After(order.ExpireTime) {
		return nil, errorx.NewInvalidArgumentError("pay order is expired")
	}
	return order, nil
}

func (m *OrderModel) UpdateOrderSuccess(ctx context.Context, channel *ent.Channel, orderExtension *ent.OrderExtension, notifyResp *pay.NotifyOrderReq) error {
	orderEnt, err := m.Get(ctx, orderExtension.OrderID)
	if err != nil {
		return errorhandler.DefaultEntError(logx.WithContext(ctx), err, notifyResp)
	}
	if orderEnt.Status == consts.SUCCESS && orderEnt.ExtensionID == orderExtension.ID {
		logx.Infof("[updateOrderExtensionSuccess][order:%d 已经是已支付，无需更新]", orderEnt.ID)
		return nil
	}
	if orderEnt.Status != consts.WAITING {
		return errorx.NewInvalidArgumentError("pay order status is not waiting")
	}
	channelFeePrice, err := money.CalculateRatePriceInternal(orderEnt.Price, channel.FeeRate)
	if err != nil {
		return err
	}
	updateCounts, err := m.Update().Where(order.IDEQ(orderEnt.ID), order.StatusEQ(consts.WAITING)).
		SetStatus(consts.SUCCESS).SetChannelID(channel.ID).SetChannelCode(channel.Code).
		SetNotNilSuccessTime(pointy.GetTimePointer(&notifyResp.SuccessTime, 0)).SetExtensionID(orderExtension.ID).SetNo(orderExtension.No).
		SetChannelOrderNo(notifyResp.ChannelOrderNo).SetNotNilChannelUserID(notifyResp.ChannelUserId).
		SetChannelFeeRate(channel.FeeRate).SetChannelFeePrice(channelFeePrice).Save(ctx)
	if err != nil {
		return errorhandler.DefaultEntError(logx.WithContext(ctx), err, notifyResp)
	}
	if updateCounts == 0 {
		return errorx.NewInvalidArgumentError("pay order status is not waiting")
	}
	logx.Infof("[updateOrderExtensionSuccess][order %v 更新为已支付]", orderEnt)
	return nil
}
