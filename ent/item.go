// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/Frosin/shoplist-api-client-go/ent/item"
	"github.com/Frosin/shoplist-api-client-go/ent/shopping"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Item is the model entity for the Item schema.
type Item struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// ProductName holds the value of the "product_name" field.
	ProductName string `json:"product_name,omitempty"`
	// Quantity holds the value of the "quantity" field.
	Quantity int `json:"quantity,omitempty"`
	// CategoryID holds the value of the "category_id" field.
	CategoryID int `json:"category_id,omitempty"`
	// Complete holds the value of the "complete" field.
	Complete bool `json:"complete,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ItemQuery when eager-loading is set.
	Edges       ItemEdges `json:"edges"`
	shopping_id *int
}

// ItemEdges holds the relations/edges for other nodes in the graph.
type ItemEdges struct {
	// Shopping holds the value of the shopping edge.
	Shopping *Shopping
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// ShoppingOrErr returns the Shopping value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ItemEdges) ShoppingOrErr() (*Shopping, error) {
	if e.loadedTypes[0] {
		if e.Shopping == nil {
			// The edge shopping was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: shopping.Label}
		}
		return e.Shopping, nil
	}
	return nil, &NotLoadedError{edge: "shopping"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Item) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // product_name
		&sql.NullInt64{},  // quantity
		&sql.NullInt64{},  // category_id
		&sql.NullBool{},   // complete
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Item) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // shopping_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Item fields.
func (i *Item) assignValues(values ...interface{}) error {
	if m, n := len(values), len(item.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	i.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field product_name", values[0])
	} else if value.Valid {
		i.ProductName = value.String
	}
	if value, ok := values[1].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field quantity", values[1])
	} else if value.Valid {
		i.Quantity = int(value.Int64)
	}
	if value, ok := values[2].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field category_id", values[2])
	} else if value.Valid {
		i.CategoryID = int(value.Int64)
	}
	if value, ok := values[3].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field complete", values[3])
	} else if value.Valid {
		i.Complete = value.Bool
	}
	values = values[4:]
	if len(values) == len(item.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field shopping_id", value)
		} else if value.Valid {
			i.shopping_id = new(int)
			*i.shopping_id = int(value.Int64)
		}
	}
	return nil
}

// QueryShopping queries the shopping edge of the Item.
func (i *Item) QueryShopping() *ShoppingQuery {
	return (&ItemClient{i.config}).QueryShopping(i)
}

// Update returns a builder for updating this Item.
// Note that, you need to call Item.Unwrap() before calling this method, if this Item
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Item) Update() *ItemUpdateOne {
	return (&ItemClient{i.config}).UpdateOne(i)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (i *Item) Unwrap() *Item {
	tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Item is not a transactional entity")
	}
	i.config.driver = tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Item) String() string {
	var builder strings.Builder
	builder.WriteString("Item(")
	builder.WriteString(fmt.Sprintf("id=%v", i.ID))
	builder.WriteString(", product_name=")
	builder.WriteString(i.ProductName)
	builder.WriteString(", quantity=")
	builder.WriteString(fmt.Sprintf("%v", i.Quantity))
	builder.WriteString(", category_id=")
	builder.WriteString(fmt.Sprintf("%v", i.CategoryID))
	builder.WriteString(", complete=")
	builder.WriteString(fmt.Sprintf("%v", i.Complete))
	builder.WriteByte(')')
	return builder.String()
}

// Items is a parsable slice of Item.
type Items []*Item

func (i Items) config(cfg config) {
	for _i := range i {
		i[_i].config = cfg
	}
}
