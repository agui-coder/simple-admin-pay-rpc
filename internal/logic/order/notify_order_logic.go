package order

import (
	"context"
	"encoding/json"
	"github.com/agui-coder/simple-admin-pay-rpc/ent"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/order"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/orderextension"
	"github.com/agui-coder/simple-admin-pay-rpc/internal/svc"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"
	"github.com/agui-coder/simple-admin-pay-rpc/payment/model"
	"github.com/agui-coder/simple-admin-pay-rpc/utils/entx"
	"github.com/agui-coder/simple-admin-pay-rpc/utils/errorhandler"
	"github.com/agui-coder/simple-admin-pay-rpc/utils/money"
	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/errorx"

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
		err := l.notifyOrderClosed(resp)
		return err
	}
	return nil
}

// notifyOrderSuccess 支付成功
func (l *NotifyOrderLogic) notifyOrderSuccess(channelCode string, resp *model.OrderResp) error {
	var id uint64
	err := entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		orderExtension, err := l.UpdateOrderExtensionSuccess(resp, tx)
		if err != nil {
			return err
		}
		err = l.UpdateOrderSuccess(tx, channelCode, orderExtension, resp)
		if err != nil {
			return err
		}
		id = orderExtension.OrderID
		return nil
	})
	if err != nil {
		return err
	}
	orderInfo, err := l.svcCtx.DB.Order.Get(l.ctx, id)
	if err != nil {
		return errorhandler.DefaultEntError(l.Logger, err, id)
	}
	// TODO 如果不引入 job 模块，typename 如何获取,消息体如何构建靠约定吗？
	paySuccessPayload, err := json.Marshal(struct {
		MerchantOrderId string `json:"merchantOrderId"`
		PayOrderId      uint64 `json:"payOrderId"`
	}{
		MerchantOrderId: orderInfo.MerchantOrderID,
		PayOrderId:      orderInfo.ID,
	})
	if err != nil {
		return err
	}
	_, err = l.svcCtx.AsynqClient.Enqueue(asynq.NewTask("pay_order_success_notify", paySuccessPayload))
	if err != nil {
		return err
	}
	return nil
}

func (l *NotifyOrderLogic) notifyOrderClosed(resp *model.OrderResp) error {
	extension, err := l.svcCtx.DB.OrderExtension.Query().Where(orderextension.NoEQ(resp.OutTradeNo)).Only(l.ctx)
	if err != nil {
		return errorhandler.DefaultEntError(l.Logger, err, resp)
	}
	if extension.Status == uint8(pay.PayStatus_PAY_CLOSED) {
		logx.Infof("[notifyOrderClosed][orderExtension%d 已经是已关闭，无需更新]", extension.ID)
		return nil
	}
	if extension.Status == uint8(pay.PayStatus_PAY_SUCCESS) {
		logx.Infof("[notifyOrderClosed][orderExtension%d 已经是已支付，无需更新]", extension.ID)
		return nil
	}
	if extension.Status != uint8(pay.PayStatus_PAY_WAITING) {
		return errorx.NewInvalidArgumentError("pay order extension status is not waiting")
	}
	notify, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	err = l.svcCtx.DB.OrderExtension.UpdateOne(extension).SetStatus(uint8(pay.PayStatus_PAY_CLOSED)).
		SetChannelNotifyData(string(notify)).
		SetNotNilChannelErrorCode(resp.ChannelErrorCode).
		SetNotNilChannelErrorMsg(resp.ChannelErrorMsg).
		Exec(l.ctx)
	if err != nil {
		return errorhandler.DefaultEntError(l.Logger, err, resp)
	}
	logx.Infof("[notifyOrderClosed][orderExtension:%d 更新为已关闭]", extension.ID)
	return nil
}

func (l *NotifyOrderLogic) UpdateOrderExtensionSuccess(notifyResp *model.OrderResp, tx *ent.Tx) (*ent.OrderExtension, error) {
	orderExtension, err := tx.OrderExtension.Query().Where().Where(orderextension.NoEQ(notifyResp.OutTradeNo)).Only(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, notifyResp)
	}
	// 更新支付单状态
	if orderExtension.Status == uint8(pay.PayStatus_PAY_SUCCESS) {
		logx.Infof("[updateOrderExtensionSuccess][orderExtension%d 已经是已支付，无需更新]", orderExtension.ID)
		return orderExtension, nil
	}
	if orderExtension.Status != uint8(pay.PayStatus_PAY_WAITING) {
		return nil, errorx.NewInvalidArgumentError("pay order extension status is not waiting")
	}
	notifyData, err := json.Marshal(notifyResp)
	if err != nil {
		return nil, err
	}
	updateCounts, err := tx.OrderExtension.Update().
		Where(orderextension.IDEQ(orderExtension.ID),
			orderextension.StatusEQ(orderExtension.Status)).
		SetStatus(uint8(pay.PayStatus_PAY_SUCCESS)).
		SetChannelNotifyData(string(notifyData)).Save(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, notifyResp)
	}
	if updateCounts == 0 {
		return nil, errorx.NewInvalidArgumentError("pay order extension status is not waiting")
	}
	logx.Infof("[updateOrderExtensionSuccess][orderExtension:%d 更新为已支付]", orderExtension.ID)
	orderExtension.Status = uint8(pay.PayStatus_PAY_SUCCESS)
	orderExtension.ChannelNotifyData = string(notifyData)
	return orderExtension, nil
}

func (l *NotifyOrderLogic) UpdateOrderSuccess(tx *ent.Tx, channelCode string, orderExtension *ent.OrderExtension, notifyResp *model.OrderResp) error {
	orderEnt, err := tx.Order.Get(l.ctx, orderExtension.OrderID)
	if err != nil {
		return errorhandler.DefaultEntError(l.Logger, err, notifyResp)
	}
	if orderEnt.Status == uint8(pay.PayStatus_PAY_SUCCESS) && orderEnt.ExtensionID == orderExtension.ID {
		logx.Infof("[updateOrderExtensionSuccess][order:%d 已经是已支付，无需更新]", orderEnt.ID)
		return nil
	}
	if orderEnt.Status != uint8(pay.PayStatus_PAY_WAITING) {
		return errorx.NewInvalidArgumentError("pay order status is not waiting")
	}
	// TODO FeeRate 怎么定
	channelFeePrice, err := money.CalculateRatePriceInternal(orderEnt.Price, 0.0)
	if err != nil {
		return err
	}
	err = tx.Order.Update().Where(order.IDEQ(orderEnt.ID), order.StatusEQ(uint8(pay.PayStatus_PAY_WAITING))).
		SetStatus(uint8(pay.PayStatus_PAY_SUCCESS)).SetChannelCode(channelCode).
		SetSuccessTime(notifyResp.SuccessTime).SetExtensionID(orderExtension.ID).SetNo(orderExtension.No).
		SetChannelOrderNo(notifyResp.ChannelOrderNo).SetNotNilChannelUserID(notifyResp.ChannelUserId).
		SetChannelFeeRate(0.0).SetChannelFeePrice(channelFeePrice).Exec(l.ctx)
	if err != nil {
		return errorhandler.DefaultEntError(l.Logger, err, notifyResp)
	}
	logx.Infof("[updateOrderExtensionSuccess][order %v 更新为已支付]", orderEnt)
	return nil
}
