package demo

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"
	"github.com/agui-coder/simple-admin-pay-rpc/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/agui-coder/simple-admin-pay-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetDemoOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetDemoOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDemoOrderLogic {
	return &GetDemoOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetDemoOrderLogic) GetDemoOrder(in *pay.IDReq) (*pay.DemoOrderInfo, error) {
	demoOrder, err := l.svcCtx.DB.DemoOrder.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}
	return &pay.DemoOrderInfo{
		Id:             &demoOrder.ID,
		CreatedAt:      pointy.GetPointer(demoOrder.CreatedAt.Unix()),
		UpdatedAt:      pointy.GetPointer(demoOrder.UpdatedAt.Unix()),
		UserId:         &demoOrder.UserID,
		SpuId:          &demoOrder.SpuID,
		SpuName:        &demoOrder.SpuName,
		Price:          &demoOrder.Price,
		PayStatus:      &demoOrder.PayStatus,
		PayOrderId:     &demoOrder.PayOrderId,
		PayTime:        pointy.GetPointer(demoOrder.PayTime.Unix()),
		PayChannelCode: &demoOrder.PayChannelCode,
		PayRefundId:    &demoOrder.PayRefundID,
		RefundPrice:    &demoOrder.RefundPrice,
		RefundTime:     pointy.GetPointer(demoOrder.RefundTime.Unix()),
	}, nil
}
