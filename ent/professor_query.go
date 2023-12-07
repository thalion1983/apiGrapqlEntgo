// Code generated by ent, DO NOT EDIT.

package ent

import (
	"apiGrapqlEntgo/ent/course"
	"apiGrapqlEntgo/ent/predicate"
	"apiGrapqlEntgo/ent/professor"
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProfessorQuery is the builder for querying Professor entities.
type ProfessorQuery struct {
	config
	ctx         *QueryContext
	order       []professor.OrderOption
	inters      []Interceptor
	predicates  []predicate.Professor
	withCourses *CourseQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProfessorQuery builder.
func (pq *ProfessorQuery) Where(ps ...predicate.Professor) *ProfessorQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *ProfessorQuery) Limit(limit int) *ProfessorQuery {
	pq.ctx.Limit = &limit
	return pq
}

// Offset to start from.
func (pq *ProfessorQuery) Offset(offset int) *ProfessorQuery {
	pq.ctx.Offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *ProfessorQuery) Unique(unique bool) *ProfessorQuery {
	pq.ctx.Unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *ProfessorQuery) Order(o ...professor.OrderOption) *ProfessorQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryCourses chains the current query on the "courses" edge.
func (pq *ProfessorQuery) QueryCourses() *CourseQuery {
	query := (&CourseClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(professor.Table, professor.FieldID, selector),
			sqlgraph.To(course.Table, course.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, professor.CoursesTable, professor.CoursesColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Professor entity from the query.
// Returns a *NotFoundError when no Professor was found.
func (pq *ProfessorQuery) First(ctx context.Context) (*Professor, error) {
	nodes, err := pq.Limit(1).All(setContextOp(ctx, pq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{professor.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *ProfessorQuery) FirstX(ctx context.Context) *Professor {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Professor ID from the query.
// Returns a *NotFoundError when no Professor ID was found.
func (pq *ProfessorQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = pq.Limit(1).IDs(setContextOp(ctx, pq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{professor.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *ProfessorQuery) FirstIDX(ctx context.Context) string {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Professor entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Professor entity is found.
// Returns a *NotFoundError when no Professor entities are found.
func (pq *ProfessorQuery) Only(ctx context.Context) (*Professor, error) {
	nodes, err := pq.Limit(2).All(setContextOp(ctx, pq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{professor.Label}
	default:
		return nil, &NotSingularError{professor.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *ProfessorQuery) OnlyX(ctx context.Context) *Professor {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Professor ID in the query.
// Returns a *NotSingularError when more than one Professor ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *ProfessorQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = pq.Limit(2).IDs(setContextOp(ctx, pq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{professor.Label}
	default:
		err = &NotSingularError{professor.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *ProfessorQuery) OnlyIDX(ctx context.Context) string {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Professors.
func (pq *ProfessorQuery) All(ctx context.Context) ([]*Professor, error) {
	ctx = setContextOp(ctx, pq.ctx, "All")
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Professor, *ProfessorQuery]()
	return withInterceptors[[]*Professor](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *ProfessorQuery) AllX(ctx context.Context) []*Professor {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Professor IDs.
func (pq *ProfessorQuery) IDs(ctx context.Context) (ids []string, err error) {
	if pq.ctx.Unique == nil && pq.path != nil {
		pq.Unique(true)
	}
	ctx = setContextOp(ctx, pq.ctx, "IDs")
	if err = pq.Select(professor.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *ProfessorQuery) IDsX(ctx context.Context) []string {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *ProfessorQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pq.ctx, "Count")
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*ProfessorQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *ProfessorQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *ProfessorQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pq.ctx, "Exist")
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *ProfessorQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProfessorQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *ProfessorQuery) Clone() *ProfessorQuery {
	if pq == nil {
		return nil
	}
	return &ProfessorQuery{
		config:      pq.config,
		ctx:         pq.ctx.Clone(),
		order:       append([]professor.OrderOption{}, pq.order...),
		inters:      append([]Interceptor{}, pq.inters...),
		predicates:  append([]predicate.Professor{}, pq.predicates...),
		withCourses: pq.withCourses.Clone(),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

// WithCourses tells the query-builder to eager-load the nodes that are connected to
// the "courses" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *ProfessorQuery) WithCourses(opts ...func(*CourseQuery)) *ProfessorQuery {
	query := (&CourseClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withCourses = query
	return pq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Professor.Query().
//		GroupBy(professor.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pq *ProfessorQuery) GroupBy(field string, fields ...string) *ProfessorGroupBy {
	pq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ProfessorGroupBy{build: pq}
	grbuild.flds = &pq.ctx.Fields
	grbuild.label = professor.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.Professor.Query().
//		Select(professor.FieldName).
//		Scan(ctx, &v)
func (pq *ProfessorQuery) Select(fields ...string) *ProfessorSelect {
	pq.ctx.Fields = append(pq.ctx.Fields, fields...)
	sbuild := &ProfessorSelect{ProfessorQuery: pq}
	sbuild.label = professor.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ProfessorSelect configured with the given aggregations.
func (pq *ProfessorQuery) Aggregate(fns ...AggregateFunc) *ProfessorSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *ProfessorQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pq); err != nil {
				return err
			}
		}
	}
	for _, f := range pq.ctx.Fields {
		if !professor.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *ProfessorQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Professor, error) {
	var (
		nodes       = []*Professor{}
		_spec       = pq.querySpec()
		loadedTypes = [1]bool{
			pq.withCourses != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Professor).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Professor{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withCourses; query != nil {
		if err := pq.loadCourses(ctx, query, nodes,
			func(n *Professor) { n.Edges.Courses = []*Course{} },
			func(n *Professor, e *Course) { n.Edges.Courses = append(n.Edges.Courses, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *ProfessorQuery) loadCourses(ctx context.Context, query *CourseQuery, nodes []*Professor, init func(*Professor), assign func(*Professor, *Course)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[string]*Professor)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Course(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(professor.CoursesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.professor_courses
		if fk == nil {
			return fmt.Errorf(`foreign-key "professor_courses" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "professor_courses" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (pq *ProfessorQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	_spec.Node.Columns = pq.ctx.Fields
	if len(pq.ctx.Fields) > 0 {
		_spec.Unique = pq.ctx.Unique != nil && *pq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *ProfessorQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(professor.Table, professor.Columns, sqlgraph.NewFieldSpec(professor.FieldID, field.TypeString))
	_spec.From = pq.sql
	if unique := pq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pq.path != nil {
		_spec.Unique = true
	}
	if fields := pq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, professor.FieldID)
		for i := range fields {
			if fields[i] != professor.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *ProfessorQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(professor.Table)
	columns := pq.ctx.Fields
	if len(columns) == 0 {
		columns = professor.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.ctx.Unique != nil && *pq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProfessorGroupBy is the group-by builder for Professor entities.
type ProfessorGroupBy struct {
	selector
	build *ProfessorQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *ProfessorGroupBy) Aggregate(fns ...AggregateFunc) *ProfessorGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *ProfessorGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, "GroupBy")
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProfessorQuery, *ProfessorGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *ProfessorGroupBy) sqlScan(ctx context.Context, root *ProfessorQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pgb.flds)+len(pgb.fns))
		for _, f := range *pgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ProfessorSelect is the builder for selecting fields of Professor entities.
type ProfessorSelect struct {
	*ProfessorQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *ProfessorSelect) Aggregate(fns ...AggregateFunc) *ProfessorSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *ProfessorSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, "Select")
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ProfessorQuery, *ProfessorSelect](ctx, ps.ProfessorQuery, ps, ps.inters, v)
}

func (ps *ProfessorSelect) sqlScan(ctx context.Context, root *ProfessorQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}