// Code generated by entc, DO NOT EDIT.

package shop

import (
	"github.com/Frosin/shoplist-api-client-go/ent/schema"
)

const (
	// Label holds the string label denoting the shop type in the database.
	Label = "shop"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name vertex property in the database.
	FieldName = "name"

	// Table holds the table name of the shop in the database.
	Table = "shops"
	// ShoppingTable is the table the holds the shopping relation/edge.
	ShoppingTable = "shoppings"
	// ShoppingInverseTable is the table name for the Shopping entity.
	// It exists in this package in order to avoid circular dependency with the "shopping" package.
	ShoppingInverseTable = "shoppings"
	// ShoppingColumn is the table column denoting the shopping relation/edge.
	ShoppingColumn = "shop_id"
)

// Columns holds all SQL columns for shop fields.
var Columns = []string{
	FieldID,
	FieldName,
}

var (
	fields = schema.Shop{}.Fields()

	// descName is the schema descriptor for name field.
	descName = fields[0].Descriptor()
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator = descName.Validators[0].(func(string) error)
)
