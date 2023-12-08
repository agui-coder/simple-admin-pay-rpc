package model

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/ent"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/refund"
)

type RefundModel struct {
	*ent.RefundClient
}

func NewRefundModel(client *ent.RefundClient) *RefundModel {
	return &RefundModel{client}
}

func (m *RefundModel) SelectByAppIdAndNo(ctx context.Context, appId uint64, no string) (*ent.Refund, error) {
	return m.Query().Where(refund.AppIDEQ(appId), refund.NoEQ(no)).First(ctx)
}
