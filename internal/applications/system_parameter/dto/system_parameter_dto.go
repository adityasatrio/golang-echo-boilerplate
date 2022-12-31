package dto

type (
	SystemParameterCreateRequest struct {
		Key   string `json:"key" validate:"required"`
		Value string `json:"value" validate:"required"`
	}

	SystemParameterUpdateRequest struct {
		Key   string `json:"key" validate:"gte=5,lte=50"`
		Value string `json:"value" validate:"gte=5,lte=50"`
	}

	SystemParameterResponse struct {
		Key       string `json:"key" validate:"required"`
		Value     string `json:"value" validate:"required"`
		CreatedBy string `json:"created-by" validate:"required"`
		CreatedAt string `json:"created-at" validate:"required"`
		UpdatedBy string `json:"updated-by" validate:"required"`
		UpdatedAt string `json:"update-at" validate:"required"`
	}
)
