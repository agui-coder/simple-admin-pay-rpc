// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/orderextension"
)

// OrderExtension is the model entity for the OrderExtension schema.
type OrderExtension struct {
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
	// 支付订单号
	No string `json:"no,omitempty"`
	// 渠道编号
	OrderID uint64 `json:"order_id,omitempty"`
	// 渠道编码
	ChannelCode string `json:"channel_code,omitempty"`
	// 用户 IP
	UserIP string `json:"user_ip,omitempty"`
	// 支付渠道的额外参数
	ChannelExtras map[string]string `json:"channel_extras,omitempty"`
	// 调用渠道的错误码
	ChannelErrorCode string `json:"channel_error_code,omitempty"`
	// 调用渠道报错时，错误信息
	ChannelErrorMsg string `json:"channel_error_msg,omitempty"`
	// 支付渠道异步通知的内容
	ChannelNotifyData string `json:"channel_notify_data,omitempty"`
	selectValues      sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrderExtension) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case orderextension.FieldChannelExtras:
			values[i] = new([]byte)
		case orderextension.FieldID, orderextension.FieldStatus, orderextension.FieldOrderID:
			values[i] = new(sql.NullInt64)
		case orderextension.FieldNo, orderextension.FieldChannelCode, orderextension.FieldUserIP, orderextension.FieldChannelErrorCode, orderextension.FieldChannelErrorMsg, orderextension.FieldChannelNotifyData:
			values[i] = new(sql.NullString)
		case orderextension.FieldCreatedAt, orderextension.FieldUpdatedAt, orderextension.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrderExtension fields.
func (oe *OrderExtension) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case orderextension.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			oe.ID = uint64(value.Int64)
		case orderextension.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				oe.CreatedAt = value.Time
			}
		case orderextension.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				oe.UpdatedAt = value.Time
			}
		case orderextension.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				oe.Status = uint8(value.Int64)
			}
		case orderextension.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				oe.DeletedAt = value.Time
			}
		case orderextension.FieldNo:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field no", values[i])
			} else if value.Valid {
				oe.No = value.String
			}
		case orderextension.FieldOrderID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_id", values[i])
			} else if value.Valid {
				oe.OrderID = uint64(value.Int64)
			}
		case orderextension.FieldChannelCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field channel_code", values[i])
			} else if value.Valid {
				oe.ChannelCode = value.String
			}
		case orderextension.FieldUserIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_ip", values[i])
			} else if value.Valid {
				oe.UserIP = value.String
			}
		case orderextension.FieldChannelExtras:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field channel_extras", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &oe.ChannelExtras); err != nil {
					return fmt.Errorf("unmarshal field channel_extras: %w", err)
				}
			}
		case orderextension.FieldChannelErrorCode:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field channel_error_code", values[i])
			} else if value.Valid {
				oe.ChannelErrorCode = value.String
			}
		case orderextension.FieldChannelErrorMsg:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field channel_error_msg", values[i])
			} else if value.Valid {
				oe.ChannelErrorMsg = value.String
			}
		case orderextension.FieldChannelNotifyData:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field channel_notify_data", values[i])
			} else if value.Valid {
				oe.ChannelNotifyData = value.String
			}
		default:
			oe.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the OrderExtension.
// This includes values selected through modifiers, order, etc.
func (oe *OrderExtension) Value(name string) (ent.Value, error) {
	return oe.selectValues.Get(name)
}

// Update returns a builder for updating this OrderExtension.
// Note that you need to call OrderExtension.Unwrap() before calling this method if this OrderExtension
// was returned from a transaction, and the transaction was committed or rolled back.
func (oe *OrderExtension) Update() *OrderExtensionUpdateOne {
	return NewOrderExtensionClient(oe.config).UpdateOne(oe)
}

// Unwrap unwraps the OrderExtension entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (oe *OrderExtension) Unwrap() *OrderExtension {
	_tx, ok := oe.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrderExtension is not a transactional entity")
	}
	oe.config.driver = _tx.drv
	return oe
}

// String implements the fmt.Stringer.
func (oe *OrderExtension) String() string {
	var builder strings.Builder
	builder.WriteString("OrderExtension(")
	builder.WriteString(fmt.Sprintf("id=%v, ", oe.ID))
	builder.WriteString("created_at=")
	builder.WriteString(oe.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(oe.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", oe.Status))
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(oe.DeletedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("no=")
	builder.WriteString(oe.No)
	builder.WriteString(", ")
	builder.WriteString("order_id=")
	builder.WriteString(fmt.Sprintf("%v", oe.OrderID))
	builder.WriteString(", ")
	builder.WriteString("channel_code=")
	builder.WriteString(oe.ChannelCode)
	builder.WriteString(", ")
	builder.WriteString("user_ip=")
	builder.WriteString(oe.UserIP)
	builder.WriteString(", ")
	builder.WriteString("channel_extras=")
	builder.WriteString(fmt.Sprintf("%v", oe.ChannelExtras))
	builder.WriteString(", ")
	builder.WriteString("channel_error_code=")
	builder.WriteString(oe.ChannelErrorCode)
	builder.WriteString(", ")
	builder.WriteString("channel_error_msg=")
	builder.WriteString(oe.ChannelErrorMsg)
	builder.WriteString(", ")
	builder.WriteString("channel_notify_data=")
	builder.WriteString(oe.ChannelNotifyData)
	builder.WriteByte(')')
	return builder.String()
}

// OrderExtensions is a parsable slice of OrderExtension.
type OrderExtensions []*OrderExtension
