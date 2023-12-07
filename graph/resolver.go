package graph

import (
	"apiGrapqlEntgo/ent"
	"context"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Cli *ent.Client
	Ctx context.Context
}

var (
	dateInputLayout  = "2006/01/02"
	dateOutputLayout = "2006-01-02T15:04:05"
)

// marshalEntProfessor marshals an object Professor from ent into model.Professor
func marshalEntProfessor(prof *ent.Professor) *model.Professor {
	return &model.Professor{
		ID:             prof.ID,
		Name:           prof.Name,
		LastName:       prof.LastName,
		BirthDate:      prof.BirthDate.Format(dateInputLayout),
		CreatedAt:      prof.CreatedAt.Format(dateOutputLayout),
		LastModifiedAt: prof.LastModifiedAt.Format(dateOutputLayout),
	}
}
