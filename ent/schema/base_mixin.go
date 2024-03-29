package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
	"time"
)

// BaseFieldMixin type struct implements the ent.Mixin for sharing  base fields with package schemas.
type BaseFieldMixin struct {
	// We embed the `mixin.Schema` to avoid
	// implementing the rest of the methods.
	mixin.Schema
}

func (BaseFieldMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("created_by").NotEmpty(),
		field.Time("created_at").Immutable().Default(time.Now),

		field.String("updated_by").Optional(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),

		field.String("deleted_by").Optional(),
		field.Time("deleted_at").Optional(),
	}
}

// VersionMixin provides an optimistic concurrency
// control mechanism using a "versions" field.
type VersionMixin struct {
	mixin.Schema
}

// Fields of the VersionMixin.
func (VersionMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("versions").
			DefaultFunc(time.Now().UnixNano()).
			Comment("Unix time of when the latest update occurred"),
	}
}
