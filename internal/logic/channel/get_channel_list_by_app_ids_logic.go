package channel

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"

	"github.com/agui-coder/simple-admin-pay-rpc/ent/channel"
	"github.com/agui-coder/simple-admin-pay-rpc/utils/errorhandler"

	"github.com/suyuan32/simple-admin-common/utils/pointy"

	"github.com/agui-coder/simple-admin-pay-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChannelListByAppIdsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetChannelListByAppIdsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChannelListByAppIdsLogic {
	return &GetChannelListByAppIdsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetChannelListByAppIdsLogic) GetChannelListByAppIds(in *pay.IDsReq) (*pay.ChannelListResp, error) {
	channels, err := l.svcCtx.DB.Channel.Query().Where(channel.AppIDIn(in.Ids...)).All(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	channelList := make([]*pay.ChannelInfo, len(channels))
	for i, channel := range channels {
		channelList[i] = &pay.ChannelInfo{
			Id:        &channel.ID,
			CreatedAt: pointy.GetPointer(channel.CreatedAt.UnixMilli()),
			UpdatedAt: pointy.GetPointer(channel.UpdatedAt.UnixMilli()),
			Status:    pointy.GetPointer(uint32(channel.Status)),
			Code:      &channel.Code,
			Remark:    &channel.Remark,
			FeeRate:   &channel.FeeRate,
			AppId:     &channel.AppID,
			Config:    &channel.Config,
		}
	}
	return &pay.ChannelListResp{Data: channelList}, nil
}
