package order

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/consts"
	"github.com/agui-coder/simple-admin-pay-rpc/ent"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"

	"github.com/agui-coder/simple-admin-pay-rpc/internal/logic/notify"
	"github.com/agui-coder/simple-admin-pay-rpc/internal/svc"
	"github.com/agui-coder/simple-admin-pay-rpc/utils/entx"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/zeromicro/go-zero/core/logx"
)

type NotifyOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewNotifyOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NotifyOrderLogic {
	return &NotifyOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *NotifyOrderLogic) NotifyOrder(in *pay.NotifyOrderReq) (*pay.BaseResp, error) {
	channel, err := l.svcCtx.DB.Channel.Get(l.ctx, in.ChannelId)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	if channel.Status == consts.Disable {
		logx.Error("channel is disable")
	}
	err = entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {

		if *pointy.GetStatusPointer(&in.Status) == consts.SUCCESS {
			err := notifyOrderSuccess(l.ctx, l.svcCtx, channel, in)
			return err
		}
		if *pointy.GetStatusPointer(&in.Status) == consts.CLOSED {
			// TODO 失败处理
			return nil
		}
		return nil
	})
	if err != nil {
		logx.Error(err)
	}
	return &pay.BaseResp{Msg: i18n.CreateSuccess}, nil
}

// notifyOrderSuccess 支付成功
func notifyOrderSuccess(ctx context.Context, svcCtx *svc.ServiceContext, channel *ent.Channel, notifyResp *pay.NotifyOrderReq) error {
	orderExtension, err := svcCtx.Model.OrderExtension.UpdateOrderSuccess(ctx, notifyResp)
	if err != nil {
		return err
	}
	err = svcCtx.Model.Order.UpdateOrderSuccess(ctx, channel, orderExtension, notifyResp)
	if err != nil {
		return err
	}
	task, err := svcCtx.Model.CreatePayNotifyTask(ctx, consts.OrderType, orderExtension.OrderID)
	if err != nil {
		return err
	}
	go func(task *ent.NotifyTask) {
		err := notify.NewExecuteNotifyLogic(context.Background(), svcCtx).ExecuteNotify(task)
		if err != nil {
			logx.Error(err)
		}
	}(task)
	return nil
}
