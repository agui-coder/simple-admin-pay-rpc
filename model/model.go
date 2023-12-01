package model

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/consts"
	"github.com/agui-coder/simple-admin-pay-rpc/ent"
	"time"

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
}

func NewModel(client *ent.Client) *Model {
	return &Model{
		App:            NewAppModel(client),
		Channel:        NewChannelModel(client),
		NotifyLog:      NewNotifyLogModel(client),
		NotifyTask:     NewNotifyTaskModel(client),
		OrderExtension: NewOrderExtensionModel(client),
		Order:          NewOrderModel(client),
	}
}

// CreatePayNotifyTask 同时用到了 order 和 notify_task 两个表
func (m *Model) CreatePayNotifyTask(ctx context.Context, types int, dataId uint64) (*ent.NotifyTask, error) {
	// 创建通知任务
	taskCreate := m.NotifyTask.Create().
		SetType(types).
		SetDataID(dataId).
		SetStatus(consts.WAITING).
		SetNextNotifyTime(time.Now()).
		SetNotifyTimes(0).
		SetLastExecuteTime(time.Now()).
		SetMaxNotifyTimes(int8(len(payload.NotifyFrequency)) + 1)
	if types == consts.OrderType {
		order, err := m.Order.Get(ctx, dataId)
		if err != nil {
			return nil, errorhandler.DefaultEntError(logx.WithContext(ctx), err, dataId)
		}
		taskCreate.SetAppID(order.AppID).
			SetMerchantOrderID(order.MerchantOrderID).
			SetNotifyURL(order.NotifyURL)
	}
	task, err := taskCreate.Save(ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(logx.WithContext(ctx), err, dataId)
	}
	return task, nil
}
