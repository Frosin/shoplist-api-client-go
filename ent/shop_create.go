// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/Frosin/shoplist-api-client-go/ent/shop"
	"github.com/Frosin/shoplist-api-client-go/ent/shopping"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// ShopCreate is the builder for creating a Shop entity.
type ShopCreate struct {
	config
	mutation *ShopMutation
	hooks    []Hook
}

// SetName sets the name field.
func (sc *ShopCreate) SetName(s string) *ShopCreate {
	sc.mutation.SetName(s)
	return sc
}

// AddShoppingIDs adds the shopping edge to Shopping by ids.
func (sc *ShopCreate) AddShoppingIDs(ids ...int) *ShopCreate {
	sc.mutation.AddShoppingIDs(ids...)
	return sc
}

// AddShopping adds the shopping edges to Shopping.
func (sc *ShopCreate) AddShopping(s ...*Shopping) *ShopCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return sc.AddShoppingIDs(ids...)
}

// Save creates the Shop in the database.
func (sc *ShopCreate) Save(ctx context.Context) (*Shop, error) {
	if _, ok := sc.mutation.Name(); !ok {
		return nil, errors.New("ent: missing required field \"name\"")
	}
	if v, ok := sc.mutation.Name(); ok {
		if err := shop.NameValidator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
		}
	}
	var (
		err  error
		node *Shop
	)
	if len(sc.hooks) == 0 {
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ShopMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			sc.mutation = mutation
			node, err = sc.sqlSave(ctx)
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *ShopCreate) SaveX(ctx context.Context) *Shop {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sc *ShopCreate) sqlSave(ctx context.Context) (*Shop, error) {
	var (
		s     = &Shop{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: shop.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: shop.FieldID,
			},
		}
	)
	if value, ok := sc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shop.FieldName,
		})
		s.Name = value
	}
	if nodes := sc.mutation.ShoppingIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	s.ID = int(id)
	return s, nil
}
