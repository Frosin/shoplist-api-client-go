// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/Frosin/shoplist-api-client-go/ent/predicate"
	"github.com/Frosin/shoplist-api-client-go/ent/shop"
	"github.com/Frosin/shoplist-api-client-go/ent/shopping"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// ShopUpdate is the builder for updating Shop entities.
type ShopUpdate struct {
	config
	hooks      []Hook
	mutation   *ShopMutation
	predicates []predicate.Shop
}

// Where adds a new predicate for the builder.
func (su *ShopUpdate) Where(ps ...predicate.Shop) *ShopUpdate {
	su.predicates = append(su.predicates, ps...)
	return su
}

// SetName sets the name field.
func (su *ShopUpdate) SetName(s string) *ShopUpdate {
	su.mutation.SetName(s)
	return su
}

// AddShoppingIDs adds the shopping edge to Shopping by ids.
func (su *ShopUpdate) AddShoppingIDs(ids ...int) *ShopUpdate {
	su.mutation.AddShoppingIDs(ids...)
	return su
}

// AddShopping adds the shopping edges to Shopping.
func (su *ShopUpdate) AddShopping(s ...*Shopping) *ShopUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.AddShoppingIDs(ids...)
}

// RemoveShoppingIDs removes the shopping edge to Shopping by ids.
func (su *ShopUpdate) RemoveShoppingIDs(ids ...int) *ShopUpdate {
	su.mutation.RemoveShoppingIDs(ids...)
	return su
}

// RemoveShopping removes shopping edges to Shopping.
func (su *ShopUpdate) RemoveShopping(s ...*Shopping) *ShopUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return su.RemoveShoppingIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (su *ShopUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := su.mutation.Name(); ok {
		if err := shop.NameValidator(v); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
		}
	}

	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ShopMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *ShopUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *ShopUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *ShopUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *ShopUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   shop.Table,
			Columns: shop.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: shop.FieldID,
			},
		},
	}
	if ps := su.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shop.FieldName,
		})
	}
	if nodes := su.mutation.RemovedShoppingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   shop.ShoppingTable,
			Columns: []string{shop.ShoppingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: shopping.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.ShoppingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   shop.ShoppingTable,
			Columns: []string{shop.ShoppingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: shopping.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{shop.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ShopUpdateOne is the builder for updating a single Shop entity.
type ShopUpdateOne struct {
	config
	hooks    []Hook
	mutation *ShopMutation
}

// SetName sets the name field.
func (suo *ShopUpdateOne) SetName(s string) *ShopUpdateOne {
	suo.mutation.SetName(s)
	return suo
}

// AddShoppingIDs adds the shopping edge to Shopping by ids.
func (suo *ShopUpdateOne) AddShoppingIDs(ids ...int) *ShopUpdateOne {
	suo.mutation.AddShoppingIDs(ids...)
	return suo
}

// AddShopping adds the shopping edges to Shopping.
func (suo *ShopUpdateOne) AddShopping(s ...*Shopping) *ShopUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.AddShoppingIDs(ids...)
}

// RemoveShoppingIDs removes the shopping edge to Shopping by ids.
func (suo *ShopUpdateOne) RemoveShoppingIDs(ids ...int) *ShopUpdateOne {
	suo.mutation.RemoveShoppingIDs(ids...)
	return suo
}

// RemoveShopping removes shopping edges to Shopping.
func (suo *ShopUpdateOne) RemoveShopping(s ...*Shopping) *ShopUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return suo.RemoveShoppingIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (suo *ShopUpdateOne) Save(ctx context.Context) (*Shop, error) {
	if v, ok := suo.mutation.Name(); ok {
		if err := shop.NameValidator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
		}
	}

	var (
		err  error
		node *Shop
	)
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ShopMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			mut = suo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, suo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (suo *ShopUpdateOne) SaveX(ctx context.Context) *Shop {
	s, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return s
}

// Exec executes the query on the entity.
func (suo *ShopUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *ShopUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *ShopUpdateOne) sqlSave(ctx context.Context) (s *Shop, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   shop.Table,
			Columns: shop.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: shop.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, fmt.Errorf("missing Shop.ID for update")
	}
	_spec.Node.ID.Value = id
	if value, ok := suo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shop.FieldName,
		})
	}
	if nodes := suo.mutation.RemovedShoppingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   shop.ShoppingTable,
			Columns: []string{shop.ShoppingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: shopping.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.ShoppingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   shop.ShoppingTable,
			Columns: []string{shop.ShoppingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: shopping.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	s = &Shop{config: suo.config}
	_spec.Assign = s.assignValues
	_spec.ScanValues = s.scanValues()
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{shop.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return s, nil
}
