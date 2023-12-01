package app

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"

	"github.com/agui-coder/simple-admin-pay-rpc/utils/errorhandler"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/agui-coder/simple-admin-pay-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateAppLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateAppLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAppLogic {
	return &CreateAppLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// App management
func (l *CreateAppLogic) CreateApp(in *pay.AppCreateReq) (*pay.BaseIDResp, error) {
	app, err := l.svcCtx.DB.App.Create().
		SetName(in.Name).
		SetNotNilStatus(pointy.GetStatusPointer(&in.Status)).
		SetNotNilRemark(in.Remark).
		SetOrderNotifyURL(in.OrderNotifyUrl).
		SetRefundNotifyURL(in.RefundNotifyUrl).Save(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	return &pay.BaseIDResp{Id: app.ID, Msg: i18n.CreateSuccess}, nil
}
