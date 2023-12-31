// Code generated by ent, DO NOT EDIT.

package ent

import (
	"apiGrapqlEntgo/ent/professor"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Professor is the model entity for the Professor schema.
type Professor struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// BirthDate holds the value of the "birth_date" field.
	BirthDate time.Time `json:"birth_date,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// LastModifiedAt holds the value of the "last_modified_at" field.
	LastModifiedAt time.Time `json:"last_modified_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProfessorQuery when eager-loading is set.
	Edges        ProfessorEdges `json:"edges"`
	selectValues sql.SelectValues
}

// ProfessorEdges holds the relations/edges for other nodes in the graph.
type ProfessorEdges struct {
	// Courses holds the value of the courses edge.
	Courses []*Course `json:"courses,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// CoursesOrErr returns the Courses value or an error if the edge
// was not loaded in eager-loading.
func (e ProfessorEdges) CoursesOrErr() ([]*Course, error) {
	if e.loadedTypes[0] {
		return e.Courses, nil
	}
	return nil, &NotLoadedError{edge: "courses"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Professor) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case professor.FieldID, professor.FieldName, professor.FieldLastName:
			values[i] = new(sql.NullString)
		case professor.FieldBirthDate, professor.FieldCreatedAt, professor.FieldLastModifiedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Professor fields.
func (pr *Professor) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case professor.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				pr.ID = value.String
			}
		case professor.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case professor.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				pr.LastName = value.String
			}
		case professor.FieldBirthDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field birth_date", values[i])
			} else if value.Valid {
				pr.BirthDate = value.Time
			}
		case professor.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				pr.CreatedAt = value.Time
			}
		case professor.FieldLastModifiedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_modified_at", values[i])
			} else if value.Valid {
				pr.LastModifiedAt = value.Time
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Professor.
// This includes values selected through modifiers, order, etc.
func (pr *Professor) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// QueryCourses queries the "courses" edge of the Professor entity.
func (pr *Professor) QueryCourses() *CourseQuery {
	return NewProfessorClient(pr.config).QueryCourses(pr)
}

// Update returns a builder for updating this Professor.
// Note that you need to call Professor.Unwrap() before calling this method if this Professor
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Professor) Update() *ProfessorUpdateOne {
	return NewProfessorClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Professor entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Professor) Unwrap() *Professor {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Professor is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Professor) String() string {
	var builder strings.Builder
	builder.WriteString("Professor(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", ")
	builder.WriteString("last_name=")
	builder.WriteString(pr.LastName)
	builder.WriteString(", ")
	builder.WriteString("birth_date=")
	builder.WriteString(pr.BirthDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(pr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("last_modified_at=")
	builder.WriteString(pr.LastModifiedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Professors is a parsable slice of Professor.
type Professors []*Professor
