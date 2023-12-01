// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"github.com/agui-coder/simple-admin-pay-rpc/ent/app"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/channel"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/demoorder"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/notifylog"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/notifytask"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/order"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/orderextension"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/refund"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/schema"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	appMixin := schema.App{}.Mixin()
	appMixinHooks2 := appMixin[2].Hooks()
	app.Hooks[0] = appMixinHooks2[0]
	appMixinInters2 := appMixin[2].Interceptors()
	app.Interceptors[0] = appMixinInters2[0]
	appMixinFields0 := appMixin[0].Fields()
	_ = appMixinFields0
	appMixinFields1 := appMixin[1].Fields()
	_ = appMixinFields1
	appFields := schema.App{}.Fields()
	_ = appFields
	// appDescCreatedAt is the schema descriptor for created_at field.
	appDescCreatedAt := appMixinFields0[1].Descriptor()
	// app.DefaultCreatedAt holds the default value on creation for the created_at field.
	app.DefaultCreatedAt = appDescCreatedAt.Default.(func() time.Time)
	// appDescUpdatedAt is the schema descriptor for updated_at field.
	appDescUpdatedAt := appMixinFields0[2].Descriptor()
	// app.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	app.DefaultUpdatedAt = appDescUpdatedAt.Default.(func() time.Time)
	// app.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	app.UpdateDefaultUpdatedAt = appDescUpdatedAt.UpdateDefault.(func() time.Time)
	// appDescStatus is the schema descriptor for status field.
	appDescStatus := appMixinFields1[0].Descriptor()
	// app.DefaultStatus holds the default value on creation for the status field.
	app.DefaultStatus = appDescStatus.Default.(uint8)
	channelMixin := schema.Channel{}.Mixin()
	channelMixinHooks2 := channelMixin[2].Hooks()
	channel.Hooks[0] = channelMixinHooks2[0]
	channelMixinInters2 := channelMixin[2].Interceptors()
	channel.Interceptors[0] = channelMixinInters2[0]
	channelMixinFields0 := channelMixin[0].Fields()
	_ = channelMixinFields0
	channelMixinFields1 := channelMixin[1].Fields()
	_ = channelMixinFields1
	channelFields := schema.Channel{}.Fields()
	_ = channelFields
	// channelDescCreatedAt is the schema descriptor for created_at field.
	channelDescCreatedAt := channelMixinFields0[1].Descriptor()
	// channel.DefaultCreatedAt holds the default value on creation for the created_at field.
	channel.DefaultCreatedAt = channelDescCreatedAt.Default.(func() time.Time)
	// channelDescUpdatedAt is the schema descriptor for updated_at field.
	channelDescUpdatedAt := channelMixinFields0[2].Descriptor()
	// channel.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	channel.DefaultUpdatedAt = channelDescUpdatedAt.Default.(func() time.Time)
	// channel.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	channel.UpdateDefaultUpdatedAt = channelDescUpdatedAt.UpdateDefault.(func() time.Time)
	// channelDescStatus is the schema descriptor for status field.
	channelDescStatus := channelMixinFields1[0].Descriptor()
	// channel.DefaultStatus holds the default value on creation for the status field.
	channel.DefaultStatus = channelDescStatus.Default.(uint8)
	demoorderMixin := schema.DemoOrder{}.Mixin()
	demoorderMixinHooks1 := demoorderMixin[1].Hooks()
	demoorder.Hooks[0] = demoorderMixinHooks1[0]
	demoorderMixinInters1 := demoorderMixin[1].Interceptors()
	demoorder.Interceptors[0] = demoorderMixinInters1[0]
	demoorderMixinFields0 := demoorderMixin[0].Fields()
	_ = demoorderMixinFields0
	demoorderFields := schema.DemoOrder{}.Fields()
	_ = demoorderFields
	// demoorderDescCreatedAt is the schema descriptor for created_at field.
	demoorderDescCreatedAt := demoorderMixinFields0[1].Descriptor()
	// demoorder.DefaultCreatedAt holds the default value on creation for the created_at field.
	demoorder.DefaultCreatedAt = demoorderDescCreatedAt.Default.(func() time.Time)
	// demoorderDescUpdatedAt is the schema descriptor for updated_at field.
	demoorderDescUpdatedAt := demoorderMixinFields0[2].Descriptor()
	// demoorder.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	demoorder.DefaultUpdatedAt = demoorderDescUpdatedAt.Default.(func() time.Time)
	// demoorder.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	demoorder.UpdateDefaultUpdatedAt = demoorderDescUpdatedAt.UpdateDefault.(func() time.Time)
	notifylogMixin := schema.NotifyLog{}.Mixin()
	notifylogMixinHooks2 := notifylogMixin[2].Hooks()
	notifylog.Hooks[0] = notifylogMixinHooks2[0]
	notifylogMixinInters2 := notifylogMixin[2].Interceptors()
	notifylog.Interceptors[0] = notifylogMixinInters2[0]
	notifylogMixinFields0 := notifylogMixin[0].Fields()
	_ = notifylogMixinFields0
	notifylogMixinFields1 := notifylogMixin[1].Fields()
	_ = notifylogMixinFields1
	notifylogFields := schema.NotifyLog{}.Fields()
	_ = notifylogFields
	// notifylogDescCreatedAt is the schema descriptor for created_at field.
	notifylogDescCreatedAt := notifylogMixinFields0[1].Descriptor()
	// notifylog.DefaultCreatedAt holds the default value on creation for the created_at field.
	notifylog.DefaultCreatedAt = notifylogDescCreatedAt.Default.(func() time.Time)
	// notifylogDescUpdatedAt is the schema descriptor for updated_at field.
	notifylogDescUpdatedAt := notifylogMixinFields0[2].Descriptor()
	// notifylog.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	notifylog.DefaultUpdatedAt = notifylogDescUpdatedAt.Default.(func() time.Time)
	// notifylog.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	notifylog.UpdateDefaultUpdatedAt = notifylogDescUpdatedAt.UpdateDefault.(func() time.Time)
	// notifylogDescStatus is the schema descriptor for status field.
	notifylogDescStatus := notifylogMixinFields1[0].Descriptor()
	// notifylog.DefaultStatus holds the default value on creation for the status field.
	notifylog.DefaultStatus = notifylogDescStatus.Default.(uint8)
	notifytaskMixin := schema.NotifyTask{}.Mixin()
	notifytaskMixinHooks2 := notifytaskMixin[2].Hooks()
	notifytask.Hooks[0] = notifytaskMixinHooks2[0]
	notifytaskMixinInters2 := notifytaskMixin[2].Interceptors()
	notifytask.Interceptors[0] = notifytaskMixinInters2[0]
	notifytaskMixinFields0 := notifytaskMixin[0].Fields()
	_ = notifytaskMixinFields0
	notifytaskMixinFields1 := notifytaskMixin[1].Fields()
	_ = notifytaskMixinFields1
	notifytaskFields := schema.NotifyTask{}.Fields()
	_ = notifytaskFields
	// notifytaskDescCreatedAt is the schema descriptor for created_at field.
	notifytaskDescCreatedAt := notifytaskMixinFields0[1].Descriptor()
	// notifytask.DefaultCreatedAt holds the default value on creation for the created_at field.
	notifytask.DefaultCreatedAt = notifytaskDescCreatedAt.Default.(func() time.Time)
	// notifytaskDescUpdatedAt is the schema descriptor for updated_at field.
	notifytaskDescUpdatedAt := notifytaskMixinFields0[2].Descriptor()
	// notifytask.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	notifytask.DefaultUpdatedAt = notifytaskDescUpdatedAt.Default.(func() time.Time)
	// notifytask.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	notifytask.UpdateDefaultUpdatedAt = notifytaskDescUpdatedAt.UpdateDefault.(func() time.Time)
	// notifytaskDescStatus is the schema descriptor for status field.
	notifytaskDescStatus := notifytaskMixinFields1[0].Descriptor()
	// notifytask.DefaultStatus holds the default value on creation for the status field.
	notifytask.DefaultStatus = notifytaskDescStatus.Default.(uint8)
	orderMixin := schema.Order{}.Mixin()
	orderMixinHooks2 := orderMixin[2].Hooks()
	order.Hooks[0] = orderMixinHooks2[0]
	orderMixinInters2 := orderMixin[2].Interceptors()
	order.Interceptors[0] = orderMixinInters2[0]
	orderMixinFields0 := orderMixin[0].Fields()
	_ = orderMixinFields0
	orderMixinFields1 := orderMixin[1].Fields()
	_ = orderMixinFields1
	orderFields := schema.Order{}.Fields()
	_ = orderFields
	// orderDescCreatedAt is the schema descriptor for created_at field.
	orderDescCreatedAt := orderMixinFields0[1].Descriptor()
	// order.DefaultCreatedAt holds the default value on creation for the created_at field.
	order.DefaultCreatedAt = orderDescCreatedAt.Default.(func() time.Time)
	// orderDescUpdatedAt is the schema descriptor for updated_at field.
	orderDescUpdatedAt := orderMixinFields0[2].Descriptor()
	// order.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	order.DefaultUpdatedAt = orderDescUpdatedAt.Default.(func() time.Time)
	// order.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	order.UpdateDefaultUpdatedAt = orderDescUpdatedAt.UpdateDefault.(func() time.Time)
	// orderDescStatus is the schema descriptor for status field.
	orderDescStatus := orderMixinFields1[0].Descriptor()
	// order.DefaultStatus holds the default value on creation for the status field.
	order.DefaultStatus = orderDescStatus.Default.(uint8)
	orderextensionMixin := schema.OrderExtension{}.Mixin()
	orderextensionMixinHooks2 := orderextensionMixin[2].Hooks()
	orderextension.Hooks[0] = orderextensionMixinHooks2[0]
	orderextensionMixinInters2 := orderextensionMixin[2].Interceptors()
	orderextension.Interceptors[0] = orderextensionMixinInters2[0]
	orderextensionMixinFields0 := orderextensionMixin[0].Fields()
	_ = orderextensionMixinFields0
	orderextensionMixinFields1 := orderextensionMixin[1].Fields()
	_ = orderextensionMixinFields1
	orderextensionFields := schema.OrderExtension{}.Fields()
	_ = orderextensionFields
	// orderextensionDescCreatedAt is the schema descriptor for created_at field.
	orderextensionDescCreatedAt := orderextensionMixinFields0[1].Descriptor()
	// orderextension.DefaultCreatedAt holds the default value on creation for the created_at field.
	orderextension.DefaultCreatedAt = orderextensionDescCreatedAt.Default.(func() time.Time)
	// orderextensionDescUpdatedAt is the schema descriptor for updated_at field.
	orderextensionDescUpdatedAt := orderextensionMixinFields0[2].Descriptor()
	// orderextension.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	orderextension.DefaultUpdatedAt = orderextensionDescUpdatedAt.Default.(func() time.Time)
	// orderextension.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	orderextension.UpdateDefaultUpdatedAt = orderextensionDescUpdatedAt.UpdateDefault.(func() time.Time)
	// orderextensionDescStatus is the schema descriptor for status field.
	orderextensionDescStatus := orderextensionMixinFields1[0].Descriptor()
	// orderextension.DefaultStatus holds the default value on creation for the status field.
	orderextension.DefaultStatus = orderextensionDescStatus.Default.(uint8)
	refundMixin := schema.Refund{}.Mixin()
	refundMixinHooks2 := refundMixin[2].Hooks()
	refund.Hooks[0] = refundMixinHooks2[0]
	refundMixinInters2 := refundMixin[2].Interceptors()
	refund.Interceptors[0] = refundMixinInters2[0]
	refundMixinFields0 := refundMixin[0].Fields()
	_ = refundMixinFields0
	refundMixinFields1 := refundMixin[1].Fields()
	_ = refundMixinFields1
	refundFields := schema.Refund{}.Fields()
	_ = refundFields
	// refundDescCreatedAt is the schema descriptor for created_at field.
	refundDescCreatedAt := refundMixinFields0[1].Descriptor()
	// refund.DefaultCreatedAt holds the default value on creation for the created_at field.
	refund.DefaultCreatedAt = refundDescCreatedAt.Default.(func() time.Time)
	// refundDescUpdatedAt is the schema descriptor for updated_at field.
	refundDescUpdatedAt := refundMixinFields0[2].Descriptor()
	// refund.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	refund.DefaultUpdatedAt = refundDescUpdatedAt.Default.(func() time.Time)
	// refund.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	refund.UpdateDefaultUpdatedAt = refundDescUpdatedAt.UpdateDefault.(func() time.Time)
	// refundDescStatus is the schema descriptor for status field.
	refundDescStatus := refundMixinFields1[0].Descriptor()
	// refund.DefaultStatus holds the default value on creation for the status field.
	refund.DefaultStatus = refundDescStatus.Default.(uint8)
}

const (
	Version = "v0.12.4"                                         // Version of ent codegen.
	Sum     = "h1:LddPnAyxls/O7DTXZvUGDj0NZIdGSu317+aoNLJWbD8=" // Sum of ent codegen.
)
