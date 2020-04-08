// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/Frosin/shoplist-api-client-go/ent/shopping"
	"github.com/Frosin/shoplist-api-client-go/ent/user"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetTelegramID sets the telegram_id field.
func (uc *UserCreate) SetTelegramID(i int64) *UserCreate {
	uc.mutation.SetTelegramID(i)
	return uc
}

// SetTelegramUsername sets the telegram_username field.
func (uc *UserCreate) SetTelegramUsername(s string) *UserCreate {
	uc.mutation.SetTelegramUsername(s)
	return uc
}

// SetComunityID sets the comunity_id field.
func (uc *UserCreate) SetComunityID(s string) *UserCreate {
	uc.mutation.SetComunityID(s)
	return uc
}

// SetToken sets the token field.
func (uc *UserCreate) SetToken(s string) *UserCreate {
	uc.mutation.SetToken(s)
	return uc
}

// SetChatID sets the chat_id field.
func (uc *UserCreate) SetChatID(i int64) *UserCreate {
	uc.mutation.SetChatID(i)
	return uc
}

// AddShoppingIDs adds the shopping edge to Shopping by ids.
func (uc *UserCreate) AddShoppingIDs(ids ...int) *UserCreate {
	uc.mutation.AddShoppingIDs(ids...)
	return uc
}

// AddShopping adds the shopping edges to Shopping.
func (uc *UserCreate) AddShopping(s ...*Shopping) *UserCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return uc.AddShoppingIDs(ids...)
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	if _, ok := uc.mutation.TelegramID(); !ok {
		return nil, errors.New("ent: missing required field \"telegram_id\"")
	}
	if _, ok := uc.mutation.TelegramUsername(); !ok {
		return nil, errors.New("ent: missing required field \"telegram_username\"")
	}
	if _, ok := uc.mutation.ComunityID(); !ok {
		return nil, errors.New("ent: missing required field \"comunity_id\"")
	}
	if v, ok := uc.mutation.ComunityID(); ok {
		if err := user.ComunityIDValidator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"comunity_id\": %v", err)
		}
	}
	if _, ok := uc.mutation.Token(); !ok {
		return nil, errors.New("ent: missing required field \"token\"")
	}
	if v, ok := uc.mutation.Token(); ok {
		if err := user.TokenValidator(v); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"token\": %v", err)
		}
	}
	if _, ok := uc.mutation.ChatID(); !ok {
		return nil, errors.New("ent: missing required field \"chat_id\"")
	}
	var (
		err  error
		node *User
	)
	if len(uc.hooks) == 0 {
		node, err = uc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uc.mutation = mutation
			node, err = uc.sqlSave(ctx)
			return node, err
		})
		for i := len(uc.hooks) - 1; i >= 0; i-- {
			mut = uc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	var (
		u     = &User{config: uc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: user.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		}
	)
	if value, ok := uc.mutation.TelegramID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: user.FieldTelegramID,
		})
		u.TelegramID = value
	}
	if value, ok := uc.mutation.TelegramUsername(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldTelegramUsername,
		})
		u.TelegramUsername = value
	}
	if value, ok := uc.mutation.ComunityID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldComunityID,
		})
		u.ComunityID = value
	}
	if value, ok := uc.mutation.Token(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: user.FieldToken,
		})
		u.Token = value
	}
	if value, ok := uc.mutation.ChatID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: user.FieldChatID,
		})
		u.ChatID = value
	}
	if nodes := uc.mutation.ShoppingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.ShoppingTable,
			Columns: []string{user.ShoppingColumn},
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
	if err := sqlgraph.CreateNode(ctx, uc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	u.ID = int(id)
	return u, nil
}
