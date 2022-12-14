package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"myapp/helper"
)

// System_parameter holds the schema definition for the System_parameter entity.
type System_parameter struct {
	ent.Schema
}

// Fields of the System_parameter.
func (System_parameter) Fields() []ent.Field {
	schema := []ent.Field{
		field.String("key").NotEmpty().Unique(),
		field.String("value").NotEmpty(),
	}

	return helper.InitBaseSchema(schema)
}

// Edges of the System_parameter.
func (System_parameter) Edges() []ent.Edge {
	return nil
}

func (System_parameter) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("key"),
		index.Fields("value"),
	}
}
