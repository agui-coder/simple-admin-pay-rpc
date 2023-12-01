package app

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"

	"github.com/agui-coder/simple-admin-pay-rpc/utils/errorhandler"

	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/agui-coder/simple-admin-pay-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAppPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAppPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAppPageLogic {
	return &GetAppPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAppPageLogic) GetAppPage(in *pay.AppPageReq) (*pay.AppListResp, error) {
	appPage, err := l.svcCtx.DB.App.Query().Where().Page(l.ctx, in.Page, in.PageSize)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	apps := appPage.List
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
	return &pay.AppListResp{AppList: appList, Total: appPage.PageDetails.Total}, nil
}
