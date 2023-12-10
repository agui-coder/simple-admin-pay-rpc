// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/notifytask"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/predicate"
)

// NotifyTaskQuery is the builder for querying NotifyTask entities.
type NotifyTaskQuery struct {
	config
	ctx        *QueryContext
	order      []notifytask.OrderOption
	inters     []Interceptor
	predicates []predicate.NotifyTask
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the NotifyTaskQuery builder.
func (ntq *NotifyTaskQuery) Where(ps ...predicate.NotifyTask) *NotifyTaskQuery {
	ntq.predicates = append(ntq.predicates, ps...)
	return ntq
}

// Limit the number of records to be returned by this query.
func (ntq *NotifyTaskQuery) Limit(limit int) *NotifyTaskQuery {
	ntq.ctx.Limit = &limit
	return ntq
}

// Offset to start from.
func (ntq *NotifyTaskQuery) Offset(offset int) *NotifyTaskQuery {
	ntq.ctx.Offset = &offset
	return ntq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ntq *NotifyTaskQuery) Unique(unique bool) *NotifyTaskQuery {
	ntq.ctx.Unique = &unique
	return ntq
}

// Order specifies how the records should be ordered.
func (ntq *NotifyTaskQuery) Order(o ...notifytask.OrderOption) *NotifyTaskQuery {
	ntq.order = append(ntq.order, o...)
	return ntq
}

// First returns the first NotifyTask entity from the query.
// Returns a *NotFoundError when no NotifyTask was found.
func (ntq *NotifyTaskQuery) First(ctx context.Context) (*NotifyTask, error) {
	nodes, err := ntq.Limit(1).All(setContextOp(ctx, ntq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{notifytask.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ntq *NotifyTaskQuery) FirstX(ctx context.Context) *NotifyTask {
	node, err := ntq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first NotifyTask ID from the query.
// Returns a *NotFoundError when no NotifyTask ID was found.
func (ntq *NotifyTaskQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = ntq.Limit(1).IDs(setContextOp(ctx, ntq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{notifytask.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ntq *NotifyTaskQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := ntq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single NotifyTask entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one NotifyTask entity is found.
// Returns a *NotFoundError when no NotifyTask entities are found.
func (ntq *NotifyTaskQuery) Only(ctx context.Context) (*NotifyTask, error) {
	nodes, err := ntq.Limit(2).All(setContextOp(ctx, ntq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{notifytask.Label}
	default:
		return nil, &NotSingularError{notifytask.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ntq *NotifyTaskQuery) OnlyX(ctx context.Context) *NotifyTask {
	node, err := ntq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only NotifyTask ID in the query.
// Returns a *NotSingularError when more than one NotifyTask ID is found.
// Returns a *NotFoundError when no entities are found.
func (ntq *NotifyTaskQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = ntq.Limit(2).IDs(setContextOp(ctx, ntq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{notifytask.Label}
	default:
		err = &NotSingularError{notifytask.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ntq *NotifyTaskQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := ntq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of NotifyTasks.
func (ntq *NotifyTaskQuery) All(ctx context.Context) ([]*NotifyTask, error) {
	ctx = setContextOp(ctx, ntq.ctx, "All")
	if err := ntq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*NotifyTask, *NotifyTaskQuery]()
	return withInterceptors[[]*NotifyTask](ctx, ntq, qr, ntq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ntq *NotifyTaskQuery) AllX(ctx context.Context) []*NotifyTask {
	nodes, err := ntq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of NotifyTask IDs.
func (ntq *NotifyTaskQuery) IDs(ctx context.Context) (ids []uint64, err error) {
	if ntq.ctx.Unique == nil && ntq.path != nil {
		ntq.Unique(true)
	}
	ctx = setContextOp(ctx, ntq.ctx, "IDs")
	if err = ntq.Select(notifytask.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ntq *NotifyTaskQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := ntq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ntq *NotifyTaskQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ntq.ctx, "Count")
	if err := ntq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ntq, querierCount[*NotifyTaskQuery](), ntq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ntq *NotifyTaskQuery) CountX(ctx context.Context) int {
	count, err := ntq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ntq *NotifyTaskQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ntq.ctx, "Exist")
	switch _, err := ntq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ntq *NotifyTaskQuery) ExistX(ctx context.Context) bool {
	exist, err := ntq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the NotifyTaskQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ntq *NotifyTaskQuery) Clone() *NotifyTaskQuery {
	if ntq == nil {
		return nil
	}
	return &NotifyTaskQuery{
		config:     ntq.config,
		ctx:        ntq.ctx.Clone(),
		order:      append([]notifytask.OrderOption{}, ntq.order...),
		inters:     append([]Interceptor{}, ntq.inters...),
		predicates: append([]predicate.NotifyTask{}, ntq.predicates...),
		// clone intermediate query.
		sql:  ntq.sql.Clone(),
		path: ntq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.NotifyTask.Query().
//		GroupBy(notifytask.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ntq *NotifyTaskQuery) GroupBy(field string, fields ...string) *NotifyTaskGroupBy {
	ntq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &NotifyTaskGroupBy{build: ntq}
	grbuild.flds = &ntq.ctx.Fields
	grbuild.label = notifytask.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.NotifyTask.Query().
//		Select(notifytask.FieldCreatedAt).
//		Scan(ctx, &v)
func (ntq *NotifyTaskQuery) Select(fields ...string) *NotifyTaskSelect {
	ntq.ctx.Fields = append(ntq.ctx.Fields, fields...)
	sbuild := &NotifyTaskSelect{NotifyTaskQuery: ntq}
	sbuild.label = notifytask.Label
	sbuild.flds, sbuild.scan = &ntq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a NotifyTaskSelect configured with the given aggregations.
func (ntq *NotifyTaskQuery) Aggregate(fns ...AggregateFunc) *NotifyTaskSelect {
	return ntq.Select().Aggregate(fns...)
}

func (ntq *NotifyTaskQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ntq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ntq); err != nil {
				return err
			}
		}
	}
	for _, f := range ntq.ctx.Fields {
		if !notifytask.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ntq.path != nil {
		prev, err := ntq.path(ctx)
		if err != nil {
			return err
		}
		ntq.sql = prev
	}
	return nil
}

func (ntq *NotifyTaskQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*NotifyTask, error) {
	var (
		nodes = []*NotifyTask{}
		_spec = ntq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*NotifyTask).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &NotifyTask{config: ntq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ntq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ntq *NotifyTaskQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ntq.querySpec()
	_spec.Node.Columns = ntq.ctx.Fields
	if len(ntq.ctx.Fields) > 0 {
		_spec.Unique = ntq.ctx.Unique != nil && *ntq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ntq.driver, _spec)
}

func (ntq *NotifyTaskQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(notifytask.Table, notifytask.Columns, sqlgraph.NewFieldSpec(notifytask.FieldID, field.TypeUint64))
	_spec.From = ntq.sql
	if unique := ntq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ntq.path != nil {
		_spec.Unique = true
	}
	if fields := ntq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, notifytask.FieldID)
		for i := range fields {
			if fields[i] != notifytask.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ntq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ntq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ntq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ntq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ntq *NotifyTaskQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ntq.driver.Dialect())
	t1 := builder.Table(notifytask.Table)
	columns := ntq.ctx.Fields
	if len(columns) == 0 {
		columns = notifytask.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ntq.sql != nil {
		selector = ntq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ntq.ctx.Unique != nil && *ntq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ntq.predicates {
		p(selector)
	}
	for _, p := range ntq.order {
		p(selector)
	}
	if offset := ntq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ntq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// NotifyTaskGroupBy is the group-by builder for NotifyTask entities.
type NotifyTaskGroupBy struct {
	selector
	build *NotifyTaskQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ntgb *NotifyTaskGroupBy) Aggregate(fns ...AggregateFunc) *NotifyTaskGroupBy {
	ntgb.fns = append(ntgb.fns, fns...)
	return ntgb
}

// Scan applies the selector query and scans the result into the given value.
func (ntgb *NotifyTaskGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ntgb.build.ctx, "GroupBy")
	if err := ntgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NotifyTaskQuery, *NotifyTaskGroupBy](ctx, ntgb.build, ntgb, ntgb.build.inters, v)
}

func (ntgb *NotifyTaskGroupBy) sqlScan(ctx context.Context, root *NotifyTaskQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ntgb.fns))
	for _, fn := range ntgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ntgb.flds)+len(ntgb.fns))
		for _, f := range *ntgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ntgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ntgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// NotifyTaskSelect is the builder for selecting fields of NotifyTask entities.
type NotifyTaskSelect struct {
	*NotifyTaskQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (nts *NotifyTaskSelect) Aggregate(fns ...AggregateFunc) *NotifyTaskSelect {
	nts.fns = append(nts.fns, fns...)
	return nts
}

// Scan applies the selector query and scans the result into the given value.
func (nts *NotifyTaskSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, nts.ctx, "Select")
	if err := nts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*NotifyTaskQuery, *NotifyTaskSelect](ctx, nts.NotifyTaskQuery, nts, nts.inters, v)
}

func (nts *NotifyTaskSelect) sqlScan(ctx context.Context, root *NotifyTaskQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(nts.fns))
	for _, fn := range nts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*nts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := nts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
