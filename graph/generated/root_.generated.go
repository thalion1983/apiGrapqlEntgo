// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"apiGrapqlEntgo/graph/model"
	"bytes"
	"context"
	"errors"
	"sync/atomic"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/introspection"
	gqlparser "github.com/vektah/gqlparser/v2"
	"github.com/vektah/gqlparser/v2/ast"
)

// NewExecutableSchema creates an ExecutableSchema from the ResolverRoot interface.
func NewExecutableSchema(cfg Config) graphql.ExecutableSchema {
	return &executableSchema{
		schema:     cfg.Schema,
		resolvers:  cfg.Resolvers,
		directives: cfg.Directives,
		complexity: cfg.Complexity,
	}
}

type Config struct {
	Schema     *ast.Schema
	Resolvers  ResolverRoot
	Directives DirectiveRoot
	Complexity ComplexityRoot
}

type ResolverRoot interface {
	Mutation() MutationResolver
	Query() QueryResolver
}

type DirectiveRoot struct {
}

type ComplexityRoot struct {
	Mutation struct {
		CreateProfessor func(childComplexity int, input model.NewProfessor) int
		RemoveProfessor func(childComplexity int, id string) int
	}

	Professor struct {
		BirthDate      func(childComplexity int) int
		CreatedAt      func(childComplexity int) int
		ID             func(childComplexity int) int
		LastModifiedAt func(childComplexity int) int
		LastName       func(childComplexity int) int
		Name           func(childComplexity int) int
	}

	Query struct {
		Professor  func(childComplexity int, id string) int
		Professors func(childComplexity int) int
	}
}

type executableSchema struct {
	schema     *ast.Schema
	resolvers  ResolverRoot
	directives DirectiveRoot
	complexity ComplexityRoot
}

func (e *executableSchema) Schema() *ast.Schema {
	if e.schema != nil {
		return e.schema
	}
	return parsedSchema
}

func (e *executableSchema) Complexity(typeName, field string, childComplexity int, rawArgs map[string]interface{}) (int, bool) {
	ec := executionContext{nil, e, 0, 0, nil}
	_ = ec
	switch typeName + "." + field {

	case "Mutation.createProfessor":
		if e.complexity.Mutation.CreateProfessor == nil {
			break
		}

		args, err := ec.field_Mutation_createProfessor_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.CreateProfessor(childComplexity, args["input"].(model.NewProfessor)), true

	case "Mutation.removeProfessor":
		if e.complexity.Mutation.RemoveProfessor == nil {
			break
		}

		args, err := ec.field_Mutation_removeProfessor_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Mutation.RemoveProfessor(childComplexity, args["id"].(string)), true

	case "Professor.birth_date":
		if e.complexity.Professor.BirthDate == nil {
			break
		}

		return e.complexity.Professor.BirthDate(childComplexity), true

	case "Professor.created_at":
		if e.complexity.Professor.CreatedAt == nil {
			break
		}

		return e.complexity.Professor.CreatedAt(childComplexity), true

	case "Professor.id":
		if e.complexity.Professor.ID == nil {
			break
		}

		return e.complexity.Professor.ID(childComplexity), true

	case "Professor.last_modified_at":
		if e.complexity.Professor.LastModifiedAt == nil {
			break
		}

		return e.complexity.Professor.LastModifiedAt(childComplexity), true

	case "Professor.last_name":
		if e.complexity.Professor.LastName == nil {
			break
		}

		return e.complexity.Professor.LastName(childComplexity), true

	case "Professor.name":
		if e.complexity.Professor.Name == nil {
			break
		}

		return e.complexity.Professor.Name(childComplexity), true

	case "Query.professor":
		if e.complexity.Query.Professor == nil {
			break
		}

		args, err := ec.field_Query_professor_args(context.TODO(), rawArgs)
		if err != nil {
			return 0, false
		}

		return e.complexity.Query.Professor(childComplexity, args["id"].(string)), true

	case "Query.professors":
		if e.complexity.Query.Professors == nil {
			break
		}

		return e.complexity.Query.Professors(childComplexity), true

	}
	return 0, false
}

func (e *executableSchema) Exec(ctx context.Context) graphql.ResponseHandler {
	rc := graphql.GetOperationContext(ctx)
	ec := executionContext{rc, e, 0, 0, make(chan graphql.DeferredResult)}
	inputUnmarshalMap := graphql.BuildUnmarshalerMap(
		ec.unmarshalInputNewProfessor,
	)
	first := true

	switch rc.Operation.Operation {
	case ast.Query:
		return func(ctx context.Context) *graphql.Response {
			var response graphql.Response
			var data graphql.Marshaler
			if first {
				first = false
				ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
				data = ec._Query(ctx, rc.Operation.SelectionSet)
			} else {
				if atomic.LoadInt32(&ec.pendingDeferred) > 0 {
					result := <-ec.deferredResults
					atomic.AddInt32(&ec.pendingDeferred, -1)
					data = result.Result
					response.Path = result.Path
					response.Label = result.Label
					response.Errors = result.Errors
				} else {
					return nil
				}
			}
			var buf bytes.Buffer
			data.MarshalGQL(&buf)
			response.Data = buf.Bytes()
			if atomic.LoadInt32(&ec.deferred) > 0 {
				hasNext := atomic.LoadInt32(&ec.pendingDeferred) > 0
				response.HasNext = &hasNext
			}

			return &response
		}
	case ast.Mutation:
		return func(ctx context.Context) *graphql.Response {
			if !first {
				return nil
			}
			first = false
			ctx = graphql.WithUnmarshalerMap(ctx, inputUnmarshalMap)
			data := ec._Mutation(ctx, rc.Operation.SelectionSet)
			var buf bytes.Buffer
			data.MarshalGQL(&buf)

			return &graphql.Response{
				Data: buf.Bytes(),
			}
		}

	default:
		return graphql.OneShot(graphql.ErrorResponse(ctx, "unsupported GraphQL operation"))
	}
}

type executionContext struct {
	*graphql.OperationContext
	*executableSchema
	deferred        int32
	pendingDeferred int32
	deferredResults chan graphql.DeferredResult
}

func (ec *executionContext) processDeferredGroup(dg graphql.DeferredGroup) {
	atomic.AddInt32(&ec.pendingDeferred, 1)
	go func() {
		ctx := graphql.WithFreshResponseContext(dg.Context)
		dg.FieldSet.Dispatch(ctx)
		ds := graphql.DeferredResult{
			Path:   dg.Path,
			Label:  dg.Label,
			Result: dg.FieldSet,
			Errors: graphql.GetErrors(ctx),
		}
		// null fields should bubble up
		if dg.FieldSet.Invalids > 0 {
			ds.Result = graphql.Null
		}
		ec.deferredResults <- ds
	}()
}

func (ec *executionContext) introspectSchema() (*introspection.Schema, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapSchema(ec.Schema()), nil
}

func (ec *executionContext) introspectType(name string) (*introspection.Type, error) {
	if ec.DisableIntrospection {
		return nil, errors.New("introspection disabled")
	}
	return introspection.WrapTypeFromDef(ec.Schema(), ec.Schema().Types[name]), nil
}

var sources = []*ast.Source{
	{Name: "../schema/mutation.graphql", Input: `type Mutation {
	createProfessor(input: NewProfessor!): Professor!
	removeProfessor(id: String!): Professor!
}
`, BuiltIn: false},
	{Name: "../schema/query.graphql", Input: `type Query {
	professors: [Professor!]!
	professor(id: String!): Professor!
}
`, BuiltIn: false},
	{Name: "../schema/types.graphql", Input: `input NewProfessor {
	id: String!
	name: String!
	last_name: String!
	birth_date: String!
}

type Professor {
	id: String!
	name: String!
	last_name: String!
	birth_date: String!
	created_at: String!
	last_modified_at: String!
}
`, BuiltIn: false},
}
var parsedSchema = gqlparser.MustLoadSchema(sources...)
