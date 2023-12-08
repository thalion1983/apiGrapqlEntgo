package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
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
		field.String("professor_id"),
		field.String("subject_id").Unique().
			SchemaType(map[string]string{
				dialect.Postgres: "char(6)",
			}),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("last_modified_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Course.
func (Course) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("subject", Subject.Type).Field("subject_id").Unique().Required(),
		edge.To("professor", Professor.Type).Field("professor_id").Unique().Required(),
	}
}

// Indexes of the Course.
func (Course) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("year", "period", "subject_id").Unique(),
	}
}
