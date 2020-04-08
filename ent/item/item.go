// Code generated by entc, DO NOT EDIT.

package item

const (
	// Label holds the string label denoting the item type in the database.
	Label = "item"
	// FieldID holds the string denoting the id field in the database.
	FieldID          = "id"           // FieldProductName holds the string denoting the product_name vertex property in the database.
	FieldProductName = "product_name" // FieldQuantity holds the string denoting the quantity vertex property in the database.
	FieldQuantity    = "quantity"     // FieldCategoryID holds the string denoting the category_id vertex property in the database.
	FieldCategoryID  = "category_id"  // FieldComplete holds the string denoting the complete vertex property in the database.
	FieldComplete    = "complete"

	// EdgeShopping holds the string denoting the shopping edge name in mutations.
	EdgeShopping = "shopping"

	// Table holds the table name of the item in the database.
	Table = "items"
	// ShoppingTable is the table the holds the shopping relation/edge.
	ShoppingTable = "items"
	// ShoppingInverseTable is the table name for the Shopping entity.
	// It exists in this package in order to avoid circular dependency with the "shopping" package.
	ShoppingInverseTable = "shoppings"
	// ShoppingColumn is the table column denoting the shopping relation/edge.
	ShoppingColumn = "shopping_item"
)

// Columns holds all SQL columns for item fields.
var Columns = []string{
	FieldID,
	FieldProductName,
	FieldQuantity,
	FieldCategoryID,
	FieldComplete,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Item type.
var ForeignKeys = []string{
	"shopping_item",
}

var (
	// ProductNameValidator is a validator for the "product_name" field. It is called by the builders before save.
	ProductNameValidator func(string) error
	// DefaultQuantity holds the default value on creation for the quantity field.
	DefaultQuantity int
	// DefaultCategoryID holds the default value on creation for the category_id field.
	DefaultCategoryID int
	// DefaultComplete holds the default value on creation for the complete field.
	DefaultComplete bool
)
