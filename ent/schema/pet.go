package schema

import "entgo.io/ent"

// Pet holds the schema definition for the Pet entity.
type Pet struct {
	ent.Schema
}

// Fields of the Pet.
func (Pet) Fields() []ent.Field {
	return nil
}

// Edges of the Pet.
func (Pet) Edges() []ent.Edge {
	return nil
}
