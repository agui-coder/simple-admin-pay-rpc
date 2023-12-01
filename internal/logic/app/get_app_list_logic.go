package app

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"

	"github.com/agui-coder/simple-admin-pay-rpc/utils/errorhandler"

	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/agui-coder/simple-admin-pay-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAppListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAppListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAppListLogic {
	return &GetAppListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAppListLogic) GetAppList(in *pay.Empty) (*pay.AppListResp, error) {
	apps, err := l.svcCtx.DB.App.Query().Where().All(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	appList := make([]*pay.AppInfo, len(apps))
	for i, app := range apps {
		appList[i] = &pay.AppInfo{
			Id:              &app.ID,
			CreatedAt:       pointy.GetPointer(app.CreatedAt.UnixMilli()),
			UpdatedAt:       pointy.GetPointer(app.UpdatedAt.UnixMilli()),
			Status:          pointy.GetPointer(uint32(app.Status)),
			Name:            &app.Name,
			Remark:          &app.Remark,
			OrderNotifyUrl:  &app.OrderNotifyURL,
			RefundNotifyUrl: &app.RefundNotifyURL,
		}
	}
	return &pay.AppListResp{AppList: appList}, nil
}
