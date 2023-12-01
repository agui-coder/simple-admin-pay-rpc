package schema

import (
	mixins2 "github.com/agui-coder/simple-admin-pay-rpc/ent/schema/localmixin"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

type NotifyLog struct {
	ent.Schema
}

func (NotifyLog) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("task_id").
			Annotations(entsql.WithComments(true)).
			Comment("通知任务编号"),
		field.Int8("notify_times").
			Annotations(entsql.WithComments(true)).
			Comment("第几次被通知"),
		field.Text("response").
			Annotations(entsql.WithComments(true)).
			Comment("请求参数"),
	}
}

func (NotifyLog) Edges() []ent.Edge {
	return nil
}

func (NotifyLog) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.StatusMixin{},
		mixins2.SoftDeleteMixin{},
	}
}

func (NotifyLog) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "pay_notify_log"}}
}
