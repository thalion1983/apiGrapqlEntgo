// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CoursesColumns holds the columns for the "courses" table.
	CoursesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "year", Type: field.TypeInt},
		{Name: "period", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "last_modified_at", Type: field.TypeTime},
		{Name: "subject_id", Type: field.TypeString, SchemaType: map[string]string{"postgres": "char(6)"}},
		{Name: "professor_id", Type: field.TypeString},
	}
	// CoursesTable holds the schema information for the "courses" table.
	CoursesTable = &schema.Table{
		Name:       "courses",
		Columns:    CoursesColumns,
		PrimaryKey: []*schema.Column{CoursesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "courses_subjects_subject",
				Columns:    []*schema.Column{CoursesColumns[5]},
				RefColumns: []*schema.Column{SubjectsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "courses_professors_professor",
				Columns:    []*schema.Column{CoursesColumns[6]},
				RefColumns: []*schema.Column{ProfessorsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "course_year_period_subject_id",
				Unique:  true,
				Columns: []*schema.Column{CoursesColumns[1], CoursesColumns[2], CoursesColumns[5]},
			},
		},
	}
	// ProfessorsColumns holds the columns for the "professors" table.
	ProfessorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "last_name", Type: field.TypeString},
		{Name: "birth_date", Type: field.TypeTime},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "last_modified_at", Type: field.TypeTime},
	}
	// ProfessorsTable holds the schema information for the "professors" table.
	ProfessorsTable = &schema.Table{
		Name:       "professors",
		Columns:    ProfessorsColumns,
		PrimaryKey: []*schema.Column{ProfessorsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "professor_id",
				Unique:  true,
				Columns: []*schema.Column{ProfessorsColumns[0]},
			},
		},
	}
	// SubjectsColumns holds the columns for the "subjects" table.
	SubjectsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true, SchemaType: map[string]string{"postgres": "char(6)"}},
		{Name: "name", Type: field.TypeString},
		{Name: "description", Type: field.TypeString},
		{Name: "active", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "last_modified_at", Type: field.TypeTime},
	}
	// SubjectsTable holds the schema information for the "subjects" table.
	SubjectsTable = &schema.Table{
		Name:       "subjects",
		Columns:    SubjectsColumns,
		PrimaryKey: []*schema.Column{SubjectsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "subject_id",
				Unique:  true,
				Columns: []*schema.Column{SubjectsColumns[0]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CoursesTable,
		ProfessorsTable,
		SubjectsTable,
	}
)

func init() {
	CoursesTable.ForeignKeys[0].RefTable = SubjectsTable
	CoursesTable.ForeignKeys[1].RefTable = ProfessorsTable
}
