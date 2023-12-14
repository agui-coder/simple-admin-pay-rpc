package order

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"

	"github.com/agui-coder/simple-admin-pay-rpc/utils/errorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/agui-coder/simple-admin-pay-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrderLogic {
	return &CreateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// CreateOrder Order management
func (l *CreateOrderLogic) CreateOrder(in *pay.OrderCreateReq) (*pay.BaseIDResp, error) {
	app, err := l.svcCtx.Model.App.ValidPayApp(l.ctx, in.AppId)

	if err != nil {
		return nil, err
	}
	order, err := l.svcCtx.Model.Order.QueryByAppIdAndMerchantOrderId(l.ctx, in.AppId, in.MerchantOrderId)
	if err != nil {
		return nil, err
	}
	if order != nil {
		logx.Infof("[createOrder][appId(%d) merchantOrderId(%s) 已经存在对应的支付单(%s)]", in.AppId, in.MerchantOrderId, order)
		return &pay.BaseIDResp{Id: order.ID, Msg: "订单已存在"}, nil
	}
	order, err = l.svcCtx.DB.Order.Create().
		SetUserIP(in.UserIp).
		SetMerchantOrderID(in.MerchantOrderId).
		SetSubject(in.Subject).
		SetBody(in.Body).
		SetPrice(in.Price).
		SetNotNilExpireTime(pointy.GetTimePointer(&in.ExpireTime, 0)).
		SetAppID(app.ID).
		SetNotifyURL(app.OrderNotifyURL).
		SetStatus(uint8(pay.PayStatus_PAY_WAITING)).
		SetRefundPrice(0).Save(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &pay.BaseIDResp{Id: order.ID, Msg: i18n.CreateSuccess}, nil
}
