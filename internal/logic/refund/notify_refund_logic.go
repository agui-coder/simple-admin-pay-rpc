package refund

import (
	"context"
	"encoding/json"
	"github.com/agui-coder/simple-admin-pay-rpc/ent"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/refund"
	"github.com/agui-coder/simple-admin-pay-rpc/payment/model"
	"github.com/agui-coder/simple-admin-pay-rpc/utils/dberrorhandler"
	"github.com/agui-coder/simple-admin-pay-rpc/utils/entx"
	"github.com/hibiken/asynq"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/agui-coder/simple-admin-pay-rpc/internal/svc"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"

	"github.com/zeromicro/go-zero/core/logx"
)

type NotifyRefundLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewNotifyRefundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *NotifyRefundLogic {
	return &NotifyRefundLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *NotifyRefundLogic) NotifyRefund(in *pay.NotifyRefundReq) (*pay.BaseResp, error) {
	refundResp, err := l.svcCtx.PayClient[in.ChannelCode].ParseRefundNotify(in.R)
	if err != nil {
		return nil, errorx.NewInvalidArgumentError("parse refund notify error")
	}
	err = l.ProcessRefundStatus(refundResp)
	if err != nil {
		return nil, err
	}
	return &pay.BaseResp{Msg: i18n.CreateSuccess}, nil
}

func (l *NotifyRefundLogic) ProcessRefundStatus(notify *model.RefundResp) error {
	if notify.Status == uint8(pay.PayStatus_PAY_SUCCESS) {
		err := l.notifyRefundSuccess(notify)
		if err != nil {
			return err
		}
	}
	if notify.Status == uint8(pay.PayStatus_PAY_FAILURE) {
		return l.notifyRefundFailure(notify)
	}
	return nil
}

func (l *NotifyRefundLogic) notifyRefundSuccess(resp *model.RefundResp) error {
	refundInfo, err := l.svcCtx.DB.Refund.Query().Where(refund.NoEQ(resp.OutRefundNo)).First(l.ctx)
	if err != nil {
		return dberrorhandler.DefaultEntError(l.Logger, err, resp)
	}
	if refundInfo.Status == uint8(pay.PayStatus_PAY_SUCCESS) {
		logx.Infof("refund success, refundId: %d", refundInfo.ID)
		return errorx.NewInvalidArgumentError("refund is success")
	}
	if refundInfo.Status != uint8(pay.PayStatus_PAY_WAITING) {
		return errorx.NewInvalidArgumentError("refund status is not waiting")
	}
	err = entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		channelNotifyData, err := json.Marshal(resp.RawData)
		if err != nil {
			return err
		}
		err = tx.Refund.UpdateOneID(refundInfo.ID).
			SetSuccessTime(resp.SuccessTime).
			SetChannelRefundNo(resp.ChannelRefundNo).
			SetStatus(uint8(pay.PayStatus_PAY_SUCCESS)).
			SetChannelNotifyData(string(channelNotifyData)).
			Exec(l.ctx)
		if err != nil {
			return dberrorhandler.DefaultEntError(l.Logger, err, refundInfo.ID)
		}
		logx.Infof("refund success, refundId: %d", refundInfo.ID)
		orderInfo, err := tx.Order.Get(l.ctx, refundInfo.OrderID)
		if err != nil {
			return dberrorhandler.DefaultEntError(l.Logger, err, refundInfo.OrderID)
		}
		if !(orderInfo.Status == uint8(pay.PayStatus_PAY_SUCCESS) || orderInfo.Status == uint8(pay.PayStatus_PAY_REFUND)) {
			return errorx.NewInvalidArgumentError("pay order refund is fail status error")
		}
		if orderInfo.RefundPrice+refundInfo.RefundPrice > orderInfo.Price {
			return errorx.NewInvalidArgumentError("refund price is error")
		}
		err = tx.Order.UpdateOneID(refundInfo.OrderID).
			SetRefundPrice(orderInfo.RefundPrice + refundInfo.RefundPrice).
			SetStatus(uint8(pay.PayStatus_PAY_REFUND)).Exec(l.ctx)
		if err != nil {
			return dberrorhandler.DefaultEntError(l.Logger, err, orderInfo)
		}
		return nil
	})
	if err != nil {
		return err
	}
	notifyRep, err := json.Marshal(struct {
		MerchantOrderId string `json:"merchantOrderId"`
		PayRefundId     uint64 `json:"payRefundId"`
	}{
		PayRefundId:     refundInfo.ID,
		MerchantOrderId: refundInfo.MerchantOrderID,
	})
	if err != nil {
		return err
	}
	// TODO 如果不引入 job 模块，typename 如何获取
	_, err = l.svcCtx.AsynqClient.Enqueue(asynq.NewTask("pay_refund_success_notify", notifyRep))
	if err != nil {
		return err
	}
	return err
}

func (l *NotifyRefundLogic) notifyRefundFailure(resp *model.RefundResp) error {
	refundInfo, err := l.svcCtx.DB.Refund.Query().Where(refund.NoEQ(resp.OutRefundNo)).First(l.ctx)
	if err != nil {
		return dberrorhandler.DefaultEntError(l.Logger, err, resp)
	}
	if refundInfo.Status == uint8(pay.PayStatus_PAY_FAILURE) {
		logx.Infof("refund failure, refundId: %d", refundInfo.ID)
		return errorx.NewInvalidArgumentError("refund is failure")
	}
	if refundInfo.Status != uint8(pay.PayStatus_PAY_WAITING) {
		return errorx.NewInvalidArgumentError("refund status is not waiting")
	}

	channelNotifyData, err := json.Marshal(resp.RawData)
	if err != nil {
		return err
	}
	err = l.svcCtx.DB.Refund.UpdateOneID(refundInfo.ID).
		SetSuccessTime(resp.SuccessTime).
		SetChannelRefundNo(resp.ChannelRefundNo).
		SetStatus(uint8(pay.PayStatus_PAY_FAILURE)).
		SetChannelNotifyData(string(channelNotifyData)).
		Exec(l.ctx)
	if err != nil {
		return dberrorhandler.DefaultEntError(l.Logger, err, refundInfo.ID)
	}
	logx.Infof("refund failure, refundId: %d", refundInfo.ID)

	notifyRep, err := json.Marshal(struct {
		MerchantOrderId string `json:"merchantOrderId"`
		PayRefundId     uint64 `json:"payRefundId"`
	}{
		PayRefundId:     refundInfo.ID,
		MerchantOrderId: refundInfo.MerchantOrderID,
	})
	if err != nil {
		return err
	}
	// TODO 如果不引入 job 模块，typename 如何获取
	_, err = l.svcCtx.AsynqClient.Enqueue(asynq.NewTask("pay_refund_success_notify", notifyRep))
	if err != nil {
		return err
	}
	return nil
}
