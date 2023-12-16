// Code generated by ent, DO NOT EDIT.

//go:build tools
// +build tools

// Package internal holds a loadable version of the latest schema.
package internal

const Schema = `{"Schema":"github.com/agui-coder/simple-admin-pay-rpc/ent/schema","Package":"github.com/agui-coder/simple-admin-pay-rpc/ent","Schemas":[{"name":"DemoOrder","config":{"Table":""},"fields":[{"name":"id","type":{"Type":18,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":true,"MixinIndex":0}},{"name":"created_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"immutable":true,"position":{"Index":1,"MixedIn":true,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"Create Time | 创建日期"},{"name":"updated_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"update_default":true,"position":{"Index":2,"MixedIn":true,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"Update Time | 修改日期"},{"name":"deleted_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":1},"annotations":{"EntSQL":{"with_comments":true}},"comment":"Delete Time | 删除日期"},{"name":"user_id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"用户编号"},{"name":"spu_id","type":{"Type":18,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"商品编号"},{"name":"spu_name","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":2,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"商品名称"},{"name":"price","type":{"Type":11,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":3,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"价格，单位：分"},{"name":"pay_status","type":{"Type":1,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":4,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"是否支付"},{"name":"pay_orderId","type":{"Type":18,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":5,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"支付订单编号"},{"name":"pay_time","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":6,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"付款时间"},{"name":"pay_channel_code","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":7,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"支付渠道"},{"name":"pay_refund_id","type":{"Type":18,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":8,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"支付退款单号"},{"name":"refund_price","type":{"Type":11,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":9,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"退款金额，单位：分"},{"name":"refund_time","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":10,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"退款完成时间"}],"hooks":[{"Index":0,"MixedIn":true,"MixinIndex":1}],"interceptors":[{"Index":0,"MixedIn":true,"MixinIndex":1}],"annotations":{"EntSQL":{"table":"pay_demo_order"}}},{"name":"Order","config":{"Table":""},"fields":[{"name":"id","type":{"Type":18,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":true,"MixinIndex":0}},{"name":"created_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"immutable":true,"position":{"Index":1,"MixedIn":true,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"Create Time | 创建日期"},{"name":"updated_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"update_default":true,"position":{"Index":2,"MixedIn":true,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"Update Time | 修改日期"},{"name":"status","type":{"Type":14,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"default":true,"default_value":1,"default_kind":8,"position":{"Index":0,"MixedIn":true,"MixinIndex":1},"annotations":{"EntSQL":{"with_comments":true}},"comment":"Status 1: normal 2: ban | 状态 1 正常 2 禁用"},{"name":"deleted_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":2},"annotations":{"EntSQL":{"with_comments":true}},"comment":"Delete Time | 删除日期"},{"name":"channel_code","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"渠道编码"},{"name":"merchant_order_id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"商户订单编号"},{"name":"subject","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":2,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"商品标题"},{"name":"body","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":3,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"商品描述"},{"name":"price","type":{"Type":11,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":4,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"支付金额，单位：分"},{"name":"channel_fee_rate","type":{"Type":20,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":5,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"渠道手续费，单位：百分比"},{"name":"channel_fee_price","type":{"Type":11,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":6,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"渠道手续金额，单位：分"},{"name":"user_ip","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":7,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"用户 IP"},{"name":"expire_time","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"position":{"Index":8,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"订单失效时间"},{"name":"success_time","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":9,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"订单支付成功时间"},{"name":"notify_time","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":10,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"订单支付通知时间"},{"name":"extension_id","type":{"Type":18,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":11,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"支付成功的订单拓展单编号"},{"name":"no","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":12,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"订单号"},{"name":"refund_price","type":{"Type":11,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":13,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"退款总金额，单位：分"},{"name":"channel_user_id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":14,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"渠道用户编号"},{"name":"channel_order_no","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":15,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"渠道订单号"}],"hooks":[{"Index":0,"MixedIn":true,"MixinIndex":2}],"interceptors":[{"Index":0,"MixedIn":true,"MixinIndex":2}],"annotations":{"EntSQL":{"table":"pay_order"}}},{"name":"OrderExtension","config":{"Table":""},"fields":[{"name":"id","type":{"Type":18,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":true,"MixinIndex":0}},{"name":"created_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"immutable":true,"position":{"Index":1,"MixedIn":true,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"Create Time | 创建日期"},{"name":"updated_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"update_default":true,"position":{"Index":2,"MixedIn":true,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"Update Time | 修改日期"},{"name":"status","type":{"Type":14,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"default":true,"default_value":1,"default_kind":8,"position":{"Index":0,"MixedIn":true,"MixinIndex":1},"annotations":{"EntSQL":{"with_comments":true}},"comment":"Status 1: normal 2: ban | 状态 1 正常 2 禁用"},{"name":"deleted_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":2},"annotations":{"EntSQL":{"with_comments":true}},"comment":"Delete Time | 删除日期"},{"name":"no","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"支付订单号"},{"name":"order_id","type":{"Type":18,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"渠道编号"},{"name":"channel_code","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":2,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"渠道编码"},{"name":"user_ip","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":3,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"用户 IP"},{"name":"channel_extras","type":{"Type":3,"Ident":"map[string]string","PkgPath":"","PkgName":"","Nillable":true,"RType":{"Name":"","Ident":"map[string]string","Kind":21,"PkgPath":"","Methods":{}}},"optional":true,"position":{"Index":4,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"支付渠道的额外参数"},{"name":"channel_error_code","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":5,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"调用渠道的错误码"},{"name":"channel_error_msg","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":6,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"调用渠道报错时，错误信息"},{"name":"channel_notify_data","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":2147483647,"optional":true,"position":{"Index":7,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"支付渠道异步通知的内容"}],"hooks":[{"Index":0,"MixedIn":true,"MixinIndex":2}],"interceptors":[{"Index":0,"MixedIn":true,"MixinIndex":2}],"annotations":{"EntSQL":{"table":"pay_order_extension"}}},{"name":"Refund","config":{"Table":""},"fields":[{"name":"id","type":{"Type":18,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":true,"MixinIndex":0}},{"name":"created_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"immutable":true,"position":{"Index":1,"MixedIn":true,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"Create Time | 创建日期"},{"name":"updated_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"default":true,"default_kind":19,"update_default":true,"position":{"Index":2,"MixedIn":true,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"Update Time | 修改日期"},{"name":"status","type":{"Type":14,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"default":true,"default_value":1,"default_kind":8,"position":{"Index":0,"MixedIn":true,"MixinIndex":1},"annotations":{"EntSQL":{"with_comments":true}},"comment":"Status 1: normal 2: ban | 状态 1 正常 2 禁用"},{"name":"deleted_at","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":0,"MixedIn":true,"MixinIndex":2},"annotations":{"EntSQL":{"with_comments":true}},"comment":"Delete Time | 删除日期"},{"name":"no","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":0,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"退款单号"},{"name":"channel_code","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":1,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"渠道编码"},{"name":"order_id","type":{"Type":18,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":2,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"支付订单编号 pay_order 表id"},{"name":"order_no","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":3,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"支付订单 no"},{"name":"merchant_order_id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":4,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"商户订单编号（商户系统生成）"},{"name":"merchant_refund_id","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":5,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"商户退款订单号（商户系统生成）"},{"name":"pay_price","type":{"Type":11,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":6,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"支付金额,单位分"},{"name":"refund_price","type":{"Type":11,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":7,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"退款金额,单位分"},{"name":"reason","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":8,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"退款原因"},{"name":"user_ip","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":9,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"用户 IP"},{"name":"channel_order_no","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"position":{"Index":10,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"渠道订单号，pay_order 中的 channel_order_no 对应"},{"name":"channel_refund_no","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":11,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"渠道退款单号，渠道返回"},{"name":"success_time","type":{"Type":2,"Ident":"","PkgPath":"time","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":12,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"退款成功时间"},{"name":"channel_error_code","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":13,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"渠道调用报错时，错误码"},{"name":"channel_error_msg","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"optional":true,"position":{"Index":14,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"渠道调用报错时，错误信息"},{"name":"channel_notify_data","type":{"Type":7,"Ident":"","PkgPath":"","PkgName":"","Nillable":false,"RType":null},"size":2147483647,"optional":true,"position":{"Index":15,"MixedIn":false,"MixinIndex":0},"annotations":{"EntSQL":{"with_comments":true}},"comment":"支付渠道异步通知的内容"}],"hooks":[{"Index":0,"MixedIn":true,"MixinIndex":2}],"interceptors":[{"Index":0,"MixedIn":true,"MixinIndex":2}],"annotations":{"EntSQL":{"table":"pay_refund"}}}],"Features":["intercept","schema/snapshot"]}`
