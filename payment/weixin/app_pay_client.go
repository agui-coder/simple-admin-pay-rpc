package weixin

import (
	"context"
	"encoding/json"

	"github.com/agui-coder/simple-admin-pay-rpc/payment/model"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

// appPayStrategy 结构体继承了 PayStrategy 接口
type appPayStrategy struct {
	*Client
}

func (w *appPayStrategy) UnifiedOrder(ctx context.Context, req model.OrderUnifiedReq) (*model.OrderResp, error) {
	bm := w.buildPayUnifiedOrderBm(req)
	wxRsp, err := w.client.V3TransactionApp(ctx, bm)
	if err != nil {
		logx.Error(err)
		return nil, err
	}
	if wxRsp.Code == wechat.Success {
		result, err := w.client.PaySignOfApp(w.Config.AppId, wxRsp.Response.PrepayId)
		if err != nil {
			return nil, err
		}
		jsonData, err := json.Marshal(result)
		if err != nil {
			return nil, err
		}
		return model.WaitingOf(pointy.GetPointer(model.App),
			pointy.GetPointer(string(jsonData)),
			req.OutTradeNo, result), nil
	}
	logx.Errorf("wxRsp:%s", wxRsp.Error)
	return nil, errorx.NewInvalidArgumentError(wxRsp.Error)
}
