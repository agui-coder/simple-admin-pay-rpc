package order

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/ent"
	"github.com/agui-coder/simple-admin-pay-rpc/model"
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
	if channel.Status == uint8(pay.CommonStatus_Disable) {
		logx.Error("channel is disable")
	}

	if *pointy.GetStatusPointer(&in.Status) == uint8(pay.PayStatus_PAY_SUCCESS) {
		err := l.notifyOrderSuccess(channel, in)
		return &pay.BaseResp{Msg: i18n.Failed}, err
	}
	if *pointy.GetStatusPointer(&in.Status) == uint8(pay.PayStatus_PAY_CLOSED) {
		// TODO 失败处理
		return &pay.BaseResp{Msg: i18n.Failed}, err
	}

	if err != nil {
		logx.Error(err)
	}
	return &pay.BaseResp{Msg: i18n.CreateSuccess}, nil
}

// notifyOrderSuccess 支付成功
func (l *NotifyOrderLogic) notifyOrderSuccess(channel *ent.Channel, notifyResp *pay.NotifyOrderReq) error {
	var task *ent.NotifyTask
	err := entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		newModel := model.NewModel(tx.Client())
		orderExtension, err := newModel.OrderExtension.UpdateOrderSuccess(l.ctx, notifyResp)
		if err != nil {
			return err
		}
		err = newModel.Order.UpdateOrderSuccess(l.ctx, channel, orderExtension, notifyResp)
		if err != nil {
			return err
		}
		task, err = newModel.CreatePayNotifyTask(l.ctx, int(pay.PayType_PAY_ORDER), orderExtension.OrderID)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	err = notify.NewExecuteNotifyLogic(context.Background(), l.svcCtx).ExecuteNotify(task)
	if err != nil {
		return err
	}
	return nil
}
