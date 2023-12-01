package schema

import (
	mixins2 "github.com/agui-coder/simple-admin-pay-rpc/ent/schema/localmixin"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

type Order struct {
	ent.Schema
}

func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("app_id").
			Annotations(entsql.WithComments(true)).
			Comment("应用编号"),
		field.Uint64("channel_id").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("渠道编号"),
		field.String("channel_code").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("渠道编码"),
		field.String("merchant_order_id").
			Annotations(entsql.WithComments(true)).
			Comment("商户订单编号"),
		field.String("subject").
			Annotations(entsql.WithComments(true)).
			Comment("商品标题"),
		field.String("body").
			Annotations(entsql.WithComments(true)).
			Comment("商品描述"),
		field.Text("notify_url").
			Annotations(entsql.WithComments(true)).
			Comment("异步通知地址"),
		field.Int32("price").
			Annotations(entsql.WithComments(true)).
			Comment("支付金额，单位：分"),
		field.Float("channel_fee_rate").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("渠道手续费，单位：百分比"),
		field.Int32("channel_fee_price").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("渠道手续金额，单位：分"),
		field.String("user_ip").
			Annotations(entsql.WithComments(true)).
			Comment("用户 IP"),
		field.Time("expire_time").
			Annotations(entsql.WithComments(true)).
			Comment("订单失效时间"),
		field.Time("success_time").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("订单支付成功时间"),
		field.Time("notify_time").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("订单支付通知时间"),
		field.Uint64("extension_id").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("支付成功的订单拓展单编号"),
		field.String("no").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("订单号"),
		field.Int32("refund_price").
			Annotations(entsql.WithComments(true)).
			Comment("退款总金额，单位：分"),
		field.String("channel_user_id").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("渠道用户编号"),
		field.String("channel_order_no").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("渠道订单号"),
	}
}

func (Order) Edges() []ent.Edge {
	return nil
}

func (Order) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.StatusMixin{},
		mixins2.SoftDeleteMixin{},
	}
}

func (Order) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "pay_order"}}
}
