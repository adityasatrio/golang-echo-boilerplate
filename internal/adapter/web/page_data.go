package web

// LoginPageData is the template data for templates/login.html.
type LoginPageData struct {
	Error string
}

// RoleOption is a <select> entry sourced from the role service.
type RoleOption struct {
	ID   uint64
	Name string
}

// UserRow is a single row in templates/partials/user_row.html.
// OOBMode controls how the row is swapped when returned from a mutation
// endpoint: "" for a normal full-page render, "true" to replace an existing
// row in place (update), "beforeend:#user-table-body" to append a new row
// (create).
type UserRow struct {
	ID       uint64
	Name     string
	Email    string
	RoleName string
	OOBMode  string
}

// UsersPageData is the template data for templates/users.html.
type UsersPageData struct {
	Active  string
	IsAdmin bool
	Users   []UserRow
}

// UserFormData is the template data for templates/partials/user_form.html.
type UserFormData struct {
	IsEdit       bool
	ID           uint64
	Name         string
	Email        string
	RoleID       uint64
	Roles        []RoleOption
	GeneralError string
}

// SystemParameterRow is a single row in
// templates/partials/system_parameter_row.html.
type SystemParameterRow struct {
	ID      int
	Key     string
	Value   string
	IsAdmin bool
	OOBMode string
}

// SystemParametersPageData is the template data for
// templates/system_parameters.html.
type SystemParametersPageData struct {
	Active  string
	IsAdmin bool
	Params  []SystemParameterRow
}

// SystemParameterFormData is the template data for
// templates/partials/system_parameter_form.html.
type SystemParameterFormData struct {
	IsEdit       bool
	ID           int
	Key          string
	Value        string
	GeneralError string
}
