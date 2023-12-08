// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package generated

import (
	"apiGrapqlEntgo/graph/model"
	"context"
	"fmt"
	"strconv"
	"sync/atomic"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

type MutationResolver interface {
	CreateProfessor(ctx context.Context, input model.NewProfessor) (*model.Professor, error)
	RemoveProfessor(ctx context.Context, id string) (*model.Professor, error)
	UpdateProfessor(ctx context.Context, id string, input model.NewProfessor) (*model.Professor, error)
	CreateSubject(ctx context.Context, input model.NewSubject) (*model.Subject, error)
	RemoveSubject(ctx context.Context, id string) (*model.Subject, error)
	UpdateSubject(ctx context.Context, id string, input model.NewSubject) (*model.Subject, error)
}

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

func (ec *executionContext) field_Mutation_createProfessor_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 model.NewProfessor
	if tmp, ok := rawArgs["input"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
		arg0, err = ec.unmarshalNNewProfessor2apiGrapqlEntgoᚋgraphᚋmodelᚐNewProfessor(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg0
	return args, nil
}

func (ec *executionContext) field_Mutation_createSubject_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 model.NewSubject
	if tmp, ok := rawArgs["input"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
		arg0, err = ec.unmarshalNNewSubject2apiGrapqlEntgoᚋgraphᚋmodelᚐNewSubject(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg0
	return args, nil
}

func (ec *executionContext) field_Mutation_removeProfessor_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
		arg0, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	return args, nil
}

func (ec *executionContext) field_Mutation_removeSubject_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
		arg0, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	return args, nil
}

func (ec *executionContext) field_Mutation_updateProfessor_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
		arg0, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	var arg1 model.NewProfessor
	if tmp, ok := rawArgs["input"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
		arg1, err = ec.unmarshalNNewProfessor2apiGrapqlEntgoᚋgraphᚋmodelᚐNewProfessor(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg1
	return args, nil
}

func (ec *executionContext) field_Mutation_updateSubject_args(ctx context.Context, rawArgs map[string]interface{}) (map[string]interface{}, error) {
	var err error
	args := map[string]interface{}{}
	var arg0 string
	if tmp, ok := rawArgs["id"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("id"))
		arg0, err = ec.unmarshalNString2string(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["id"] = arg0
	var arg1 model.NewSubject
	if tmp, ok := rawArgs["input"]; ok {
		ctx := graphql.WithPathContext(ctx, graphql.NewPathWithField("input"))
		arg1, err = ec.unmarshalNNewSubject2apiGrapqlEntgoᚋgraphᚋmodelᚐNewSubject(ctx, tmp)
		if err != nil {
			return nil, err
		}
	}
	args["input"] = arg1
	return args, nil
}

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _Mutation_createProfessor(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_createProfessor(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().CreateProfessor(rctx, fc.Args["input"].(model.NewProfessor))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*model.Professor)
	fc.Result = res
	return ec.marshalNProfessor2ᚖapiGrapqlEntgoᚋgraphᚋmodelᚐProfessor(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_createProfessor(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Professor_id(ctx, field)
			case "name":
				return ec.fieldContext_Professor_name(ctx, field)
			case "last_name":
				return ec.fieldContext_Professor_last_name(ctx, field)
			case "birth_date":
				return ec.fieldContext_Professor_birth_date(ctx, field)
			case "created_at":
				return ec.fieldContext_Professor_created_at(ctx, field)
			case "last_modified_at":
				return ec.fieldContext_Professor_last_modified_at(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Professor", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_createProfessor_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_removeProfessor(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_removeProfessor(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().RemoveProfessor(rctx, fc.Args["id"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*model.Professor)
	fc.Result = res
	return ec.marshalNProfessor2ᚖapiGrapqlEntgoᚋgraphᚋmodelᚐProfessor(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_removeProfessor(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Professor_id(ctx, field)
			case "name":
				return ec.fieldContext_Professor_name(ctx, field)
			case "last_name":
				return ec.fieldContext_Professor_last_name(ctx, field)
			case "birth_date":
				return ec.fieldContext_Professor_birth_date(ctx, field)
			case "created_at":
				return ec.fieldContext_Professor_created_at(ctx, field)
			case "last_modified_at":
				return ec.fieldContext_Professor_last_modified_at(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Professor", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_removeProfessor_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_updateProfessor(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_updateProfessor(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().UpdateProfessor(rctx, fc.Args["id"].(string), fc.Args["input"].(model.NewProfessor))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*model.Professor)
	fc.Result = res
	return ec.marshalNProfessor2ᚖapiGrapqlEntgoᚋgraphᚋmodelᚐProfessor(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_updateProfessor(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Professor_id(ctx, field)
			case "name":
				return ec.fieldContext_Professor_name(ctx, field)
			case "last_name":
				return ec.fieldContext_Professor_last_name(ctx, field)
			case "birth_date":
				return ec.fieldContext_Professor_birth_date(ctx, field)
			case "created_at":
				return ec.fieldContext_Professor_created_at(ctx, field)
			case "last_modified_at":
				return ec.fieldContext_Professor_last_modified_at(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Professor", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_updateProfessor_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_createSubject(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_createSubject(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().CreateSubject(rctx, fc.Args["input"].(model.NewSubject))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*model.Subject)
	fc.Result = res
	return ec.marshalNSubject2ᚖapiGrapqlEntgoᚋgraphᚋmodelᚐSubject(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_createSubject(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Subject_id(ctx, field)
			case "name":
				return ec.fieldContext_Subject_name(ctx, field)
			case "description":
				return ec.fieldContext_Subject_description(ctx, field)
			case "active":
				return ec.fieldContext_Subject_active(ctx, field)
			case "created_at":
				return ec.fieldContext_Subject_created_at(ctx, field)
			case "last_modified_at":
				return ec.fieldContext_Subject_last_modified_at(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Subject", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_createSubject_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_removeSubject(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_removeSubject(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().RemoveSubject(rctx, fc.Args["id"].(string))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*model.Subject)
	fc.Result = res
	return ec.marshalNSubject2ᚖapiGrapqlEntgoᚋgraphᚋmodelᚐSubject(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_removeSubject(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Subject_id(ctx, field)
			case "name":
				return ec.fieldContext_Subject_name(ctx, field)
			case "description":
				return ec.fieldContext_Subject_description(ctx, field)
			case "active":
				return ec.fieldContext_Subject_active(ctx, field)
			case "created_at":
				return ec.fieldContext_Subject_created_at(ctx, field)
			case "last_modified_at":
				return ec.fieldContext_Subject_last_modified_at(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Subject", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_removeSubject_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

func (ec *executionContext) _Mutation_updateSubject(ctx context.Context, field graphql.CollectedField) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Mutation_updateSubject(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp, err := ec.ResolverMiddleware(ctx, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Mutation().UpdateSubject(rctx, fc.Args["id"].(string), fc.Args["input"].(model.NewSubject))
	})
	if err != nil {
		ec.Error(ctx, err)
		return graphql.Null
	}
	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*model.Subject)
	fc.Result = res
	return ec.marshalNSubject2ᚖapiGrapqlEntgoᚋgraphᚋmodelᚐSubject(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Mutation_updateSubject(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Mutation",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Subject_id(ctx, field)
			case "name":
				return ec.fieldContext_Subject_name(ctx, field)
			case "description":
				return ec.fieldContext_Subject_description(ctx, field)
			case "active":
				return ec.fieldContext_Subject_active(ctx, field)
			case "created_at":
				return ec.fieldContext_Subject_created_at(ctx, field)
			case "last_modified_at":
				return ec.fieldContext_Subject_last_modified_at(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Subject", field.Name)
		},
	}
	defer func() {
		if r := recover(); r != nil {
			err = ec.Recover(ctx, r)
			ec.Error(ctx, err)
		}
	}()
	ctx = graphql.WithFieldContext(ctx, fc)
	if fc.Args, err = ec.field_Mutation_updateSubject_args(ctx, field.ArgumentMap(ec.Variables)); err != nil {
		ec.Error(ctx, err)
		return fc, err
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var mutationImplementors = []string{"Mutation"}

func (ec *executionContext) _Mutation(ctx context.Context, sel ast.SelectionSet) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, mutationImplementors)
	ctx = graphql.WithFieldContext(ctx, &graphql.FieldContext{
		Object: "Mutation",
	})

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		innerCtx := graphql.WithRootFieldContext(ctx, &graphql.RootFieldContext{
			Object: field.Name,
			Field:  field,
		})

		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Mutation")
		case "createProfessor":
			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_createProfessor(ctx, field)
			})
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "removeProfessor":
			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_removeProfessor(ctx, field)
			})
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "updateProfessor":
			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_updateProfessor(ctx, field)
			})
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "createSubject":
			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_createSubject(ctx, field)
			})
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "removeSubject":
			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_removeSubject(ctx, field)
			})
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "updateSubject":
			out.Values[i] = ec.OperationContext.RootResolverMiddleware(innerCtx, func(ctx context.Context) (res graphql.Marshaler) {
				return ec._Mutation_updateSubject(ctx, field)
			})
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	atomic.AddInt32(&ec.deferred, int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

// endregion ***************************** type.gotpl *****************************
