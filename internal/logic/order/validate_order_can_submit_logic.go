package order

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/consts"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"
	"time"

	"github.com/agui-coder/simple-admin-pay-rpc/utils/errorhandler"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/agui-coder/simple-admin-pay-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type ValidateOrderCanSubmitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewValidateOrderCanSubmitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ValidateOrderCanSubmitLogic {
	return &ValidateOrderCanSubmitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ValidateOrderCanSubmitLogic) ValidateOrderCanSubmit(in *pay.IDReq) (*pay.OrderInfo, error) {
	order, err := l.svcCtx.Model.Order.Get(l.ctx, in.Id)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in.Id)
	}
	if consts.SUCCESS == order.Status { // 校验状态，发现已支付
		return nil, errorx.NewInvalidArgumentError("pay order status is success")
	}
	if consts.WAITING != order.Status { // 校验状态，必须是待支付
		return nil, errorx.NewInvalidArgumentError("pay order status is not waiting")
	}
	if time.Now().After(order.ExpireTime) {
		return nil, errorx.NewInvalidArgumentError("pay order is expired")
	}
	return &pay.OrderInfo{
		Id:              &order.ID,
		CreatedAt:       pointy.GetPointer(order.CreatedAt.UnixMilli()),
		UpdatedAt:       pointy.GetPointer(order.UpdatedAt.UnixMilli()),
		Status:          pointy.GetPointer(uint32(order.Status)),
		AppId:           &order.AppID,
		ChannelId:       &order.ChannelID,
		ChannelCode:     &order.ChannelCode,
		MerchantOrderId: &order.MerchantOrderID,
		Subject:         &order.Subject,
		Body:            &order.Body,
		NotifyUrl:       &order.NotifyURL,
		Price:           &order.Price,
		ChannelFeeRate:  &order.ChannelFeeRate,
		ChannelFeePrice: &order.ChannelFeePrice,
		UserIp:          &order.UserIP,
		ExpireTime:      pointy.GetPointer(order.ExpireTime.UnixMilli()),
		SuccessTime:     pointy.GetPointer(order.SuccessTime.UnixMilli()),
		NotifyTime:      pointy.GetPointer(order.NotifyTime.UnixMilli()),
		ExtensionId:     &order.ExtensionID,
		No:              &order.No,
		RefundPrice:     &order.RefundPrice,
		ChannelUserId:   &order.ChannelUserID,
		ChannelOrderNo:  &order.ChannelOrderNo,
	}, nil
}
