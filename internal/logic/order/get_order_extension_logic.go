package order

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"

	"github.com/agui-coder/simple-admin-pay-rpc/utils/errorhandler"

	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/agui-coder/simple-admin-pay-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderExtensionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetOrderExtensionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderExtensionLogic {
	return &GetOrderExtensionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetOrderExtensionLogic) GetOrderExtension(in *pay.IDReq) (*pay.OrderExtensionInfo, error) {
	orderExtension, err := l.svcCtx.DB.OrderExtension.Get(l.ctx, in.Id)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	return &pay.OrderExtensionInfo{
		Id:                &orderExtension.ID,
		CreatedAt:         pointy.GetPointer(orderExtension.CreatedAt.UnixMilli()),
		UpdatedAt:         pointy.GetPointer(orderExtension.UpdatedAt.UnixMilli()),
		Status:            pointy.GetPointer(uint32(orderExtension.Status)),
		No:                &orderExtension.No,
		OrderId:           &orderExtension.OrderID,
		ChannelCode:       &orderExtension.ChannelCode,
		UserIp:            &orderExtension.UserIP,
		ChannelExtras:     orderExtension.ChannelExtras,
		ChannelErrorCode:  &orderExtension.ChannelErrorCode,
		ChannelErrorMsg:   &orderExtension.ChannelErrorMsg,
		ChannelNotifyData: &orderExtension.ChannelNotifyData,
	}, nil
}
