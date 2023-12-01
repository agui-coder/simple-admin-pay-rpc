package schema

import (
	mixins2 "github.com/agui-coder/simple-admin-pay-rpc/ent/schema/localmixin"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/suyuan32/simple-admin-common/orm/ent/mixins"
)

type NotifyTask struct {
	ent.Schema
}

func (NotifyTask) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").
			Annotations(entsql.WithComments(true)).
			Comment("任务编号"),
		field.Uint64("app_id").
			Annotations(entsql.WithComments(true)).
			Comment("应用编号"),
		field.Int("type").
			Annotations(entsql.WithComments(true)).
			Comment("通知类型"),
		field.Uint64("data_id").
			Annotations(entsql.WithComments(true)).
			Comment("数据编号"),
		field.String("merchant_order_id").
			Annotations(entsql.WithComments(true)).
			Comment("商户订单编号"),
		field.Time("next_notify_time").
			Annotations(entsql.WithComments(true)).
			Comment("下一次通知时间"),
		field.Time("last_execute_time").
			Annotations(entsql.WithComments(true)).
			Comment("最后一次执行时间"),
		field.Int8("notify_times").
			Annotations(entsql.WithComments(true)).
			Comment("当前通知次数"),
		field.Int8("max_notify_times").
			Annotations(entsql.WithComments(true)).
			Comment("最大可通知次数"),
		field.String("notify_url").
			Annotations(entsql.WithComments(true)).
			Comment("异步通知地址")}
}
func (NotifyTask) Edges() []ent.Edge {
	return nil
}

func (NotifyTask) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.IDMixin{},
		mixins.StatusMixin{},
		mixins2.SoftDeleteMixin{},
	}
}

func (NotifyTask) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "pay_notify_task"}}
}
