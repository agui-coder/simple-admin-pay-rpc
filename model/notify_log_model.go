package model

import (
	"context"
	"encoding/json"
	"github.com/agui-coder/simple-admin-pay-rpc/ent"

	"github.com/agui-coder/simple-admin-pay-rpc/internal/mqs/amq/types/payload"
	"github.com/agui-coder/simple-admin-pay-rpc/utils/errorhandler"

	"github.com/zeromicro/go-zero/core/logx"
)

type NotifyLogModel struct {
	*ent.NotifyLogClient
}

func NewNotifyLogModel(client *ent.Client) *NotifyLogModel {
	return &NotifyLogModel{client.NotifyLog}
}

// CreateNotifyLog 创建通知日志
func (l *NotifyLogModel) CreateNotifyLog(ctx context.Context, task *ent.NotifyTask, newStatus uint8, err error, resp *payload.PayOrderNotifyResp) error {
	create := l.Create().SetTaskID(task.ID).SetStatus(newStatus).
		SetNotifyTimes(task.NotifyTimes + 1)
	if err != nil {
		create.SetResponse(err.Error())
	}
	respJson, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	create.SetResponse(string(respJson))
	_, err = create.Save(ctx)
	if err != nil {
		return errorhandler.DefaultEntError(logx.WithContext(ctx), err, task)
	}
	return nil
}
