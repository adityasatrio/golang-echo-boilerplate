package dto

type (
	SystemParameterCreateRequest struct {
		Key   string `json:"key" form:"key" validate:"required"`
		Value string `json:"value" form:"value" validate:"required"`
	}

	SystemParameterUpdateRequest struct {
		Key   string `json:"key" form:"key" validate:"gte=3,lte=50"`
		Value string `json:"value" form:"value" validate:"gte=3,lte=50"`
	}

	SystemParameterResponse struct {
		ID    int    `json:"ID" validate:"required"`
		Key   string `json:"Key" validate:"required"`
		Value string `json:"Value" validate:"required"`
		// CreatedBy time.Time `json:"CreatedBy" validate:"required"`
		// CreatedAt time.Time `json:"CreatedAt" validate:"required"`
		// UpdatedBy time.Time `json:"UpdatedBy" validate:"required"`
		// UpdatedAt time.Time `json:"UpdateAt" validate:"required"`
	}
)
