package demo

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/utils/dberrorhandler"
	"strconv"
	"time"

	"github.com/agui-coder/simple-admin-pay-rpc/internal/logic/refund"
	"github.com/agui-coder/simple-admin-pay-rpc/internal/svc"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"
	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDemoRefundPaidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDemoRefundPaidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDemoRefundPaidLogic {
	return &UpdateDemoRefundPaidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateDemoRefundPaidLogic) UpdateDemoRefundPaid(in *pay.UpdateDemoRefundPaidReq) (*pay.BaseResp, error) {
	demoOrder, err := l.svcCtx.DB.DemoOrder.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}
	if demoOrder.PayOrderId == in.PayRefundId {
		logx.Errorf("[UpdateDemoRefundPaid][order(%d) 退款单不匹配(%d)，请进行处理！order 数据是：%s]",
			in.Id, in.PayRefundId, demoOrder.String())
		return nil, errorx.NewInvalidArgumentError("demo order refund fail refund order id error")
	}
	refundInfo, err := refund.NewGetRefundByIdLogic(l.ctx, l.svcCtx).GetRefundById(&pay.IDReq{Id: in.PayRefundId})
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}
	if uint8(*refundInfo.Status) != uint8(pay.PayStatus_PAY_SUCCESS) {
		return nil, errorx.NewInvalidArgumentError("demo order refund fail refund not success")
	}
	if *refundInfo.RefundPrice != demoOrder.Price {
		logx.Errorf("[order(%d) payRefund(%d) 退款金额不匹配，请进行处理！order 数据是：%s，payRefund 数据是：%s]",
			in.Id, in.PayRefundId, demoOrder.String(), refundInfo.String())
		return nil, errorx.NewInvalidArgumentError("demo order refund fail refund price not match")
	}
	id := strconv.FormatUint(in.Id, 10) + "-refund"
	if *refundInfo.MerchantRefundId != id {
		logx.Errorf("[UpdateDemoRefundPaid][order(%d) 退款单不匹配(%s)，请进行处理！payRefund 数据是：%s]",
			in.Id, *refundInfo.MerchantRefundId, refundInfo.String())
		return nil, errorx.NewInvalidArgumentError("demo order refund fail refund id not match")
	}
	times := new(time.Time)

	*times = time.UnixMilli(*refundInfo.SuccessTime)
	err = l.svcCtx.DB.DemoOrder.UpdateOne(demoOrder).
		SetNotNilRefundTime(times).
		Exec(l.ctx)
	if err != nil {
		return nil, err
	}
	return &pay.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
