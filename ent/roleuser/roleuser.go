// Code generated by ent, DO NOT EDIT.

package roleuser

const (
	// Label holds the string label denoting the roleuser type in the database.
	Label = "role_user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldRoleID holds the string denoting the role_id field in the database.
	FieldRoleID = "role_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// Table holds the table name of the roleuser in the database.
	Table = "role_users"
)

// Columns holds all SQL columns for roleuser fields.
var Columns = []string{
	FieldID,
	FieldUserID,
	FieldRoleID,
	FieldCreatedAt,
	FieldUpdatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}