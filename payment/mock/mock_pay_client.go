package mock

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/payment/model"
	"github.com/zeromicro/go-zero/core/logx"
	"time"

	"github.com/zeromicro/go-zero/core/errorx"
)

// ClientConfig 实现了 model.ClientConfig 接口 模拟支付方便调试
type ClientConfig struct {
	Name string
}

func (p ClientConfig) Validate() error {
	return nil
}

const respSuccessData = "MOCK_SUCCESS"

// Client 结构体实现了 model.Client 接口
type Client struct {
	ChannelId uint64
	Config    *ClientConfig
}

func (c *Client) ParseOrderNotify(r []byte) (*model.OrderResp, error) {
	return nil, errorx.NewInvalidArgumentError("mock no parse order notify")
}

func (c *Client) ParseRefundNotify(r []byte) (*model.RefundResp, error) {
	return nil, errorx.NewInvalidArgumentError("mock no parse refund notify")
}

// 编译时接口实现的检查
var _ model.Client = (*Client)(nil)

func NewMockPayClient(channelId uint64, config model.ClientConfig) model.Client {
	mockConfig, ok := config.(ClientConfig)
	if !ok {
		logx.Error("config is not of type mock.ClientConfig")
		return nil
	}
	return &Client{
		Config: &mockConfig, ChannelId: channelId,
	}
}

func (c *Client) Init() error {
	return nil
}

func (c *Client) UnifiedOrder(ctx context.Context, code string, req model.OrderUnifiedReq) (*model.OrderResp, error) {
	return model.SuccessOf("MOCK-P-"+req.OutTradeNo, "", time.Now(), req.OutTradeNo, respSuccessData), nil
}

func (c *Client) GetId() uint64 {
	return c.ChannelId
}

func (c *Client) GetOrder(ctx context.Context, outTradeNo string) (*model.OrderResp, error) {
	return model.SuccessOf("MOCK-P-"+outTradeNo, "", time.Now(),
		outTradeNo, respSuccessData), nil
}

func (c *Client) Refresh(config model.ClientConfig) error {
	return nil
}

func (c *Client) UnifiedRefund(ctx context.Context, req model.RefundUnifiedReq) (*model.RefundResp, error) {
	return nil, nil
}
