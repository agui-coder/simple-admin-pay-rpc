package payjob

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/ent"
	"sync"
	"sync/atomic"
	"time"

	"github.com/agui-coder/simple-admin-pay-rpc/internal/logic/notify"
	"github.com/agui-coder/simple-admin-pay-rpc/internal/mqs/amq/types/payload"
	"github.com/agui-coder/simple-admin-pay-rpc/internal/svc"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
)

type PayNotifyHandler struct {
	svcCtx *svc.ServiceContext
}

func NewPayNotifyWorldHandler(svcCtx *svc.ServiceContext) *PayNotifyHandler {
	return &PayNotifyHandler{
		svcCtx: svcCtx,
	}
}

// ProcessTask if return err != nil , asynq will retry | 如果返回错误不为空则会重试
func (l *PayNotifyHandler) ProcessTask(ctx context.Context, t *asynq.Task) error {
	notifyTasks, err := l.svcCtx.Model.NotifyTask.QueryListByNotify(ctx)
	if err != nil {
		logx.Errorf("[executeNotify][查询任务失败，原因是%s]", err)
		return nil
	}
	if len(notifyTasks) <= 0 {
		return nil
	}
	// 任务数量
	var size = len(notifyTasks)
	// 定义工作线程的数量
	const numWorkers = 10
	tasks := make(chan *ent.NotifyTask, len(notifyTasks))
	// 记录剩余任务数量
	remainingTasks := int32(size)
	logx.Infof("执行支付通知 %d 个", size)
	var wg sync.WaitGroup
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range tasks {
				err := notify.NewExecuteNotifyLogic(ctx, l.svcCtx).ExecuteNotify(task)
				if err != nil {
					logx.Errorf("[executeNotify][task%s 任务处理失败，原因是%s]", task, err)
				}
				// 减少剩余任务数量
				atomic.AddInt32(&remainingTasks, -1)
			}
		}()
	}
	// 将任务分配给工作线程，并在完成后关闭通道
	for _, task := range notifyTasks {
		tasks <- task
	}
	close(tasks)
	// 启动一个 goroutine 每秒钟打印剩余任务数量
	done := make(chan struct{})
	go func() {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				logx.Infof("[executeNotify][任务处理中， 总任务数%d 剩余任务数%d]", size, atomic.LoadInt32(&remainingTasks))
			case <-ctx.Done():
				logx.Errorf("[executeNotify][任务处理失败，原因是%s]", ctx.Err())
				return
			case <-time.After(time.Duration(payload.NotifyTimeoutMillis) * time.Second):
				logx.Infof("[executeNotify][任务未处理完， 总任务数%d 剩余任务数%d]", size, atomic.LoadInt32(&remainingTasks))
				return
			case <-done:
				logx.Infof("[executeNotify][任务完成， 总任务数%d 剩余任务数%d]", size, atomic.LoadInt32(&remainingTasks))
				return
			}
		}
	}()
	wg.Wait()
	close(done)
	return nil
}
