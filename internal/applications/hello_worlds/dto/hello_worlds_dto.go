package dto

type (
	HelloWorldsResponse struct {
		message   string `json:"message" validate:"required"`
		CreatedBy string `json:"created-by" validate:"required"`
		CreatedAt string `json:"created-at" validate:"required"`
		UpdatedBy string `json:"updated-by" validate:"required"`
		UpdatedAt string `json:"update-at" validate:"required"`
	}
)
