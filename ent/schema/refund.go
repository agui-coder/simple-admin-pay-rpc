package schema

import (
	mixins2 "github.com/agui-coder/simple-admin-pay-rpc/ent/schema/localmixin"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

type Refund struct {
	ent.Schema
}

func (Refund) Fields() []ent.Field {
	return []ent.Field{
		field.String("no").
			Annotations(entsql.WithComments(true)).
			Comment("退款单号"),
		field.String("channel_code").
			Annotations(entsql.WithComments(true)).
			Comment("渠道编码"),
		field.Uint64("order_id").
			Annotations(entsql.WithComments(true)).
			Comment("支付订单编号 pay_order 表id"),
		field.String("order_no").
			Annotations(entsql.WithComments(true)).
			Comment("支付订单 no"),
		field.String("merchant_order_id").
			Annotations(entsql.WithComments(true)).
			Comment("商户订单编号（商户系统生成）"),
		field.String("merchant_refund_id").
			Annotations(entsql.WithComments(true)).
			Comment("商户退款订单号（商户系统生成）"),
		field.Int32("pay_price").
			Annotations(entsql.WithComments(true)).
			Comment("支付金额,单位分"),
		field.Int32("refund_price").
			Annotations(entsql.WithComments(true)).
			Comment("退款金额,单位分"),
		field.String("reason").
			Annotations(entsql.WithComments(true)).
			Comment("退款原因"),
		field.String("user_ip").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("用户 IP"),
		field.String("channel_order_no").
			Annotations(entsql.WithComments(true)).
			Comment("渠道订单号，pay_order 中的 channel_order_no 对应"),
		field.String("channel_refund_no").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("渠道退款单号，渠道返回"),
		field.Time("success_time").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("退款成功时间"),
		field.String("channel_error_code").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("渠道调用报错时，错误码"),
		field.String("channel_error_msg").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("渠道调用报错时，错误信息"),
		field.Text("channel_notify_data").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("支付渠道异步通知的内容")}
}
func (Refund) Edges() []ent.Edge {
	return nil
}

func (Refund) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.StatusMixin{},
		mixins2.SoftDeleteMixin{},
	}
}

func (Refund) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "pay_refund"}}
}
