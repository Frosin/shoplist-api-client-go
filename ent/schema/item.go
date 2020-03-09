package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Item holds the schema definition for the Item entity.
type Item struct {
	ent.Schema
}

// Fields of the Item.
func (Item) Fields() []ent.Field {
	return []ent.Field{
		field.String("product_name").NotEmpty(),
		field.Int("quantity").Default(1),
		field.Int("category_id").Default(0),
		field.Bool("complete").Default(false),
	}
}

// Edges of the Item.
func (Item) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("shopping", Shopping.Type).Ref("item").Unique(),
	}
}
