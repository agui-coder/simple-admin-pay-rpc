package demo

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"

	"github.com/agui-coder/simple-admin-pay-rpc/utils/errorhandler"

	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/agui-coder/simple-admin-pay-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetListDemoOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetListDemoOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetListDemoOrderLogic {
	return &GetListDemoOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetListDemoOrderLogic) GetListDemoOrder(in *pay.DemoOrderPageReq) (*pay.DemoOrderListResp, error) {
	page, err := l.svcCtx.DB.DemoOrder.Query().Page(l.ctx, in.Page, in.PageSize)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	infos := make([]*pay.DemoOrderInfo, len(page.List))
	for i, demoOrder := range page.List {
		infos[i] = &pay.DemoOrderInfo{
			Id:             &demoOrder.ID,
			CreatedAt:      pointy.GetPointer(demoOrder.CreatedAt.UnixMilli()),
			UpdatedAt:      pointy.GetPointer(demoOrder.UpdatedAt.UnixMilli()),
			UserId:         &demoOrder.UserID,
			SpuId:          &demoOrder.SpuID,
			SpuName:        &demoOrder.SpuName,
			Price:          &demoOrder.Price,
			PayStatus:      &demoOrder.PayStatus,
			PayOrderId:     &demoOrder.PayOrderId,
			PayTime:        pointy.GetPointer(demoOrder.PayTime.UnixMilli()),
			PayChannelCode: &demoOrder.PayChannelCode,
			PayRefundId:    &demoOrder.PayRefundID,
			RefundPrice:    &demoOrder.RefundPrice,
			RefundTime:     pointy.GetPointer(demoOrder.RefundTime.UnixMilli()),
		}
	}
	return &pay.DemoOrderListResp{
		DemoOrderList: infos,
		Total:         page.PageDetails.Total,
	}, nil
}
