// Code generated by ent, DO NOT EDIT.

package ent

import (
	"apiGrapqlEntgo/ent/course"
	"apiGrapqlEntgo/ent/professor"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProfessorCreate is the builder for creating a Professor entity.
type ProfessorCreate struct {
	config
	mutation *ProfessorMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (pc *ProfessorCreate) SetName(s string) *ProfessorCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetLastName sets the "last_name" field.
func (pc *ProfessorCreate) SetLastName(s string) *ProfessorCreate {
	pc.mutation.SetLastName(s)
	return pc
}

// SetBirthDate sets the "birth_date" field.
func (pc *ProfessorCreate) SetBirthDate(t time.Time) *ProfessorCreate {
	pc.mutation.SetBirthDate(t)
	return pc
}

// SetCreatedAt sets the "created_at" field.
func (pc *ProfessorCreate) SetCreatedAt(t time.Time) *ProfessorCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (pc *ProfessorCreate) SetNillableCreatedAt(t *time.Time) *ProfessorCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetLastModifiedAt sets the "last_modified_at" field.
func (pc *ProfessorCreate) SetLastModifiedAt(t time.Time) *ProfessorCreate {
	pc.mutation.SetLastModifiedAt(t)
	return pc
}

// SetNillableLastModifiedAt sets the "last_modified_at" field if the given value is not nil.
func (pc *ProfessorCreate) SetNillableLastModifiedAt(t *time.Time) *ProfessorCreate {
	if t != nil {
		pc.SetLastModifiedAt(*t)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *ProfessorCreate) SetID(s string) *ProfessorCreate {
	pc.mutation.SetID(s)
	return pc
}

// AddCourseIDs adds the "courses" edge to the Course entity by IDs.
func (pc *ProfessorCreate) AddCourseIDs(ids ...int) *ProfessorCreate {
	pc.mutation.AddCourseIDs(ids...)
	return pc
}

// AddCourses adds the "courses" edges to the Course entity.
func (pc *ProfessorCreate) AddCourses(c ...*Course) *ProfessorCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pc.AddCourseIDs(ids...)
}

// Mutation returns the ProfessorMutation object of the builder.
func (pc *ProfessorCreate) Mutation() *ProfessorMutation {
	return pc.mutation
}

// Save creates the Professor in the database.
func (pc *ProfessorCreate) Save(ctx context.Context) (*Professor, error) {
	pc.defaults()
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProfessorCreate) SaveX(ctx context.Context) *Professor {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProfessorCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProfessorCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *ProfessorCreate) defaults() {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		v := professor.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.LastModifiedAt(); !ok {
		v := professor.DefaultLastModifiedAt()
		pc.mutation.SetLastModifiedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProfessorCreate) check() error {
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Professor.name"`)}
	}
	if v, ok := pc.mutation.Name(); ok {
		if err := professor.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Professor.name": %w`, err)}
		}
	}
	if _, ok := pc.mutation.LastName(); !ok {
		return &ValidationError{Name: "last_name", err: errors.New(`ent: missing required field "Professor.last_name"`)}
	}
	if v, ok := pc.mutation.LastName(); ok {
		if err := professor.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`ent: validator failed for field "Professor.last_name": %w`, err)}
		}
	}
	if _, ok := pc.mutation.BirthDate(); !ok {
		return &ValidationError{Name: "birth_date", err: errors.New(`ent: missing required field "Professor.birth_date"`)}
	}
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Professor.created_at"`)}
	}
	if _, ok := pc.mutation.LastModifiedAt(); !ok {
		return &ValidationError{Name: "last_modified_at", err: errors.New(`ent: missing required field "Professor.last_modified_at"`)}
	}
	if v, ok := pc.mutation.ID(); ok {
		if err := professor.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Professor.id": %w`, err)}
		}
	}
	return nil
}

func (pc *ProfessorCreate) sqlSave(ctx context.Context) (*Professor, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Professor.ID type: %T", _spec.ID.Value)
		}
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *ProfessorCreate) createSpec() (*Professor, *sqlgraph.CreateSpec) {
	var (
		_node = &Professor{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(professor.Table, sqlgraph.NewFieldSpec(professor.FieldID, field.TypeString))
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(professor.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.LastName(); ok {
		_spec.SetField(professor.FieldLastName, field.TypeString, value)
		_node.LastName = value
	}
	if value, ok := pc.mutation.BirthDate(); ok {
		_spec.SetField(professor.FieldBirthDate, field.TypeTime, value)
		_node.BirthDate = value
	}
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(professor.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.LastModifiedAt(); ok {
		_spec.SetField(professor.FieldLastModifiedAt, field.TypeTime, value)
		_node.LastModifiedAt = value
	}
	if nodes := pc.mutation.CoursesIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ProfessorCreateBulk is the builder for creating many Professor entities in bulk.
type ProfessorCreateBulk struct {
	config
	err      error
	builders []*ProfessorCreate
}

// Save creates the Professor entities in the database.
func (pcb *ProfessorCreateBulk) Save(ctx context.Context) ([]*Professor, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Professor, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProfessorMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProfessorCreateBulk) SaveX(ctx context.Context) []*Professor {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProfessorCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProfessorCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
