// Code generated by ent, DO NOT EDIT.

package ent

import (
	"apiGrapqlEntgo/ent/course"
	"apiGrapqlEntgo/ent/predicate"
	"apiGrapqlEntgo/ent/professor"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProfessorUpdate is the builder for updating Professor entities.
type ProfessorUpdate struct {
	config
	hooks    []Hook
	mutation *ProfessorMutation
}

// Where appends a list predicates to the ProfessorUpdate builder.
func (pu *ProfessorUpdate) Where(ps ...predicate.Professor) *ProfessorUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetName sets the "name" field.
func (pu *ProfessorUpdate) SetName(s string) *ProfessorUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (pu *ProfessorUpdate) SetNillableName(s *string) *ProfessorUpdate {
	if s != nil {
		pu.SetName(*s)
	}
	return pu
}

// SetLastName sets the "last_name" field.
func (pu *ProfessorUpdate) SetLastName(s string) *ProfessorUpdate {
	pu.mutation.SetLastName(s)
	return pu
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (pu *ProfessorUpdate) SetNillableLastName(s *string) *ProfessorUpdate {
	if s != nil {
		pu.SetLastName(*s)
	}
	return pu
}

// SetBirthDate sets the "birth_date" field.
func (pu *ProfessorUpdate) SetBirthDate(t time.Time) *ProfessorUpdate {
	pu.mutation.SetBirthDate(t)
	return pu
}

// SetNillableBirthDate sets the "birth_date" field if the given value is not nil.
func (pu *ProfessorUpdate) SetNillableBirthDate(t *time.Time) *ProfessorUpdate {
	if t != nil {
		pu.SetBirthDate(*t)
	}
	return pu
}

// SetLastModifiedAt sets the "last_modified_at" field.
func (pu *ProfessorUpdate) SetLastModifiedAt(t time.Time) *ProfessorUpdate {
	pu.mutation.SetLastModifiedAt(t)
	return pu
}

// AddCourseIDs adds the "courses" edge to the Course entity by IDs.
func (pu *ProfessorUpdate) AddCourseIDs(ids ...int) *ProfessorUpdate {
	pu.mutation.AddCourseIDs(ids...)
	return pu
}

// AddCourses adds the "courses" edges to the Course entity.
func (pu *ProfessorUpdate) AddCourses(c ...*Course) *ProfessorUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pu.AddCourseIDs(ids...)
}

// Mutation returns the ProfessorMutation object of the builder.
func (pu *ProfessorUpdate) Mutation() *ProfessorMutation {
	return pu.mutation
}

// ClearCourses clears all "courses" edges to the Course entity.
func (pu *ProfessorUpdate) ClearCourses() *ProfessorUpdate {
	pu.mutation.ClearCourses()
	return pu
}

// RemoveCourseIDs removes the "courses" edge to Course entities by IDs.
func (pu *ProfessorUpdate) RemoveCourseIDs(ids ...int) *ProfessorUpdate {
	pu.mutation.RemoveCourseIDs(ids...)
	return pu
}

// RemoveCourses removes "courses" edges to Course entities.
func (pu *ProfessorUpdate) RemoveCourses(c ...*Course) *ProfessorUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pu.RemoveCourseIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProfessorUpdate) Save(ctx context.Context) (int, error) {
	pu.defaults()
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProfessorUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProfessorUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProfessorUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *ProfessorUpdate) defaults() {
	if _, ok := pu.mutation.LastModifiedAt(); !ok {
		v := professor.UpdateDefaultLastModifiedAt()
		pu.mutation.SetLastModifiedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *ProfessorUpdate) check() error {
	if v, ok := pu.mutation.Name(); ok {
		if err := professor.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Professor.name": %w`, err)}
		}
	}
	if v, ok := pu.mutation.LastName(); ok {
		if err := professor.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`ent: validator failed for field "Professor.last_name": %w`, err)}
		}
	}
	return nil
}

func (pu *ProfessorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(professor.Table, professor.Columns, sqlgraph.NewFieldSpec(professor.FieldID, field.TypeString))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(professor.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.LastName(); ok {
		_spec.SetField(professor.FieldLastName, field.TypeString, value)
	}
	if value, ok := pu.mutation.BirthDate(); ok {
		_spec.SetField(professor.FieldBirthDate, field.TypeTime, value)
	}
	if value, ok := pu.mutation.LastModifiedAt(); ok {
		_spec.SetField(professor.FieldLastModifiedAt, field.TypeTime, value)
	}
	if pu.mutation.CoursesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   professor.CoursesTable,
			Columns: []string{professor.CoursesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedCoursesIDs(); len(nodes) > 0 && !pu.mutation.CoursesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   professor.CoursesTable,
			Columns: []string{professor.CoursesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.CoursesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   professor.CoursesTable,
			Columns: []string{professor.CoursesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{professor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// ProfessorUpdateOne is the builder for updating a single Professor entity.
type ProfessorUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProfessorMutation
}

// SetName sets the "name" field.
func (puo *ProfessorUpdateOne) SetName(s string) *ProfessorUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (puo *ProfessorUpdateOne) SetNillableName(s *string) *ProfessorUpdateOne {
	if s != nil {
		puo.SetName(*s)
	}
	return puo
}

// SetLastName sets the "last_name" field.
func (puo *ProfessorUpdateOne) SetLastName(s string) *ProfessorUpdateOne {
	puo.mutation.SetLastName(s)
	return puo
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (puo *ProfessorUpdateOne) SetNillableLastName(s *string) *ProfessorUpdateOne {
	if s != nil {
		puo.SetLastName(*s)
	}
	return puo
}

// SetBirthDate sets the "birth_date" field.
func (puo *ProfessorUpdateOne) SetBirthDate(t time.Time) *ProfessorUpdateOne {
	puo.mutation.SetBirthDate(t)
	return puo
}

// SetNillableBirthDate sets the "birth_date" field if the given value is not nil.
func (puo *ProfessorUpdateOne) SetNillableBirthDate(t *time.Time) *ProfessorUpdateOne {
	if t != nil {
		puo.SetBirthDate(*t)
	}
	return puo
}

// SetLastModifiedAt sets the "last_modified_at" field.
func (puo *ProfessorUpdateOne) SetLastModifiedAt(t time.Time) *ProfessorUpdateOne {
	puo.mutation.SetLastModifiedAt(t)
	return puo
}

// AddCourseIDs adds the "courses" edge to the Course entity by IDs.
func (puo *ProfessorUpdateOne) AddCourseIDs(ids ...int) *ProfessorUpdateOne {
	puo.mutation.AddCourseIDs(ids...)
	return puo
}

// AddCourses adds the "courses" edges to the Course entity.
func (puo *ProfessorUpdateOne) AddCourses(c ...*Course) *ProfessorUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return puo.AddCourseIDs(ids...)
}

// Mutation returns the ProfessorMutation object of the builder.
func (puo *ProfessorUpdateOne) Mutation() *ProfessorMutation {
	return puo.mutation
}

// ClearCourses clears all "courses" edges to the Course entity.
func (puo *ProfessorUpdateOne) ClearCourses() *ProfessorUpdateOne {
	puo.mutation.ClearCourses()
	return puo
}

// RemoveCourseIDs removes the "courses" edge to Course entities by IDs.
func (puo *ProfessorUpdateOne) RemoveCourseIDs(ids ...int) *ProfessorUpdateOne {
	puo.mutation.RemoveCourseIDs(ids...)
	return puo
}

// RemoveCourses removes "courses" edges to Course entities.
func (puo *ProfessorUpdateOne) RemoveCourses(c ...*Course) *ProfessorUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return puo.RemoveCourseIDs(ids...)
}

// Where appends a list predicates to the ProfessorUpdate builder.
func (puo *ProfessorUpdateOne) Where(ps ...predicate.Professor) *ProfessorUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProfessorUpdateOne) Select(field string, fields ...string) *ProfessorUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Professor entity.
func (puo *ProfessorUpdateOne) Save(ctx context.Context) (*Professor, error) {
	puo.defaults()
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProfessorUpdateOne) SaveX(ctx context.Context) *Professor {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProfessorUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProfessorUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *ProfessorUpdateOne) defaults() {
	if _, ok := puo.mutation.LastModifiedAt(); !ok {
		v := professor.UpdateDefaultLastModifiedAt()
		puo.mutation.SetLastModifiedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *ProfessorUpdateOne) check() error {
	if v, ok := puo.mutation.Name(); ok {
		if err := professor.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Professor.name": %w`, err)}
		}
	}
	if v, ok := puo.mutation.LastName(); ok {
		if err := professor.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`ent: validator failed for field "Professor.last_name": %w`, err)}
		}
	}
	return nil
}

func (puo *ProfessorUpdateOne) sqlSave(ctx context.Context) (_node *Professor, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(professor.Table, professor.Columns, sqlgraph.NewFieldSpec(professor.FieldID, field.TypeString))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Professor.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, professor.FieldID)
		for _, f := range fields {
			if !professor.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != professor.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(professor.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.LastName(); ok {
		_spec.SetField(professor.FieldLastName, field.TypeString, value)
	}
	if value, ok := puo.mutation.BirthDate(); ok {
		_spec.SetField(professor.FieldBirthDate, field.TypeTime, value)
	}
	if value, ok := puo.mutation.LastModifiedAt(); ok {
		_spec.SetField(professor.FieldLastModifiedAt, field.TypeTime, value)
	}
	if puo.mutation.CoursesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   professor.CoursesTable,
			Columns: []string{professor.CoursesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedCoursesIDs(); len(nodes) > 0 && !puo.mutation.CoursesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   professor.CoursesTable,
			Columns: []string{professor.CoursesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.CoursesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   professor.CoursesTable,
			Columns: []string{professor.CoursesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(course.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Professor{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{professor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}