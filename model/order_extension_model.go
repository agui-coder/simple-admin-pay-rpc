package model

import (
	"context"
	"encoding/json"
	"github.com/agui-coder/simple-admin-pay-rpc/payment/model"

	"github.com/agui-coder/simple-admin-pay-rpc/ent"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"

	"github.com/agui-coder/simple-admin-pay-rpc/ent/orderextension"
	"github.com/agui-coder/simple-admin-pay-rpc/utils/errorhandler"

	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type OrderExtensionModel struct {
	*ent.OrderExtensionClient
}

func NewOrderExtensionModel(client *ent.OrderExtensionClient) *OrderExtensionModel {
	return &OrderExtensionModel{client}
}

func (m *OrderExtensionModel) QueryByNo(ctx context.Context, no string) (*ent.OrderExtension, error) {
	orderExtension, err := m.Query().Where().Where(orderextension.NoEQ(no)).Only(ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(logx.WithContext(ctx), err, no)
	}
	return orderExtension, nil
}

func (m *OrderExtensionModel) UpdateOrderSuccess(ctx context.Context, notifyResp *model.OrderResp) (*ent.OrderExtension, error) {
	orderExtension, err := m.Query().Where().Where(orderextension.NoEQ(notifyResp.OutTradeNo)).Only(ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(logx.WithContext(ctx), err, notifyResp)
	}
	// 更新支付单状态
	if orderExtension.Status == uint8(pay.PayStatus_PAY_SUCCESS) {
		logx.Infof("[updateOrderExtensionSuccess][orderExtension%d 已经是已支付，无需更新]", orderExtension.ID)
		return orderExtension, nil
	}
	if orderExtension.Status != uint8(pay.PayStatus_PAY_WAITING) {
		return nil, errorx.NewInvalidArgumentError("pay order extension status is not waiting")
	}
	notifyData, err := json.Marshal(notifyResp)
	if err != nil {
		return nil, err
	}
	updateCounts, err := m.Update().
		Where(orderextension.IDEQ(orderExtension.ID),
			orderextension.StatusEQ(orderExtension.Status)).
		SetStatus(uint8(pay.PayStatus_PAY_SUCCESS)).
		SetChannelNotifyData(string(notifyData)).Save(ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(logx.WithContext(ctx), err, notifyResp)
	}
	if updateCounts == 0 {
		return nil, errorx.NewInvalidArgumentError("pay order extension status is not waiting")
	}
	logx.Infof("[updateOrderExtensionSuccess][orderExtension:%d 更新为已支付]", orderExtension.ID)
	orderExtension.Status = uint8(pay.PayStatus_PAY_SUCCESS)
	orderExtension.ChannelNotifyData = string(notifyData)
	return orderExtension, nil
}
