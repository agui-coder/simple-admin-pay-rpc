package model

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/ent"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/order"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/refund"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"
	"github.com/agui-coder/simple-admin-pay-rpc/utils/errorhandler"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type Model struct {
	App            *AppModel
	Channel        *ChannelModel
	OrderExtension *OrderExtensionModel
	Order          *OrderModel
	Refund         *RefundModel
}

func NewModel(client *ent.Client) *Model {
	return &Model{
		App:            NewAppModel(client.App),
		Channel:        NewChannelModel(client.Channel),
		OrderExtension: NewOrderExtensionModel(client.OrderExtension),
		Order:          NewOrderModel(client.Order),
		Refund:         NewRefundModel(client.Refund),
	}
}

func (m *Model) ValidatePayOrderCanRefund(ctx context.Context, in *pay.RefundCreateReq) (*ent.Order, error) {
	order, err := m.Order.Query().Where(order.AppIDEQ(in.AppId), order.MerchantOrderIDEQ(in.MerchantOrderId)).Only(ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(logx.WithContext(ctx), err, in)
	}
	if order.Status != uint8(pay.PayStatus_PAY_SUCCESS) && order.Status != uint8(pay.PayStatus_PAY_REFUND) {
		return nil, errorx.NewInvalidArgumentError("pay order refund fail status error")
	}
	if in.Price+order.RefundPrice > order.Price {
		return nil, errorx.NewInvalidArgumentError("refund price exceed")
	}
	count, err := m.Refund.Query().Where(refund.AppIDEQ(in.AppId), refund.OrderIDEQ(order.ID),
		refund.StatusEQ(uint8(pay.PayStatus_PAY_WAITING))).Count(ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(logx.WithContext(ctx), err, in)
	}
	if count > 0 {
		return nil, errorx.NewInvalidArgumentError("refund has refunding")
	}
	return order, nil
}
