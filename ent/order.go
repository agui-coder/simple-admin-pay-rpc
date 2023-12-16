// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/order"
)

// Order is the model entity for the Order schema.
type Order struct {
	config `json:"-"`
	// ID of the ent.
	ID uint64 `json:"id,omitempty"`
	// Create Time | 创建日期
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Update Time | 修改日期
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Status 1: normal 2: ban | 状态 1 正常 2 禁用
	Status uint8 `json:"status,omitempty"`
	// Delete Time | 删除日期
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// 渠道编码
	ChannelCode string `json:"channel_code,omitempty"`
	// 商户订单编号
	MerchantOrderID string `json:"merchant_order_id,omitempty"`
	// 商品标题
	Subject string `json:"subject,omitempty"`
	// 商品描述
	Body string `json:"body,omitempty"`
	// 支付金额，单位：分
	Price int32 `json:"price,omitempty"`
	// 渠道手续费，单位：百分比
	ChannelFeeRate float64 `json:"channel_fee_rate,omitempty"`
	// 渠道手续金额，单位：分
	ChannelFeePrice int32 `json:"channel_fee_price,omitempty"`
	// 用户 IP
	UserIP string `json:"user_ip,omitempty"`
	// 订单失效时间
	ExpireTime time.Time `json:"expire_time,omitempty"`
	// 订单支付成功时间
	SuccessTime time.Time `json:"success_time,omitempty"`
	// 订单支付通知时间
	NotifyTime time.Time `json:"notify_time,omitempty"`
	// 支付成功的订单拓展单编号
	ExtensionID uint64 `json:"extension_id,omitempty"`
	// 订单号
	No string `json:"no,omitempty"`
	// 退款总金额，单位：分
	RefundPrice int32 `json:"refund_price,omitempty"`
	// 渠道用户编号
	ChannelUserID string `json:"channel_user_id,omitempty"`
	// 渠道订单号
	ChannelOrderNo string `json:"channel_order_no,omitempty"`
	selectValues   sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Order) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case order.FieldChannelFeeRate:
			values[i] = new(sql.NullFloat64)
		case order.FieldID, order.FieldStatus, order.FieldPrice, order.FieldChannelFeePrice, order.FieldExtensionID, order.FieldRefundPrice:
			values[i] = new(sql.NullInt64)
		case order.FieldChannelCode, order.FieldMerchantOrderID, order.FieldSubject, order.FieldBody, order.FieldUserIP, order.FieldNo, order.FieldChannelUserID, order.FieldChannelOrderNo:
			values[i] = new(sql.NullString)
		case order.FieldCreatedAt, order.FieldUpdatedAt, order.FieldDeletedAt, order.FieldExpireTime, order.FieldSuccessTime, order.FieldNotifyTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Order fields.
func (o *Order) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case order.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			o.ID = uint64(value.Int64)
		case order.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				o.CreatedAt = value.Time
			}
		case order.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				o.UpdatedAt = value.Time
			}
		case order.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				o.Status = uint8(value.Int64)
			}
		case order.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				o.DeletedAt = value.Time
			}
		case order.FieldChannelCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field channel_code", values[i])
			} else if value.Valid {
				o.ChannelCode = value.String
			}
		case order.FieldMerchantOrderID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field merchant_order_id", values[i])
			} else if value.Valid {
				o.MerchantOrderID = value.String
			}
		case order.FieldSubject:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field subject", values[i])
			} else if value.Valid {
				o.Subject = value.String
			}
		case order.FieldBody:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field body", values[i])
			} else if value.Valid {
				o.Body = value.String
			}
		case order.FieldPrice:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[i])
			} else if value.Valid {
				o.Price = int32(value.Int64)
			}
		case order.FieldChannelFeeRate:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field channel_fee_rate", values[i])
			} else if value.Valid {
				o.ChannelFeeRate = value.Float64
			}
		case order.FieldChannelFeePrice:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field channel_fee_price", values[i])
			} else if value.Valid {
				o.ChannelFeePrice = int32(value.Int64)
			}
		case order.FieldUserIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_ip", values[i])
			} else if value.Valid {
				o.UserIP = value.String
			}
		case order.FieldExpireTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field expire_time", values[i])
			} else if value.Valid {
				o.ExpireTime = value.Time
			}
		case order.FieldSuccessTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field success_time", values[i])
			} else if value.Valid {
				o.SuccessTime = value.Time
			}
		case order.FieldNotifyTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field notify_time", values[i])
			} else if value.Valid {
				o.NotifyTime = value.Time
			}
		case order.FieldExtensionID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field extension_id", values[i])
			} else if value.Valid {
				o.ExtensionID = uint64(value.Int64)
			}
		case order.FieldNo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field no", values[i])
			} else if value.Valid {
				o.No = value.String
			}
		case order.FieldRefundPrice:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field refund_price", values[i])
			} else if value.Valid {
				o.RefundPrice = int32(value.Int64)
			}
		case order.FieldChannelUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field channel_user_id", values[i])
			} else if value.Valid {
				o.ChannelUserID = value.String
			}
		case order.FieldChannelOrderNo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field channel_order_no", values[i])
			} else if value.Valid {
				o.ChannelOrderNo = value.String
			}
		default:
			o.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Order.
// This includes values selected through modifiers, order, etc.
func (o *Order) Value(name string) (ent.Value, error) {
	return o.selectValues.Get(name)
}

// Update returns a builder for updating this Order.
// Note that you need to call Order.Unwrap() before calling this method if this Order
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Order) Update() *OrderUpdateOne {
	return NewOrderClient(o.config).UpdateOne(o)
}

// Unwrap unwraps the Order entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Order) Unwrap() *Order {
	_tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Order is not a transactional entity")
	}
	o.config.driver = _tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Order) String() string {
	var builder strings.Builder
	builder.WriteString("Order(")
	builder.WriteString(fmt.Sprintf("id=%v, ", o.ID))
	builder.WriteString("created_at=")
	builder.WriteString(o.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(o.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", o.Status))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(o.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("channel_code=")
	builder.WriteString(o.ChannelCode)
	builder.WriteString(", ")
	builder.WriteString("merchant_order_id=")
	builder.WriteString(o.MerchantOrderID)
	builder.WriteString(", ")
	builder.WriteString("subject=")
	builder.WriteString(o.Subject)
	builder.WriteString(", ")
	builder.WriteString("body=")
	builder.WriteString(o.Body)
	builder.WriteString(", ")
	builder.WriteString("price=")
	builder.WriteString(fmt.Sprintf("%v", o.Price))
	builder.WriteString(", ")
	builder.WriteString("channel_fee_rate=")
	builder.WriteString(fmt.Sprintf("%v", o.ChannelFeeRate))
	builder.WriteString(", ")
	builder.WriteString("channel_fee_price=")
	builder.WriteString(fmt.Sprintf("%v", o.ChannelFeePrice))
	builder.WriteString(", ")
	builder.WriteString("user_ip=")
	builder.WriteString(o.UserIP)
	builder.WriteString(", ")
	builder.WriteString("expire_time=")
	builder.WriteString(o.ExpireTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("success_time=")
	builder.WriteString(o.SuccessTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("notify_time=")
	builder.WriteString(o.NotifyTime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("extension_id=")
	builder.WriteString(fmt.Sprintf("%v", o.ExtensionID))
	builder.WriteString(", ")
	builder.WriteString("no=")
	builder.WriteString(o.No)
	builder.WriteString(", ")
	builder.WriteString("refund_price=")
	builder.WriteString(fmt.Sprintf("%v", o.RefundPrice))
	builder.WriteString(", ")
	builder.WriteString("channel_user_id=")
	builder.WriteString(o.ChannelUserID)
	builder.WriteString(", ")
	builder.WriteString("channel_order_no=")
	builder.WriteString(o.ChannelOrderNo)
	builder.WriteByte(')')
	return builder.String()
}

// Orders is a parsable slice of Order.
type Orders []*Order
