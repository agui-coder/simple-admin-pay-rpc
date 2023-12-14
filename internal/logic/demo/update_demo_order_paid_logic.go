package demo

import (
	"context"
	"strconv"
	"time"

	"github.com/agui-coder/simple-admin-pay-rpc/pay"

	"github.com/agui-coder/simple-admin-pay-rpc/internal/logic/order"
	"github.com/agui-coder/simple-admin-pay-rpc/utils/errorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/agui-coder/simple-admin-pay-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateDemoOrderPaidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateDemoOrderPaidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateDemoOrderPaidLogic {
	return &UpdateDemoOrderPaidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateDemoOrderPaidLogic) UpdateDemoOrderPaid(in *pay.UpdateDemoOrderPaidReq) (*pay.BaseResp, error) {
	demoOrder, err := l.svcCtx.DB.DemoOrder.Get(l.ctx, in.Id)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	if demoOrder.PayStatus {
		logx.Errorf("[validateDemoOrderCanPaid][order(%d) 不处于待支付状态，请进行处理！order 数据是：%s]", demoOrder.ID, demoOrder)
		return nil, errorx.NewInvalidArgumentError("DEMO_ORDER_UPDATE_PAID_STATUS_NOT_UNPAID")
	}
	if demoOrder.PayOrderId != in.PayOrderId {
		logx.Errorf("[validateDemoOrderCanPaid][order(%d) 支付订单号不匹配，请进行处理！order 数据是：%s]", demoOrder.ID, demoOrder)
		return nil, errorx.NewInvalidArgumentError("DEMO_ORDER_UPDATE_PAID_PAY_ORDER_ID_NOT_MATCH")
	}
	payOrder, err := order.NewGetOrderLogic(l.ctx, l.svcCtx).GetOrder(&pay.IDReq{Id: in.PayOrderId})
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	if uint8(pay.PayStatus_PAY_SUCCESS) != *pointy.GetStatusPointer(payOrder.Status) {
		logx.Errorf("[validateDemoOrderCanPaid][order(%d) 支付订单未支付，请进行处理！order 数据是：%s]", demoOrder.ID, demoOrder)
		return nil, errorx.NewInvalidArgumentError("DEMO_ORDER_UPDATE_PAID_PAY_ORDER_PAID")
	}
	if *payOrder.Price != demoOrder.Price {
		logx.Errorf("[validateDemoOrderCanPaid][order(%d) 支付订单金额不匹配，请进行处理！order 数据是：%s]", demoOrder.ID, demoOrder)
		return nil, errorx.NewInvalidArgumentError("DEMO_ORDER_UPDATE_PAID_PAY_ORDER_PRICE_NOT_MATCH")
	}
	if *payOrder.MerchantOrderId != strconv.FormatUint(demoOrder.ID, 10) {
		logx.Errorf("[validateDemoOrderCanPaid][order(%d) 支付订单商户订单号不匹配，请进行处理！order 数据是：%s]", demoOrder.ID, demoOrder)
		return nil, errorx.NewInvalidArgumentError("DEMO_ORDER_UPDATE_PAID_PAY_ORDER_MERCHANT_ORDER_ID_NOT_MATCH")
	}
	err = l.svcCtx.DB.DemoOrder.UpdateOneID(in.Id).
		SetPayStatus(true).
		SetPayTime(time.Now()).
		SetPayChannelCode(*payOrder.ChannelCode).Exec(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	return &pay.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
