package notify

import (
	"context"
	"encoding/json"
	"github.com/agui-coder/simple-admin-pay-common/consts"
	"github.com/agui-coder/simple-admin-pay-rpc/ent"
	"github.com/agui-coder/simple-admin-pay-rpc/model"
	"net/http"
	"strconv"

	"github.com/agui-coder/simple-admin-pay-rpc/internal/mqs/amq/types/payload"
	"github.com/agui-coder/simple-admin-pay-rpc/internal/svc"
	"github.com/agui-coder/simple-admin-pay-rpc/utils/entx"
	"github.com/agui-coder/simple-admin-pay-rpc/utils/errorhandler"

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
	lock := redis.NewRedisLock(l.svcCtx.Redis, "pay_task_lock:"+strconv.FormatUint(task.ID, 10))
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
		err = l.executeNotify0(task)
		return err
	}
	return err
}

// executeNotifyOrder 执行订单通知
func (l *ExecuteNotifyLogic) executeNotify0(task *ent.NotifyTask) error {
	resp, err := l.executeNotifyInvoke(task)
	err = entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		txModel := model.NewModel(tx.Client())
		status, err := txModel.NotifyTask.ProcessNotifyResult(l.ctx, task, resp, err)
		err = txModel.NotifyLog.CreateNotifyLog(l.ctx, task, status, err, resp)
		if err != nil {
			return err
		}
		return nil
	})
	return err
}

func (l *ExecuteNotifyLogic) executeNotifyInvoke(task *ent.NotifyTask) (resp payload.PayOrderNotifyResp, err error) {
	var request any
	if consts.OrderType == task.Status {
		request = payload.PayOrderNotifyReq{
			MerchantOrderId: task.MerchantOrderID,
			PayOrderId:      task.DataID,
		}
	} else if consts.RefundType == task.Status {
		request = payload.PayRefundNotifyReq{
			MerchantOrderId: task.MerchantOrderID,
			PayRefundId:     task.DataID,
		}
	}
	response, err := httpc.Do(l.ctx, http.MethodPost, task.NotifyURL, request)
	if err == nil {
		return payload.PayOrderNotifyResp{}, err
	}
	err = httpc.Parse(response, &resp)
	return resp, err
}
