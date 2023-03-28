// Code generated by entimport, DO NOT EDIT.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type RoleUser struct {
	ent.Schema
}

func (RoleUser) Fields() []ent.Field {
	return []ent.Field{field.Uint64("id"), field.Uint64("user_id").Optional(), field.Uint64("role_id").Optional(), field.Time("created_at").Optional(), field.Time("updated_at").Optional()}
}
func (RoleUser) Edges() []ent.Edge {
	return nil
}
func (RoleUser) Annotations() []schema.Annotation {
	return nil
}