// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/orderextension"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/predicate"
)

// OrderExtensionQuery is the builder for querying OrderExtension entities.
type OrderExtensionQuery struct {
	config
	ctx        *QueryContext
	order      []orderextension.OrderOption
	inters     []Interceptor
	predicates []predicate.OrderExtension
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrderExtensionQuery builder.
func (oeq *OrderExtensionQuery) Where(ps ...predicate.OrderExtension) *OrderExtensionQuery {
	oeq.predicates = append(oeq.predicates, ps...)
	return oeq
}

// Limit the number of records to be returned by this query.
func (oeq *OrderExtensionQuery) Limit(limit int) *OrderExtensionQuery {
	oeq.ctx.Limit = &limit
	return oeq
}

// Offset to start from.
func (oeq *OrderExtensionQuery) Offset(offset int) *OrderExtensionQuery {
	oeq.ctx.Offset = &offset
	return oeq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (oeq *OrderExtensionQuery) Unique(unique bool) *OrderExtensionQuery {
	oeq.ctx.Unique = &unique
	return oeq
}

// Order specifies how the records should be ordered.
func (oeq *OrderExtensionQuery) Order(o ...orderextension.OrderOption) *OrderExtensionQuery {
	oeq.order = append(oeq.order, o...)
	return oeq
}

// First returns the first OrderExtension entity from the query.
// Returns a *NotFoundError when no OrderExtension was found.
func (oeq *OrderExtensionQuery) First(ctx context.Context) (*OrderExtension, error) {
	nodes, err := oeq.Limit(1).All(setContextOp(ctx, oeq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{orderextension.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (oeq *OrderExtensionQuery) FirstX(ctx context.Context) *OrderExtension {
	node, err := oeq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrderExtension ID from the query.
// Returns a *NotFoundError when no OrderExtension ID was found.
func (oeq *OrderExtensionQuery) FirstID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = oeq.Limit(1).IDs(setContextOp(ctx, oeq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{orderextension.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (oeq *OrderExtensionQuery) FirstIDX(ctx context.Context) uint64 {
	id, err := oeq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrderExtension entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OrderExtension entity is found.
// Returns a *NotFoundError when no OrderExtension entities are found.
func (oeq *OrderExtensionQuery) Only(ctx context.Context) (*OrderExtension, error) {
	nodes, err := oeq.Limit(2).All(setContextOp(ctx, oeq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{orderextension.Label}
	default:
		return nil, &NotSingularError{orderextension.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (oeq *OrderExtensionQuery) OnlyX(ctx context.Context) *OrderExtension {
	node, err := oeq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrderExtension ID in the query.
// Returns a *NotSingularError when more than one OrderExtension ID is found.
// Returns a *NotFoundError when no entities are found.
func (oeq *OrderExtensionQuery) OnlyID(ctx context.Context) (id uint64, err error) {
	var ids []uint64
	if ids, err = oeq.Limit(2).IDs(setContextOp(ctx, oeq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{orderextension.Label}
	default:
		err = &NotSingularError{orderextension.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (oeq *OrderExtensionQuery) OnlyIDX(ctx context.Context) uint64 {
	id, err := oeq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrderExtensions.
func (oeq *OrderExtensionQuery) All(ctx context.Context) ([]*OrderExtension, error) {
	ctx = setContextOp(ctx, oeq.ctx, "All")
	if err := oeq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*OrderExtension, *OrderExtensionQuery]()
	return withInterceptors[[]*OrderExtension](ctx, oeq, qr, oeq.inters)
}

// AllX is like All, but panics if an error occurs.
func (oeq *OrderExtensionQuery) AllX(ctx context.Context) []*OrderExtension {
	nodes, err := oeq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrderExtension IDs.
func (oeq *OrderExtensionQuery) IDs(ctx context.Context) (ids []uint64, err error) {
	if oeq.ctx.Unique == nil && oeq.path != nil {
		oeq.Unique(true)
	}
	ctx = setContextOp(ctx, oeq.ctx, "IDs")
	if err = oeq.Select(orderextension.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (oeq *OrderExtensionQuery) IDsX(ctx context.Context) []uint64 {
	ids, err := oeq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (oeq *OrderExtensionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, oeq.ctx, "Count")
	if err := oeq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, oeq, querierCount[*OrderExtensionQuery](), oeq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (oeq *OrderExtensionQuery) CountX(ctx context.Context) int {
	count, err := oeq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (oeq *OrderExtensionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, oeq.ctx, "Exist")
	switch _, err := oeq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (oeq *OrderExtensionQuery) ExistX(ctx context.Context) bool {
	exist, err := oeq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrderExtensionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (oeq *OrderExtensionQuery) Clone() *OrderExtensionQuery {
	if oeq == nil {
		return nil
	}
	return &OrderExtensionQuery{
		config:     oeq.config,
		ctx:        oeq.ctx.Clone(),
		order:      append([]orderextension.OrderOption{}, oeq.order...),
		inters:     append([]Interceptor{}, oeq.inters...),
		predicates: append([]predicate.OrderExtension{}, oeq.predicates...),
		// clone intermediate query.
		sql:  oeq.sql.Clone(),
		path: oeq.path,
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
//	client.OrderExtension.Query().
//		GroupBy(orderextension.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (oeq *OrderExtensionQuery) GroupBy(field string, fields ...string) *OrderExtensionGroupBy {
	oeq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OrderExtensionGroupBy{build: oeq}
	grbuild.flds = &oeq.ctx.Fields
	grbuild.label = orderextension.Label
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
//	client.OrderExtension.Query().
//		Select(orderextension.FieldCreatedAt).
//		Scan(ctx, &v)
func (oeq *OrderExtensionQuery) Select(fields ...string) *OrderExtensionSelect {
	oeq.ctx.Fields = append(oeq.ctx.Fields, fields...)
	sbuild := &OrderExtensionSelect{OrderExtensionQuery: oeq}
	sbuild.label = orderextension.Label
	sbuild.flds, sbuild.scan = &oeq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OrderExtensionSelect configured with the given aggregations.
func (oeq *OrderExtensionQuery) Aggregate(fns ...AggregateFunc) *OrderExtensionSelect {
	return oeq.Select().Aggregate(fns...)
}

func (oeq *OrderExtensionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range oeq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, oeq); err != nil {
				return err
			}
		}
	}
	for _, f := range oeq.ctx.Fields {
		if !orderextension.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if oeq.path != nil {
		prev, err := oeq.path(ctx)
		if err != nil {
			return err
		}
		oeq.sql = prev
	}
	return nil
}

func (oeq *OrderExtensionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OrderExtension, error) {
	var (
		nodes = []*OrderExtension{}
		_spec = oeq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*OrderExtension).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &OrderExtension{config: oeq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, oeq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (oeq *OrderExtensionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := oeq.querySpec()
	_spec.Node.Columns = oeq.ctx.Fields
	if len(oeq.ctx.Fields) > 0 {
		_spec.Unique = oeq.ctx.Unique != nil && *oeq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, oeq.driver, _spec)
}

func (oeq *OrderExtensionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(orderextension.Table, orderextension.Columns, sqlgraph.NewFieldSpec(orderextension.FieldID, field.TypeUint64))
	_spec.From = oeq.sql
	if unique := oeq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if oeq.path != nil {
		_spec.Unique = true
	}
	if fields := oeq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, orderextension.FieldID)
		for i := range fields {
			if fields[i] != orderextension.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := oeq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := oeq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := oeq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := oeq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (oeq *OrderExtensionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(oeq.driver.Dialect())
	t1 := builder.Table(orderextension.Table)
	columns := oeq.ctx.Fields
	if len(columns) == 0 {
		columns = orderextension.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if oeq.sql != nil {
		selector = oeq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if oeq.ctx.Unique != nil && *oeq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range oeq.predicates {
		p(selector)
	}
	for _, p := range oeq.order {
		p(selector)
	}
	if offset := oeq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := oeq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// OrderExtensionGroupBy is the group-by builder for OrderExtension entities.
type OrderExtensionGroupBy struct {
	selector
	build *OrderExtensionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (oegb *OrderExtensionGroupBy) Aggregate(fns ...AggregateFunc) *OrderExtensionGroupBy {
	oegb.fns = append(oegb.fns, fns...)
	return oegb
}

// Scan applies the selector query and scans the result into the given value.
func (oegb *OrderExtensionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, oegb.build.ctx, "GroupBy")
	if err := oegb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrderExtensionQuery, *OrderExtensionGroupBy](ctx, oegb.build, oegb, oegb.build.inters, v)
}

func (oegb *OrderExtensionGroupBy) sqlScan(ctx context.Context, root *OrderExtensionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(oegb.fns))
	for _, fn := range oegb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*oegb.flds)+len(oegb.fns))
		for _, f := range *oegb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*oegb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := oegb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OrderExtensionSelect is the builder for selecting fields of OrderExtension entities.
type OrderExtensionSelect struct {
	*OrderExtensionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (oes *OrderExtensionSelect) Aggregate(fns ...AggregateFunc) *OrderExtensionSelect {
	oes.fns = append(oes.fns, fns...)
	return oes
}

// Scan applies the selector query and scans the result into the given value.
func (oes *OrderExtensionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, oes.ctx, "Select")
	if err := oes.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrderExtensionQuery, *OrderExtensionSelect](ctx, oes.OrderExtensionQuery, oes, oes.inters, v)
}

func (oes *OrderExtensionSelect) sqlScan(ctx context.Context, root *OrderExtensionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(oes.fns))
	for _, fn := range oes.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*oes.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := oes.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
