package refund

import (
	"context"
	"encoding/json"
	"github.com/hibiken/asynq"
	"github.com/suyuan32/simple-admin-job/types/pattern"
	"github.com/suyuan32/simple-admin-job/types/payload"

	payModel "github.com/agui-coder/simple-admin-pay-common/payment/model"
	"github.com/agui-coder/simple-admin-pay-rpc/ent"
	"github.com/agui-coder/simple-admin-pay-rpc/utils/entx"
	"github.com/agui-coder/simple-admin-pay-rpc/utils/errorhandler"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
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
	channel, err := l.svcCtx.Model.Channel.ValidPayChannelById(l.ctx, in.ChannelId)
	if err != nil {
		return nil, err
	}
	go func() {
		err = l.ProcessRefundStatus(channel, &payModel.RefundResp{
			OutRefundNo:     in.OutRefundNo,
			ChannelRefundNo: in.ChannelRefundNo,
			Status:          uint8(in.Status),
			SuccessTime:     *pointy.GetTimePointer(&in.SuccessTime, 0),
			RawData:         in.ChannelNotifyData,
		})
		logx.Error(err.Error())
	}()
	return &pay.BaseResp{}, nil
}

func (l *NotifyRefundLogic) ProcessRefundStatus(channel *ent.Channel, notify *payModel.RefundResp) error {
	if notify.Status == uint8(pay.PayStatus_PAY_SUCCESS) {
		err := l.notifyRefundSuccess(channel, notify)
		if err != nil {
			return err
		}
	}
	if notify.Status == uint8(pay.PayStatus_PAY_FAILURE) {
		return nil
	}
	return nil
}

func (l *NotifyRefundLogic) notifyRefundSuccess(channel *ent.Channel, resp *payModel.RefundResp) error {
	refundInfo, err := l.svcCtx.Model.Refund.SelectByAppIdAndNo(l.ctx, channel.AppID, resp.OutRefundNo)
	if err != nil {
		return errorhandler.DefaultEntError(l.Logger, err, resp)
	}
	if refundInfo.Status == uint8(pay.PayStatus_PAY_SUCCESS) {
		logx.Infof("refund success, refundId: %d", refundInfo.ID)
		return errorhandler.DefaultEntError(l.Logger, err, resp)
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
			return errorhandler.DefaultEntError(l.Logger, err, refundInfo.ID)
		}
		logx.Infof("refund success, refundId: %d", refundInfo.ID)
		orderInfo, err := tx.Order.Get(l.ctx, refundInfo.OrderID)
		if err != nil {
			return errorhandler.DefaultEntError(l.Logger, err, refundInfo.OrderID)
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
			return errorhandler.DefaultEntError(l.Logger, err, orderInfo)
		}
		return nil
	})
	if err != nil {
		return err
	}
	notifyRep, err := json.Marshal(payload.PayRefundNotifyReq{
		PayRefundId:     refundInfo.ID,
		MerchantOrderId: refundInfo.MerchantOrderID,
	})
	if err != nil {
		return err
	}
	_, err = l.svcCtx.AsynqClient.Enqueue(asynq.NewTask(pattern.PayRefundSuccessNotify, notifyRep))
	if err != nil {
		return err
	}
	return err
}
