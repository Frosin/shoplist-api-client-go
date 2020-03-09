// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/Frosin/shoplist-api-client-go/ent/predicate"
	"github.com/Frosin/shoplist-api-client-go/ent/shop"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// ShopDelete is the builder for deleting a Shop entity.
type ShopDelete struct {
	config
	predicates []predicate.Shop
}

// Where adds a new predicate to the delete builder.
func (sd *ShopDelete) Where(ps ...predicate.Shop) *ShopDelete {
	sd.predicates = append(sd.predicates, ps...)
	return sd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sd *ShopDelete) Exec(ctx context.Context) (int, error) {
	return sd.sqlExec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (sd *ShopDelete) ExecX(ctx context.Context) int {
	n, err := sd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sd *ShopDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: shop.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: shop.FieldID,
			},
		},
	}
	if ps := sd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, sd.driver, _spec)
}

// ShopDeleteOne is the builder for deleting a single Shop entity.
type ShopDeleteOne struct {
	sd *ShopDelete
}

// Exec executes the deletion query.
func (sdo *ShopDeleteOne) Exec(ctx context.Context) error {
	n, err := sdo.sd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{shop.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sdo *ShopDeleteOne) ExecX(ctx context.Context) {
	sdo.sd.ExecX(ctx)
}
