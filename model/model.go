package model

import (
	"context"
	"time"

	"github.com/agui-coder/simple-admin-pay-rpc/ent"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/order"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/refund"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/agui-coder/simple-admin-pay-rpc/internal/mqs/amq/types/payload"
	"github.com/agui-coder/simple-admin-pay-rpc/utils/errorhandler"

	"github.com/zeromicro/go-zero/core/logx"
)

type Model struct {
	App            *AppModel
	Channel        *ChannelModel
	NotifyLog      *NotifyLogModel
	NotifyTask     *NotifyTaskModel
	OrderExtension *OrderExtensionModel
	Order          *OrderModel
	Refund         *RefundModel
}

func NewModel(client *ent.Client) *Model {
	return &Model{
		App:            NewAppModel(client.App),
		Channel:        NewChannelModel(client.Channel),
		NotifyLog:      NewNotifyLogModel(client.NotifyLog),
		NotifyTask:     NewNotifyTaskModel(client.NotifyTask),
		OrderExtension: NewOrderExtensionModel(client.OrderExtension),
		Order:          NewOrderModel(client.Order),
		Refund:         NewRefundModel(client.Refund),
	}
}

// CreatePayNotifyTask 同时用到了 order 和 notify_task 两个表
func (m *Model) CreatePayNotifyTask(ctx context.Context, types int, dataId uint64) (*ent.NotifyTask, error) {
	// 创建通知任务
	taskCreate := m.NotifyTask.Create().
		SetType(types).
		SetDataID(dataId).
		SetStatus(uint8(pay.PayStatus_PAY_WAITING)).
		SetNextNotifyTime(time.Now()).
		SetNotifyTimes(0).
		SetLastExecuteTime(time.Now()).
		SetMaxNotifyTimes(int8(len(payload.NotifyFrequency)) + 1)
	if types == int(pay.PayType_PAY_ORDER) {
		order, err := m.Order.Get(ctx, dataId)
		if err != nil {
			return nil, errorhandler.DefaultEntError(logx.WithContext(ctx), err, dataId)
		}
		taskCreate.SetAppID(order.AppID).
			SetMerchantOrderID(order.MerchantOrderID).
			SetNotifyURL(order.NotifyURL)
	} else if types == int(pay.PayType_PAY_RETURN) {
		refundInfo, err := m.Refund.Get(ctx, dataId)
		if err != nil {
			return nil, errorhandler.DefaultEntError(logx.WithContext(ctx), err, dataId)
		}
		taskCreate.SetAppID(refundInfo.AppID).
			SetMerchantOrderID(refundInfo.MerchantOrderID).
			SetNotifyURL(refundInfo.NotifyURL)
	}
	task, err := taskCreate.Save(ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(logx.WithContext(ctx), err, dataId)
	}
	return task, nil
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
