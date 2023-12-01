package refund

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/consts"
	"github.com/agui-coder/simple-admin-pay-rpc/ent"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"

	"github.com/agui-coder/simple-admin-pay-rpc/ent/order"
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
	app, err := l.validPayApp(in.AppId)
	if err != nil {
		return nil, err
	}
	order, err := l.validatePayOrderCanRefund(in)
	if err != nil {
		return nil, err
	}
	channel, err := l.svcCtx.DB.Channel.Get(l.ctx, order.ChannelID)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errorx.NewInvalidArgumentError("channel not found")
		}
		return nil, errorhandler.DefaultEntError(l.Logger, err, order.ChannelID)
	}
	if consts.Disable == channel.Status {
		return nil, errorx.NewInvalidArgumentError("channel is disable")
	}
	//_, err = l.svcCtx.GetPayClient(l.ctx, channel.ID)
	//if err != nil {
	//	logx.Errorf("[validatePayChannelCanSubmit][渠道编号(%d) 找不到对应的支付客户端]", channel.ID)
	//	return nil, err
	//}
	exit, err := l.svcCtx.DB.Refund.Query().Where(refund.AppIDEQ(app.ID), refund.MerchantOrderIDEQ(in.MerchantRefundId)).Exist(l.ctx)
	if err != nil {
		// TODO no 生成
		return nil, errorhandler.DefaultEntError(l.Logger, err, order.ChannelID)
	}
	if exit {
		return nil, errorx.NewInvalidArgumentError("refund exists")
	}
	refund, err := l.svcCtx.DB.Refund.Create().
		SetNo(in.MerchantRefundId).
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
	//unifiedReq, err := client.UnifiedRefund(l.ctx, payClient.RefundUnifiedReq{
	//	OutTradeNo:  order.No,
	//	OutRefundNo: refund.No,
	//	Reason:      in.Reason,
	//	PayPrice:    order.Price,
	//	RefundPrice: refund.RefundPrice,
	//	NotifyUrl:   l.svcCtx.Config.PayProperties.RefundNotifyUrl + "/" + strconv.FormatUint(channel.ID, 10),
	//})
	//if err != nil {
	//	return nil, err
	//}
	//err = l.notifyRefund(channel, unifiedReq)
	if err != nil {
		return nil, err
	}
	return &pay.BaseIDResp{Id: refund.ID, Msg: errormsg.CreateSuccess}, nil
}

func (l *CreateRefundLogic) validPayApp(id uint64) (*ent.App, error) {
	app, err := l.svcCtx.DB.App.Get(l.ctx, id)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, id)
	}
	if app == nil {
		return nil, errorx.NewInvalidArgumentError("app.notFound")
	}
	if consts.Disable == app.Status {
		return nil, errorx.NewInvalidArgumentError("app.isDisable")
	}
	return app, nil
}

func (l *CreateRefundLogic) validatePayOrderCanRefund(in *pay.RefundCreateReq) (*ent.Order, error) {
	order, err := l.svcCtx.DB.Order.Query().Where(order.AppIDEQ(in.AppId), order.MerchantOrderIDEQ(in.MerchantOrderId)).Only(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	if order.Status != consts.SUCCESS && order.Status != consts.REFUND {
		return nil, errorx.NewInvalidArgumentError("pay order refund fail status error")
	}
	if in.Price+order.RefundPrice > order.Price {
		return nil, errorx.NewInvalidArgumentError("refund price exceed")
	}
	count, err := l.svcCtx.DB.Refund.Query().Where(refund.AppIDEQ(in.AppId), refund.OrderIDEQ(order.ID), refund.StatusEQ(consts.WAITING)).Count(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	if count > 0 {
		return nil, errorx.NewInvalidArgumentError("refund has refunding")
	}
	return order, nil
}
