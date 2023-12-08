package graph

import (
	"apiGrapqlEntgo/ent"
	"apiGrapqlEntgo/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Cli *ent.Client
}

var (
	dateInputLayout  = "2006/01/02"
	dateOutputLayout = "2006/01/02T15:04:05"
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

// marshalEntSubject marshals an object Subject from ent into model.Subject
func marshalEntSubject(subj *ent.Subject) *model.Subject {
	return &model.Subject{
		ID:             subj.ID,
		Name:           subj.Name,
		Description:    subj.Description,
		Active:         subj.Active,
		CreatedAt:      subj.CreatedAt.Format(dateOutputLayout),
		LastModifiedAt: subj.LastModifiedAt.Format(dateOutputLayout),
	}
}

// marshalEntCourse marshals an object Course from ent into model.Course
func marshalEntCourse(cour *ent.Course) *model.Course {
	return &model.Course{
		ID:             cour.ID,
		Year:           cour.Year,
		Period:         cour.Period,
		ProfessorID:    cour.ProfessorID,
		SubjectID:      cour.SubjectID,
		CreatedAt:      cour.CreatedAt.Format(dateOutputLayout),
		LastModifiedAt: cour.LastModifiedAt.Format(dateOutputLayout),
	}
}
