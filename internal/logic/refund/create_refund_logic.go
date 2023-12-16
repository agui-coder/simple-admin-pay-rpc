package refund

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/refund"
	"github.com/agui-coder/simple-admin-pay-rpc/internal/svc"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"
	"github.com/agui-coder/simple-admin-pay-rpc/payment/model"
	"github.com/agui-coder/simple-admin-pay-rpc/utils/errorhandler"
	"github.com/agui-coder/simple-admin-pay-rpc/utils/payno"

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
	order, err := l.svcCtx.Model.ValidatePayOrderCanRefund(l.ctx, in)
	if err != nil {
		return nil, err
	}
	client := l.svcCtx.PayClient[order.ChannelCode]
	exit, err := l.svcCtx.DB.Refund.Query().Where(refund.MerchantOrderIDEQ(in.MerchantRefundId)).Exist(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in.MerchantRefundId)
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
		SetChannelCode(order.ChannelCode).
		SetOrderID(order.ID).
		SetOrderNo(order.No).
		SetMerchantOrderID(in.MerchantOrderId).
		SetMerchantRefundID(in.MerchantRefundId).
		SetReason(in.Reason).
		SetUserIP(in.UserIp).
		SetChannelOrderNo(order.ChannelOrderNo).
		SetNotNilStatus(pointy.GetPointer(uint8(pay.PayStatus_PAY_WAITING))).
		SetPayPrice(order.Price).SetRefundPrice(in.Price).Save(l.ctx)

	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	go func() {
		newCtx := context.Background()
		refundUnifiedResp, err := client.UnifiedRefund(newCtx, model.RefundUnifiedReq{
			OutTradeNo:  refundInfo.OrderNo,
			OutRefundNo: refundInfo.No,
			Reason:      refundInfo.Reason,
			PayPrice:    refundInfo.PayPrice,
			RefundPrice: refundInfo.RefundPrice,
			NotifyUrl:   l.svcCtx.Config.PayProperties.RefundNotifyUrl + "/" + refundInfo.ChannelCode,
		})
		if err != nil {
			logx.Errorf("[CreateRefund][退款 id:%d err:%s", refundInfo.ID, err.Error())
		}
		err = NewNotifyRefundLogic(newCtx, l.svcCtx).ProcessRefundStatus(refundUnifiedResp)
		if err != nil {
			logx.Errorf(err.Error())
		}
	}()
	return &pay.BaseIDResp{Id: refundInfo.ID, Msg: errormsg.CreateSuccess}, nil
}
