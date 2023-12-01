package notify

import (
	"context"
	"encoding/json"
	"github.com/agui-coder/simple-admin-pay-rpc/consts"
	"github.com/agui-coder/simple-admin-pay-rpc/ent"
	"net/http"
	"strconv"

	"github.com/agui-coder/simple-admin-pay-rpc/internal/mqs/amq/types/payload"
	"github.com/agui-coder/simple-admin-pay-rpc/internal/svc"
	"github.com/agui-coder/simple-admin-pay-rpc/utils/entx"
	"github.com/agui-coder/simple-admin-pay-rpc/utils/errorhandler"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest/httpc"
)

type ExecuteNotifyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExecuteNotifyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExecuteNotifyLogic {
	return &ExecuteNotifyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ExecuteNotify 通知锁核心逻辑 redis 实现
func (l *ExecuteNotifyLogic) ExecuteNotify(task *ent.NotifyTask) error {
	lock := redis.NewRedisLock(l.svcCtx.Redis, strconv.FormatUint(task.ID, 10))
	lock.SetExpire(payload.NotifyTimeoutMillis)
	// 尝试获取锁
	acquire, err := lock.Acquire()
	defer func(lock *redis.RedisLock) {
		_, err := lock.Release()
		if err != nil {
			logx.Error(err)
		}
	}(lock) // 释放锁
	if err != nil {
		return err
	}
	if acquire {
		dbTask, err := l.svcCtx.DB.NotifyTask.Get(l.ctx, task.ID)
		if err != nil {
			return errorhandler.DefaultEntError(l.Logger, err, dbTask)
		}
		if task.NotifyTimes != dbTask.NotifyTimes {
			taskJson, err := json.Marshal(task)
			if err != nil {
				return err
			}
			logx.Errorf("[executeNotifySync][task%s 任务被忽略，原因是它的通知不是第 %d 次，可能是因为并发执行了]", string(taskJson), dbTask.NotifyTimes)
			return nil
		}
		if task.Type == consts.OrderType {
			err = l.executeNotifyOrder(dbTask)
			return err
		} else if task.Type == consts.RefundType {
			return nil
		}
	}
	return err
}

// executeNotifyOrder 执行订单通知
func (l *ExecuteNotifyLogic) executeNotifyOrder(task *ent.NotifyTask) error {
	err := entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		var resp *payload.PayOrderNotifyResp
		response, err := httpc.Do(l.ctx, http.MethodPost, task.NotifyURL, payload.PayOrderNotifyReq{
			MerchantOrderId: task.MerchantOrderID,
			PayOrderId:      task.DataID,
		})
		if err == nil {
			newErr := httpc.ParseJsonBody(response, resp)
			if newErr != nil {
				err = errors.Wrap(err, newErr.Error())
			}
		}
		status, newErr := l.svcCtx.Model.NotifyTask.ProcessNotifyResult(l.ctx, task, resp)
		if newErr != nil {
			err = errors.Wrap(err, newErr.Error())
		}
		err = l.svcCtx.Model.NotifyLog.CreateNotifyLog(l.ctx, task, status, err, resp)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}
