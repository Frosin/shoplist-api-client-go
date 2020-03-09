package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Shopping holds the schema definition for the Shopping entity.
type Shopping struct {
	ent.Schema
}

// Fields of the Shopping.
func (Shopping) Fields() []ent.Field {
	return []ent.Field{
		field.Time("date").Default(time.Now),
		field.Int("sum").Default(0),
		field.Bool("complete").Default(false),
	}
}

// Edges of the Shopping.
func (Shopping) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("item", Item.Type),
		edge.From("shop", Shop.Type).Ref("shopping").Unique(),
		edge.From("user", User.Type).Ref("shopping").Unique(),
	}
}
