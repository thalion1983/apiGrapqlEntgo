// Code generated by ent, DO NOT EDIT.

package ent

import (
	"apiGrapqlEntgo/ent/course"
	"apiGrapqlEntgo/ent/professor"
	"apiGrapqlEntgo/ent/schema"
	"apiGrapqlEntgo/ent/subject"
	"time"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	courseFields := schema.Course{}.Fields()
	_ = courseFields
	// courseDescCreatedAt is the schema descriptor for created_at field.
	courseDescCreatedAt := courseFields[2].Descriptor()
	// course.DefaultCreatedAt holds the default value on creation for the created_at field.
	course.DefaultCreatedAt = courseDescCreatedAt.Default.(func() time.Time)
	// courseDescLastModifiedAt is the schema descriptor for last_modified_at field.
	courseDescLastModifiedAt := courseFields[3].Descriptor()
	// course.DefaultLastModifiedAt holds the default value on creation for the last_modified_at field.
	course.DefaultLastModifiedAt = courseDescLastModifiedAt.Default.(func() time.Time)
	// course.UpdateDefaultLastModifiedAt holds the default value on update for the last_modified_at field.
	course.UpdateDefaultLastModifiedAt = courseDescLastModifiedAt.UpdateDefault.(func() time.Time)
	professorFields := schema.Professor{}.Fields()
	_ = professorFields
	// professorDescName is the schema descriptor for name field.
	professorDescName := professorFields[1].Descriptor()
	// professor.NameValidator is a validator for the "name" field. It is called by the builders before save.
	professor.NameValidator = professorDescName.Validators[0].(func(string) error)
	// professorDescLastName is the schema descriptor for last_name field.
	professorDescLastName := professorFields[2].Descriptor()
	// professor.LastNameValidator is a validator for the "last_name" field. It is called by the builders before save.
	professor.LastNameValidator = professorDescLastName.Validators[0].(func(string) error)
	// professorDescCreatedAt is the schema descriptor for created_at field.
	professorDescCreatedAt := professorFields[4].Descriptor()
	// professor.DefaultCreatedAt holds the default value on creation for the created_at field.
	professor.DefaultCreatedAt = professorDescCreatedAt.Default.(func() time.Time)
	// professorDescLastModifiedAt is the schema descriptor for last_modified_at field.
	professorDescLastModifiedAt := professorFields[5].Descriptor()
	// professor.DefaultLastModifiedAt holds the default value on creation for the last_modified_at field.
	professor.DefaultLastModifiedAt = professorDescLastModifiedAt.Default.(func() time.Time)
	// professor.UpdateDefaultLastModifiedAt holds the default value on update for the last_modified_at field.
	professor.UpdateDefaultLastModifiedAt = professorDescLastModifiedAt.UpdateDefault.(func() time.Time)
	// professorDescID is the schema descriptor for id field.
	professorDescID := professorFields[0].Descriptor()
	// professor.IDValidator is a validator for the "id" field. It is called by the builders before save.
	professor.IDValidator = professorDescID.Validators[0].(func(string) error)
	subjectFields := schema.Subject{}.Fields()
	_ = subjectFields
	// subjectDescName is the schema descriptor for name field.
	subjectDescName := subjectFields[1].Descriptor()
	// subject.NameValidator is a validator for the "name" field. It is called by the builders before save.
	subject.NameValidator = subjectDescName.Validators[0].(func(string) error)
	// subjectDescDescription is the schema descriptor for description field.
	subjectDescDescription := subjectFields[2].Descriptor()
	// subject.DescriptionValidator is a validator for the "description" field. It is called by the builders before save.
	subject.DescriptionValidator = subjectDescDescription.Validators[0].(func(string) error)
	// subjectDescActive is the schema descriptor for active field.
	subjectDescActive := subjectFields[3].Descriptor()
	// subject.DefaultActive holds the default value on creation for the active field.
	subject.DefaultActive = subjectDescActive.Default.(bool)
	// subjectDescCreatedAt is the schema descriptor for created_at field.
	subjectDescCreatedAt := subjectFields[4].Descriptor()
	// subject.DefaultCreatedAt holds the default value on creation for the created_at field.
	subject.DefaultCreatedAt = subjectDescCreatedAt.Default.(func() time.Time)
	// subjectDescLastModifiedAt is the schema descriptor for last_modified_at field.
	subjectDescLastModifiedAt := subjectFields[5].Descriptor()
	// subject.DefaultLastModifiedAt holds the default value on creation for the last_modified_at field.
	subject.DefaultLastModifiedAt = subjectDescLastModifiedAt.Default.(func() time.Time)
	// subject.UpdateDefaultLastModifiedAt holds the default value on update for the last_modified_at field.
	subject.UpdateDefaultLastModifiedAt = subjectDescLastModifiedAt.UpdateDefault.(func() time.Time)
}
