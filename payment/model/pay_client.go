package model

import (
	"context"
	"encoding/json"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/zeromicro/go-zero/core/errorx"
	"net/http"
	"time"
)

// 支付方式
const (
	Url       = "url"
	Iframe    = "iframe"
	From      = "from"
	QrCode    = "qr_code"
	QrCodeUrl = "qr_code_url"
	BarCode   = "bar_code"
	App       = "app"
)

// 支付状态
const (
	WAITING uint8 = iota + 1
	SUCCESS
	REFUND
	CLOSED
	ERROR
)

// 支付渠道
const (
	Wx       = "Wx"
	WxPub    = "WxPub"
	WxLite   = "WxLite"
	WxApp    = "WxApp"
	WxNative = "WxNative"

	Ali       = "Ali"
	AlipayPc  = "AlipayPc"
	AlipayWap = "AlipayWap"
	AlipayQr  = "AlipayQr"
	AlipayBar = "AlipayBar"
	Mock      = "mock"
)

type (
	// ClientConfig 支付客户端配置
	ClientConfig interface {
		// Validate 参数校验
		Validate() error
	}

	// Properties 支付配置
	Properties struct {
		OrderNotifyUrl  string `json:",env=ORDER_NOTIFY_URL"`
		RefundNotifyUrl string `json:",env=REFUND_NOTIFY_URL"`
		OrderNoPrefix   string `json:",default=P"`
		RefundNoPrefix  string `json:",default=R"`
	}

	PayStrategy interface {
		UnifiedOrder(ctx context.Context, req OrderUnifiedReq) (*OrderResp, error)
	}

	// Client 支付客户端
	Client interface {
		// Init 初始化
		Init() error
		// UnifiedOrder 统一下单
		UnifiedOrder(context.Context, string, OrderUnifiedReq) (*OrderResp, error)
		// GetOrder 获得支付订单
		GetOrder(context.Context, string) (*OrderResp, error)
		// Refresh 刷新配置
		Refresh(config ClientConfig) error
		// UnifiedRefund 退款 返回 WAIT 状态. 后续 job 会轮询
		UnifiedRefund(context.Context, RefundUnifiedReq) (*RefundResp, error)
		// ParseOrderNotify 解析支付回调
		ParseOrderNotify([]byte) (*OrderResp, error)
	}

	OrderUnifiedReq struct {
		UserIp        string
		OutTradeNo    string
		Subject       string
		Body          string
		NotifyUrl     string
		ReturnUrl     string
		Price         int32
		ExpireTime    time.Time
		ChannelExtras map[string]string
		DisplayMode   string
	}

	OrderResp struct {
		Status           uint8
		OutTradeNo       string
		ChannelOrderNo   string
		ChannelUserId    *string
		SuccessTime      time.Time
		RawData          any
		DisplayMode      *string
		DisplayContent   *string
		ChannelErrorCode *string
		ChannelErrorMsg  *string
	}
	RefundUnifiedReq struct {
		OutTradeNo  string
		OutRefundNo string
		Reason      string
		PayPrice    int32
		RefundPrice int32
		NotifyUrl   string
	}

	RefundResp struct {
		Status           uint8
		OutRefundNo      string
		ChannelRefundNo  string
		SuccessTime      time.Time
		RawData          any
		ChannelErrorCode string
		ChannelErrorMsg  string
	}
)

// WaitingOf 创建等待支付订单
func WaitingOf(displayMode, displayContent *string, outTradeNo string, rawData any) *OrderResp {
	return &OrderResp{
		Status:         WAITING,
		DisplayMode:    displayMode,
		DisplayContent: displayContent,
		OutTradeNo:     outTradeNo,
		RawData:        rawData,
	}
}

// SuccessOf 创建支付成功订单
func SuccessOf(channelOrderNo string, channelUserId string, successTime time.Time, outTradeNo string, rawData any) *OrderResp {
	return &OrderResp{
		Status:         SUCCESS,
		ChannelOrderNo: channelOrderNo,
		ChannelUserId:  &channelUserId,
		SuccessTime:    successTime,
		OutTradeNo:     outTradeNo,
		RawData:        rawData,
	}
}

// Of 创建支付订单
func Of(status uint8, channelOrderNo string, channelUserId *string, successTime time.Time, outTradeNo string, rawData any) *OrderResp {
	return &OrderResp{
		Status:         status,
		ChannelOrderNo: channelOrderNo,
		ChannelUserId:  channelUserId,
		SuccessTime:    successTime,
		OutTradeNo:     outTradeNo,
		RawData:        rawData,
	}
}

func CloseOf(channelErrorCode, channelErrorMsg, outTradeNo string, rawData any) *OrderResp {
	return &OrderResp{
		Status:           CLOSED,
		ChannelErrorCode: &channelErrorCode,
		ChannelErrorMsg:  &channelErrorMsg,
		OutTradeNo:       outTradeNo,
		RawData:          rawData,
	}
}

func ParseDate(timeStr string) time.Time {
	parsedTime, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		// 用当前时间代替?
		return time.Now()
	}
	return parsedTime
}

func ParseOrderNotify(code string, r *http.Request) (req []byte, err error) {
	switch code {
	case AlipayPc, AlipayWap,
		AlipayQr, AlipayBar:
		bm, err := alipay.ParseNotifyToBodyMap(r)
		if err != nil {
			return nil, err
		}
		req, err = json.Marshal(bm)
		if err != nil {
			return nil, err
		}
	case WxLite, WxPub,
		WxNative, WxApp:
		notifyReq, err := wechat.V3ParseNotify(r)
		if err != nil {
			return nil, err
		}
		req, err = json.Marshal(notifyReq)
		if err != nil {
			return nil, err
		}
	default:
		return nil, errorx.NewApiBadRequestError("channel code error")
	}
	return
}
