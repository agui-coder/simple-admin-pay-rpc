package model

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/payment/model"
	"time"

	"github.com/agui-coder/simple-admin-pay-rpc/pay"
	"github.com/agui-coder/simple-admin-pay-rpc/utils/money"

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

func (m *OrderModel) QueryByAppIdAndMerchantOrderId(ctx context.Context, merchantOrderId string) (*ent.Order, error) {
	order, err := m.Query().Where(order.MerchantOrderIDEQ(merchantOrderId)).Only(ctx)
	if ent.IsNotFound(err) {
		return nil, nil
	}
	if err != nil {
		return nil, errorhandler.DefaultEntError(logx.WithContext(ctx), err, merchantOrderId)
	}
	return order, nil
}

func (m *OrderModel) QueryPage(ctx context.Context, in *pay.OrderPageReq) (*ent.OrderPageList, error) {
	query := m.Query().Where()
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
	if uint8(pay.PayStatus_PAY_SUCCESS) == order.Status { // 校验状态，发现已支付
		return nil, errorx.NewInvalidArgumentError("pay order status is success")
	}
	if uint8(pay.PayStatus_PAY_WAITING) != order.Status { // 校验状态，必须是待支付
		return nil, errorx.NewInvalidArgumentError("pay order status is not waiting")
	}
	if time.Now().After(order.ExpireTime) {
		return nil, errorx.NewInvalidArgumentError("pay order is expired")
	}
	return order, nil
}

func (m *OrderModel) UpdateOrderSuccess(ctx context.Context, channelCode string, orderExtension *ent.OrderExtension, notifyResp *model.OrderResp) error {
	orderEnt, err := m.Get(ctx, orderExtension.OrderID)
	if err != nil {
		return errorhandler.DefaultEntError(logx.WithContext(ctx), err, notifyResp)
	}
	if orderEnt.Status == uint8(pay.PayStatus_PAY_SUCCESS) && orderEnt.ExtensionID == orderExtension.ID {
		logx.Infof("[updateOrderExtensionSuccess][order:%d 已经是已支付，无需更新]", orderEnt.ID)
		return nil
	}
	if orderEnt.Status != uint8(pay.PayStatus_PAY_WAITING) {
		return errorx.NewInvalidArgumentError("pay order status is not waiting")
	}
	// TODO FeeRate 怎么定
	channelFeePrice, err := money.CalculateRatePriceInternal(orderEnt.Price, 0.0)
	if err != nil {
		return err
	}
	//
	err = m.Update().Where(order.IDEQ(orderEnt.ID), order.StatusEQ(uint8(pay.PayStatus_PAY_WAITING))).
		SetStatus(uint8(pay.PayStatus_PAY_SUCCESS)).SetChannelCode(channelCode).
		SetSuccessTime(notifyResp.SuccessTime).SetExtensionID(orderExtension.ID).SetNo(orderExtension.No).
		SetChannelOrderNo(notifyResp.ChannelOrderNo).SetNotNilChannelUserID(notifyResp.ChannelUserId).
		SetChannelFeeRate(0.0).SetChannelFeePrice(channelFeePrice).Exec(ctx)
	if err != nil {
		return errorhandler.DefaultEntError(logx.WithContext(ctx), err, notifyResp)
	}
	logx.Infof("[updateOrderExtensionSuccess][order %v 更新为已支付]", orderEnt)
	return nil
}
