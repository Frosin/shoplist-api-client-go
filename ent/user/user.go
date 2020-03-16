// Code generated by entc, DO NOT EDIT.

package user

import (
	"github.com/Frosin/shoplist-api-client-go/ent/schema"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTelegramID holds the string denoting the telegram_id vertex property in the database.
	FieldTelegramID = "telegram_id"
	// FieldTelegramUsername holds the string denoting the telegram_username vertex property in the database.
	FieldTelegramUsername = "telegram_username"
	// FieldComunityID holds the string denoting the comunity_id vertex property in the database.
	FieldComunityID = "comunity_id"
	// FieldToken holds the string denoting the token vertex property in the database.
	FieldToken = "token"
	// FieldChatID holds the string denoting the chat_id vertex property in the database.
	FieldChatID = "chat_id"

	// Table holds the table name of the user in the database.
	Table = "users"
	// ShoppingTable is the table the holds the shopping relation/edge.
	ShoppingTable = "shoppings"
	// ShoppingInverseTable is the table name for the Shopping entity.
	// It exists in this package in order to avoid circular dependency with the "shopping" package.
	ShoppingInverseTable = "shoppings"
	// ShoppingColumn is the table column denoting the shopping relation/edge.
	ShoppingColumn = "user_id"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldTelegramID,
	FieldTelegramUsername,
	FieldComunityID,
	FieldToken,
	FieldChatID,
}

var (
	fields = schema.User{}.Fields()

	// descTelegramUsername is the schema descriptor for telegram_username field.
	descTelegramUsername = fields[1].Descriptor()
	// TelegramUsernameValidator is a validator for the "telegram_username" field. It is called by the builders before save.
	TelegramUsernameValidator = descTelegramUsername.Validators[0].(func(string) error)

	// descComunityID is the schema descriptor for comunity_id field.
	descComunityID = fields[2].Descriptor()
	// ComunityIDValidator is a validator for the "comunity_id" field. It is called by the builders before save.
	ComunityIDValidator = descComunityID.Validators[0].(func(string) error)

	// descToken is the schema descriptor for token field.
	descToken = fields[3].Descriptor()
	// TokenValidator is a validator for the "token" field. It is called by the builders before save.
	TokenValidator = descToken.Validators[0].(func(string) error)
)