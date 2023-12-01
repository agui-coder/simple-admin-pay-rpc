package app

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"

	"github.com/agui-coder/simple-admin-pay-rpc/utils/errorhandler"

	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/agui-coder/simple-admin-pay-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAppLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAppLogic {
	return &GetAppLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAppLogic) GetApp(in *pay.IDReq) (*pay.AppInfo, error) {
	app, err := l.svcCtx.DB.App.Get(l.ctx, in.Id)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	return &pay.AppInfo{
		Id:              &app.ID,
		CreatedAt:       pointy.GetPointer(app.CreatedAt.UnixMilli()),
		UpdatedAt:       pointy.GetPointer(app.UpdatedAt.UnixMilli()),
		Status:          pointy.GetPointer(uint32(app.Status)),
		Name:            &app.Name,
		Remark:          &app.Remark,
		OrderNotifyUrl:  &app.OrderNotifyURL,
		RefundNotifyUrl: &app.RefundNotifyURL,
	}, nil
}
