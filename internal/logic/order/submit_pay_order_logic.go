package order

import (
	"context"
	"fmt"
	"github.com/agui-coder/simple-admin-pay-rpc/ent"
	"github.com/agui-coder/simple-admin-pay-rpc/payment/model"
	"github.com/agui-coder/simple-admin-pay-rpc/utils/errorhandler"
	"github.com/agui-coder/simple-admin-pay-rpc/utils/payno"
	"github.com/zeromicro/go-zero/core/errorx"
	"time"

	"github.com/agui-coder/simple-admin-pay-rpc/internal/svc"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubmitPayOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubmitPayOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubmitPayOrderLogic {
	return &SubmitPayOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SubmitPayOrderLogic) SubmitPayOrder(in *pay.OrderSubmitReq) (*pay.OrderSubmitResp, error) {
	order, err := l.ValidateOrderCanSubmit(in.Id)
	if err != nil {
		return nil, err
	}
	no, err := payno.Generate(l.svcCtx.Redis, payno.OrderNoPrefix)
	if err != nil {
		return nil, err
	}
	_, err = NewCreateOrderExtensionLogic(l.ctx, l.svcCtx).CreateOrderExtension(&pay.OrderCreateExtensionReq{
		OrderID:       order.ID,
		ChannelCode:   in.ChannelCode,
		ChannelExtras: in.ChannelExtras,
		No:            no,
		Status:        0,
		UserIP:        in.UserIP,
	})
	if err != nil {
		return nil, err
	}
	unifiedOrderResp, err := l.svcCtx.PayClient[in.ChannelCode].
		UnifiedOrder(l.ctx, in.ChannelCode, model.OrderUnifiedReq{
			DisplayMode:   in.DisplayMode,
			UserIp:        in.UserIP,
			OutTradeNo:    no,
			Subject:       order.Subject,
			Body:          order.Body,
			NotifyUrl:     l.svcCtx.Config.PayProperties.OrderNotifyUrl + "/" + in.ChannelCode,
			ReturnUrl:     in.ReturnUrl,
			Price:         order.Price,
			ExpireTime:    order.ExpireTime,
			ChannelExtras: in.ChannelExtras,
		})
	if err != nil {
		return nil, err
	}
	if unifiedOrderResp != nil {
		go func() {
			err = NewNotifyOrderLogic(context.Background(), l.svcCtx).NotifyOrder0(in.ChannelCode, unifiedOrderResp)
			if err != nil {
				logx.Error(err)
			}
		}()
		if unifiedOrderResp.ChannelErrorCode != nil {
			return nil, errorx.NewInvalidArgumentError(fmt.Sprintf("发起支付报错，错误码：%d，错误提示：%d",
				unifiedOrderResp.ChannelErrorCode, unifiedOrderResp.ChannelErrorMsg))
		}
	}
	return &pay.OrderSubmitResp{Status: uint32(order.Status), DisplayMode: unifiedOrderResp.DisplayMode, DisplayContent: unifiedOrderResp.DisplayContent}, nil
}

func (l *SubmitPayOrderLogic) ValidateOrderCanSubmit(id uint64) (*ent.Order, error) {
	order, err := l.svcCtx.DB.Order.Get(l.ctx, id)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, id)
	}
	if uint8(pay.PayStatus_PAY_SUCCESS) == order.Status { // 校验状态，发现已支付
		return nil, errorx.NewInvalidArgumentError("pay order status is success")
	}
	if uint8(pay.PayStatus_PAY_WAITING) != order.Status { // 校验状态，必须是待支付
		return nil, errorx.NewInvalidArgumentError("pay order status is not waiting")
	}
	if time.Now().After(order.ExpireTime) {
		return nil, errorx.NewInvalidArgumentError("pay order is expired")
	}
	return order, nil
}
