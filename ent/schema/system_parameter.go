package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// System_parameter holds the schema definition for the System_parameter entity.
type System_parameter struct {
	ent.Schema
}

// Fields of the System_parameter.
func (System_parameter) Fields() []ent.Field {
	return []ent.Field{
		field.String("key").NotEmpty().Unique(),
		field.String("value").NotEmpty(),
	}
}

// Edges of the System_parameter.
func (System_parameter) Edges() []ent.Edge {
	return nil
}
