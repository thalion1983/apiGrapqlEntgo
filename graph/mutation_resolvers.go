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
	"time"
)

// CreateProfessor is the resolver for the createProfessor field.
func (r *mutationResolver) CreateProfessor(ctx context.Context, input model.NewProfessor) (*model.Professor, error) {
	birthDate, err := time.Parse(dateInputLayout, input.BirthDate)
	if err != nil {
		return nil, fmt.Errorf("parsing birthdate: %w", err)
	}

	prof, err := r.Cli.Professor.Create().
		SetID(input.ID).
		SetName(input.Name).
		SetLastName(input.LastName).
		SetBirthDate(birthDate).
		Save(r.Ctx)

	if err != nil {
		return nil, fmt.Errorf("creating new professor: %w", err)
	}

	return marshalEntProfessor(prof), nil
}

// RemoveProfessor is the resolver for the removeProfessor field.
func (r *mutationResolver) RemoveProfessor(ctx context.Context, id string) (*model.Professor, error) {
	prof, err := r.Cli.Professor.Query().
		Where(professor.ID(id)).
		Only(r.Ctx)
	if err != nil {
		return nil, fmt.Errorf("getting professor by ID: %w", err)
	}

	res := marshalEntProfessor(prof)
	if err = r.Cli.Professor.DeleteOneID(id).Exec(r.Ctx); err != nil {
		return nil, fmt.Errorf("removing professor: %w", err)
	}

	return res, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
