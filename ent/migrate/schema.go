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
	// RolesColumns holds the columns for the "roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "text", Type: field.TypeString},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// RolesTable holds the schema information for the "roles" table.
	RolesTable = &schema.Table{
		Name:       "roles",
		Columns:    RolesColumns,
		PrimaryKey: []*schema.Column{RolesColumns[0]},
	}
	// RoleUsersColumns holds the columns for the "role_users" table.
	RoleUsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "user_id", Type: field.TypeUint64, Nullable: true},
		{Name: "role_id", Type: field.TypeUint64, Nullable: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
	}
	// RoleUsersTable holds the schema information for the "role_users" table.
	RoleUsersTable = &schema.Table{
		Name:       "role_users",
		Columns:    RoleUsersColumns,
		PrimaryKey: []*schema.Column{RoleUsersColumns[0]},
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
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "is_verified", Type: field.TypeBool},
		{Name: "email_verified_at", Type: field.TypeTime, Nullable: true},
		{Name: "password", Type: field.TypeString, Nullable: true},
		{Name: "remember_token", Type: field.TypeString, Nullable: true},
		{Name: "social_media_id", Type: field.TypeString, Nullable: true},
		{Name: "avatar", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "role_id", Type: field.TypeInt32},
		{Name: "login_type", Type: field.TypeString, Nullable: true},
		{Name: "sub_specialist", Type: field.TypeString, Nullable: true},
		{Name: "firebase_token", Type: field.TypeString, Nullable: true},
		{Name: "info", Type: field.TypeString, Nullable: true},
		{Name: "description", Type: field.TypeString, Nullable: true},
		{Name: "specialist", Type: field.TypeString, Nullable: true},
		{Name: "phone", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "last_access_at", Type: field.TypeTime},
		{Name: "pregnancy_mode", Type: field.TypeBool},
		{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
		{Name: "latest_skip_update", Type: field.TypeTime, Nullable: true},
		{Name: "latest_deleted_at", Type: field.TypeTime, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserDevicesColumns holds the columns for the "user_devices" table.
	UserDevicesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUint64, Increment: true},
		{Name: "user_id", Type: field.TypeUint64},
		{Name: "version", Type: field.TypeString},
		{Name: "platform", Type: field.TypeString},
		{Name: "latest_skip_update", Type: field.TypeTime, Nullable: true},
		{Name: "created_at", Type: field.TypeTime, Nullable: true},
		{Name: "updated_at", Type: field.TypeTime, Nullable: true},
		{Name: "device_id", Type: field.TypeString, Nullable: true},
	}
	// UserDevicesTable holds the schema information for the "user_devices" table.
	UserDevicesTable = &schema.Table{
		Name:       "user_devices",
		Columns:    UserDevicesColumns,
		PrimaryKey: []*schema.Column{UserDevicesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		PetsTable,
		RolesTable,
		RoleUsersTable,
		SystemParametersTable,
		UsersTable,
		UserDevicesTable,
	}
)

func init() {
}
