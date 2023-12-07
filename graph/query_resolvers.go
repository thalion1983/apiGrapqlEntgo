package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"apiGrapqlEntgo/ent/professor"
	"apiGrapqlEntgo/graph/generated"
	"apiGrapqlEntgo/graph/model"
	"context"
	"fmt"
)

// Professors is the resolver for the professors field.
func (r *queryResolver) Professors(ctx context.Context) ([]*model.Professor, error) {
	profList, err := r.Cli.Professor.Query().All(r.Ctx)
	if err != nil {
		return nil, fmt.Errorf("getting professors list: %w", err)
	}

	var res []*model.Professor
	for _, prof := range profList {
		res = append(res, marshalEntProfessor(prof))
	}
	return res, nil
}

// Professor is the resolver for the professor field.
func (r *queryResolver) Professor(ctx context.Context, id string) (*model.Professor, error) {
	prof, err := r.Cli.Professor.Query().
		Where(professor.ID(id)).
		Only(r.Ctx)
	if err != nil {
		return nil, fmt.Errorf("getting professor by ID: %w", err)
	}

	return marshalEntProfessor(prof), nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
