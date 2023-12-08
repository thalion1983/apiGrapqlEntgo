package graph

import (
	"apiGrapqlEntgo/ent"
	"apiGrapqlEntgo/ent/course"
	"apiGrapqlEntgo/ent/professor"
	"apiGrapqlEntgo/ent/subject"
	"apiGrapqlEntgo/graph/model"
	"context"
	"fmt"
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

// queryProfessor queries a Professor by ID
func queryProfessor(cli *ent.Client, ctx context.Context, id string) (*ent.Professor, error) {
	prof, err := cli.Professor.Query().
		Where(professor.ID(id)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting professor by ID: %w", err)
	}

	return prof, nil
}

// querySubject queries a Subject by ID
func querySubject(cli *ent.Client, ctx context.Context, id string) (*ent.Subject, error) {
	subj, err := cli.Subject.Query().
		Where(subject.ID(id)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting subject by ID: %w", err)
	}

	return subj, nil
}

// queryCourseByID queries a Course by ID
func queryCourseByID(cli *ent.Client, ctx context.Context, id int) (*ent.Course, error) {
	cour, err := cli.Course.Query().
		Where(course.ID(id)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting course by ID: %w", err)
	}

	return cour, nil
}

// queryCourse queries a Course by its data: Year, Period and SubjectID
func queryCourse(cli *ent.Client, ctx context.Context, year int, period int, subjectID string) (*ent.Course, error) {
	cour, err := cli.Course.Query().
		Where(
			course.Year(year),
			course.Period(period),
			course.SubjectID(subjectID),
		).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("getting course for subject_id %s in period %d-%d: %w", subjectID, year, period, err)
	}

	return cour, nil
}
