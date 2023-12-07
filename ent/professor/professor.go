// Code generated by ent, DO NOT EDIT.

package professor

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the professor type in the database.
	Label = "professor"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldLastName holds the string denoting the last_name field in the database.
	FieldLastName = "last_name"
	// FieldBirthDate holds the string denoting the birth_date field in the database.
	FieldBirthDate = "birth_date"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldLastModifiedAt holds the string denoting the last_modified_at field in the database.
	FieldLastModifiedAt = "last_modified_at"
	// EdgeCourses holds the string denoting the courses edge name in mutations.
	EdgeCourses = "courses"
	// Table holds the table name of the professor in the database.
	Table = "professors"
	// CoursesTable is the table that holds the courses relation/edge.
	CoursesTable = "courses"
	// CoursesInverseTable is the table name for the Course entity.
	// It exists in this package in order to avoid circular dependency with the "course" package.
	CoursesInverseTable = "courses"
	// CoursesColumn is the table column denoting the courses relation/edge.
	CoursesColumn = "professor_courses"
)

// Columns holds all SQL columns for professor fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldLastName,
	FieldBirthDate,
	FieldCreatedAt,
	FieldLastModifiedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	LastNameValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultLastModifiedAt holds the default value on creation for the "last_modified_at" field.
	DefaultLastModifiedAt func() time.Time
	// UpdateDefaultLastModifiedAt holds the default value on update for the "last_modified_at" field.
	UpdateDefaultLastModifiedAt func() time.Time
	// IDValidator is a validator for the "id" field. It is called by the builders before save.
	IDValidator func(string) error
)

// OrderOption defines the ordering options for the Professor queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByLastName orders the results by the last_name field.
func ByLastName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastName, opts...).ToFunc()
}

// ByBirthDate orders the results by the birth_date field.
func ByBirthDate(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBirthDate, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByLastModifiedAt orders the results by the last_modified_at field.
func ByLastModifiedAt(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLastModifiedAt, opts...).ToFunc()
}

// ByCoursesCount orders the results by courses count.
func ByCoursesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCoursesStep(), opts...)
	}
}

// ByCourses orders the results by courses terms.
func ByCourses(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCoursesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newCoursesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CoursesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CoursesTable, CoursesColumn),
	)
}
