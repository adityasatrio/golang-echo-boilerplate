package dto

type (
	PostRequest struct {
		Title   string `json:"title" validate:"required"`
		Content string `json:"content" validate:"required"`
		Slug    string `json:"slug" validate:"required"`
		Status  int    `json:"status"`
	}

	Post struct {
		Title     string `json:"title" validate:"required"`
		Content   string `json:"content" validate:"required"`
		Slug      string `json:"slug" validate:"required"`
		Status    string `json:"status" validate:"required"`
		CreatedBy string `json:"created-by" validate:"required"`
		CreatedAt string `json:"created-at" validate:"required"`
		UpdatedBy string `json:"updated-by" validate:"required"`
		UpdatedAt string `json:"update-at" validate:"required"`
	}
)
