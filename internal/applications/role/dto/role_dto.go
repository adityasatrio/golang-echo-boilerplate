package dto

type (
	RoleRequest struct {
		Name string `json:"name" validate:"required"`
		Text string `json:"text" validate:"required"`
	}

	RoleResponse struct {
		Name      string `json:"name" validate:"required"`
		Text      string `json:"text" validate:"required"`
		CreatedAt string `json:"created_at" validate:"required"`
		UpdatedAt string `json:"update_at" validate:"required"`
		DeletedAt string `json:"deleted_at" validate:"required"`
	}
)
