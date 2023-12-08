package refund

import (
	"context"
	"encoding/json"
	"github.com/agui-coder/simple-admin-pay-common/consts"
	payModel "github.com/agui-coder/simple-admin-pay-common/payment/model"
	"github.com/agui-coder/simple-admin-pay-rpc/ent"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/refund"
	"github.com/agui-coder/simple-admin-pay-rpc/internal/logic/notify"
	"github.com/agui-coder/simple-admin-pay-rpc/model"
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
	if notify.Status == consts.SUCCESS {
		err := l.notifyRefundSuccess(channel, notify)
		if err != nil {
			return err
		}
	}
	if notify.Status == consts.FAILURE {
		return nil
	}
	return nil
}

func (l *NotifyRefundLogic) notifyRefundSuccess(channel *ent.Channel, resp *payModel.RefundResp) error {
	refundInfo, err := l.svcCtx.Model.Refund.SelectByAppIdAndNo(l.ctx, channel.AppID, resp.OutRefundNo)
	if refundInfo.Status == consts.SUCCESS {
		logx.Infof("refund success, refundId: %d", refundInfo.ID)
		return errorhandler.DefaultEntError(l.Logger, err, resp)
	}
	if refundInfo.Status != consts.WAITING {
		return errorx.NewInvalidArgumentError("refund status is not waiting")
	}
	var task *ent.NotifyTask
	err = entx.WithTx(l.ctx, l.svcCtx.DB, func(tx *ent.Tx) error {
		channelNotifyData, err := json.Marshal(resp.RawData)
		if err != nil {
			return err
		}
		err = tx.Refund.Update().Where(refund.IDEQ(refundInfo.ID),
			refund.StatusEQ(resp.Status)).
			SetSuccessTime(resp.SuccessTime).
			SetChannelRefundNo(resp.ChannelRefundNo).
			SetStatus(consts.SUCCESS).
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
		if orderInfo.Status == consts.SUCCESS || orderInfo.Status == consts.REFUND {
			return errorx.NewInvalidArgumentError("pay order refund is fail status error")
		}
		if orderInfo.RefundPrice+refundInfo.RefundPrice > orderInfo.Price {
			return errorx.NewInvalidArgumentError("refund price is error")
		}
		err = tx.Order.UpdateOneID(refundInfo.OrderID).
			SetRefundPrice(orderInfo.RefundPrice + refundInfo.RefundPrice).
			SetStatus(consts.REFUND).Exec(l.ctx)
		if err != nil {
			return errorhandler.DefaultEntError(l.Logger, err, orderInfo)
		}
		task, err = model.NewModel(tx.Client()).CreatePayNotifyTask(l.ctx, consts.RefundType, refundInfo.ID)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	err = notify.NewExecuteNotifyLogic(l.ctx, l.svcCtx).ExecuteNotify(task)
	if err != nil {
		return err
	}
	return err
}
