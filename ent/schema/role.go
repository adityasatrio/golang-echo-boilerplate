// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Role struct {
	ent.Schema
}

// Mixin of the User.
func (Role) Mixin() []ent.Mixin {
	return []ent.Mixin{
		VersionMixin{},
		BaseFieldMixin{},
	}
}

func (Role) Fields() []ent.Field {
	return []ent.Field{field.Uint64("id"),
		field.String("name"),
		field.String("text"),
	}
}
func (Role) Edges() []ent.Edge {
	return nil
}
func (Role) Annotations() []schema.Annotation {
	return nil
}
