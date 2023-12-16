package schema

import (
	mixins2 "github.com/agui-coder/simple-admin-pay-rpc/ent/schema/localmixin"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

type OrderExtension struct {
	ent.Schema
}

func (OrderExtension) Fields() []ent.Field {
	return []ent.Field{
		field.String("no").
			Annotations(entsql.WithComments(true)).
			Comment("支付订单号"),
		field.Uint64("order_id").
			Annotations(entsql.WithComments(true)).
			Comment("渠道编号"),
		field.String("channel_code").
			Annotations(entsql.WithComments(true)).
			Comment("渠道编码"),
		field.String("user_ip").
			Annotations(entsql.WithComments(true)).
			Comment("用户 IP"),
		field.JSON("channel_extras", map[string]string{}).Optional().
			Annotations(entsql.WithComments(true)).
			Comment("支付渠道的额外参数"),
		field.String("channel_error_code").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("调用渠道的错误码"),
		field.String("channel_error_msg").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("调用渠道报错时，错误信息"),
		field.Text("channel_notify_data").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("支付渠道异步通知的内容"),
	}
}
func (OrderExtension) Edges() []ent.Edge {
	return nil
}

func (OrderExtension) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.StatusMixin{},
		mixins2.SoftDeleteMixin{},
	}
}

func (OrderExtension) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "pay_order_extension"}}
}
