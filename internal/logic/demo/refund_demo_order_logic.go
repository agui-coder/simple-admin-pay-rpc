package demo

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/internal/logic/refund"
	"github.com/agui-coder/simple-admin-pay-rpc/utils/dberrorhandler"
	"github.com/zeromicro/go-zero/core/errorx"
	"strconv"

	"github.com/agui-coder/simple-admin-pay-rpc/internal/svc"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"

	"github.com/zeromicro/go-zero/core/logx"
)

type RefundDemoOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRefundDemoOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RefundDemoOrderLogic {
	return &RefundDemoOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RefundDemoOrderLogic) RefundDemoOrder(in *pay.RefundDemoOrderReq) (*pay.BaseResp, error) {
	order, err := l.svcCtx.DB.DemoOrder.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, "demo order not found")
	}
	if !order.PayStatus {
		return nil, errorx.NewInvalidArgumentError("demo order not pay")
	}
	if order.PayRefundID != 0 {
		return nil, errorx.NewInvalidArgumentError("demo order already refund")
	}
	refundId := strconv.FormatUint(order.ID, 10) + "-refund"
	resp, err := refund.NewCreateRefundLogic(l.ctx, l.svcCtx).CreateRefund(&pay.RefundCreateReq{
		UserIp:           in.UserIp,
		MerchantOrderId:  strconv.FormatUint(in.Id, 10),
		MerchantRefundId: refundId,
		Reason:           "想退钱",
		Price:            order.Price,
	})
	if err != nil {
		return nil, err
	}
	err = l.svcCtx.DB.DemoOrder.UpdateOneID(in.Id).SetPayRefundID(resp.Id).
		SetRefundPrice(order.Price).Exec(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, "update demo order error")
	}
	return &pay.BaseResp{Msg: resp.Msg}, nil
}
