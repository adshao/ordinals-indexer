// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"github.com/adshao/ordinals-indexer/internal/data/ent/inscription"
	"github.com/adshao/ordinals-indexer/internal/data/ent/predicate"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// InscriptionQuery is the builder for querying Inscription entities.
type InscriptionQuery struct {
	config
	ctx        *QueryContext
	order      []inscription.OrderOption
	inters     []Interceptor
	predicates []predicate.Inscription
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the InscriptionQuery builder.
func (iq *InscriptionQuery) Where(ps ...predicate.Inscription) *InscriptionQuery {
	iq.predicates = append(iq.predicates, ps...)
	return iq
}

// Limit the number of records to be returned by this query.
func (iq *InscriptionQuery) Limit(limit int) *InscriptionQuery {
	iq.ctx.Limit = &limit
	return iq
}

// Offset to start from.
func (iq *InscriptionQuery) Offset(offset int) *InscriptionQuery {
	iq.ctx.Offset = &offset
	return iq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (iq *InscriptionQuery) Unique(unique bool) *InscriptionQuery {
	iq.ctx.Unique = &unique
	return iq
}

// Order specifies how the records should be ordered.
func (iq *InscriptionQuery) Order(o ...inscription.OrderOption) *InscriptionQuery {
	iq.order = append(iq.order, o...)
	return iq
}

// First returns the first Inscription entity from the query.
// Returns a *NotFoundError when no Inscription was found.
func (iq *InscriptionQuery) First(ctx context.Context) (*Inscription, error) {
	nodes, err := iq.Limit(1).All(setContextOp(ctx, iq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{inscription.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (iq *InscriptionQuery) FirstX(ctx context.Context) *Inscription {
	node, err := iq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Inscription ID from the query.
// Returns a *NotFoundError when no Inscription ID was found.
func (iq *InscriptionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = iq.Limit(1).IDs(setContextOp(ctx, iq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{inscription.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (iq *InscriptionQuery) FirstIDX(ctx context.Context) int {
	id, err := iq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Inscription entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Inscription entity is found.
// Returns a *NotFoundError when no Inscription entities are found.
func (iq *InscriptionQuery) Only(ctx context.Context) (*Inscription, error) {
	nodes, err := iq.Limit(2).All(setContextOp(ctx, iq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{inscription.Label}
	default:
		return nil, &NotSingularError{inscription.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (iq *InscriptionQuery) OnlyX(ctx context.Context) *Inscription {
	node, err := iq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Inscription ID in the query.
// Returns a *NotSingularError when more than one Inscription ID is found.
// Returns a *NotFoundError when no entities are found.
func (iq *InscriptionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = iq.Limit(2).IDs(setContextOp(ctx, iq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{inscription.Label}
	default:
		err = &NotSingularError{inscription.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (iq *InscriptionQuery) OnlyIDX(ctx context.Context) int {
	id, err := iq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Inscriptions.
func (iq *InscriptionQuery) All(ctx context.Context) ([]*Inscription, error) {
	ctx = setContextOp(ctx, iq.ctx, "All")
	if err := iq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Inscription, *InscriptionQuery]()
	return withInterceptors[[]*Inscription](ctx, iq, qr, iq.inters)
}

// AllX is like All, but panics if an error occurs.
func (iq *InscriptionQuery) AllX(ctx context.Context) []*Inscription {
	nodes, err := iq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Inscription IDs.
func (iq *InscriptionQuery) IDs(ctx context.Context) (ids []int, err error) {
	if iq.ctx.Unique == nil && iq.path != nil {
		iq.Unique(true)
	}
	ctx = setContextOp(ctx, iq.ctx, "IDs")
	if err = iq.Select(inscription.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (iq *InscriptionQuery) IDsX(ctx context.Context) []int {
	ids, err := iq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (iq *InscriptionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, iq.ctx, "Count")
	if err := iq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, iq, querierCount[*InscriptionQuery](), iq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (iq *InscriptionQuery) CountX(ctx context.Context) int {
	count, err := iq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (iq *InscriptionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, iq.ctx, "Exist")
	switch _, err := iq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (iq *InscriptionQuery) ExistX(ctx context.Context) bool {
	exist, err := iq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the InscriptionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (iq *InscriptionQuery) Clone() *InscriptionQuery {
	if iq == nil {
		return nil
	}
	return &InscriptionQuery{
		config:     iq.config,
		ctx:        iq.ctx.Clone(),
		order:      append([]inscription.OrderOption{}, iq.order...),
		inters:     append([]Interceptor{}, iq.inters...),
		predicates: append([]predicate.Inscription{}, iq.predicates...),
		// clone intermediate query.
		sql:  iq.sql.Clone(),
		path: iq.path,
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
//	client.Inscription.Query().
//		GroupBy(inscription.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (iq *InscriptionQuery) GroupBy(field string, fields ...string) *InscriptionGroupBy {
	iq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &InscriptionGroupBy{build: iq}
	grbuild.flds = &iq.ctx.Fields
	grbuild.label = inscription.Label
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
//	client.Inscription.Query().
//		Select(inscription.FieldCreatedAt).
//		Scan(ctx, &v)
func (iq *InscriptionQuery) Select(fields ...string) *InscriptionSelect {
	iq.ctx.Fields = append(iq.ctx.Fields, fields...)
	sbuild := &InscriptionSelect{InscriptionQuery: iq}
	sbuild.label = inscription.Label
	sbuild.flds, sbuild.scan = &iq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a InscriptionSelect configured with the given aggregations.
func (iq *InscriptionQuery) Aggregate(fns ...AggregateFunc) *InscriptionSelect {
	return iq.Select().Aggregate(fns...)
}

func (iq *InscriptionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range iq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, iq); err != nil {
				return err
			}
		}
	}
	for _, f := range iq.ctx.Fields {
		if !inscription.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if iq.path != nil {
		prev, err := iq.path(ctx)
		if err != nil {
			return err
		}
		iq.sql = prev
	}
	return nil
}

func (iq *InscriptionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Inscription, error) {
	var (
		nodes = []*Inscription{}
		_spec = iq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Inscription).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Inscription{config: iq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, iq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (iq *InscriptionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := iq.querySpec()
	_spec.Node.Columns = iq.ctx.Fields
	if len(iq.ctx.Fields) > 0 {
		_spec.Unique = iq.ctx.Unique != nil && *iq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, iq.driver, _spec)
}

func (iq *InscriptionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(inscription.Table, inscription.Columns, sqlgraph.NewFieldSpec(inscription.FieldID, field.TypeInt))
	_spec.From = iq.sql
	if unique := iq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if iq.path != nil {
		_spec.Unique = true
	}
	if fields := iq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, inscription.FieldID)
		for i := range fields {
			if fields[i] != inscription.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := iq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := iq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := iq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := iq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (iq *InscriptionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(iq.driver.Dialect())
	t1 := builder.Table(inscription.Table)
	columns := iq.ctx.Fields
	if len(columns) == 0 {
		columns = inscription.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if iq.sql != nil {
		selector = iq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if iq.ctx.Unique != nil && *iq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range iq.predicates {
		p(selector)
	}
	for _, p := range iq.order {
		p(selector)
	}
	if offset := iq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := iq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// InscriptionGroupBy is the group-by builder for Inscription entities.
type InscriptionGroupBy struct {
	selector
	build *InscriptionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (igb *InscriptionGroupBy) Aggregate(fns ...AggregateFunc) *InscriptionGroupBy {
	igb.fns = append(igb.fns, fns...)
	return igb
}

// Scan applies the selector query and scans the result into the given value.
func (igb *InscriptionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, igb.build.ctx, "GroupBy")
	if err := igb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*InscriptionQuery, *InscriptionGroupBy](ctx, igb.build, igb, igb.build.inters, v)
}

func (igb *InscriptionGroupBy) sqlScan(ctx context.Context, root *InscriptionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(igb.fns))
	for _, fn := range igb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*igb.flds)+len(igb.fns))
		for _, f := range *igb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*igb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := igb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// InscriptionSelect is the builder for selecting fields of Inscription entities.
type InscriptionSelect struct {
	*InscriptionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (is *InscriptionSelect) Aggregate(fns ...AggregateFunc) *InscriptionSelect {
	is.fns = append(is.fns, fns...)
	return is
}

// Scan applies the selector query and scans the result into the given value.
func (is *InscriptionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, is.ctx, "Select")
	if err := is.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*InscriptionQuery, *InscriptionSelect](ctx, is.InscriptionQuery, is, is.inters, v)
}

func (is *InscriptionSelect) sqlScan(ctx context.Context, root *InscriptionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(is.fns))
	for _, fn := range is.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*is.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := is.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}