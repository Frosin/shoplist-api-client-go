// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/Frosin/shoplist-api-client-go/ent/shop"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Shop is the model entity for the Shop schema.
type Shop struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ShopQuery when eager-loading is set.
	Edges ShopEdges `json:"edges"`
}

// ShopEdges holds the relations/edges for other nodes in the graph.
type ShopEdges struct {
	// Shopping holds the value of the shopping edge.
	Shopping []*Shopping
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ShoppingOrErr returns the Shopping value or an error if the edge
// was not loaded in eager-loading.
func (e ShopEdges) ShoppingOrErr() ([]*Shopping, error) {
	if e.loadedTypes[0] {
		return e.Shopping, nil
	}
	return nil, &NotLoadedError{edge: "shopping"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Shop) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Shop fields.
func (s *Shop) assignValues(values ...interface{}) error {
	if m, n := len(values), len(shop.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	s.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		s.Name = value.String
	}
	return nil
}

// QueryShopping queries the shopping edge of the Shop.
func (s *Shop) QueryShopping() *ShoppingQuery {
	return (&ShopClient{s.config}).QueryShopping(s)
}

// Update returns a builder for updating this Shop.
// Note that, you need to call Shop.Unwrap() before calling this method, if this Shop
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Shop) Update() *ShopUpdateOne {
	return (&ShopClient{s.config}).UpdateOne(s)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (s *Shop) Unwrap() *Shop {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Shop is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Shop) String() string {
	var builder strings.Builder
	builder.WriteString("Shop(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", name=")
	builder.WriteString(s.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Shops is a parsable slice of Shop.
type Shops []*Shop

func (s Shops) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}