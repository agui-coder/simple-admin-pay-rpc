// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/predicate"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/refund"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RefundDelete is the builder for deleting a Refund entity.
type RefundDelete struct {
	config
	hooks    []Hook
	mutation *RefundMutation
}

// Where appends a list predicates to the RefundDelete builder.
func (rd *RefundDelete) Where(ps ...predicate.Refund) *RefundDelete {
	rd.mutation.Where(ps...)
	return rd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rd *RefundDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, rd.sqlExec, rd.mutation, rd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (rd *RefundDelete) ExecX(ctx context.Context) int {
	n, err := rd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rd *RefundDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(refund.Table, sqlgraph.NewFieldSpec(refund.FieldID, field.TypeUint64))
	if ps := rd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	rd.mutation.done = true
	return affected, err
}

// RefundDeleteOne is the builder for deleting a single Refund entity.
type RefundDeleteOne struct {
	rd *RefundDelete
}

// Where appends a list predicates to the RefundDelete builder.
func (rdo *RefundDeleteOne) Where(ps ...predicate.Refund) *RefundDeleteOne {
	rdo.rd.mutation.Where(ps...)
	return rdo
}

// Exec executes the deletion query.
func (rdo *RefundDeleteOne) Exec(ctx context.Context) error {
	n, err := rdo.rd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{refund.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rdo *RefundDeleteOne) ExecX(ctx context.Context) {
	if err := rdo.Exec(ctx); err != nil {
		panic(err)
	}
}
