package weixin

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/agui-coder/simple-admin-pay-rpc/payment/model"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/go-playground/validator/v10"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

// ClientConfig API 版本 - V3协议说明  https://pay.weixin.qq.com/wiki/doc/apiv3/wechatpay/wechatpay-1.shtml
// ClientConfig 实现了 pay.ClientConfig 接口
type ClientConfig struct {
	AppId             string `json:"appId" validate:"required"`                      //appId
	MchId             string `json:"mchId" validate:"required"`                      //商户ID 或者服务商模式的 sp_mchid
	PrivateKeyContent []byte `json:"privateKeyContent,optional" validate:"required"` //apiclient_key.pem 证书文件的对应字符串
	SerialNumber      string `json:"serialNumber,optional" validate:"required"`      //apiclient_cert.pem 证书文件的证书号
	ApiV3Key          string `json:"apiV3Key" validate:"required"`                   //apiV3Key，商户平台获取
	Status            bool   `json:"status"`                                         // 是否启用
}

func (p ClientConfig) Validate() error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	err := validate.Struct(p)
	return err
}

type Client struct {
	Config     *ClientConfig
	client     *wechat.ClientV3
	strategies map[string]model.PayStrategy
	once       sync.Once
}

// 编译时接口实现的检查
var _ model.Client = (*Client)(nil)
var instance = &Client{}

func NewClient(config model.ClientConfig) model.Client {
	wxConfig, ok := config.(ClientConfig)
	if !ok {
		logx.Error("config is not of type weixin.ClientConfig")
		return nil
	}

	instance.once.Do(func() {
		instance = &Client{
			Config: &wxConfig,
		}
	})

	return instance
}

func (w *Client) Init() error {
	client, err := wechat.NewClientV3(w.Config.MchId, w.Config.SerialNumber, w.Config.ApiV3Key, string(w.Config.PrivateKeyContent))
	if err != nil {
		logx.Error(err)
		return err
	}
	//certs, err := wechat.GetPlatformCerts(context.Background(), w.Config.MchId, w.Config.ApiV3Key, w.Config.SerialNumber, w.Config.PrivateKeyContent, wechat.CertTypeALL)
	//if certs.Code == wechat.Success && len(certs.Certs) > 0 {
	//	client.SetPlatformCert([]byte(certs.Certs[0].PublicKey), w.Config.SerialNumber)
	//} else {
	//	logx.Error("certs:%s", certs.Error)
	//	return errorx.NewInvalidArgumentError(certs.Error)
	//}
	// 启用自动同步返回验签，并定时更新微信平台API证书（开启自动验签时，无需单独设置微信平台API证书和序列号）
	err = client.AutoVerifySign()
	if err != nil {
		logx.Error(err)
		return err
	}
	// 自定义配置http请求接收返回结果body大小，默认 10MB
	//client.SetBodySize() // 没有特殊需求，可忽略此配置
	// 打开Debug开关，输出日志，默认是关闭的
	client.DebugSwitch = gopay.DebugOn
	w.client = client
	w.strategies = map[string]model.PayStrategy{
		model.WxApp:    &appPayStrategy{w},
		model.WxPub:    &pubPayStrategy{w},
		model.WxNative: &nativePayStrategy{w},
		model.WxLite:   &litePayStrategy{pubPayStrategy{w}},
	}
	logx.Infof("[init][微信 客户端 初始化完成]\n")

	return nil
}

func (w *Client) UnifiedOrder(ctx context.Context, code string, req model.OrderUnifiedReq) (*model.OrderResp, error) {
	strategy, ok := w.strategies[code]
	if !ok {
		return nil, errorx.NewInvalidArgumentError("该客户端不能下单哦")
	}
	return strategy.UnifiedOrder(ctx, req)
}

func (w *Client) GetOrder(ctx context.Context, no string) (*model.OrderResp, error) {
	rsp, err := w.client.V3TransactionQueryOrder(ctx, wechat.OutTradeNo, no)
	if err != nil && rsp.Code != 0 {
		return nil, err
	}
	status, err := parseStatus(rsp.Response.TradeState)
	if err != nil {
		return nil, err
	}
	var openid *string
	if rsp.Response.Payer != nil {
		openid = &rsp.Response.Payer.Openid
	}
	return model.Of(status, rsp.Response.TransactionId, openid, model.ParseDate(rsp.Response.SuccessTime), no, rsp.Response), nil
}

func (w *Client) Refresh(config model.ClientConfig) error {
	if w.Config == config {
		return nil
	}
	logx.Infof("[refresh][微信客户端 发生变化，重新初始化]")
	if clientConfig, ok := config.(ClientConfig); ok {
		w.Config = &clientConfig
	} else {
		return errorx.NewInvalidArgumentError("微信客户端重新初始化失败")
	}
	return nil
}

// buildPayUnifiedOrderBm 通用返回
func (w *Client) buildPayUnifiedOrderBm(req model.OrderUnifiedReq) gopay.BodyMap {
	// 初始化 BodyMap
	bm := make(gopay.BodyMap)
	bm.Set("appid", w.Config.AppId).
		Set("description", req.Body).
		Set("out_trade_no", req.OutTradeNo).
		Set("time_expire", req.ExpireTime.Format(time.RFC3339)).
		Set("notify_url", req.NotifyUrl).
		SetBodyMap("scene_info", func(bm gopay.BodyMap) {
			bm.Set("payer_client_ip", req.UserIp)
		}).
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			bm.Set("total", req.Price).
				Set("currency", "CNY") //CNY：人民币，境内商户号仅支持人民币。
		})
	return bm
}

func (w *Client) ParseOrderNotify(r []byte) (*model.OrderResp, error) {
	req := new(wechat.V3NotifyReq)
	if err := json.Unmarshal(r, req); err != nil {
		return nil, err
	}
	cert := w.client.WxPublicKey()
	err := req.VerifySignByPK(cert)
	if err != nil {
		return nil, err
	}
	result, err := req.DecryptCipherText(string(w.client.ApiV3Key))
	if err != nil {
		return nil, err
	}
	status, err := parseStatus(result.TradeState)
	if err != nil {
		return nil, err
	}
	var openid *string
	if result.Payer != nil {
		openid = &result.Payer.Openid
	}
	return model.Of(status, result.TransactionId, openid, model.ParseDate(result.SuccessTime), result.OutTradeNo, req), nil
}

func (w *Client) UnifiedRefund(ctx context.Context, req model.RefundUnifiedReq) (*model.RefundResp, error) {
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", req.OutTradeNo).
		Set("out_refund_no", req.OutRefundNo).
		SetBodyMap("amount", func(bm gopay.BodyMap) {
			bm.Set("refund", req.RefundPrice).Set("total", req.PayPrice).
				Set("currency", "CNY") //CNY：人民币，境内商户号仅支持人民币。
		}).
		Set("reason", req.Reason).
		Set("notify_url", req.NotifyUrl)

	resp, err := w.client.V3Refund(ctx, bm)
	if err != nil {
		return nil, err
	}
	if resp.Response.Status == "SUCCESS" {
		return &model.RefundResp{
			Status:          model.SUCCESS,
			ChannelRefundNo: resp.Response.RefundId,
			SuccessTime:     model.ParseDate(resp.Response.SuccessTime),
			OutRefundNo:     resp.Response.OutRefundNo,
			RawData:         resp.Response,
		}, nil
	}
	if resp.Response.Status == "PROCESSING" {
		return &model.RefundResp{
			Status:          model.WAITING,
			ChannelRefundNo: resp.Response.RefundId,
			OutRefundNo:     resp.Response.OutRefundNo,
			RawData:         resp.Response,
		}, nil
	}
	return &model.RefundResp{
		Status:      model.ERROR,
		OutRefundNo: req.OutRefundNo,
		RawData:     resp.Response,
	}, nil
}

func parseStatus(tradeState string) (uint8, error) {
	switch tradeState {
	case "NOTPAY":
	case "USERPAYING": // 支付中，等待用户输入密码（条码支付独有）
		return model.WAITING, nil
	case "SUCCESS":
		return model.SUCCESS, nil
	case "REFUND":
		return model.REFUND, nil
	case "CLOSED":
	case "REVOKED": // 已撤销（刷卡支付独有）
	case "PAYERROR": // 支付失败（其它原因，如银行返回失败）
		return model.CLOSED, nil
	default:
		return model.ERROR, errorx.NewInvalidArgumentError(fmt.Sprintf("未知的支付状态%s", tradeState))
	}
	return model.ERROR, errorx.NewInvalidArgumentError(fmt.Sprintf("未知的支付状态%s", tradeState))
}
