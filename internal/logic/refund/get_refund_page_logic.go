package refund

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"

	"github.com/agui-coder/simple-admin-pay-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRefundPageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRefundPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRefundPageLogic {
	return &GetRefundPageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetRefundPageLogic) GetRefundPage(in *pay.RefundPageReq) (*pay.RefundPageResp, error) {
	// todo: add your logic here and delete this line

	return &pay.RefundPageResp{}, nil
}
