package schema

import (
	mixins2 "github.com/agui-coder/simple-admin-pay-rpc/ent/schema/localmixin"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

type DemoOrder struct {
	ent.Schema
}

func (DemoOrder) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id").
			Annotations(entsql.WithComments(true)).
			Comment("用户编号"),
		field.Uint64("spu_id").
			Annotations(entsql.WithComments(true)).
			Comment("商品编号"),
		field.String("spu_name").
			Annotations(entsql.WithComments(true)).
			Comment("商品名称"),
		field.Int32("price").
			Annotations(entsql.WithComments(true)).
			Comment("价格，单位：分"),
		field.Bool("pay_status").
			Annotations(entsql.WithComments(true)).
			Comment("是否支付"),
		field.Uint64("pay_orderId").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("支付订单编号"),
		field.Time("pay_time").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("付款时间"),
		field.String("pay_channel_code").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("支付渠道"),
		field.Uint64("pay_refund_id").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("支付退款单号"),
		field.Int32("refund_price").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("退款金额，单位：分"),
		field.Time("refund_time").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("退款完成时间"),
	}
}

func (DemoOrder) Edges() []ent.Edge {
	return nil
}

func (DemoOrder) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins2.SoftDeleteMixin{},
	}
}

func (DemoOrder) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "pay_demo_order"}}
}
