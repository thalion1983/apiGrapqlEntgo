package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Professor holds the schema definition for the Professor entity.
type Professor struct {
	ent.Schema
}

// Fields of the Professor.
func (Professor) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").Unique().NotEmpty(),
		field.String("name").NotEmpty(),
		field.String("last_name").NotEmpty(),
		field.Time("birth_date"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("last_modified_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Professor.
func (Professor) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("courses", Course.Type).Ref("professor"),
	}
}

// Indexes of the Professor.
func (Professor) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("id").Unique(),
	}
}
