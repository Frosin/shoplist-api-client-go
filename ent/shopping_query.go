// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"

	"github.com/Frosin/shoplist-api-client-go/ent/item"
	"github.com/Frosin/shoplist-api-client-go/ent/predicate"
	"github.com/Frosin/shoplist-api-client-go/ent/shop"
	"github.com/Frosin/shoplist-api-client-go/ent/shopping"
	"github.com/Frosin/shoplist-api-client-go/ent/user"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// ShoppingQuery is the builder for querying Shopping entities.
type ShoppingQuery struct {
	config
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.Shopping
	// eager-loading edges.
	withItem *ItemQuery
	withShop *ShopQuery
	withUser *UserQuery
	withFKs  bool
	// intermediate query.
	sql *sql.Selector
}

// Where adds a new predicate for the builder.
func (sq *ShoppingQuery) Where(ps ...predicate.Shopping) *ShoppingQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit adds a limit step to the query.
func (sq *ShoppingQuery) Limit(limit int) *ShoppingQuery {
	sq.limit = &limit
	return sq
}

// Offset adds an offset step to the query.
func (sq *ShoppingQuery) Offset(offset int) *ShoppingQuery {
	sq.offset = &offset
	return sq
}

// Order adds an order step to the query.
func (sq *ShoppingQuery) Order(o ...Order) *ShoppingQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QueryItem chains the current query on the item edge.
func (sq *ShoppingQuery) QueryItem() *ItemQuery {
	query := &ItemQuery{config: sq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(shopping.Table, shopping.FieldID, sq.sqlQuery()),
		sqlgraph.To(item.Table, item.FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, shopping.ItemTable, shopping.ItemColumn),
	)
	query.sql = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
	return query
}

// QueryShop chains the current query on the shop edge.
func (sq *ShoppingQuery) QueryShop() *ShopQuery {
	query := &ShopQuery{config: sq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(shopping.Table, shopping.FieldID, sq.sqlQuery()),
		sqlgraph.To(shop.Table, shop.FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, shopping.ShopTable, shopping.ShopColumn),
	)
	query.sql = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
	return query
}

// QueryUser chains the current query on the user edge.
func (sq *ShoppingQuery) QueryUser() *UserQuery {
	query := &UserQuery{config: sq.config}
	step := sqlgraph.NewStep(
		sqlgraph.From(shopping.Table, shopping.FieldID, sq.sqlQuery()),
		sqlgraph.To(user.Table, user.FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, shopping.UserTable, shopping.UserColumn),
	)
	query.sql = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
	return query
}

// First returns the first Shopping entity in the query. Returns *NotFoundError when no shopping was found.
func (sq *ShoppingQuery) First(ctx context.Context) (*Shopping, error) {
	sSlice, err := sq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(sSlice) == 0 {
		return nil, &NotFoundError{shopping.Label}
	}
	return sSlice[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *ShoppingQuery) FirstX(ctx context.Context) *Shopping {
	s, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return s
}

// FirstID returns the first Shopping id in the query. Returns *NotFoundError when no id was found.
func (sq *ShoppingQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{shopping.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (sq *ShoppingQuery) FirstXID(ctx context.Context) int {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only Shopping entity in the query, returns an error if not exactly one entity was returned.
func (sq *ShoppingQuery) Only(ctx context.Context) (*Shopping, error) {
	sSlice, err := sq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(sSlice) {
	case 1:
		return sSlice[0], nil
	case 0:
		return nil, &NotFoundError{shopping.Label}
	default:
		return nil, &NotSingularError{shopping.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *ShoppingQuery) OnlyX(ctx context.Context) *Shopping {
	s, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return s
}

// OnlyID returns the only Shopping id in the query, returns an error if not exactly one id was returned.
func (sq *ShoppingQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{shopping.Label}
	default:
		err = &NotSingularError{shopping.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (sq *ShoppingQuery) OnlyXID(ctx context.Context) int {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Shoppings.
func (sq *ShoppingQuery) All(ctx context.Context) ([]*Shopping, error) {
	return sq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (sq *ShoppingQuery) AllX(ctx context.Context) []*Shopping {
	sSlice, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return sSlice
}

// IDs executes the query and returns a list of Shopping ids.
func (sq *ShoppingQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := sq.Select(shopping.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *ShoppingQuery) IDsX(ctx context.Context) []int {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *ShoppingQuery) Count(ctx context.Context) (int, error) {
	return sq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (sq *ShoppingQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *ShoppingQuery) Exist(ctx context.Context) (bool, error) {
	return sq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *ShoppingQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *ShoppingQuery) Clone() *ShoppingQuery {
	return &ShoppingQuery{
		config:     sq.config,
		limit:      sq.limit,
		offset:     sq.offset,
		order:      append([]Order{}, sq.order...),
		unique:     append([]string{}, sq.unique...),
		predicates: append([]predicate.Shopping{}, sq.predicates...),
		// clone intermediate query.
		sql: sq.sql.Clone(),
	}
}

//  WithItem tells the query-builder to eager-loads the nodes that are connected to
// the "item" edge. The optional arguments used to configure the query builder of the edge.
func (sq *ShoppingQuery) WithItem(opts ...func(*ItemQuery)) *ShoppingQuery {
	query := &ItemQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withItem = query
	return sq
}

//  WithShop tells the query-builder to eager-loads the nodes that are connected to
// the "shop" edge. The optional arguments used to configure the query builder of the edge.
func (sq *ShoppingQuery) WithShop(opts ...func(*ShopQuery)) *ShoppingQuery {
	query := &ShopQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withShop = query
	return sq
}

//  WithUser tells the query-builder to eager-loads the nodes that are connected to
// the "user" edge. The optional arguments used to configure the query builder of the edge.
func (sq *ShoppingQuery) WithUser(opts ...func(*UserQuery)) *ShoppingQuery {
	query := &UserQuery{config: sq.config}
	for _, opt := range opts {
		opt(query)
	}
	sq.withUser = query
	return sq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Date time.Time `json:"date,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Shopping.Query().
//		GroupBy(shopping.FieldDate).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (sq *ShoppingQuery) GroupBy(field string, fields ...string) *ShoppingGroupBy {
	group := &ShoppingGroupBy{config: sq.config}
	group.fields = append([]string{field}, fields...)
	group.sql = sq.sqlQuery()
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Date time.Time `json:"date,omitempty"`
//	}
//
//	client.Shopping.Query().
//		Select(shopping.FieldDate).
//		Scan(ctx, &v)
//
func (sq *ShoppingQuery) Select(field string, fields ...string) *ShoppingSelect {
	selector := &ShoppingSelect{config: sq.config}
	selector.fields = append([]string{field}, fields...)
	selector.sql = sq.sqlQuery()
	return selector
}

func (sq *ShoppingQuery) sqlAll(ctx context.Context) ([]*Shopping, error) {
	var (
		nodes       = []*Shopping{}
		withFKs     = sq.withFKs
		_spec       = sq.querySpec()
		loadedTypes = [3]bool{
			sq.withItem != nil,
			sq.withShop != nil,
			sq.withUser != nil,
		}
	)
	if sq.withShop != nil || sq.withUser != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, shopping.ForeignKeys...)
	}
	_spec.ScanValues = func() []interface{} {
		node := &Shopping{config: sq.config}
		nodes = append(nodes, node)
		values := node.scanValues()
		if withFKs {
			values = append(values, node.fkValues()...)
		}
		return values
	}
	_spec.Assign = func(values ...interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(values...)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := sq.withItem; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int]*Shopping)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
		}
		query.withFKs = true
		query.Where(predicate.Item(func(s *sql.Selector) {
			s.Where(sql.InValues(shopping.ItemColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.shopping_item
			if fk == nil {
				return nil, fmt.Errorf(`foreign-key "shopping_item" is nil for node %v`, n.ID)
			}
			node, ok := nodeids[*fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "shopping_item" returned %v for node %v`, *fk, n.ID)
			}
			node.Edges.Item = append(node.Edges.Item, n)
		}
	}

	if query := sq.withShop; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Shopping)
		for i := range nodes {
			if fk := nodes[i].shop_shopping; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(shop.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "shop_shopping" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Shop = n
			}
		}
	}

	if query := sq.withUser; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*Shopping)
		for i := range nodes {
			if fk := nodes[i].user_shopping; fk != nil {
				ids = append(ids, *fk)
				nodeids[*fk] = append(nodeids[*fk], nodes[i])
			}
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_shopping" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.User = n
			}
		}
	}

	return nodes, nil
}

func (sq *ShoppingQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *ShoppingQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := sq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %v", err)
	}
	return n > 0, nil
}

func (sq *ShoppingQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   shopping.Table,
			Columns: shopping.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: shopping.FieldID,
			},
		},
		From:   sq.sql,
		Unique: true,
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *ShoppingQuery) sqlQuery() *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(shopping.Table)
	selector := builder.Select(t1.Columns(shopping.Columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(shopping.Columns...)...)
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ShoppingGroupBy is the builder for group-by Shopping entities.
type ShoppingGroupBy struct {
	config
	fields []string
	fns    []Aggregate
	// intermediate query.
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *ShoppingGroupBy) Aggregate(fns ...Aggregate) *ShoppingGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the group-by query and scan the result into the given value.
func (sgb *ShoppingGroupBy) Scan(ctx context.Context, v interface{}) error {
	return sgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (sgb *ShoppingGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := sgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (sgb *ShoppingGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(sgb.fields) > 1 {
		return nil, errors.New("ent: ShoppingGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := sgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (sgb *ShoppingGroupBy) StringsX(ctx context.Context) []string {
	v, err := sgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (sgb *ShoppingGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(sgb.fields) > 1 {
		return nil, errors.New("ent: ShoppingGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := sgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (sgb *ShoppingGroupBy) IntsX(ctx context.Context) []int {
	v, err := sgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (sgb *ShoppingGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(sgb.fields) > 1 {
		return nil, errors.New("ent: ShoppingGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := sgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (sgb *ShoppingGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := sgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (sgb *ShoppingGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(sgb.fields) > 1 {
		return nil, errors.New("ent: ShoppingGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := sgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (sgb *ShoppingGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := sgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sgb *ShoppingGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := sgb.sqlQuery().Query()
	if err := sgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (sgb *ShoppingGroupBy) sqlQuery() *sql.Selector {
	selector := sgb.sql
	columns := make([]string, 0, len(sgb.fields)+len(sgb.fns))
	columns = append(columns, sgb.fields...)
	for _, fn := range sgb.fns {
		columns = append(columns, fn(selector))
	}
	return selector.Select(columns...).GroupBy(sgb.fields...)
}

// ShoppingSelect is the builder for select fields of Shopping entities.
type ShoppingSelect struct {
	config
	fields []string
	// intermediate queries.
	sql *sql.Selector
}

// Scan applies the selector query and scan the result into the given value.
func (ss *ShoppingSelect) Scan(ctx context.Context, v interface{}) error {
	return ss.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ss *ShoppingSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ss.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (ss *ShoppingSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ss.fields) > 1 {
		return nil, errors.New("ent: ShoppingSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ss *ShoppingSelect) StringsX(ctx context.Context) []string {
	v, err := ss.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (ss *ShoppingSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ss.fields) > 1 {
		return nil, errors.New("ent: ShoppingSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ss *ShoppingSelect) IntsX(ctx context.Context) []int {
	v, err := ss.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (ss *ShoppingSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ss.fields) > 1 {
		return nil, errors.New("ent: ShoppingSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ss *ShoppingSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ss.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (ss *ShoppingSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ss.fields) > 1 {
		return nil, errors.New("ent: ShoppingSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ss.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ss *ShoppingSelect) BoolsX(ctx context.Context) []bool {
	v, err := ss.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ss *ShoppingSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ss.sqlQuery().Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ss *ShoppingSelect) sqlQuery() sql.Querier {
	selector := ss.sql
	selector.Select(selector.Columns(ss.fields...)...)
	return selector
}
