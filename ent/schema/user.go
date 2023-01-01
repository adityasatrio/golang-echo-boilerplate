package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"myapp/helper"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	schema := []ent.Field{
		field.String("name").NotEmpty().Unique(),
		field.String("value").NotEmpty(),
	}

	return helper.InitBaseSchema(schema)
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}
