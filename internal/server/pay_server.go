// Code generated by goctl. DO NOT EDIT.
// Source: pay.proto

package server

import (
	"context"

	"github.com/agui-coder/simple-admin-pay-rpc/internal/logic/base"
	"github.com/agui-coder/simple-admin-pay-rpc/internal/logic/demo"
	"github.com/agui-coder/simple-admin-pay-rpc/internal/logic/order"
	"github.com/agui-coder/simple-admin-pay-rpc/internal/logic/refund"
	"github.com/agui-coder/simple-admin-pay-rpc/internal/svc"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"
)

type PayServer struct {
	svcCtx *svc.ServiceContext
	pay.UnimplementedPayServer
}

func NewPayServer(svcCtx *svc.ServiceContext) *PayServer {
	return &PayServer{
		svcCtx: svcCtx,
	}
}

func (s *PayServer) InitDatabase(ctx context.Context, in *pay.Empty) (*pay.BaseResp, error) {
	l := base.NewInitDatabaseLogic(ctx, s.svcCtx)
	return l.InitDatabase(in)
}

// Demo management
func (s *PayServer) CreateDemoOrder(ctx context.Context, in *pay.PayDemoOrderCreateReq) (*pay.BaseIDResp, error) {
	l := demo.NewCreateDemoOrderLogic(ctx, s.svcCtx)
	return l.CreateDemoOrder(in)
}

func (s *PayServer) GetDemoOrder(ctx context.Context, in *pay.IDReq) (*pay.DemoOrderInfo, error) {
	l := demo.NewGetDemoOrderLogic(ctx, s.svcCtx)
	return l.GetDemoOrder(in)
}

func (s *PayServer) GetListDemoOrder(ctx context.Context, in *pay.DemoOrderPageReq) (*pay.DemoOrderListResp, error) {
	l := demo.NewGetListDemoOrderLogic(ctx, s.svcCtx)
	return l.GetListDemoOrder(in)
}

func (s *PayServer) UpdateDemoOrderPaid(ctx context.Context, in *pay.UpdateDemoOrderPaidReq) (*pay.BaseResp, error) {
	l := demo.NewUpdateDemoOrderPaidLogic(ctx, s.svcCtx)
	return l.UpdateDemoOrderPaid(in)
}

func (s *PayServer) RefundDemoOrder(ctx context.Context, in *pay.RefundDemoOrderReq) (*pay.BaseResp, error) {
	l := demo.NewRefundDemoOrderLogic(ctx, s.svcCtx)
	return l.RefundDemoOrder(in)
}

func (s *PayServer) UpdateDemoRefundPaid(ctx context.Context, in *pay.UpdateDemoRefundPaidReq) (*pay.BaseResp, error) {
	l := demo.NewUpdateDemoRefundPaidLogic(ctx, s.svcCtx)
	return l.UpdateDemoRefundPaid(in)
}

// Order management
func (s *PayServer) CreateOrder(ctx context.Context, in *pay.OrderCreateReq) (*pay.BaseIDResp, error) {
	l := order.NewCreateOrderLogic(ctx, s.svcCtx)
	return l.CreateOrder(in)
}

func (s *PayServer) GetOrder(ctx context.Context, in *pay.IDReq) (*pay.OrderInfo, error) {
	l := order.NewGetOrderLogic(ctx, s.svcCtx)
	return l.GetOrder(in)
}

func (s *PayServer) GetOrderPage(ctx context.Context, in *pay.OrderPageReq) (*pay.OrderListResp, error) {
	l := order.NewGetOrderPageLogic(ctx, s.svcCtx)
	return l.GetOrderPage(in)
}

func (s *PayServer) CreateOrderExtension(ctx context.Context, in *pay.OrderCreateExtensionReq) (*pay.BaseIDResp, error) {
	l := order.NewCreateOrderExtensionLogic(ctx, s.svcCtx)
	return l.CreateOrderExtension(in)
}

func (s *PayServer) GetOrderExtension(ctx context.Context, in *pay.IDReq) (*pay.OrderExtensionInfo, error) {
	l := order.NewGetOrderExtensionLogic(ctx, s.svcCtx)
	return l.GetOrderExtension(in)
}

func (s *PayServer) NotifyOrder(ctx context.Context, in *pay.NotifyOrderReq) (*pay.BaseResp, error) {
	l := order.NewNotifyOrderLogic(ctx, s.svcCtx)
	return l.NotifyOrder(in)
}

func (s *PayServer) SubmitPayOrder(ctx context.Context, in *pay.OrderSubmitReq) (*pay.OrderSubmitResp, error) {
	l := order.NewSubmitPayOrderLogic(ctx, s.svcCtx)
	return l.SubmitPayOrder(in)
}

// refund management
func (s *PayServer) CreateRefund(ctx context.Context, in *pay.RefundCreateReq) (*pay.BaseIDResp, error) {
	l := refund.NewCreateRefundLogic(ctx, s.svcCtx)
	return l.CreateRefund(in)
}

func (s *PayServer) GetRefundPage(ctx context.Context, in *pay.RefundPageReq) (*pay.RefundPageResp, error) {
	l := refund.NewGetRefundPageLogic(ctx, s.svcCtx)
	return l.GetRefundPage(in)
}

func (s *PayServer) GetRefundById(ctx context.Context, in *pay.IDReq) (*pay.RefundInfo, error) {
	l := refund.NewGetRefundByIdLogic(ctx, s.svcCtx)
	return l.GetRefundById(in)
}

func (s *PayServer) NotifyRefund(ctx context.Context, in *pay.NotifyRefundReq) (*pay.BaseResp, error) {
	l := refund.NewNotifyRefundLogic(ctx, s.svcCtx)
	return l.NotifyRefund(in)
}
