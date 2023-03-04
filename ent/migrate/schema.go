// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// PetsColumns holds the columns for the "pets" table.
	PetsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString, Size: 10},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"DOG", "CAT"}},
		{Name: "code", Type: field.TypeString, Unique: true},
		{Name: "age_month", Type: field.TypeInt},
		{Name: "is_deleted", Type: field.TypeBool, Default: false},
		{Name: "created_by", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
	}
	// PetsTable holds the schema information for the "pets" table.
	PetsTable = &schema.Table{
		Name:       "pets",
		Columns:    PetsColumns,
		PrimaryKey: []*schema.Column{PetsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "pet_name",
				Unique:  false,
				Columns: []*schema.Column{PetsColumns[1]},
			},
			{
				Name:    "pet_type",
				Unique:  false,
				Columns: []*schema.Column{PetsColumns[2]},
			},
		},
	}
	// PostsColumns holds the columns for the "posts" table.
	PostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString},
		{Name: "content", Type: field.TypeString},
		{Name: "slug", Type: field.TypeString},
		{Name: "status", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// PostsTable holds the schema information for the "posts" table.
	PostsTable = &schema.Table{
		Name:       "posts",
		Columns:    PostsColumns,
		PrimaryKey: []*schema.Column{PostsColumns[0]},
	}
	// SystemParametersColumns holds the columns for the "system_parameters" table.
	SystemParametersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "key", Type: field.TypeString, Unique: true},
		{Name: "value", Type: field.TypeString},
		{Name: "is_deleted", Type: field.TypeBool, Default: false},
		{Name: "created_by", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
	}
	// SystemParametersTable holds the schema information for the "system_parameters" table.
	SystemParametersTable = &schema.Table{
		Name:       "system_parameters",
		Columns:    SystemParametersColumns,
		PrimaryKey: []*schema.Column{SystemParametersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "system_parameter_key",
				Unique:  false,
				Columns: []*schema.Column{SystemParametersColumns[1]},
			},
			{
				Name:    "system_parameter_value",
				Unique:  false,
				Columns: []*schema.Column{SystemParametersColumns[2]},
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "name", Type: field.TypeString, Size: 10},
		{Name: "email", Type: field.TypeString},
		{Name: "phone", Type: field.TypeString},
		{Name: "is_deleted", Type: field.TypeBool, Default: false},
		{Name: "created_by", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_by", Type: field.TypeString, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "user_email",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[2]},
			},
			{
				Name:    "user_phone",
				Unique:  true,
				Columns: []*schema.Column{UsersColumns[3]},
			},
			{
				Name:    "user_email_phone",
				Unique:  false,
				Columns: []*schema.Column{UsersColumns[2], UsersColumns[3]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		PetsTable,
		PostsTable,
		SystemParametersTable,
		UsersTable,
	}
)

func init() {
}
