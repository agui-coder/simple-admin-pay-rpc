package ali

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/payment/model"
	"github.com/go-pay/gopay"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
)

// pcPayStrategy
type pcPayStrategy struct {
	*Client
}

func (a *pcPayStrategy) UnifiedOrder(ctx context.Context, req model.OrderUnifiedReq) (*model.OrderResp, error) {
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", req.OutTradeNo)
	bm.Set("subject", req.Subject)
	bm.Set("body", req.Body)
	bm.Set("total_amount", formatAmount(req.Price))
	bm.Set("product_code", "FAST_INSTANT_TRADE_PAY") // 销售产品码. 目前 PC 支付场景下仅支持 FAST_INSTANT_TRADE_PAY,已内置
	// ② 个性化的参数
	// 如果想弄更多个性化的参数，可参考 https://www.pingxx.com/api/支付渠道 extra 参数说明.html 的 alipay_pc_direct 部分进行拓展
	bm.Set("qr_pay_mode", "2") // 跳转模式 - 订单码，效果参见：https://help.pingxx.com/article/1137360/
	bm.Set("time_expire", req.ExpireTime.Format("2006-01-02 15:04:05"))
	if req.DisplayMode == "" {
		req.DisplayMode = model.Url
	}
	a.client.SetReturnUrl(req.ReturnUrl)
	a.client.SetNotifyUrl(req.NotifyUrl)
	payUrl, err := a.client.TradePagePay(ctx, bm)
	if err != nil {
		return nil, err
	}
	return model.WaitingOf(&req.DisplayMode, pointy.GetPointer(payUrl), req.OutTradeNo, payUrl), nil
}
