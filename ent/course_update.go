// Code generated by ent, DO NOT EDIT.

package ent

import (
	"apiGrapqlEntgo/ent/course"
	"apiGrapqlEntgo/ent/predicate"
	"apiGrapqlEntgo/ent/professor"
	"apiGrapqlEntgo/ent/subject"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CourseUpdate is the builder for updating Course entities.
type CourseUpdate struct {
	config
	hooks    []Hook
	mutation *CourseMutation
}

// Where appends a list predicates to the CourseUpdate builder.
func (cu *CourseUpdate) Where(ps ...predicate.Course) *CourseUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetYear sets the "year" field.
func (cu *CourseUpdate) SetYear(i int) *CourseUpdate {
	cu.mutation.ResetYear()
	cu.mutation.SetYear(i)
	return cu
}

// SetNillableYear sets the "year" field if the given value is not nil.
func (cu *CourseUpdate) SetNillableYear(i *int) *CourseUpdate {
	if i != nil {
		cu.SetYear(*i)
	}
	return cu
}

// AddYear adds i to the "year" field.
func (cu *CourseUpdate) AddYear(i int) *CourseUpdate {
	cu.mutation.AddYear(i)
	return cu
}

// SetPeriod sets the "period" field.
func (cu *CourseUpdate) SetPeriod(i int) *CourseUpdate {
	cu.mutation.ResetPeriod()
	cu.mutation.SetPeriod(i)
	return cu
}

// SetNillablePeriod sets the "period" field if the given value is not nil.
func (cu *CourseUpdate) SetNillablePeriod(i *int) *CourseUpdate {
	if i != nil {
		cu.SetPeriod(*i)
	}
	return cu
}

// AddPeriod adds i to the "period" field.
func (cu *CourseUpdate) AddPeriod(i int) *CourseUpdate {
	cu.mutation.AddPeriod(i)
	return cu
}

// SetProfessorID sets the "professor_id" field.
func (cu *CourseUpdate) SetProfessorID(s string) *CourseUpdate {
	cu.mutation.SetProfessorID(s)
	return cu
}

// SetNillableProfessorID sets the "professor_id" field if the given value is not nil.
func (cu *CourseUpdate) SetNillableProfessorID(s *string) *CourseUpdate {
	if s != nil {
		cu.SetProfessorID(*s)
	}
	return cu
}

// SetSubjectID sets the "subject_id" field.
func (cu *CourseUpdate) SetSubjectID(s string) *CourseUpdate {
	cu.mutation.SetSubjectID(s)
	return cu
}

// SetNillableSubjectID sets the "subject_id" field if the given value is not nil.
func (cu *CourseUpdate) SetNillableSubjectID(s *string) *CourseUpdate {
	if s != nil {
		cu.SetSubjectID(*s)
	}
	return cu
}

// SetLastModifiedAt sets the "last_modified_at" field.
func (cu *CourseUpdate) SetLastModifiedAt(t time.Time) *CourseUpdate {
	cu.mutation.SetLastModifiedAt(t)
	return cu
}

// SetSubject sets the "subject" edge to the Subject entity.
func (cu *CourseUpdate) SetSubject(s *Subject) *CourseUpdate {
	return cu.SetSubjectID(s.ID)
}

// SetProfessor sets the "professor" edge to the Professor entity.
func (cu *CourseUpdate) SetProfessor(p *Professor) *CourseUpdate {
	return cu.SetProfessorID(p.ID)
}

// Mutation returns the CourseMutation object of the builder.
func (cu *CourseUpdate) Mutation() *CourseMutation {
	return cu.mutation
}

// ClearSubject clears the "subject" edge to the Subject entity.
func (cu *CourseUpdate) ClearSubject() *CourseUpdate {
	cu.mutation.ClearSubject()
	return cu
}

// ClearProfessor clears the "professor" edge to the Professor entity.
func (cu *CourseUpdate) ClearProfessor() *CourseUpdate {
	cu.mutation.ClearProfessor()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *CourseUpdate) Save(ctx context.Context) (int, error) {
	cu.defaults()
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CourseUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CourseUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CourseUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *CourseUpdate) defaults() {
	if _, ok := cu.mutation.LastModifiedAt(); !ok {
		v := course.UpdateDefaultLastModifiedAt()
		cu.mutation.SetLastModifiedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *CourseUpdate) check() error {
	if _, ok := cu.mutation.SubjectID(); cu.mutation.SubjectCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Course.subject"`)
	}
	if _, ok := cu.mutation.ProfessorID(); cu.mutation.ProfessorCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Course.professor"`)
	}
	return nil
}

func (cu *CourseUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(course.Table, course.Columns, sqlgraph.NewFieldSpec(course.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Year(); ok {
		_spec.SetField(course.FieldYear, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedYear(); ok {
		_spec.AddField(course.FieldYear, field.TypeInt, value)
	}
	if value, ok := cu.mutation.Period(); ok {
		_spec.SetField(course.FieldPeriod, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedPeriod(); ok {
		_spec.AddField(course.FieldPeriod, field.TypeInt, value)
	}
	if value, ok := cu.mutation.LastModifiedAt(); ok {
		_spec.SetField(course.FieldLastModifiedAt, field.TypeTime, value)
	}
	if cu.mutation.SubjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   course.SubjectTable,
			Columns: []string{course.SubjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subject.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.SubjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   course.SubjectTable,
			Columns: []string{course.SubjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subject.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.ProfessorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   course.ProfessorTable,
			Columns: []string{course.ProfessorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(professor.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ProfessorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   course.ProfessorTable,
			Columns: []string{course.ProfessorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(professor.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{course.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// CourseUpdateOne is the builder for updating a single Course entity.
type CourseUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CourseMutation
}

// SetYear sets the "year" field.
func (cuo *CourseUpdateOne) SetYear(i int) *CourseUpdateOne {
	cuo.mutation.ResetYear()
	cuo.mutation.SetYear(i)
	return cuo
}

// SetNillableYear sets the "year" field if the given value is not nil.
func (cuo *CourseUpdateOne) SetNillableYear(i *int) *CourseUpdateOne {
	if i != nil {
		cuo.SetYear(*i)
	}
	return cuo
}

// AddYear adds i to the "year" field.
func (cuo *CourseUpdateOne) AddYear(i int) *CourseUpdateOne {
	cuo.mutation.AddYear(i)
	return cuo
}

// SetPeriod sets the "period" field.
func (cuo *CourseUpdateOne) SetPeriod(i int) *CourseUpdateOne {
	cuo.mutation.ResetPeriod()
	cuo.mutation.SetPeriod(i)
	return cuo
}

// SetNillablePeriod sets the "period" field if the given value is not nil.
func (cuo *CourseUpdateOne) SetNillablePeriod(i *int) *CourseUpdateOne {
	if i != nil {
		cuo.SetPeriod(*i)
	}
	return cuo
}

// AddPeriod adds i to the "period" field.
func (cuo *CourseUpdateOne) AddPeriod(i int) *CourseUpdateOne {
	cuo.mutation.AddPeriod(i)
	return cuo
}

// SetProfessorID sets the "professor_id" field.
func (cuo *CourseUpdateOne) SetProfessorID(s string) *CourseUpdateOne {
	cuo.mutation.SetProfessorID(s)
	return cuo
}

// SetNillableProfessorID sets the "professor_id" field if the given value is not nil.
func (cuo *CourseUpdateOne) SetNillableProfessorID(s *string) *CourseUpdateOne {
	if s != nil {
		cuo.SetProfessorID(*s)
	}
	return cuo
}

// SetSubjectID sets the "subject_id" field.
func (cuo *CourseUpdateOne) SetSubjectID(s string) *CourseUpdateOne {
	cuo.mutation.SetSubjectID(s)
	return cuo
}

// SetNillableSubjectID sets the "subject_id" field if the given value is not nil.
func (cuo *CourseUpdateOne) SetNillableSubjectID(s *string) *CourseUpdateOne {
	if s != nil {
		cuo.SetSubjectID(*s)
	}
	return cuo
}

// SetLastModifiedAt sets the "last_modified_at" field.
func (cuo *CourseUpdateOne) SetLastModifiedAt(t time.Time) *CourseUpdateOne {
	cuo.mutation.SetLastModifiedAt(t)
	return cuo
}

// SetSubject sets the "subject" edge to the Subject entity.
func (cuo *CourseUpdateOne) SetSubject(s *Subject) *CourseUpdateOne {
	return cuo.SetSubjectID(s.ID)
}

// SetProfessor sets the "professor" edge to the Professor entity.
func (cuo *CourseUpdateOne) SetProfessor(p *Professor) *CourseUpdateOne {
	return cuo.SetProfessorID(p.ID)
}

// Mutation returns the CourseMutation object of the builder.
func (cuo *CourseUpdateOne) Mutation() *CourseMutation {
	return cuo.mutation
}

// ClearSubject clears the "subject" edge to the Subject entity.
func (cuo *CourseUpdateOne) ClearSubject() *CourseUpdateOne {
	cuo.mutation.ClearSubject()
	return cuo
}

// ClearProfessor clears the "professor" edge to the Professor entity.
func (cuo *CourseUpdateOne) ClearProfessor() *CourseUpdateOne {
	cuo.mutation.ClearProfessor()
	return cuo
}

// Where appends a list predicates to the CourseUpdate builder.
func (cuo *CourseUpdateOne) Where(ps ...predicate.Course) *CourseUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *CourseUpdateOne) Select(field string, fields ...string) *CourseUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Course entity.
func (cuo *CourseUpdateOne) Save(ctx context.Context) (*Course, error) {
	cuo.defaults()
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CourseUpdateOne) SaveX(ctx context.Context) *Course {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CourseUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CourseUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *CourseUpdateOne) defaults() {
	if _, ok := cuo.mutation.LastModifiedAt(); !ok {
		v := course.UpdateDefaultLastModifiedAt()
		cuo.mutation.SetLastModifiedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *CourseUpdateOne) check() error {
	if _, ok := cuo.mutation.SubjectID(); cuo.mutation.SubjectCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Course.subject"`)
	}
	if _, ok := cuo.mutation.ProfessorID(); cuo.mutation.ProfessorCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Course.professor"`)
	}
	return nil
}

func (cuo *CourseUpdateOne) sqlSave(ctx context.Context) (_node *Course, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(course.Table, course.Columns, sqlgraph.NewFieldSpec(course.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Course.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, course.FieldID)
		for _, f := range fields {
			if !course.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != course.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Year(); ok {
		_spec.SetField(course.FieldYear, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedYear(); ok {
		_spec.AddField(course.FieldYear, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.Period(); ok {
		_spec.SetField(course.FieldPeriod, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedPeriod(); ok {
		_spec.AddField(course.FieldPeriod, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.LastModifiedAt(); ok {
		_spec.SetField(course.FieldLastModifiedAt, field.TypeTime, value)
	}
	if cuo.mutation.SubjectCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   course.SubjectTable,
			Columns: []string{course.SubjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subject.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.SubjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   course.SubjectTable,
			Columns: []string{course.SubjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subject.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.ProfessorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   course.ProfessorTable,
			Columns: []string{course.ProfessorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(professor.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ProfessorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   course.ProfessorTable,
			Columns: []string{course.ProfessorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(professor.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Course{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{course.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
