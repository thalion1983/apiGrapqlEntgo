// Code generated by ent, DO NOT EDIT.

package professor

import (
	"apiGrapqlEntgo/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Professor {
	return predicate.Professor(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Professor {
	return predicate.Professor(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Professor {
	return predicate.Professor(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Professor {
	return predicate.Professor(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Professor {
	return predicate.Professor(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Professor {
	return predicate.Professor(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Professor {
	return predicate.Professor(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Professor {
	return predicate.Professor(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Professor {
	return predicate.Professor(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Professor {
	return predicate.Professor(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Professor {
	return predicate.Professor(sql.FieldContainsFold(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Professor {
	return predicate.Professor(sql.FieldEQ(FieldName, v))
}

// LastName applies equality check predicate on the "last_name" field. It's identical to LastNameEQ.
func LastName(v string) predicate.Professor {
	return predicate.Professor(sql.FieldEQ(FieldLastName, v))
}

// BirthDate applies equality check predicate on the "birth_date" field. It's identical to BirthDateEQ.
func BirthDate(v time.Time) predicate.Professor {
	return predicate.Professor(sql.FieldEQ(FieldBirthDate, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Professor {
	return predicate.Professor(sql.FieldEQ(FieldCreatedAt, v))
}

// LastModifiedAt applies equality check predicate on the "last_modified_at" field. It's identical to LastModifiedAtEQ.
func LastModifiedAt(v time.Time) predicate.Professor {
	return predicate.Professor(sql.FieldEQ(FieldLastModifiedAt, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Professor {
	return predicate.Professor(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Professor {
	return predicate.Professor(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Professor {
	return predicate.Professor(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Professor {
	return predicate.Professor(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Professor {
	return predicate.Professor(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Professor {
	return predicate.Professor(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Professor {
	return predicate.Professor(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Professor {
	return predicate.Professor(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Professor {
	return predicate.Professor(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Professor {
	return predicate.Professor(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Professor {
	return predicate.Professor(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Professor {
	return predicate.Professor(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Professor {
	return predicate.Professor(sql.FieldContainsFold(FieldName, v))
}

// LastNameEQ applies the EQ predicate on the "last_name" field.
func LastNameEQ(v string) predicate.Professor {
	return predicate.Professor(sql.FieldEQ(FieldLastName, v))
}

// LastNameNEQ applies the NEQ predicate on the "last_name" field.
func LastNameNEQ(v string) predicate.Professor {
	return predicate.Professor(sql.FieldNEQ(FieldLastName, v))
}

// LastNameIn applies the In predicate on the "last_name" field.
func LastNameIn(vs ...string) predicate.Professor {
	return predicate.Professor(sql.FieldIn(FieldLastName, vs...))
}

// LastNameNotIn applies the NotIn predicate on the "last_name" field.
func LastNameNotIn(vs ...string) predicate.Professor {
	return predicate.Professor(sql.FieldNotIn(FieldLastName, vs...))
}

// LastNameGT applies the GT predicate on the "last_name" field.
func LastNameGT(v string) predicate.Professor {
	return predicate.Professor(sql.FieldGT(FieldLastName, v))
}

// LastNameGTE applies the GTE predicate on the "last_name" field.
func LastNameGTE(v string) predicate.Professor {
	return predicate.Professor(sql.FieldGTE(FieldLastName, v))
}

// LastNameLT applies the LT predicate on the "last_name" field.
func LastNameLT(v string) predicate.Professor {
	return predicate.Professor(sql.FieldLT(FieldLastName, v))
}

// LastNameLTE applies the LTE predicate on the "last_name" field.
func LastNameLTE(v string) predicate.Professor {
	return predicate.Professor(sql.FieldLTE(FieldLastName, v))
}

// LastNameContains applies the Contains predicate on the "last_name" field.
func LastNameContains(v string) predicate.Professor {
	return predicate.Professor(sql.FieldContains(FieldLastName, v))
}

// LastNameHasPrefix applies the HasPrefix predicate on the "last_name" field.
func LastNameHasPrefix(v string) predicate.Professor {
	return predicate.Professor(sql.FieldHasPrefix(FieldLastName, v))
}

// LastNameHasSuffix applies the HasSuffix predicate on the "last_name" field.
func LastNameHasSuffix(v string) predicate.Professor {
	return predicate.Professor(sql.FieldHasSuffix(FieldLastName, v))
}

// LastNameEqualFold applies the EqualFold predicate on the "last_name" field.
func LastNameEqualFold(v string) predicate.Professor {
	return predicate.Professor(sql.FieldEqualFold(FieldLastName, v))
}

// LastNameContainsFold applies the ContainsFold predicate on the "last_name" field.
func LastNameContainsFold(v string) predicate.Professor {
	return predicate.Professor(sql.FieldContainsFold(FieldLastName, v))
}

// BirthDateEQ applies the EQ predicate on the "birth_date" field.
func BirthDateEQ(v time.Time) predicate.Professor {
	return predicate.Professor(sql.FieldEQ(FieldBirthDate, v))
}

// BirthDateNEQ applies the NEQ predicate on the "birth_date" field.
func BirthDateNEQ(v time.Time) predicate.Professor {
	return predicate.Professor(sql.FieldNEQ(FieldBirthDate, v))
}

// BirthDateIn applies the In predicate on the "birth_date" field.
func BirthDateIn(vs ...time.Time) predicate.Professor {
	return predicate.Professor(sql.FieldIn(FieldBirthDate, vs...))
}

// BirthDateNotIn applies the NotIn predicate on the "birth_date" field.
func BirthDateNotIn(vs ...time.Time) predicate.Professor {
	return predicate.Professor(sql.FieldNotIn(FieldBirthDate, vs...))
}

// BirthDateGT applies the GT predicate on the "birth_date" field.
func BirthDateGT(v time.Time) predicate.Professor {
	return predicate.Professor(sql.FieldGT(FieldBirthDate, v))
}

// BirthDateGTE applies the GTE predicate on the "birth_date" field.
func BirthDateGTE(v time.Time) predicate.Professor {
	return predicate.Professor(sql.FieldGTE(FieldBirthDate, v))
}

// BirthDateLT applies the LT predicate on the "birth_date" field.
func BirthDateLT(v time.Time) predicate.Professor {
	return predicate.Professor(sql.FieldLT(FieldBirthDate, v))
}

// BirthDateLTE applies the LTE predicate on the "birth_date" field.
func BirthDateLTE(v time.Time) predicate.Professor {
	return predicate.Professor(sql.FieldLTE(FieldBirthDate, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Professor {
	return predicate.Professor(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Professor {
	return predicate.Professor(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Professor {
	return predicate.Professor(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Professor {
	return predicate.Professor(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Professor {
	return predicate.Professor(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Professor {
	return predicate.Professor(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Professor {
	return predicate.Professor(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Professor {
	return predicate.Professor(sql.FieldLTE(FieldCreatedAt, v))
}

// LastModifiedAtEQ applies the EQ predicate on the "last_modified_at" field.
func LastModifiedAtEQ(v time.Time) predicate.Professor {
	return predicate.Professor(sql.FieldEQ(FieldLastModifiedAt, v))
}

// LastModifiedAtNEQ applies the NEQ predicate on the "last_modified_at" field.
func LastModifiedAtNEQ(v time.Time) predicate.Professor {
	return predicate.Professor(sql.FieldNEQ(FieldLastModifiedAt, v))
}

// LastModifiedAtIn applies the In predicate on the "last_modified_at" field.
func LastModifiedAtIn(vs ...time.Time) predicate.Professor {
	return predicate.Professor(sql.FieldIn(FieldLastModifiedAt, vs...))
}

// LastModifiedAtNotIn applies the NotIn predicate on the "last_modified_at" field.
func LastModifiedAtNotIn(vs ...time.Time) predicate.Professor {
	return predicate.Professor(sql.FieldNotIn(FieldLastModifiedAt, vs...))
}

// LastModifiedAtGT applies the GT predicate on the "last_modified_at" field.
func LastModifiedAtGT(v time.Time) predicate.Professor {
	return predicate.Professor(sql.FieldGT(FieldLastModifiedAt, v))
}

// LastModifiedAtGTE applies the GTE predicate on the "last_modified_at" field.
func LastModifiedAtGTE(v time.Time) predicate.Professor {
	return predicate.Professor(sql.FieldGTE(FieldLastModifiedAt, v))
}

// LastModifiedAtLT applies the LT predicate on the "last_modified_at" field.
func LastModifiedAtLT(v time.Time) predicate.Professor {
	return predicate.Professor(sql.FieldLT(FieldLastModifiedAt, v))
}

// LastModifiedAtLTE applies the LTE predicate on the "last_modified_at" field.
func LastModifiedAtLTE(v time.Time) predicate.Professor {
	return predicate.Professor(sql.FieldLTE(FieldLastModifiedAt, v))
}

// HasCourses applies the HasEdge predicate on the "courses" edge.
func HasCourses() predicate.Professor {
	return predicate.Professor(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, CoursesTable, CoursesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCoursesWith applies the HasEdge predicate on the "courses" edge with a given conditions (other predicates).
func HasCoursesWith(preds ...predicate.Course) predicate.Professor {
	return predicate.Professor(func(s *sql.Selector) {
		step := newCoursesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Professor) predicate.Professor {
	return predicate.Professor(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Professor) predicate.Professor {
	return predicate.Professor(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Professor) predicate.Professor {
	return predicate.Professor(sql.NotPredicates(p))
}
