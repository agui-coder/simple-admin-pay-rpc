package refund

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/ent"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/order"
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
	orderInfo, err := l.ValidatePayOrderCanRefund(in)
	if err != nil {
		return nil, err
	}
	client := l.svcCtx.PayClient[orderInfo.ChannelCode]
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
		SetChannelCode(orderInfo.ChannelCode).
		SetOrderID(orderInfo.ID).
		SetOrderNo(orderInfo.No).
		SetMerchantOrderID(in.MerchantOrderId).
		SetMerchantRefundID(in.MerchantRefundId).
		SetReason(in.Reason).
		SetUserIP(in.UserIp).
		SetChannelOrderNo(orderInfo.ChannelOrderNo).
		SetNotNilStatus(pointy.GetPointer(uint8(pay.PayStatus_PAY_WAITING))).
		SetPayPrice(orderInfo.Price).SetRefundPrice(in.Price).Save(l.ctx)

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

func (l *CreateRefundLogic) ValidatePayOrderCanRefund(in *pay.RefundCreateReq) (*ent.Order, error) {
	orderInfo, err := l.svcCtx.DB.Order.Query().Where(order.MerchantOrderIDEQ(in.MerchantOrderId)).Only(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	if orderInfo.Status != uint8(pay.PayStatus_PAY_SUCCESS) && orderInfo.Status != uint8(pay.PayStatus_PAY_REFUND) {
		return nil, errorx.NewInvalidArgumentError("pay order refund fail status error")
	}
	if in.Price+orderInfo.RefundPrice > orderInfo.Price {
		return nil, errorx.NewInvalidArgumentError("refund price exceed")
	}
	count, err := l.svcCtx.DB.Refund.Query().Where(refund.OrderIDEQ(orderInfo.ID),
		refund.StatusEQ(uint8(pay.PayStatus_PAY_WAITING))).Count(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(logx.WithContext(l.ctx), err, in)
	}
	if count > 0 {
		return nil, errorx.NewInvalidArgumentError("refund has refunding")
	}
	return orderInfo, nil
}
