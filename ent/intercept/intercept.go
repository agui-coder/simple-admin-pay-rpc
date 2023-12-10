// Code generated by ent, DO NOT EDIT.

package intercept

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/agui-coder/simple-admin-pay-rpc/ent"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/app"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/channel"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/demoorder"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/notifylog"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/notifytask"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/order"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/orderextension"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/predicate"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/refund"
)

// The Query interface represents an operation that queries a graph.
// By using this interface, users can write generic code that manipulates
// query builders of different types.
type Query interface {
	// Type returns the string representation of the query type.
	Type() string
	// Limit the number of records to be returned by this query.
	Limit(int)
	// Offset to start from.
	Offset(int)
	// Unique configures the query builder to filter duplicate records.
	Unique(bool)
	// Order specifies how the records should be ordered.
	Order(...func(*sql.Selector))
	// WhereP appends storage-level predicates to the query builder. Using this method, users
	// can use type-assertion to append predicates that do not depend on any generated package.
	WhereP(...func(*sql.Selector))
}

// The Func type is an adapter that allows ordinary functions to be used as interceptors.
// Unlike traversal functions, interceptors are skipped during graph traversals. Note that the
// implementation of Func is different from the one defined in entgo.io/ent.InterceptFunc.
type Func func(context.Context, Query) error

// Intercept calls f(ctx, q) and then applied the next Querier.
func (f Func) Intercept(next ent.Querier) ent.Querier {
	return ent.QuerierFunc(func(ctx context.Context, q ent.Query) (ent.Value, error) {
		query, err := NewQuery(q)
		if err != nil {
			return nil, err
		}
		if err := f(ctx, query); err != nil {
			return nil, err
		}
		return next.Query(ctx, q)
	})
}

// The TraverseFunc type is an adapter to allow the use of ordinary function as Traverser.
// If f is a function with the appropriate signature, TraverseFunc(f) is a Traverser that calls f.
type TraverseFunc func(context.Context, Query) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseFunc) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseFunc) Traverse(ctx context.Context, q ent.Query) error {
	query, err := NewQuery(q)
	if err != nil {
		return err
	}
	return f(ctx, query)
}

// The AppFunc type is an adapter to allow the use of ordinary function as a Querier.
type AppFunc func(context.Context, *ent.AppQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f AppFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.AppQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.AppQuery", q)
}

// The TraverseApp type is an adapter to allow the use of ordinary function as Traverser.
type TraverseApp func(context.Context, *ent.AppQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseApp) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseApp) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.AppQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.AppQuery", q)
}

// The ChannelFunc type is an adapter to allow the use of ordinary function as a Querier.
type ChannelFunc func(context.Context, *ent.ChannelQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f ChannelFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.ChannelQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.ChannelQuery", q)
}

// The TraverseChannel type is an adapter to allow the use of ordinary function as Traverser.
type TraverseChannel func(context.Context, *ent.ChannelQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseChannel) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseChannel) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.ChannelQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.ChannelQuery", q)
}

// The DemoOrderFunc type is an adapter to allow the use of ordinary function as a Querier.
type DemoOrderFunc func(context.Context, *ent.DemoOrderQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f DemoOrderFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.DemoOrderQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.DemoOrderQuery", q)
}

// The TraverseDemoOrder type is an adapter to allow the use of ordinary function as Traverser.
type TraverseDemoOrder func(context.Context, *ent.DemoOrderQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseDemoOrder) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseDemoOrder) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.DemoOrderQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.DemoOrderQuery", q)
}

// The NotifyLogFunc type is an adapter to allow the use of ordinary function as a Querier.
type NotifyLogFunc func(context.Context, *ent.NotifyLogQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f NotifyLogFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.NotifyLogQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.NotifyLogQuery", q)
}

// The TraverseNotifyLog type is an adapter to allow the use of ordinary function as Traverser.
type TraverseNotifyLog func(context.Context, *ent.NotifyLogQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseNotifyLog) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseNotifyLog) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.NotifyLogQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.NotifyLogQuery", q)
}

// The NotifyTaskFunc type is an adapter to allow the use of ordinary function as a Querier.
type NotifyTaskFunc func(context.Context, *ent.NotifyTaskQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f NotifyTaskFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.NotifyTaskQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.NotifyTaskQuery", q)
}

// The TraverseNotifyTask type is an adapter to allow the use of ordinary function as Traverser.
type TraverseNotifyTask func(context.Context, *ent.NotifyTaskQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseNotifyTask) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseNotifyTask) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.NotifyTaskQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.NotifyTaskQuery", q)
}

// The OrderFunc type is an adapter to allow the use of ordinary function as a Querier.
type OrderFunc func(context.Context, *ent.OrderQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f OrderFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.OrderQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.OrderQuery", q)
}

// The TraverseOrder type is an adapter to allow the use of ordinary function as Traverser.
type TraverseOrder func(context.Context, *ent.OrderQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseOrder) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseOrder) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.OrderQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.OrderQuery", q)
}

// The OrderExtensionFunc type is an adapter to allow the use of ordinary function as a Querier.
type OrderExtensionFunc func(context.Context, *ent.OrderExtensionQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f OrderExtensionFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.OrderExtensionQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.OrderExtensionQuery", q)
}

// The TraverseOrderExtension type is an adapter to allow the use of ordinary function as Traverser.
type TraverseOrderExtension func(context.Context, *ent.OrderExtensionQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseOrderExtension) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseOrderExtension) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.OrderExtensionQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.OrderExtensionQuery", q)
}

// The RefundFunc type is an adapter to allow the use of ordinary function as a Querier.
type RefundFunc func(context.Context, *ent.RefundQuery) (ent.Value, error)

// Query calls f(ctx, q).
func (f RefundFunc) Query(ctx context.Context, q ent.Query) (ent.Value, error) {
	if q, ok := q.(*ent.RefundQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *ent.RefundQuery", q)
}

// The TraverseRefund type is an adapter to allow the use of ordinary function as Traverser.
type TraverseRefund func(context.Context, *ent.RefundQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseRefund) Intercept(next ent.Querier) ent.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseRefund) Traverse(ctx context.Context, q ent.Query) error {
	if q, ok := q.(*ent.RefundQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *ent.RefundQuery", q)
}

// NewQuery returns the generic Query interface for the given typed query.
func NewQuery(q ent.Query) (Query, error) {
	switch q := q.(type) {
	case *ent.AppQuery:
		return &query[*ent.AppQuery, predicate.App, app.OrderOption]{typ: ent.TypeApp, tq: q}, nil
	case *ent.ChannelQuery:
		return &query[*ent.ChannelQuery, predicate.Channel, channel.OrderOption]{typ: ent.TypeChannel, tq: q}, nil
	case *ent.DemoOrderQuery:
		return &query[*ent.DemoOrderQuery, predicate.DemoOrder, demoorder.OrderOption]{typ: ent.TypeDemoOrder, tq: q}, nil
	case *ent.NotifyLogQuery:
		return &query[*ent.NotifyLogQuery, predicate.NotifyLog, notifylog.OrderOption]{typ: ent.TypeNotifyLog, tq: q}, nil
	case *ent.NotifyTaskQuery:
		return &query[*ent.NotifyTaskQuery, predicate.NotifyTask, notifytask.OrderOption]{typ: ent.TypeNotifyTask, tq: q}, nil
	case *ent.OrderQuery:
		return &query[*ent.OrderQuery, predicate.Order, order.OrderOption]{typ: ent.TypeOrder, tq: q}, nil
	case *ent.OrderExtensionQuery:
		return &query[*ent.OrderExtensionQuery, predicate.OrderExtension, orderextension.OrderOption]{typ: ent.TypeOrderExtension, tq: q}, nil
	case *ent.RefundQuery:
		return &query[*ent.RefundQuery, predicate.Refund, refund.OrderOption]{typ: ent.TypeRefund, tq: q}, nil
	default:
		return nil, fmt.Errorf("unknown query type %T", q)
	}
}

type query[T any, P ~func(*sql.Selector), R ~func(*sql.Selector)] struct {
	typ string
	tq  interface {
		Limit(int) T
		Offset(int) T
		Unique(bool) T
		Order(...R) T
		Where(...P) T
	}
}

func (q query[T, P, R]) Type() string {
	return q.typ
}

func (q query[T, P, R]) Limit(limit int) {
	q.tq.Limit(limit)
}

func (q query[T, P, R]) Offset(offset int) {
	q.tq.Offset(offset)
}

func (q query[T, P, R]) Unique(unique bool) {
	q.tq.Unique(unique)
}

func (q query[T, P, R]) Order(orders ...func(*sql.Selector)) {
	rs := make([]R, len(orders))
	for i := range orders {
		rs[i] = orders[i]
	}
	q.tq.Order(rs...)
}

func (q query[T, P, R]) WhereP(ps ...func(*sql.Selector)) {
	p := make([]P, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	q.tq.Where(p...)
}
