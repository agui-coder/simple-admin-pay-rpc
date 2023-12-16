// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// PayDemoOrderColumns holds the columns for the "pay_demo_order" table.
	PayDemoOrderColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Comment: "Create Time | 创建日期"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "Update Time | 修改日期"},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true, Comment: "Delete Time | 删除日期"},
		{Name: "user_id", Type: field.TypeString, Comment: "用户编号"},
		{Name: "spu_id", Type: field.TypeUint64, Comment: "商品编号"},
		{Name: "spu_name", Type: field.TypeString, Comment: "商品名称"},
		{Name: "price", Type: field.TypeInt32, Comment: "价格，单位：分"},
		{Name: "pay_status", Type: field.TypeBool, Comment: "是否支付"},
		{Name: "pay_order_id", Type: field.TypeUint64, Nullable: true, Comment: "支付订单编号"},
		{Name: "pay_time", Type: field.TypeTime, Nullable: true, Comment: "付款时间"},
		{Name: "pay_channel_code", Type: field.TypeString, Nullable: true, Comment: "支付渠道"},
		{Name: "pay_refund_id", Type: field.TypeUint64, Nullable: true, Comment: "支付退款单号"},
		{Name: "refund_price", Type: field.TypeInt32, Nullable: true, Comment: "退款金额，单位：分"},
		{Name: "refund_time", Type: field.TypeTime, Nullable: true, Comment: "退款完成时间"},
	}
	// PayDemoOrderTable holds the schema information for the "pay_demo_order" table.
	PayDemoOrderTable = &schema.Table{
		Name:       "pay_demo_order",
		Columns:    PayDemoOrderColumns,
		PrimaryKey: []*schema.Column{PayDemoOrderColumns[0]},
	}
	// PayOrderColumns holds the columns for the "pay_order" table.
	PayOrderColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Comment: "Create Time | 创建日期"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "Update Time | 修改日期"},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Comment: "Status 1: normal 2: ban | 状态 1 正常 2 禁用", Default: 1},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true, Comment: "Delete Time | 删除日期"},
		{Name: "channel_code", Type: field.TypeString, Nullable: true, Comment: "渠道编码"},
		{Name: "merchant_order_id", Type: field.TypeString, Comment: "商户订单编号"},
		{Name: "subject", Type: field.TypeString, Comment: "商品标题"},
		{Name: "body", Type: field.TypeString, Comment: "商品描述"},
		{Name: "price", Type: field.TypeInt32, Comment: "支付金额，单位：分"},
		{Name: "channel_fee_rate", Type: field.TypeFloat64, Nullable: true, Comment: "渠道手续费，单位：百分比"},
		{Name: "channel_fee_price", Type: field.TypeInt32, Nullable: true, Comment: "渠道手续金额，单位：分"},
		{Name: "user_ip", Type: field.TypeString, Comment: "用户 IP"},
		{Name: "expire_time", Type: field.TypeTime, Comment: "订单失效时间"},
		{Name: "success_time", Type: field.TypeTime, Nullable: true, Comment: "订单支付成功时间"},
		{Name: "notify_time", Type: field.TypeTime, Nullable: true, Comment: "订单支付通知时间"},
		{Name: "extension_id", Type: field.TypeUint64, Nullable: true, Comment: "支付成功的订单拓展单编号"},
		{Name: "no", Type: field.TypeString, Nullable: true, Comment: "订单号"},
		{Name: "refund_price", Type: field.TypeInt32, Comment: "退款总金额，单位：分"},
		{Name: "channel_user_id", Type: field.TypeString, Nullable: true, Comment: "渠道用户编号"},
		{Name: "channel_order_no", Type: field.TypeString, Nullable: true, Comment: "渠道订单号"},
	}
	// PayOrderTable holds the schema information for the "pay_order" table.
	PayOrderTable = &schema.Table{
		Name:       "pay_order",
		Columns:    PayOrderColumns,
		PrimaryKey: []*schema.Column{PayOrderColumns[0]},
	}
	// PayOrderExtensionColumns holds the columns for the "pay_order_extension" table.
	PayOrderExtensionColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Comment: "Create Time | 创建日期"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "Update Time | 修改日期"},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Comment: "Status 1: normal 2: ban | 状态 1 正常 2 禁用", Default: 1},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true, Comment: "Delete Time | 删除日期"},
		{Name: "no", Type: field.TypeString, Comment: "支付订单号"},
		{Name: "order_id", Type: field.TypeUint64, Comment: "渠道编号"},
		{Name: "channel_code", Type: field.TypeString, Comment: "渠道编码"},
		{Name: "user_ip", Type: field.TypeString, Comment: "用户 IP"},
		{Name: "channel_extras", Type: field.TypeJSON, Nullable: true, Comment: "支付渠道的额外参数"},
		{Name: "channel_error_code", Type: field.TypeString, Nullable: true, Comment: "调用渠道的错误码"},
		{Name: "channel_error_msg", Type: field.TypeString, Nullable: true, Comment: "调用渠道报错时，错误信息"},
		{Name: "channel_notify_data", Type: field.TypeString, Nullable: true, Size: 2147483647, Comment: "支付渠道异步通知的内容"},
	}
	// PayOrderExtensionTable holds the schema information for the "pay_order_extension" table.
	PayOrderExtensionTable = &schema.Table{
		Name:       "pay_order_extension",
		Columns:    PayOrderExtensionColumns,
		PrimaryKey: []*schema.Column{PayOrderExtensionColumns[0]},
	}
	// PayRefundColumns holds the columns for the "pay_refund" table.
	PayRefundColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "created_at", Type: field.TypeTime, Comment: "Create Time | 创建日期"},
		{Name: "updated_at", Type: field.TypeTime, Comment: "Update Time | 修改日期"},
		{Name: "status", Type: field.TypeUint8, Nullable: true, Comment: "Status 1: normal 2: ban | 状态 1 正常 2 禁用", Default: 1},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true, Comment: "Delete Time | 删除日期"},
		{Name: "no", Type: field.TypeString, Comment: "退款单号"},
		{Name: "channel_code", Type: field.TypeString, Comment: "渠道编码"},
		{Name: "order_id", Type: field.TypeUint64, Comment: "支付订单编号 pay_order 表id"},
		{Name: "order_no", Type: field.TypeString, Comment: "支付订单 no"},
		{Name: "merchant_order_id", Type: field.TypeString, Comment: "商户订单编号（商户系统生成）"},
		{Name: "merchant_refund_id", Type: field.TypeString, Comment: "商户退款订单号（商户系统生成）"},
		{Name: "pay_price", Type: field.TypeInt32, Comment: "支付金额,单位分"},
		{Name: "refund_price", Type: field.TypeInt32, Comment: "退款金额,单位分"},
		{Name: "reason", Type: field.TypeString, Comment: "退款原因"},
		{Name: "user_ip", Type: field.TypeString, Nullable: true, Comment: "用户 IP"},
		{Name: "channel_order_no", Type: field.TypeString, Comment: "渠道订单号，pay_order 中的 channel_order_no 对应"},
		{Name: "channel_refund_no", Type: field.TypeString, Nullable: true, Comment: "渠道退款单号，渠道返回"},
		{Name: "success_time", Type: field.TypeTime, Nullable: true, Comment: "退款成功时间"},
		{Name: "channel_error_code", Type: field.TypeString, Nullable: true, Comment: "渠道调用报错时，错误码"},
		{Name: "channel_error_msg", Type: field.TypeString, Nullable: true, Comment: "渠道调用报错时，错误信息"},
		{Name: "channel_notify_data", Type: field.TypeString, Nullable: true, Size: 2147483647, Comment: "支付渠道异步通知的内容"},
	}
	// PayRefundTable holds the schema information for the "pay_refund" table.
	PayRefundTable = &schema.Table{
		Name:       "pay_refund",
		Columns:    PayRefundColumns,
		PrimaryKey: []*schema.Column{PayRefundColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		PayDemoOrderTable,
		PayOrderTable,
		PayOrderExtensionTable,
		PayRefundTable,
	}
)

func init() {
	PayDemoOrderTable.Annotation = &entsql.Annotation{
		Table: "pay_demo_order",
	}
	PayOrderTable.Annotation = &entsql.Annotation{
		Table: "pay_order",
	}
	PayOrderExtensionTable.Annotation = &entsql.Annotation{
		Table: "pay_order_extension",
	}
	PayRefundTable.Annotation = &entsql.Annotation{
		Table: "pay_refund",
	}
}
