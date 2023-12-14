package model

import (
	"context"
	"time"

	"github.com/agui-coder/simple-admin-pay-rpc/ent"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"
	"github.com/pkg/errors"

	"github.com/agui-coder/simple-admin-pay-rpc/ent/notifytask"
	"github.com/agui-coder/simple-admin-pay-rpc/internal/mqs/amq/types/payload"
	"github.com/agui-coder/simple-admin-pay-rpc/utils/errorhandler"

	"github.com/zeromicro/go-zero/core/logx"
)

type NotifyTaskModel struct {
	*ent.NotifyTaskClient
}

func NewNotifyTaskModel(client *ent.NotifyTaskClient) *NotifyTaskModel {
	return &NotifyTaskModel{client}
}

func (m *NotifyTaskModel) QueryListByNotify(ctx context.Context) ([]*ent.NotifyTask, error) {
	notifyTasks, err := m.Query().Where(notifytask.StatusIn(uint8(pay.PayStatus_PAY_WAITING),
		uint8(pay.PayStatus_PAY_SUCCESS),
		uint8(pay.PayStatus_PAY_FAILURE)),
		notifytask.NextNotifyTimeLTE(time.Now())).All(ctx)
	if err != nil {
		return nil, err
	}
	return notifyTasks, nil
}

func (m *NotifyTaskModel) ProcessNotifyResult(ctx context.Context, task *ent.NotifyTask, resp payload.PayOrderNotifyResp, err error) (uint8, error) {
	// 处理并更新通知结果
	updateTask := m.UpdateOne(task).SetLastExecuteTime(time.Now()).SetNotifyTimes(task.NotifyTimes + 1)
	switch {
	case resp.Code == int(payload.SUCCESS):
		updateTask.SetStatus(uint8(pay.PayStatus_PAY_SUCCESS))
	case task.NotifyTimes >= int8(len(payload.NotifyFrequency)-1):
		updateTask.SetStatus(uint8(pay.PayStatus_PAY_FAILURE))
	default:
		updateTask.SetNextNotifyTime(time.Now().Add(time.Duration(payload.NotifyFrequency[task.NotifyTimes]) * time.Second))
		if err != nil {
			updateTask.SetStatus(uint8(pay.PayStatus_PAY_FAILURE))
		} else {
			updateTask.SetStatus(uint8(pay.PayStatus_PAY_SUCCESS))
		}
	}
	task, newErr := updateTask.Save(ctx)
	if newErr != nil {
		return 0, errors.Wrap(err, errorhandler.DefaultEntError(logx.WithContext(ctx), err, task).Error())
	}
	return task.Status, nil
}
