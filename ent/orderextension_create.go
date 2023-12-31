// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/orderextension"
)

// OrderExtensionCreate is the builder for creating a OrderExtension entity.
type OrderExtensionCreate struct {
	config
	mutation *OrderExtensionMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (oec *OrderExtensionCreate) SetCreatedAt(t time.Time) *OrderExtensionCreate {
	oec.mutation.SetCreatedAt(t)
	return oec
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (oec *OrderExtensionCreate) SetNillableCreatedAt(t *time.Time) *OrderExtensionCreate {
	if t != nil {
		oec.SetCreatedAt(*t)
	}
	return oec
}

// SetUpdatedAt sets the "updated_at" field.
func (oec *OrderExtensionCreate) SetUpdatedAt(t time.Time) *OrderExtensionCreate {
	oec.mutation.SetUpdatedAt(t)
	return oec
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (oec *OrderExtensionCreate) SetNillableUpdatedAt(t *time.Time) *OrderExtensionCreate {
	if t != nil {
		oec.SetUpdatedAt(*t)
	}
	return oec
}

// SetStatus sets the "status" field.
func (oec *OrderExtensionCreate) SetStatus(u uint8) *OrderExtensionCreate {
	oec.mutation.SetStatus(u)
	return oec
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (oec *OrderExtensionCreate) SetNillableStatus(u *uint8) *OrderExtensionCreate {
	if u != nil {
		oec.SetStatus(*u)
	}
	return oec
}

// SetDeletedAt sets the "deleted_at" field.
func (oec *OrderExtensionCreate) SetDeletedAt(t time.Time) *OrderExtensionCreate {
	oec.mutation.SetDeletedAt(t)
	return oec
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (oec *OrderExtensionCreate) SetNillableDeletedAt(t *time.Time) *OrderExtensionCreate {
	if t != nil {
		oec.SetDeletedAt(*t)
	}
	return oec
}

// SetNo sets the "no" field.
func (oec *OrderExtensionCreate) SetNo(s string) *OrderExtensionCreate {
	oec.mutation.SetNo(s)
	return oec
}

// SetOrderID sets the "order_id" field.
func (oec *OrderExtensionCreate) SetOrderID(u uint64) *OrderExtensionCreate {
	oec.mutation.SetOrderID(u)
	return oec
}

// SetChannelCode sets the "channel_code" field.
func (oec *OrderExtensionCreate) SetChannelCode(s string) *OrderExtensionCreate {
	oec.mutation.SetChannelCode(s)
	return oec
}

// SetUserIP sets the "user_ip" field.
func (oec *OrderExtensionCreate) SetUserIP(s string) *OrderExtensionCreate {
	oec.mutation.SetUserIP(s)
	return oec
}

// SetChannelExtras sets the "channel_extras" field.
func (oec *OrderExtensionCreate) SetChannelExtras(m map[string]string) *OrderExtensionCreate {
	oec.mutation.SetChannelExtras(m)
	return oec
}

// SetChannelErrorCode sets the "channel_error_code" field.
func (oec *OrderExtensionCreate) SetChannelErrorCode(s string) *OrderExtensionCreate {
	oec.mutation.SetChannelErrorCode(s)
	return oec
}

// SetNillableChannelErrorCode sets the "channel_error_code" field if the given value is not nil.
func (oec *OrderExtensionCreate) SetNillableChannelErrorCode(s *string) *OrderExtensionCreate {
	if s != nil {
		oec.SetChannelErrorCode(*s)
	}
	return oec
}

// SetChannelErrorMsg sets the "channel_error_msg" field.
func (oec *OrderExtensionCreate) SetChannelErrorMsg(s string) *OrderExtensionCreate {
	oec.mutation.SetChannelErrorMsg(s)
	return oec
}

// SetNillableChannelErrorMsg sets the "channel_error_msg" field if the given value is not nil.
func (oec *OrderExtensionCreate) SetNillableChannelErrorMsg(s *string) *OrderExtensionCreate {
	if s != nil {
		oec.SetChannelErrorMsg(*s)
	}
	return oec
}

// SetChannelNotifyData sets the "channel_notify_data" field.
func (oec *OrderExtensionCreate) SetChannelNotifyData(s string) *OrderExtensionCreate {
	oec.mutation.SetChannelNotifyData(s)
	return oec
}

// SetNillableChannelNotifyData sets the "channel_notify_data" field if the given value is not nil.
func (oec *OrderExtensionCreate) SetNillableChannelNotifyData(s *string) *OrderExtensionCreate {
	if s != nil {
		oec.SetChannelNotifyData(*s)
	}
	return oec
}

// SetID sets the "id" field.
func (oec *OrderExtensionCreate) SetID(u uint64) *OrderExtensionCreate {
	oec.mutation.SetID(u)
	return oec
}

// Mutation returns the OrderExtensionMutation object of the builder.
func (oec *OrderExtensionCreate) Mutation() *OrderExtensionMutation {
	return oec.mutation
}

// Save creates the OrderExtension in the database.
func (oec *OrderExtensionCreate) Save(ctx context.Context) (*OrderExtension, error) {
	if err := oec.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, oec.sqlSave, oec.mutation, oec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oec *OrderExtensionCreate) SaveX(ctx context.Context) *OrderExtension {
	v, err := oec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oec *OrderExtensionCreate) Exec(ctx context.Context) error {
	_, err := oec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oec *OrderExtensionCreate) ExecX(ctx context.Context) {
	if err := oec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oec *OrderExtensionCreate) defaults() error {
	if _, ok := oec.mutation.CreatedAt(); !ok {
		if orderextension.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized orderextension.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := orderextension.DefaultCreatedAt()
		oec.mutation.SetCreatedAt(v)
	}
	if _, ok := oec.mutation.UpdatedAt(); !ok {
		if orderextension.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized orderextension.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := orderextension.DefaultUpdatedAt()
		oec.mutation.SetUpdatedAt(v)
	}
	if _, ok := oec.mutation.Status(); !ok {
		v := orderextension.DefaultStatus
		oec.mutation.SetStatus(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (oec *OrderExtensionCreate) check() error {
	if _, ok := oec.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "OrderExtension.created_at"`)}
	}
	if _, ok := oec.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "OrderExtension.updated_at"`)}
	}
	if _, ok := oec.mutation.No(); !ok {
		return &ValidationError{Name: "no", err: errors.New(`ent: missing required field "OrderExtension.no"`)}
	}
	if _, ok := oec.mutation.OrderID(); !ok {
		return &ValidationError{Name: "order_id", err: errors.New(`ent: missing required field "OrderExtension.order_id"`)}
	}
	if _, ok := oec.mutation.ChannelCode(); !ok {
		return &ValidationError{Name: "channel_code", err: errors.New(`ent: missing required field "OrderExtension.channel_code"`)}
	}
	if _, ok := oec.mutation.UserIP(); !ok {
		return &ValidationError{Name: "user_ip", err: errors.New(`ent: missing required field "OrderExtension.user_ip"`)}
	}
	return nil
}

func (oec *OrderExtensionCreate) sqlSave(ctx context.Context) (*OrderExtension, error) {
	if err := oec.check(); err != nil {
		return nil, err
	}
	_node, _spec := oec.createSpec()
	if err := sqlgraph.CreateNode(ctx, oec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	oec.mutation.id = &_node.ID
	oec.mutation.done = true
	return _node, nil
}

func (oec *OrderExtensionCreate) createSpec() (*OrderExtension, *sqlgraph.CreateSpec) {
	var (
		_node = &OrderExtension{config: oec.config}
		_spec = sqlgraph.NewCreateSpec(orderextension.Table, sqlgraph.NewFieldSpec(orderextension.FieldID, field.TypeUint64))
	)
	if id, ok := oec.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := oec.mutation.CreatedAt(); ok {
		_spec.SetField(orderextension.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := oec.mutation.UpdatedAt(); ok {
		_spec.SetField(orderextension.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := oec.mutation.Status(); ok {
		_spec.SetField(orderextension.FieldStatus, field.TypeUint8, value)
		_node.Status = value
	}
	if value, ok := oec.mutation.DeletedAt(); ok {
		_spec.SetField(orderextension.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := oec.mutation.No(); ok {
		_spec.SetField(orderextension.FieldNo, field.TypeString, value)
		_node.No = value
	}
	if value, ok := oec.mutation.OrderID(); ok {
		_spec.SetField(orderextension.FieldOrderID, field.TypeUint64, value)
		_node.OrderID = value
	}
	if value, ok := oec.mutation.ChannelCode(); ok {
		_spec.SetField(orderextension.FieldChannelCode, field.TypeString, value)
		_node.ChannelCode = value
	}
	if value, ok := oec.mutation.UserIP(); ok {
		_spec.SetField(orderextension.FieldUserIP, field.TypeString, value)
		_node.UserIP = value
	}
	if value, ok := oec.mutation.ChannelExtras(); ok {
		_spec.SetField(orderextension.FieldChannelExtras, field.TypeJSON, value)
		_node.ChannelExtras = value
	}
	if value, ok := oec.mutation.ChannelErrorCode(); ok {
		_spec.SetField(orderextension.FieldChannelErrorCode, field.TypeString, value)
		_node.ChannelErrorCode = value
	}
	if value, ok := oec.mutation.ChannelErrorMsg(); ok {
		_spec.SetField(orderextension.FieldChannelErrorMsg, field.TypeString, value)
		_node.ChannelErrorMsg = value
	}
	if value, ok := oec.mutation.ChannelNotifyData(); ok {
		_spec.SetField(orderextension.FieldChannelNotifyData, field.TypeString, value)
		_node.ChannelNotifyData = value
	}
	return _node, _spec
}

// OrderExtensionCreateBulk is the builder for creating many OrderExtension entities in bulk.
type OrderExtensionCreateBulk struct {
	config
	err      error
	builders []*OrderExtensionCreate
}

// Save creates the OrderExtension entities in the database.
func (oecb *OrderExtensionCreateBulk) Save(ctx context.Context) ([]*OrderExtension, error) {
	if oecb.err != nil {
		return nil, oecb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(oecb.builders))
	nodes := make([]*OrderExtension, len(oecb.builders))
	mutators := make([]Mutator, len(oecb.builders))
	for i := range oecb.builders {
		func(i int, root context.Context) {
			builder := oecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrderExtensionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, oecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, oecb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = uint64(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, oecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (oecb *OrderExtensionCreateBulk) SaveX(ctx context.Context) []*OrderExtension {
	v, err := oecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oecb *OrderExtensionCreateBulk) Exec(ctx context.Context) error {
	_, err := oecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oecb *OrderExtensionCreateBulk) ExecX(ctx context.Context) {
	if err := oecb.Exec(ctx); err != nil {
		panic(err)
	}
}
