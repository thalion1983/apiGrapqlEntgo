// Code generated by ent, DO NOT EDIT.

package course

import (
	"apiGrapqlEntgo/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Course {
	return predicate.Course(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Course {
	return predicate.Course(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Course {
	return predicate.Course(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Course {
	return predicate.Course(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Course {
	return predicate.Course(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Course {
	return predicate.Course(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Course {
	return predicate.Course(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Course {
	return predicate.Course(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Course {
	return predicate.Course(sql.FieldLTE(FieldID, id))
}

// Year applies equality check predicate on the "year" field. It's identical to YearEQ.
func Year(v int) predicate.Course {
	return predicate.Course(sql.FieldEQ(FieldYear, v))
}

// Period applies equality check predicate on the "period" field. It's identical to PeriodEQ.
func Period(v int) predicate.Course {
	return predicate.Course(sql.FieldEQ(FieldPeriod, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Course {
	return predicate.Course(sql.FieldEQ(FieldCreatedAt, v))
}

// LastModifiedAt applies equality check predicate on the "last_modified_at" field. It's identical to LastModifiedAtEQ.
func LastModifiedAt(v time.Time) predicate.Course {
	return predicate.Course(sql.FieldEQ(FieldLastModifiedAt, v))
}

// YearEQ applies the EQ predicate on the "year" field.
func YearEQ(v int) predicate.Course {
	return predicate.Course(sql.FieldEQ(FieldYear, v))
}

// YearNEQ applies the NEQ predicate on the "year" field.
func YearNEQ(v int) predicate.Course {
	return predicate.Course(sql.FieldNEQ(FieldYear, v))
}

// YearIn applies the In predicate on the "year" field.
func YearIn(vs ...int) predicate.Course {
	return predicate.Course(sql.FieldIn(FieldYear, vs...))
}

// YearNotIn applies the NotIn predicate on the "year" field.
func YearNotIn(vs ...int) predicate.Course {
	return predicate.Course(sql.FieldNotIn(FieldYear, vs...))
}

// YearGT applies the GT predicate on the "year" field.
func YearGT(v int) predicate.Course {
	return predicate.Course(sql.FieldGT(FieldYear, v))
}

// YearGTE applies the GTE predicate on the "year" field.
func YearGTE(v int) predicate.Course {
	return predicate.Course(sql.FieldGTE(FieldYear, v))
}

// YearLT applies the LT predicate on the "year" field.
func YearLT(v int) predicate.Course {
	return predicate.Course(sql.FieldLT(FieldYear, v))
}

// YearLTE applies the LTE predicate on the "year" field.
func YearLTE(v int) predicate.Course {
	return predicate.Course(sql.FieldLTE(FieldYear, v))
}

// PeriodEQ applies the EQ predicate on the "period" field.
func PeriodEQ(v int) predicate.Course {
	return predicate.Course(sql.FieldEQ(FieldPeriod, v))
}

// PeriodNEQ applies the NEQ predicate on the "period" field.
func PeriodNEQ(v int) predicate.Course {
	return predicate.Course(sql.FieldNEQ(FieldPeriod, v))
}

// PeriodIn applies the In predicate on the "period" field.
func PeriodIn(vs ...int) predicate.Course {
	return predicate.Course(sql.FieldIn(FieldPeriod, vs...))
}

// PeriodNotIn applies the NotIn predicate on the "period" field.
func PeriodNotIn(vs ...int) predicate.Course {
	return predicate.Course(sql.FieldNotIn(FieldPeriod, vs...))
}

// PeriodGT applies the GT predicate on the "period" field.
func PeriodGT(v int) predicate.Course {
	return predicate.Course(sql.FieldGT(FieldPeriod, v))
}

// PeriodGTE applies the GTE predicate on the "period" field.
func PeriodGTE(v int) predicate.Course {
	return predicate.Course(sql.FieldGTE(FieldPeriod, v))
}

// PeriodLT applies the LT predicate on the "period" field.
func PeriodLT(v int) predicate.Course {
	return predicate.Course(sql.FieldLT(FieldPeriod, v))
}

// PeriodLTE applies the LTE predicate on the "period" field.
func PeriodLTE(v int) predicate.Course {
	return predicate.Course(sql.FieldLTE(FieldPeriod, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Course {
	return predicate.Course(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Course {
	return predicate.Course(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Course {
	return predicate.Course(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Course {
	return predicate.Course(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Course {
	return predicate.Course(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Course {
	return predicate.Course(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Course {
	return predicate.Course(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Course {
	return predicate.Course(sql.FieldLTE(FieldCreatedAt, v))
}

// LastModifiedAtEQ applies the EQ predicate on the "last_modified_at" field.
func LastModifiedAtEQ(v time.Time) predicate.Course {
	return predicate.Course(sql.FieldEQ(FieldLastModifiedAt, v))
}

// LastModifiedAtNEQ applies the NEQ predicate on the "last_modified_at" field.
func LastModifiedAtNEQ(v time.Time) predicate.Course {
	return predicate.Course(sql.FieldNEQ(FieldLastModifiedAt, v))
}

// LastModifiedAtIn applies the In predicate on the "last_modified_at" field.
func LastModifiedAtIn(vs ...time.Time) predicate.Course {
	return predicate.Course(sql.FieldIn(FieldLastModifiedAt, vs...))
}

// LastModifiedAtNotIn applies the NotIn predicate on the "last_modified_at" field.
func LastModifiedAtNotIn(vs ...time.Time) predicate.Course {
	return predicate.Course(sql.FieldNotIn(FieldLastModifiedAt, vs...))
}

// LastModifiedAtGT applies the GT predicate on the "last_modified_at" field.
func LastModifiedAtGT(v time.Time) predicate.Course {
	return predicate.Course(sql.FieldGT(FieldLastModifiedAt, v))
}

// LastModifiedAtGTE applies the GTE predicate on the "last_modified_at" field.
func LastModifiedAtGTE(v time.Time) predicate.Course {
	return predicate.Course(sql.FieldGTE(FieldLastModifiedAt, v))
}

// LastModifiedAtLT applies the LT predicate on the "last_modified_at" field.
func LastModifiedAtLT(v time.Time) predicate.Course {
	return predicate.Course(sql.FieldLT(FieldLastModifiedAt, v))
}

// LastModifiedAtLTE applies the LTE predicate on the "last_modified_at" field.
func LastModifiedAtLTE(v time.Time) predicate.Course {
	return predicate.Course(sql.FieldLTE(FieldLastModifiedAt, v))
}

// HasSubject applies the HasEdge predicate on the "subject" edge.
func HasSubject() predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, SubjectTable, SubjectPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSubjectWith applies the HasEdge predicate on the "subject" edge with a given conditions (other predicates).
func HasSubjectWith(preds ...predicate.Subject) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		step := newSubjectStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProfessor applies the HasEdge predicate on the "professor" edge.
func HasProfessor() predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProfessorTable, ProfessorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProfessorWith applies the HasEdge predicate on the "professor" edge with a given conditions (other predicates).
func HasProfessorWith(preds ...predicate.Professor) predicate.Course {
	return predicate.Course(func(s *sql.Selector) {
		step := newProfessorStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Course) predicate.Course {
	return predicate.Course(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Course) predicate.Course {
	return predicate.Course(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Course) predicate.Course {
	return predicate.Course(sql.NotPredicates(p))
}
