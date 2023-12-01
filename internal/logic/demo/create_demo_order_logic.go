package demo

import (
	"context"
	"fmt"
	"github.com/agui-coder/simple-admin-pay-rpc/pay"
	"strconv"
	"time"

	"github.com/agui-coder/simple-admin-pay-rpc/internal/logic/order"
	"github.com/agui-coder/simple-admin-pay-rpc/utils/errorhandler"

	"github.com/zeromicro/go-zero/core/errorx"

	"github.com/agui-coder/simple-admin-pay-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateDemoOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
	spuNames map[uint64]spu
}

type spu struct {
	name  string
	price int32
}

func NewCreateDemoOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateDemoOrderLogic {
	spuNames := make(map[uint64]spu)
	spuNames[1] = spu{name: "金蓉颗粒", price: 1}
	spuNames[2] = spu{name: "小米手机", price: 2}
	spuNames[3] = spu{name: "苹果手机", price: 3}
	spuNames[4] = spu{name: "oppo手机", price: 4}
	spuNames[5] = spu{name: "vivo手机", price: 5}
	return &CreateDemoOrderLogic{
		ctx:      ctx,
		svcCtx:   svcCtx,
		Logger:   logx.WithContext(ctx),
		spuNames: spuNames,
	}
}

const PayAppId = 2

// Demo management
func (l *CreateDemoOrderLogic) CreateDemoOrder(in *pay.PayDemoOrderCreateReq) (*pay.BaseIDResp, error) {
	spus, ok := l.spuNames[in.SpuId]
	if !ok {
		return nil, errorx.NewInvalidArgumentError(fmt.Sprintf("spuId：%d 商品不存在", in.SpuId))
	}
	demoOrder, err := l.svcCtx.DB.DemoOrder.Create().SetUserID(in.UserId).
		SetSpuID(in.SpuId).
		SetSpuName(spus.name).
		SetPrice(spus.price).
		SetPayStatus(false).
		SetRefundPrice(0).Save(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	payOrder, err := order.NewCreateOrderLogic(l.ctx, l.svcCtx).CreateOrder(&pay.OrderCreateReq{
		AppId:           PayAppId,
		UserIp:          in.UserIp,
		Subject:         spus.name,
		Price:           spus.price,
		MerchantOrderId: strconv.FormatUint(demoOrder.ID, 10),
		Body:            spus.name,
		ExpireTime:      time.Now().Add(time.Hour * 2).Unix(),
	})
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	err = l.svcCtx.DB.DemoOrder.UpdateOne(demoOrder).SetPayOrderId(payOrder.Id).Exec(l.ctx)
	if err != nil {
		return nil, errorhandler.DefaultEntError(l.Logger, err, in)
	}
	return &pay.BaseIDResp{Id: demoOrder.ID}, nil
}
