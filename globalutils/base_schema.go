package globalutils

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"time"
)

func InitBaseSchema(fields []ent.Field) []ent.Field {
	return append(fields,
		field.String("created_by").NotEmpty(),
		field.Time("created_at").Default(time.Now()).Immutable(),
		field.String("updated_by").Optional(),
		field.Time("updated_at").Optional().Default(time.Now()),
		field.String("deleted_by").Optional(),
		field.Time("deleted_at").Optional(),
	)
}
