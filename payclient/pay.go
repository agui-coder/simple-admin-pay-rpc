// Code generated by goctl. DO NOT EDIT.
// Source: pay.proto

package payclient

import (
	"context"

	"github.com/agui-coder/simple-admin-pay-rpc/pay"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AppCreateReq            = pay.AppCreateReq
	AppInfo                 = pay.AppInfo
	AppListResp             = pay.AppListResp
	AppPageReq              = pay.AppPageReq
	AppUpdateReq            = pay.AppUpdateReq
	AppUpdateStatusReq      = pay.AppUpdateStatusReq
	BaseIDResp              = pay.BaseIDResp
	BaseMsg                 = pay.BaseMsg
	BaseResp                = pay.BaseResp
	BaseUUIDResp            = pay.BaseUUIDResp
	ByAppIdAndCodeReq       = pay.ByAppIdAndCodeReq
	ChannelCreateReq        = pay.ChannelCreateReq
	ChannelInfo             = pay.ChannelInfo
	ChannelListReq          = pay.ChannelListReq
	ChannelListResp         = pay.ChannelListResp
	ChannelUpdateReq        = pay.ChannelUpdateReq
	DemoOrderInfo           = pay.DemoOrderInfo
	DemoOrderListResp       = pay.DemoOrderListResp
	DemoOrderPageReq        = pay.DemoOrderPageReq
	Empty                   = pay.Empty
	IDReq                   = pay.IDReq
	IDsReq                  = pay.IDsReq
	NoReq                   = pay.NoReq
	NotifyOrderReq          = pay.NotifyOrderReq
	OrderCreateExtensionReq = pay.OrderCreateExtensionReq
	OrderCreateReq          = pay.OrderCreateReq
	OrderExtensionInfo      = pay.OrderExtensionInfo
	OrderInfo               = pay.OrderInfo
	OrderListResp           = pay.OrderListResp
	OrderPageReq            = pay.OrderPageReq
	PageInfoReq             = pay.PageInfoReq
	PayDemoOrderCreateReq   = pay.PayDemoOrderCreateReq
	RefundCountResp         = pay.RefundCountResp
	RefundCreateReq         = pay.RefundCreateReq
	RefundInfo              = pay.RefundInfo
	RefundListReq           = pay.RefundListReq
	RefundListResp          = pay.RefundListResp
	RefundPageReq           = pay.RefundPageReq
	RefundPageResp          = pay.RefundPageResp
	StringList              = pay.StringList
	UUIDReq                 = pay.UUIDReq
	UUIDsReq                = pay.UUIDsReq
	UpdateDemoOrderPaidReq  = pay.UpdateDemoOrderPaidReq
	ValidateChannelReq      = pay.ValidateChannelReq

	Pay interface {
		// App management
		CreateApp(ctx context.Context, in *AppCreateReq, opts ...grpc.CallOption) (*BaseIDResp, error)
		UpdateApp(ctx context.Context, in *AppUpdateReq, opts ...grpc.CallOption) (*BaseResp, error)
		UpdateAppStatus(ctx context.Context, in *AppUpdateStatusReq, opts ...grpc.CallOption) (*BaseResp, error)
		DeleteApp(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error)
		GetApp(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*AppInfo, error)
		GetAppList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AppListResp, error)
		GetAppPage(ctx context.Context, in *AppPageReq, opts ...grpc.CallOption) (*AppListResp, error)
		InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error)
		// Channel management
		CreateChannel(ctx context.Context, in *ChannelCreateReq, opts ...grpc.CallOption) (*BaseIDResp, error)
		UpdateChannel(ctx context.Context, in *ChannelUpdateReq, opts ...grpc.CallOption) (*BaseResp, error)
		DeleteChannel(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error)
		GetChannelById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*ChannelInfo, error)
		GetChannelListByAppIds(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*ChannelListResp, error)
		GetChannelListByAppIdAndCode(ctx context.Context, in *ByAppIdAndCodeReq, opts ...grpc.CallOption) (*ChannelInfo, error)
		GetEnableChannelList(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*ChannelListResp, error)
		ValidateChannelCanSubmit(ctx context.Context, in *ValidateChannelReq, opts ...grpc.CallOption) (*ChannelInfo, error)
		// Demo management
		CreateDemoOrder(ctx context.Context, in *PayDemoOrderCreateReq, opts ...grpc.CallOption) (*BaseIDResp, error)
		GetDemoOrder(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*DemoOrderInfo, error)
		GetListDemoOrder(ctx context.Context, in *DemoOrderPageReq, opts ...grpc.CallOption) (*DemoOrderListResp, error)
		UpdateDemoOrderPaid(ctx context.Context, in *UpdateDemoOrderPaidReq, opts ...grpc.CallOption) (*BaseResp, error)
		// Order management
		CreateOrder(ctx context.Context, in *OrderCreateReq, opts ...grpc.CallOption) (*BaseIDResp, error)
		GetOrder(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*OrderInfo, error)
		GetOrderPage(ctx context.Context, in *OrderPageReq, opts ...grpc.CallOption) (*OrderListResp, error)
		CreateOrderExtension(ctx context.Context, in *OrderCreateExtensionReq, opts ...grpc.CallOption) (*BaseIDResp, error)
		GetOrderExtension(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*OrderExtensionInfo, error)
		NotifyOrder(ctx context.Context, in *NotifyOrderReq, opts ...grpc.CallOption) (*BaseResp, error)
		ValidateOrderCanSubmit(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*OrderInfo, error)
		// refund management
		CreateRefund(ctx context.Context, in *RefundCreateReq, opts ...grpc.CallOption) (*BaseIDResp, error)
		GetRefundPage(ctx context.Context, in *RefundPageReq, opts ...grpc.CallOption) (*RefundPageResp, error)
		GetRefundById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*RefundInfo, error)
	}

	defaultPay struct {
		cli zrpc.Client
	}
)

func NewPay(cli zrpc.Client) Pay {
	return &defaultPay{
		cli: cli,
	}
}

// App management
func (m *defaultPay) CreateApp(ctx context.Context, in *AppCreateReq, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := pay.NewPayClient(m.cli.Conn())
	return client.CreateApp(ctx, in, opts...)
}

func (m *defaultPay) UpdateApp(ctx context.Context, in *AppUpdateReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := pay.NewPayClient(m.cli.Conn())
	return client.UpdateApp(ctx, in, opts...)
}

func (m *defaultPay) UpdateAppStatus(ctx context.Context, in *AppUpdateStatusReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := pay.NewPayClient(m.cli.Conn())
	return client.UpdateAppStatus(ctx, in, opts...)
}

func (m *defaultPay) DeleteApp(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := pay.NewPayClient(m.cli.Conn())
	return client.DeleteApp(ctx, in, opts...)
}

func (m *defaultPay) GetApp(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*AppInfo, error) {
	client := pay.NewPayClient(m.cli.Conn())
	return client.GetApp(ctx, in, opts...)
}

func (m *defaultPay) GetAppList(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*AppListResp, error) {
	client := pay.NewPayClient(m.cli.Conn())
	return client.GetAppList(ctx, in, opts...)
}

func (m *defaultPay) GetAppPage(ctx context.Context, in *AppPageReq, opts ...grpc.CallOption) (*AppListResp, error) {
	client := pay.NewPayClient(m.cli.Conn())
	return client.GetAppPage(ctx, in, opts...)
}

func (m *defaultPay) InitDatabase(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BaseResp, error) {
	client := pay.NewPayClient(m.cli.Conn())
	return client.InitDatabase(ctx, in, opts...)
}

// Channel management
func (m *defaultPay) CreateChannel(ctx context.Context, in *ChannelCreateReq, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := pay.NewPayClient(m.cli.Conn())
	return client.CreateChannel(ctx, in, opts...)
}

func (m *defaultPay) UpdateChannel(ctx context.Context, in *ChannelUpdateReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := pay.NewPayClient(m.cli.Conn())
	return client.UpdateChannel(ctx, in, opts...)
}

func (m *defaultPay) DeleteChannel(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := pay.NewPayClient(m.cli.Conn())
	return client.DeleteChannel(ctx, in, opts...)
}

func (m *defaultPay) GetChannelById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*ChannelInfo, error) {
	client := pay.NewPayClient(m.cli.Conn())
	return client.GetChannelById(ctx, in, opts...)
}

func (m *defaultPay) GetChannelListByAppIds(ctx context.Context, in *IDsReq, opts ...grpc.CallOption) (*ChannelListResp, error) {
	client := pay.NewPayClient(m.cli.Conn())
	return client.GetChannelListByAppIds(ctx, in, opts...)
}

func (m *defaultPay) GetChannelListByAppIdAndCode(ctx context.Context, in *ByAppIdAndCodeReq, opts ...grpc.CallOption) (*ChannelInfo, error) {
	client := pay.NewPayClient(m.cli.Conn())
	return client.GetChannelListByAppIdAndCode(ctx, in, opts...)
}

func (m *defaultPay) GetEnableChannelList(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*ChannelListResp, error) {
	client := pay.NewPayClient(m.cli.Conn())
	return client.GetEnableChannelList(ctx, in, opts...)
}

func (m *defaultPay) ValidateChannelCanSubmit(ctx context.Context, in *ValidateChannelReq, opts ...grpc.CallOption) (*ChannelInfo, error) {
	client := pay.NewPayClient(m.cli.Conn())
	return client.ValidateChannelCanSubmit(ctx, in, opts...)
}

// Demo management
func (m *defaultPay) CreateDemoOrder(ctx context.Context, in *PayDemoOrderCreateReq, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := pay.NewPayClient(m.cli.Conn())
	return client.CreateDemoOrder(ctx, in, opts...)
}

func (m *defaultPay) GetDemoOrder(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*DemoOrderInfo, error) {
	client := pay.NewPayClient(m.cli.Conn())
	return client.GetDemoOrder(ctx, in, opts...)
}

func (m *defaultPay) GetListDemoOrder(ctx context.Context, in *DemoOrderPageReq, opts ...grpc.CallOption) (*DemoOrderListResp, error) {
	client := pay.NewPayClient(m.cli.Conn())
	return client.GetListDemoOrder(ctx, in, opts...)
}

func (m *defaultPay) UpdateDemoOrderPaid(ctx context.Context, in *UpdateDemoOrderPaidReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := pay.NewPayClient(m.cli.Conn())
	return client.UpdateDemoOrderPaid(ctx, in, opts...)
}

// Order management
func (m *defaultPay) CreateOrder(ctx context.Context, in *OrderCreateReq, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := pay.NewPayClient(m.cli.Conn())
	return client.CreateOrder(ctx, in, opts...)
}

func (m *defaultPay) GetOrder(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*OrderInfo, error) {
	client := pay.NewPayClient(m.cli.Conn())
	return client.GetOrder(ctx, in, opts...)
}

func (m *defaultPay) GetOrderPage(ctx context.Context, in *OrderPageReq, opts ...grpc.CallOption) (*OrderListResp, error) {
	client := pay.NewPayClient(m.cli.Conn())
	return client.GetOrderPage(ctx, in, opts...)
}

func (m *defaultPay) CreateOrderExtension(ctx context.Context, in *OrderCreateExtensionReq, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := pay.NewPayClient(m.cli.Conn())
	return client.CreateOrderExtension(ctx, in, opts...)
}

func (m *defaultPay) GetOrderExtension(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*OrderExtensionInfo, error) {
	client := pay.NewPayClient(m.cli.Conn())
	return client.GetOrderExtension(ctx, in, opts...)
}

func (m *defaultPay) NotifyOrder(ctx context.Context, in *NotifyOrderReq, opts ...grpc.CallOption) (*BaseResp, error) {
	client := pay.NewPayClient(m.cli.Conn())
	return client.NotifyOrder(ctx, in, opts...)
}

func (m *defaultPay) ValidateOrderCanSubmit(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*OrderInfo, error) {
	client := pay.NewPayClient(m.cli.Conn())
	return client.ValidateOrderCanSubmit(ctx, in, opts...)
}

// refund management
func (m *defaultPay) CreateRefund(ctx context.Context, in *RefundCreateReq, opts ...grpc.CallOption) (*BaseIDResp, error) {
	client := pay.NewPayClient(m.cli.Conn())
	return client.CreateRefund(ctx, in, opts...)
}

func (m *defaultPay) GetRefundPage(ctx context.Context, in *RefundPageReq, opts ...grpc.CallOption) (*RefundPageResp, error) {
	client := pay.NewPayClient(m.cli.Conn())
	return client.GetRefundPage(ctx, in, opts...)
}

func (m *defaultPay) GetRefundById(ctx context.Context, in *IDReq, opts ...grpc.CallOption) (*RefundInfo, error) {
	client := pay.NewPayClient(m.cli.Conn())
	return client.GetRefundById(ctx, in, opts...)
}
