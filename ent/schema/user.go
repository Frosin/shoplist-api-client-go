package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("telegram_id").Immutable(),
		field.String("telegram_username").NotEmpty(),
		field.String("comunity_id").NotEmpty(),
		field.String("token").NotEmpty().Immutable(),
		field.Int64("chat_id").Immutable(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("shopping", Shopping.Type),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		// non-unique index.
		index.Fields(
			"telegram_id",
			"comunity_id",
			"token",
		),
	}
}
