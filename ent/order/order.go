// Code generated by ent, DO NOT EDIT.

package order

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the order type in the database.
	Label = "order"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldDeletedAt holds the string denoting the deleted_at field in the database.
	FieldDeletedAt = "deleted_at"
	// FieldChannelCode holds the string denoting the channel_code field in the database.
	FieldChannelCode = "channel_code"
	// FieldMerchantOrderID holds the string denoting the merchant_order_id field in the database.
	FieldMerchantOrderID = "merchant_order_id"
	// FieldSubject holds the string denoting the subject field in the database.
	FieldSubject = "subject"
	// FieldBody holds the string denoting the body field in the database.
	FieldBody = "body"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldChannelFeeRate holds the string denoting the channel_fee_rate field in the database.
	FieldChannelFeeRate = "channel_fee_rate"
	// FieldChannelFeePrice holds the string denoting the channel_fee_price field in the database.
	FieldChannelFeePrice = "channel_fee_price"
	// FieldUserIP holds the string denoting the user_ip field in the database.
	FieldUserIP = "user_ip"
	// FieldExpireTime holds the string denoting the expire_time field in the database.
	FieldExpireTime = "expire_time"
	// FieldSuccessTime holds the string denoting the success_time field in the database.
	FieldSuccessTime = "success_time"
	// FieldNotifyTime holds the string denoting the notify_time field in the database.
	FieldNotifyTime = "notify_time"
	// FieldExtensionID holds the string denoting the extension_id field in the database.
	FieldExtensionID = "extension_id"
	// FieldNo holds the string denoting the no field in the database.
	FieldNo = "no"
	// FieldRefundPrice holds the string denoting the refund_price field in the database.
	FieldRefundPrice = "refund_price"
	// FieldChannelUserID holds the string denoting the channel_user_id field in the database.
	FieldChannelUserID = "channel_user_id"
	// FieldChannelOrderNo holds the string denoting the channel_order_no field in the database.
	FieldChannelOrderNo = "channel_order_no"
	// Table holds the table name of the order in the database.
	Table = "pay_order"
)

// Columns holds all SQL columns for order fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldStatus,
	FieldDeletedAt,
	FieldChannelCode,
	FieldMerchantOrderID,
	FieldSubject,
	FieldBody,
	FieldPrice,
	FieldChannelFeeRate,
	FieldChannelFeePrice,
	FieldUserIP,
	FieldExpireTime,
	FieldSuccessTime,
	FieldNotifyTime,
	FieldExtensionID,
	FieldNo,
	FieldRefundPrice,
	FieldChannelUserID,
	FieldChannelOrderNo,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/agui-coder/simple-admin-pay-rpc/ent/runtime"
var (
	Hooks        [1]ent.Hook
	Interceptors [1]ent.Interceptor
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultStatus holds the default value on creation for the "status" field.
	DefaultStatus uint8
)

// OrderOption defines the ordering options for the Order queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByStatus orders the results by the status field.
func ByStatus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldStatus, opts...).ToFunc()
}

// ByDeletedAt orders the results by the deleted_at field.
func ByDeletedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDeletedAt, opts...).ToFunc()
}

// ByChannelCode orders the results by the channel_code field.
func ByChannelCode(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldChannelCode, opts...).ToFunc()
}

// ByMerchantOrderID orders the results by the merchant_order_id field.
func ByMerchantOrderID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldMerchantOrderID, opts...).ToFunc()
}

// BySubject orders the results by the subject field.
func BySubject(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSubject, opts...).ToFunc()
}

// ByBody orders the results by the body field.
func ByBody(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBody, opts...).ToFunc()
}

// ByPrice orders the results by the price field.
func ByPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrice, opts...).ToFunc()
}

// ByChannelFeeRate orders the results by the channel_fee_rate field.
func ByChannelFeeRate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldChannelFeeRate, opts...).ToFunc()
}

// ByChannelFeePrice orders the results by the channel_fee_price field.
func ByChannelFeePrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldChannelFeePrice, opts...).ToFunc()
}

// ByUserIP orders the results by the user_ip field.
func ByUserIP(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUserIP, opts...).ToFunc()
}

// ByExpireTime orders the results by the expire_time field.
func ByExpireTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExpireTime, opts...).ToFunc()
}

// BySuccessTime orders the results by the success_time field.
func BySuccessTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSuccessTime, opts...).ToFunc()
}

// ByNotifyTime orders the results by the notify_time field.
func ByNotifyTime(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNotifyTime, opts...).ToFunc()
}

// ByExtensionID orders the results by the extension_id field.
func ByExtensionID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldExtensionID, opts...).ToFunc()
}

// ByNo orders the results by the no field.
func ByNo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNo, opts...).ToFunc()
}

// ByRefundPrice orders the results by the refund_price field.
func ByRefundPrice(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldRefundPrice, opts...).ToFunc()
}

// ByChannelUserID orders the results by the channel_user_id field.
func ByChannelUserID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldChannelUserID, opts...).ToFunc()
}

// ByChannelOrderNo orders the results by the channel_order_no field.
func ByChannelOrderNo(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldChannelOrderNo, opts...).ToFunc()
}
