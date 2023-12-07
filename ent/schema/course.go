package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Course holds the schema definition for the Course entity.
type Course struct {
	ent.Schema
}

// Fields of the Course.
func (Course) Fields() []ent.Field {
	return []ent.Field{
		field.Int("year"),
		field.Int("period"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("last_modified_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Course.
func (Course) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("subject", Subject.Type).Ref("courses").Unique(),
		edge.From("professor", Professor.Type).Ref("courses").Unique(),
	}
}
