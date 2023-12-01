package schema

import (
	mixins2 "github.com/agui-coder/simple-admin-pay-rpc/ent/schema/localmixin"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

type App struct {
	ent.Schema
}

func (App) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Annotations(entsql.WithComments(true)).
			Comment("应用名"),
		field.String("remark").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("应用名"),
		field.String("order_notify_url").
			Annotations(entsql.WithComments(true)).
			Comment("支付结果的回调地址"),
		field.String("refund_notify_url").
			Annotations(entsql.WithComments(true)).
			Comment("退款结果的回调地址"),
	}
}

func (App) Edges() []ent.Edge {
	return nil
}

func (App) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.StatusMixin{},
		mixins2.SoftDeleteMixin{},
	}
}

func (App) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "pay_app"}}
}
