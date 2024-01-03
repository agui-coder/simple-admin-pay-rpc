package order

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/ent"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/order"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"
	"github.com/agui-coder/simple-admin-pay-rpc/utils/dberrorhandler"

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
	orderInfo, err := l.svcCtx.DB.Order.Query().Where(order.MerchantOrderIDEQ(in.MerchantOrderId)).Only(l.ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in.MerchantOrderId)
	}
	if orderInfo != nil {
		logx.Infof("[createOrder][merchantOrderId(%s) 已经存在对应的支付单(%s)]", in.MerchantOrderId, orderInfo)
		return &pay.BaseIDResp{Id: orderInfo.ID, Msg: "订单已存在"}, nil
	}
	orderInfo, err = l.svcCtx.DB.Order.Create().
		SetUserIP(in.UserIp).
		SetMerchantOrderID(in.MerchantOrderId).
		SetSubject(in.Subject).
		SetBody(in.Body).
		SetPrice(in.Price).
		SetNotNilExpireTime(pointy.GetTimePointer(&in.ExpireTime, 0)).
		SetStatus(uint8(pay.PayStatus_PAY_WAITING)).
		SetRefundPrice(0).Save(l.ctx)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &pay.BaseIDResp{Id: orderInfo.ID, Msg: i18n.CreateSuccess}, nil
}
