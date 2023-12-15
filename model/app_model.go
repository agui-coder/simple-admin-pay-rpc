package model

import (
	"context"

	"github.com/agui-coder/simple-admin-pay-rpc/ent"
	"github.com/suyuan32/simple-admin-common/enum/common"

	"github.com/agui-coder/simple-admin-pay-rpc/utils/errorhandler"

	"github.com/zeromicro/go-zero/core/errorx"
	"github.com/zeromicro/go-zero/core/logx"
)

type AppModel struct {
	*ent.AppClient
}

func NewAppModel(client *ent.AppClient) *AppModel {
	return &AppModel{client}
}

func (m *AppModel) ValidPayApp(ctx context.Context, Id uint64) (*ent.App, error) {
	app, err := m.Get(ctx, Id)
	if err != nil {
		return nil, errorhandler.DefaultEntError(logx.WithContext(ctx), err, Id)
	}
	if common.StatusBanned == app.Status {
		return nil, errorx.NewInvalidArgumentError("app is disable")
	}
	return app, nil
}
