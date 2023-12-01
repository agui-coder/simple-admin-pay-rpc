package schema

import (
	mixins2 "github.com/agui-coder/simple-admin-pay-rpc/ent/schema/localmixin"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

type Channel struct {
	ent.Schema
}

func (Channel) Fields() []ent.Field {
	return []ent.Field{
		field.String("code").
			Annotations(entsql.WithComments(true)).
			Comment("渠道编码"),
		field.String("remark").Optional().
			Annotations(entsql.WithComments(true)).
			Comment("应用名"),
		field.Float("fee_rate").
			Annotations(entsql.WithComments(true)).
			Comment("渠道费率，单位：百分比"),
		field.Uint64("app_id").
			Annotations(entsql.WithComments(true)).
			Comment("应用编号"),
		field.Text("config").
			Annotations(entsql.WithComments(true)).
			Comment("支付渠道配置"),
	}
}

func (Channel) Edges() []ent.Edge {
	return nil
}

func (Channel) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.StatusMixin{},
		mixins2.SoftDeleteMixin{},
	}
}

func (Channel) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "pay_channel"}}
}
