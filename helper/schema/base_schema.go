package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

func InitBaseSchema(fields []ent.Field) []ent.Field {
	return append(fields,
		field.Bool("is_deleted").Default(false),
		field.String("created_by").NotEmpty(),
		field.Time("created_at").Default(time.Now()),
		field.String("updated_by").Optional(),
		field.Time("updated_at").Optional().Default(time.Now()),
	)
}
