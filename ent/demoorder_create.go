// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/demoorder"
)

// DemoOrderCreate is the builder for creating a DemoOrder entity.
type DemoOrderCreate struct {
	config
	mutation *DemoOrderMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (doc *DemoOrderCreate) SetCreatedAt(t time.Time) *DemoOrderCreate {
	doc.mutation.SetCreatedAt(t)
	return doc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (doc *DemoOrderCreate) SetNillableCreatedAt(t *time.Time) *DemoOrderCreate {
	if t != nil {
		doc.SetCreatedAt(*t)
	}
	return doc
}

// SetUpdatedAt sets the "updated_at" field.
func (doc *DemoOrderCreate) SetUpdatedAt(t time.Time) *DemoOrderCreate {
	doc.mutation.SetUpdatedAt(t)
	return doc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (doc *DemoOrderCreate) SetNillableUpdatedAt(t *time.Time) *DemoOrderCreate {
	if t != nil {
		doc.SetUpdatedAt(*t)
	}
	return doc
}

// SetDeletedAt sets the "deleted_at" field.
func (doc *DemoOrderCreate) SetDeletedAt(t time.Time) *DemoOrderCreate {
	doc.mutation.SetDeletedAt(t)
	return doc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (doc *DemoOrderCreate) SetNillableDeletedAt(t *time.Time) *DemoOrderCreate {
	if t != nil {
		doc.SetDeletedAt(*t)
	}
	return doc
}

// SetUserID sets the "user_id" field.
func (doc *DemoOrderCreate) SetUserID(s string) *DemoOrderCreate {
	doc.mutation.SetUserID(s)
	return doc
}

// SetSpuID sets the "spu_id" field.
func (doc *DemoOrderCreate) SetSpuID(u uint64) *DemoOrderCreate {
	doc.mutation.SetSpuID(u)
	return doc
}

// SetSpuName sets the "spu_name" field.
func (doc *DemoOrderCreate) SetSpuName(s string) *DemoOrderCreate {
	doc.mutation.SetSpuName(s)
	return doc
}

// SetPrice sets the "price" field.
func (doc *DemoOrderCreate) SetPrice(i int32) *DemoOrderCreate {
	doc.mutation.SetPrice(i)
	return doc
}

// SetPayStatus sets the "pay_status" field.
func (doc *DemoOrderCreate) SetPayStatus(b bool) *DemoOrderCreate {
	doc.mutation.SetPayStatus(b)
	return doc
}

// SetPayOrderId sets the "pay_orderId" field.
func (doc *DemoOrderCreate) SetPayOrderId(u uint64) *DemoOrderCreate {
	doc.mutation.SetPayOrderId(u)
	return doc
}

// SetNillablePayOrderId sets the "pay_orderId" field if the given value is not nil.
func (doc *DemoOrderCreate) SetNillablePayOrderId(u *uint64) *DemoOrderCreate {
	if u != nil {
		doc.SetPayOrderId(*u)
	}
	return doc
}

// SetPayTime sets the "pay_time" field.
func (doc *DemoOrderCreate) SetPayTime(t time.Time) *DemoOrderCreate {
	doc.mutation.SetPayTime(t)
	return doc
}

// SetNillablePayTime sets the "pay_time" field if the given value is not nil.
func (doc *DemoOrderCreate) SetNillablePayTime(t *time.Time) *DemoOrderCreate {
	if t != nil {
		doc.SetPayTime(*t)
	}
	return doc
}

// SetPayChannelCode sets the "pay_channel_code" field.
func (doc *DemoOrderCreate) SetPayChannelCode(s string) *DemoOrderCreate {
	doc.mutation.SetPayChannelCode(s)
	return doc
}

// SetNillablePayChannelCode sets the "pay_channel_code" field if the given value is not nil.
func (doc *DemoOrderCreate) SetNillablePayChannelCode(s *string) *DemoOrderCreate {
	if s != nil {
		doc.SetPayChannelCode(*s)
	}
	return doc
}

// SetPayRefundID sets the "pay_refund_id" field.
func (doc *DemoOrderCreate) SetPayRefundID(u uint64) *DemoOrderCreate {
	doc.mutation.SetPayRefundID(u)
	return doc
}

// SetNillablePayRefundID sets the "pay_refund_id" field if the given value is not nil.
func (doc *DemoOrderCreate) SetNillablePayRefundID(u *uint64) *DemoOrderCreate {
	if u != nil {
		doc.SetPayRefundID(*u)
	}
	return doc
}

// SetRefundPrice sets the "refund_price" field.
func (doc *DemoOrderCreate) SetRefundPrice(i int32) *DemoOrderCreate {
	doc.mutation.SetRefundPrice(i)
	return doc
}

// SetNillableRefundPrice sets the "refund_price" field if the given value is not nil.
func (doc *DemoOrderCreate) SetNillableRefundPrice(i *int32) *DemoOrderCreate {
	if i != nil {
		doc.SetRefundPrice(*i)
	}
	return doc
}

// SetRefundTime sets the "refund_time" field.
func (doc *DemoOrderCreate) SetRefundTime(t time.Time) *DemoOrderCreate {
	doc.mutation.SetRefundTime(t)
	return doc
}

// SetNillableRefundTime sets the "refund_time" field if the given value is not nil.
func (doc *DemoOrderCreate) SetNillableRefundTime(t *time.Time) *DemoOrderCreate {
	if t != nil {
		doc.SetRefundTime(*t)
	}
	return doc
}

// SetID sets the "id" field.
func (doc *DemoOrderCreate) SetID(u uint64) *DemoOrderCreate {
	doc.mutation.SetID(u)
	return doc
}

// Mutation returns the DemoOrderMutation object of the builder.
func (doc *DemoOrderCreate) Mutation() *DemoOrderMutation {
	return doc.mutation
}

// Save creates the DemoOrder in the database.
func (doc *DemoOrderCreate) Save(ctx context.Context) (*DemoOrder, error) {
	if err := doc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, doc.sqlSave, doc.mutation, doc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (doc *DemoOrderCreate) SaveX(ctx context.Context) *DemoOrder {
	v, err := doc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (doc *DemoOrderCreate) Exec(ctx context.Context) error {
	_, err := doc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (doc *DemoOrderCreate) ExecX(ctx context.Context) {
	if err := doc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (doc *DemoOrderCreate) defaults() error {
	if _, ok := doc.mutation.CreatedAt(); !ok {
		if demoorder.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized demoorder.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := demoorder.DefaultCreatedAt()
		doc.mutation.SetCreatedAt(v)
	}
	if _, ok := doc.mutation.UpdatedAt(); !ok {
		if demoorder.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized demoorder.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := demoorder.DefaultUpdatedAt()
		doc.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (doc *DemoOrderCreate) check() error {
	if _, ok := doc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "DemoOrder.created_at"`)}
	}
	if _, ok := doc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "DemoOrder.updated_at"`)}
	}
	if _, ok := doc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "DemoOrder.user_id"`)}
	}
	if _, ok := doc.mutation.SpuID(); !ok {
		return &ValidationError{Name: "spu_id", err: errors.New(`ent: missing required field "DemoOrder.spu_id"`)}
	}
	if _, ok := doc.mutation.SpuName(); !ok {
		return &ValidationError{Name: "spu_name", err: errors.New(`ent: missing required field "DemoOrder.spu_name"`)}
	}
	if _, ok := doc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "DemoOrder.price"`)}
	}
	if _, ok := doc.mutation.PayStatus(); !ok {
		return &ValidationError{Name: "pay_status", err: errors.New(`ent: missing required field "DemoOrder.pay_status"`)}
	}
	return nil
}

func (doc *DemoOrderCreate) sqlSave(ctx context.Context) (*DemoOrder, error) {
	if err := doc.check(); err != nil {
		return nil, err
	}
	_node, _spec := doc.createSpec()
	if err := sqlgraph.CreateNode(ctx, doc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = uint64(id)
	}
	doc.mutation.id = &_node.ID
	doc.mutation.done = true
	return _node, nil
}

func (doc *DemoOrderCreate) createSpec() (*DemoOrder, *sqlgraph.CreateSpec) {
	var (
		_node = &DemoOrder{config: doc.config}
		_spec = sqlgraph.NewCreateSpec(demoorder.Table, sqlgraph.NewFieldSpec(demoorder.FieldID, field.TypeUint64))
	)
	if id, ok := doc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := doc.mutation.CreatedAt(); ok {
		_spec.SetField(demoorder.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := doc.mutation.UpdatedAt(); ok {
		_spec.SetField(demoorder.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := doc.mutation.DeletedAt(); ok {
		_spec.SetField(demoorder.FieldDeletedAt, field.TypeTime, value)
		_node.DeletedAt = value
	}
	if value, ok := doc.mutation.UserID(); ok {
		_spec.SetField(demoorder.FieldUserID, field.TypeString, value)
		_node.UserID = value
	}
	if value, ok := doc.mutation.SpuID(); ok {
		_spec.SetField(demoorder.FieldSpuID, field.TypeUint64, value)
		_node.SpuID = value
	}
	if value, ok := doc.mutation.SpuName(); ok {
		_spec.SetField(demoorder.FieldSpuName, field.TypeString, value)
		_node.SpuName = value
	}
	if value, ok := doc.mutation.Price(); ok {
		_spec.SetField(demoorder.FieldPrice, field.TypeInt32, value)
		_node.Price = value
	}
	if value, ok := doc.mutation.PayStatus(); ok {
		_spec.SetField(demoorder.FieldPayStatus, field.TypeBool, value)
		_node.PayStatus = value
	}
	if value, ok := doc.mutation.PayOrderId(); ok {
		_spec.SetField(demoorder.FieldPayOrderId, field.TypeUint64, value)
		_node.PayOrderId = value
	}
	if value, ok := doc.mutation.PayTime(); ok {
		_spec.SetField(demoorder.FieldPayTime, field.TypeTime, value)
		_node.PayTime = value
	}
	if value, ok := doc.mutation.PayChannelCode(); ok {
		_spec.SetField(demoorder.FieldPayChannelCode, field.TypeString, value)
		_node.PayChannelCode = value
	}
	if value, ok := doc.mutation.PayRefundID(); ok {
		_spec.SetField(demoorder.FieldPayRefundID, field.TypeUint64, value)
		_node.PayRefundID = value
	}
	if value, ok := doc.mutation.RefundPrice(); ok {
		_spec.SetField(demoorder.FieldRefundPrice, field.TypeInt32, value)
		_node.RefundPrice = value
	}
	if value, ok := doc.mutation.RefundTime(); ok {
		_spec.SetField(demoorder.FieldRefundTime, field.TypeTime, value)
		_node.RefundTime = value
	}
	return _node, _spec
}

// DemoOrderCreateBulk is the builder for creating many DemoOrder entities in bulk.
type DemoOrderCreateBulk struct {
	config
	err      error
	builders []*DemoOrderCreate
}

// Save creates the DemoOrder entities in the database.
func (docb *DemoOrderCreateBulk) Save(ctx context.Context) ([]*DemoOrder, error) {
	if docb.err != nil {
		return nil, docb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(docb.builders))
	nodes := make([]*DemoOrder, len(docb.builders))
	mutators := make([]Mutator, len(docb.builders))
	for i := range docb.builders {
		func(i int, root context.Context) {
			builder := docb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*DemoOrderMutation)
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
					_, err = mutators[i+1].Mutate(root, docb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, docb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, docb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (docb *DemoOrderCreateBulk) SaveX(ctx context.Context) []*DemoOrder {
	v, err := docb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (docb *DemoOrderCreateBulk) Exec(ctx context.Context) error {
	_, err := docb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (docb *DemoOrderCreateBulk) ExecX(ctx context.Context) {
	if err := docb.Exec(ctx); err != nil {
		panic(err)
	}
}
