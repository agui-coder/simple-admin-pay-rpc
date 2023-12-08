package refund

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-common/consts"
	"github.com/agui-coder/simple-admin-pay-common/payment/model"
	"github.com/agui-coder/simple-admin-pay-common/payno"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"
	"strconv"

	"github.com/agui-coder/simple-admin-pay-rpc/ent/refund"
	"github.com/agui-coder/simple-admin-pay-rpc/internal/svc"
	"github.com/agui-coder/simple-admin-pay-rpc/utils/errorhandler"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRefundLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateRefundLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRefundLogic {
	return &CreateRefundLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateRefundLogic) CreateRefund(in *pay.RefundCreateReq) (*pay.BaseIDResp, error) {
	app, err := l.svcCtx.Model.App.ValidPayApp(l.ctx, in.AppId)
	if err != nil {
		return nil, err
	}
	order, err := l.svcCtx.Model.ValidatePayOrderCanRefund(l.ctx, in)
	if err != nil {
		return nil, err
	}
	channel, err := l.svcCtx.Model.Channel.ValidPayChannelById(l.ctx, order.ChannelID)
	if err != nil {
		return nil, err
	}
	client, err := l.svcCtx.GetPayClient(l.ctx, channel.ID)
	if err != nil {
		logx.Errorf("get pay client id:%d  error: %v", channel.ID, err)
		return nil, err
	}
	exit, err := l.svcCtx.DB.Refund.Query().Where(refund.AppIDEQ(app.ID),
		refund.MerchantOrderIDEQ(in.MerchantRefundId)).Exist(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, order.ChannelID)
	}
	if exit {
		return nil, errorx.NewInvalidArgumentError("refund exists")
	}
	refundNo, err := payno.Generate(l.svcCtx.Redis, payno.RefundNoPrefix)
	if err != nil {
		return nil, err
	}
	refundInfo, err := l.svcCtx.DB.Refund.Create().
		SetNo(refundNo).
		SetAppID(in.AppId).
		SetChannelID(channel.ID).
		SetChannelCode(channel.Code).
		SetOrderID(order.ID).
		SetOrderNo(order.No).
		SetMerchantOrderID(in.MerchantOrderId).
		SetMerchantRefundID(in.MerchantRefundId).
		SetNotifyURL(app.RefundNotifyURL).
		SetReason(in.Reason).
		SetUserIP(in.UserIp).
		SetChannelOrderNo(order.ChannelOrderNo).
		SetNotNilStatus(pointy.GetPointer(consts.WAITING)).SetPayPrice(order.Price).SetRefundPrice(in.Price).Save(l.ctx)

	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	refundUnifiedResp, err := client.UnifiedRefund(l.ctx, model.RefundUnifiedReq{
		OutTradeNo:  refundInfo.OrderNo,
		OutRefundNo: refundInfo.No,
		Reason:      refundInfo.Reason,
		PayPrice:    refundInfo.PayPrice,
		RefundPrice: refundInfo.RefundPrice,
		NotifyUrl:   l.svcCtx.Config.PayProperties.RefundNotifyUrl + "/" + strconv.FormatUint(refundInfo.ChannelID, 10),
	})
	if err != nil {
		logx.Errorf("[RefundApproved][退款 id:%d err:%s", refundInfo.ID, err.Error())
		return nil, errorx.NewInvalidArgumentError(err.Error())
	}

	err = NewNotifyRefundLogic(l.ctx, l.svcCtx).ProcessRefundStatus(channel, refundUnifiedResp)
	if err != nil {
		return nil, err
	}
	return &pay.BaseIDResp{Id: refundInfo.ID, Msg: errormsg.CreateSuccess}, nil
}
