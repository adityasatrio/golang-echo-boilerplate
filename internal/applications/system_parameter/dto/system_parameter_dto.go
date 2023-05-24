package dto

type (
	SystemParameterCreateRequest struct {
		Key   string `json:"key" validate:"required"`
		Value string `json:"value" validate:"required"`
	}

	SystemParameterUpdateRequest struct {
		Key   string `json:"key" validate:"gte=3,lte=50"`
		Value string `json:"value" validate:"gte=3,lte=50"`
	}

	SystemParameterResponse struct {
		Key   string `json:"Key" validate:"required"`
		Value string `json:"Value" validate:"required"`
		//CreatedBy string `json:"CreatedBy" validate:"required"`
		//CreatedAt string `json:"CreatedAt" validate:"required"`
		//UpdatedBy string `json:"UpdatedBy" validate:"required"`
		//UpdatedAt string `json:"UpdateAt" validate:"required"`
	}
)
