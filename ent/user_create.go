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
	telegram_id       *int64
	telegram_username *string
	comunity_id       *string
	token             *string
	chat_id           *int64
	shopping          map[int]struct{}
}

// SetTelegramID sets the telegram_id field.
func (uc *UserCreate) SetTelegramID(i int64) *UserCreate {
	uc.telegram_id = &i
	return uc
}

// SetTelegramUsername sets the telegram_username field.
func (uc *UserCreate) SetTelegramUsername(s string) *UserCreate {
	uc.telegram_username = &s
	return uc
}

// SetComunityID sets the comunity_id field.
func (uc *UserCreate) SetComunityID(s string) *UserCreate {
	uc.comunity_id = &s
	return uc
}

// SetToken sets the token field.
func (uc *UserCreate) SetToken(s string) *UserCreate {
	uc.token = &s
	return uc
}

// SetChatID sets the chat_id field.
func (uc *UserCreate) SetChatID(i int64) *UserCreate {
	uc.chat_id = &i
	return uc
}

// AddShoppingIDs adds the shopping edge to Shopping by ids.
func (uc *UserCreate) AddShoppingIDs(ids ...int) *UserCreate {
	if uc.shopping == nil {
		uc.shopping = make(map[int]struct{})
	}
	for i := range ids {
		uc.shopping[ids[i]] = struct{}{}
	}
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
	if uc.telegram_id == nil {
		return nil, errors.New("ent: missing required field \"telegram_id\"")
	}
	if uc.telegram_username == nil {
		return nil, errors.New("ent: missing required field \"telegram_username\"")
	}
	if err := user.TelegramUsernameValidator(*uc.telegram_username); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"telegram_username\": %v", err)
	}
	if uc.comunity_id == nil {
		return nil, errors.New("ent: missing required field \"comunity_id\"")
	}
	if err := user.ComunityIDValidator(*uc.comunity_id); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"comunity_id\": %v", err)
	}
	if uc.token == nil {
		return nil, errors.New("ent: missing required field \"token\"")
	}
	if err := user.TokenValidator(*uc.token); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"token\": %v", err)
	}
	if uc.chat_id == nil {
		return nil, errors.New("ent: missing required field \"chat_id\"")
	}
	return uc.sqlSave(ctx)
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
	if value := uc.telegram_id; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  *value,
			Column: user.FieldTelegramID,
		})
		u.TelegramID = *value
	}
	if value := uc.telegram_username; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldTelegramUsername,
		})
		u.TelegramUsername = *value
	}
	if value := uc.comunity_id; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldComunityID,
		})
		u.ComunityID = *value
	}
	if value := uc.token; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: user.FieldToken,
		})
		u.Token = *value
	}
	if value := uc.chat_id; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  *value,
			Column: user.FieldChatID,
		})
		u.ChatID = *value
	}
	if nodes := uc.shopping; len(nodes) > 0 {
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
		for k, _ := range nodes {
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