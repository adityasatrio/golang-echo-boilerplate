package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"myapp/helper"
)

// Pet holds the schema definition for the Pet entity.
type Pet struct {
	ent.Schema
}

// Fields of the Pet.
func (Pet) Fields() []ent.Field {
	schema := []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("name").NotEmpty().MinLen(5).MaxLen(10),
		field.Enum("type").Values("DOG", "CAT"),
		field.String("code").NotEmpty().Unique(),
		field.Int("age_month").Positive(),
	}

	return helper.InitBaseSchema(schema)
}

// Edges of the Pet.
func (Pet) Edges() []ent.Edge {
	return nil
}

func (Pet) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name"),
		index.Fields("type"),
	}
}
