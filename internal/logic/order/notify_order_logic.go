package order

import (
	"context"
	"encoding/json"
	"github.com/agui-coder/simple-admin-pay-rpc/ent"
	"github.com/agui-coder/simple-admin-pay-rpc/internal/svc"
	dbModel "github.com/agui-coder/simple-admin-pay-rpc/model"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"
	"github.com/agui-coder/simple-admin-pay-rpc/payment/model"
	"github.com/agui-coder/simple-admin-pay-rpc/utils/entx"
	"github.com/agui-coder/simple-admin-pay-rpc/utils/errorhandler"
	"github.com/hibiken/asynq"

	"github.com/suyuan32/simple-admin-common/i18n"
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
	resp, err := l.svcCtx.PayClient[in.Code].ParseOrderNotify(in.R)
	if err != nil {
		return nil, err
	}
	err = l.NotifyOrder0(in.Code, resp)
	if err != nil {
		return nil, err
	}
	return &pay.BaseResp{Msg: i18n.CreateSuccess}, nil
}

func (l *NotifyOrderLogic) NotifyOrder0(channelCode string, resp *model.OrderResp) error {
	if resp.Status == uint8(pay.PayStatus_PAY_SUCCESS) {
		err := l.notifyOrderSuccess(channelCode, resp)
		return err
	}
	if resp.Status == uint8(pay.PayStatus_PAY_CLOSED) {
		// TODO 失败处理
		return nil
	}
	return nil
}

// notifyOrderSuccess 支付成功
func (l *NotifyOrderLogic) notifyOrderSuccess(channelCode string, resp *model.OrderResp) error {
	var id uint64
	err := entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		newModel := dbModel.NewModel(tx.Client())
		orderExtension, err := newModel.OrderExtension.UpdateOrderSuccess(l.ctx, resp)
		if err != nil {
			return err
		}
		err = newModel.Order.UpdateOrderSuccess(l.ctx, channelCode, orderExtension, resp)
		if err != nil {
			return err
		}
		id = orderExtension.OrderID
		return nil
	})
	if err != nil {
		return err
	}
	order, err := l.svcCtx.Model.Order.Get(l.ctx, id)
	if err != nil {
		return errorhandler.DefaultEntError(l.Logger, err, id)
	}
	paySuccessPayload, err := json.Marshal(struct {
		MerchantOrderId string `json:"merchantOrderId"`
		PayOrderId      uint64 `json:"payOrderId"`
	}{
		MerchantOrderId: order.MerchantOrderID,
		PayOrderId:      order.ID,
	})
	if err != nil {
		return err
	}
	// TODO 如果不引入 job 模块，typename 如何获取
	_, err = l.svcCtx.AsynqClient.Enqueue(asynq.NewTask("pay_demo_order_success_notify", paySuccessPayload))
	if err != nil {
		return err
	}
	return nil
}
