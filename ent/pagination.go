// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/agui-coder/simple-admin-pay-rpc/ent/app"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/channel"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/demoorder"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/notifylog"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/notifytask"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/order"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/orderextension"
	"github.com/agui-coder/simple-admin-pay-rpc/ent/refund"
)

const errInvalidPage = "INVALID_PAGE"

const (
	listField     = "list"
	pageNumField  = "pageNum"
	pageSizeField = "pageSize"
)

type PageDetails struct {
	Page  uint64 `json:"page"`
	Size  uint64 `json:"size"`
	Total uint64 `json:"total"`
}

// OrderDirection defines the directions in which to order a list of items.
type OrderDirection string

const (
	// OrderDirectionAsc specifies an ascending order.
	OrderDirectionAsc OrderDirection = "ASC"
	// OrderDirectionDesc specifies a descending order.
	OrderDirectionDesc OrderDirection = "DESC"
)

// Validate the order direction value.
func (o OrderDirection) Validate() error {
	if o != OrderDirectionAsc && o != OrderDirectionDesc {
		return fmt.Errorf("%s is not a valid OrderDirection", o)
	}
	return nil
}

// String implements fmt.Stringer interface.
func (o OrderDirection) String() string {
	return string(o)
}

func (o OrderDirection) reverse() OrderDirection {
	if o == OrderDirectionDesc {
		return OrderDirectionAsc
	}
	return OrderDirectionDesc
}

const errInvalidPagination = "INVALID_PAGINATION"

type AppPager struct {
	Order  app.OrderOption
	Filter func(*AppQuery) (*AppQuery, error)
}

// AppPaginateOption enables pagination customization.
type AppPaginateOption func(*AppPager)

// DefaultAppOrder is the default ordering of App.
var DefaultAppOrder = Desc(app.FieldID)

func newAppPager(opts []AppPaginateOption) (*AppPager, error) {
	pager := &AppPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultAppOrder
	}
	return pager, nil
}

func (p *AppPager) ApplyFilter(query *AppQuery) (*AppQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// AppPageList is App PageList result.
type AppPageList struct {
	List        []*App       `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (a *AppQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...AppPaginateOption,
) (*AppPageList, error) {

	pager, err := newAppPager(opts)
	if err != nil {
		return nil, err
	}

	if a, err = pager.ApplyFilter(a); err != nil {
		return nil, err
	}

	ret := &AppPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := a.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		a = a.Order(pager.Order)
	} else {
		a = a.Order(DefaultAppOrder)
	}

	a = a.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := a.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type ChannelPager struct {
	Order  channel.OrderOption
	Filter func(*ChannelQuery) (*ChannelQuery, error)
}

// ChannelPaginateOption enables pagination customization.
type ChannelPaginateOption func(*ChannelPager)

// DefaultChannelOrder is the default ordering of Channel.
var DefaultChannelOrder = Desc(channel.FieldID)

func newChannelPager(opts []ChannelPaginateOption) (*ChannelPager, error) {
	pager := &ChannelPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultChannelOrder
	}
	return pager, nil
}

func (p *ChannelPager) ApplyFilter(query *ChannelQuery) (*ChannelQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// ChannelPageList is Channel PageList result.
type ChannelPageList struct {
	List        []*Channel   `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (c *ChannelQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...ChannelPaginateOption,
) (*ChannelPageList, error) {

	pager, err := newChannelPager(opts)
	if err != nil {
		return nil, err
	}

	if c, err = pager.ApplyFilter(c); err != nil {
		return nil, err
	}

	ret := &ChannelPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := c.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		c = c.Order(pager.Order)
	} else {
		c = c.Order(DefaultChannelOrder)
	}

	c = c.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := c.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type DemoOrderPager struct {
	Order  demoorder.OrderOption
	Filter func(*DemoOrderQuery) (*DemoOrderQuery, error)
}

// DemoOrderPaginateOption enables pagination customization.
type DemoOrderPaginateOption func(*DemoOrderPager)

// DefaultDemoOrderOrder is the default ordering of DemoOrder.
var DefaultDemoOrderOrder = Desc(demoorder.FieldID)

func newDemoOrderPager(opts []DemoOrderPaginateOption) (*DemoOrderPager, error) {
	pager := &DemoOrderPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultDemoOrderOrder
	}
	return pager, nil
}

func (p *DemoOrderPager) ApplyFilter(query *DemoOrderQuery) (*DemoOrderQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// DemoOrderPageList is DemoOrder PageList result.
type DemoOrderPageList struct {
	List        []*DemoOrder `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (do *DemoOrderQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...DemoOrderPaginateOption,
) (*DemoOrderPageList, error) {

	pager, err := newDemoOrderPager(opts)
	if err != nil {
		return nil, err
	}

	if do, err = pager.ApplyFilter(do); err != nil {
		return nil, err
	}

	ret := &DemoOrderPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := do.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		do = do.Order(pager.Order)
	} else {
		do = do.Order(DefaultDemoOrderOrder)
	}

	do = do.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := do.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type NotifyLogPager struct {
	Order  notifylog.OrderOption
	Filter func(*NotifyLogQuery) (*NotifyLogQuery, error)
}

// NotifyLogPaginateOption enables pagination customization.
type NotifyLogPaginateOption func(*NotifyLogPager)

// DefaultNotifyLogOrder is the default ordering of NotifyLog.
var DefaultNotifyLogOrder = Desc(notifylog.FieldID)

func newNotifyLogPager(opts []NotifyLogPaginateOption) (*NotifyLogPager, error) {
	pager := &NotifyLogPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultNotifyLogOrder
	}
	return pager, nil
}

func (p *NotifyLogPager) ApplyFilter(query *NotifyLogQuery) (*NotifyLogQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// NotifyLogPageList is NotifyLog PageList result.
type NotifyLogPageList struct {
	List        []*NotifyLog `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (nl *NotifyLogQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...NotifyLogPaginateOption,
) (*NotifyLogPageList, error) {

	pager, err := newNotifyLogPager(opts)
	if err != nil {
		return nil, err
	}

	if nl, err = pager.ApplyFilter(nl); err != nil {
		return nil, err
	}

	ret := &NotifyLogPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := nl.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		nl = nl.Order(pager.Order)
	} else {
		nl = nl.Order(DefaultNotifyLogOrder)
	}

	nl = nl.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := nl.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type NotifyTaskPager struct {
	Order  notifytask.OrderOption
	Filter func(*NotifyTaskQuery) (*NotifyTaskQuery, error)
}

// NotifyTaskPaginateOption enables pagination customization.
type NotifyTaskPaginateOption func(*NotifyTaskPager)

// DefaultNotifyTaskOrder is the default ordering of NotifyTask.
var DefaultNotifyTaskOrder = Desc(notifytask.FieldID)

func newNotifyTaskPager(opts []NotifyTaskPaginateOption) (*NotifyTaskPager, error) {
	pager := &NotifyTaskPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultNotifyTaskOrder
	}
	return pager, nil
}

func (p *NotifyTaskPager) ApplyFilter(query *NotifyTaskQuery) (*NotifyTaskQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// NotifyTaskPageList is NotifyTask PageList result.
type NotifyTaskPageList struct {
	List        []*NotifyTask `json:"list"`
	PageDetails *PageDetails  `json:"pageDetails"`
}

func (nt *NotifyTaskQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...NotifyTaskPaginateOption,
) (*NotifyTaskPageList, error) {

	pager, err := newNotifyTaskPager(opts)
	if err != nil {
		return nil, err
	}

	if nt, err = pager.ApplyFilter(nt); err != nil {
		return nil, err
	}

	ret := &NotifyTaskPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := nt.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		nt = nt.Order(pager.Order)
	} else {
		nt = nt.Order(DefaultNotifyTaskOrder)
	}

	nt = nt.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := nt.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type OrderPager struct {
	Order  order.OrderOption
	Filter func(*OrderQuery) (*OrderQuery, error)
}

// OrderPaginateOption enables pagination customization.
type OrderPaginateOption func(*OrderPager)

// DefaultOrderOrder is the default ordering of Order.
var DefaultOrderOrder = Desc(order.FieldID)

func newOrderPager(opts []OrderPaginateOption) (*OrderPager, error) {
	pager := &OrderPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultOrderOrder
	}
	return pager, nil
}

func (p *OrderPager) ApplyFilter(query *OrderQuery) (*OrderQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// OrderPageList is Order PageList result.
type OrderPageList struct {
	List        []*Order     `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (o *OrderQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...OrderPaginateOption,
) (*OrderPageList, error) {

	pager, err := newOrderPager(opts)
	if err != nil {
		return nil, err
	}

	if o, err = pager.ApplyFilter(o); err != nil {
		return nil, err
	}

	ret := &OrderPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := o.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		o = o.Order(pager.Order)
	} else {
		o = o.Order(DefaultOrderOrder)
	}

	o = o.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := o.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type OrderExtensionPager struct {
	Order  orderextension.OrderOption
	Filter func(*OrderExtensionQuery) (*OrderExtensionQuery, error)
}

// OrderExtensionPaginateOption enables pagination customization.
type OrderExtensionPaginateOption func(*OrderExtensionPager)

// DefaultOrderExtensionOrder is the default ordering of OrderExtension.
var DefaultOrderExtensionOrder = Desc(orderextension.FieldID)

func newOrderExtensionPager(opts []OrderExtensionPaginateOption) (*OrderExtensionPager, error) {
	pager := &OrderExtensionPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultOrderExtensionOrder
	}
	return pager, nil
}

func (p *OrderExtensionPager) ApplyFilter(query *OrderExtensionQuery) (*OrderExtensionQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// OrderExtensionPageList is OrderExtension PageList result.
type OrderExtensionPageList struct {
	List        []*OrderExtension `json:"list"`
	PageDetails *PageDetails      `json:"pageDetails"`
}

func (oe *OrderExtensionQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...OrderExtensionPaginateOption,
) (*OrderExtensionPageList, error) {

	pager, err := newOrderExtensionPager(opts)
	if err != nil {
		return nil, err
	}

	if oe, err = pager.ApplyFilter(oe); err != nil {
		return nil, err
	}

	ret := &OrderExtensionPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := oe.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		oe = oe.Order(pager.Order)
	} else {
		oe = oe.Order(DefaultOrderExtensionOrder)
	}

	oe = oe.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := oe.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}

type RefundPager struct {
	Order  refund.OrderOption
	Filter func(*RefundQuery) (*RefundQuery, error)
}

// RefundPaginateOption enables pagination customization.
type RefundPaginateOption func(*RefundPager)

// DefaultRefundOrder is the default ordering of Refund.
var DefaultRefundOrder = Desc(refund.FieldID)

func newRefundPager(opts []RefundPaginateOption) (*RefundPager, error) {
	pager := &RefundPager{}
	for _, opt := range opts {
		opt(pager)
	}
	if pager.Order == nil {
		pager.Order = DefaultRefundOrder
	}
	return pager, nil
}

func (p *RefundPager) ApplyFilter(query *RefundQuery) (*RefundQuery, error) {
	if p.Filter != nil {
		return p.Filter(query)
	}
	return query, nil
}

// RefundPageList is Refund PageList result.
type RefundPageList struct {
	List        []*Refund    `json:"list"`
	PageDetails *PageDetails `json:"pageDetails"`
}

func (r *RefundQuery) Page(
	ctx context.Context, pageNum uint64, pageSize uint64, opts ...RefundPaginateOption,
) (*RefundPageList, error) {

	pager, err := newRefundPager(opts)
	if err != nil {
		return nil, err
	}

	if r, err = pager.ApplyFilter(r); err != nil {
		return nil, err
	}

	ret := &RefundPageList{}

	ret.PageDetails = &PageDetails{
		Page: pageNum,
		Size: pageSize,
	}

	count, err := r.Clone().Count(ctx)

	if err != nil {
		return nil, err
	}

	ret.PageDetails.Total = uint64(count)

	if pager.Order != nil {
		r = r.Order(pager.Order)
	} else {
		r = r.Order(DefaultRefundOrder)
	}

	r = r.Offset(int((pageNum - 1) * pageSize)).Limit(int(pageSize))
	list, err := r.All(ctx)
	if err != nil {
		return nil, err
	}
	ret.List = list

	return ret, nil
}
