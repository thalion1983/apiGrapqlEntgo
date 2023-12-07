package schema

import "entgo.io/ent"

// Professor holds the schema definition for the Professor entity.
type Professor struct {
	ent.Schema
}

// Fields of the Professor.
func (Professor) Fields() []ent.Field {
	return nil
}

// Edges of the Professor.
func (Professor) Edges() []ent.Edge {
	return nil
}
