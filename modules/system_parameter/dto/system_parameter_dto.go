package dto

type (
	SystemParameterRequest struct {
		Key   string `json:"key" validate:"required"`
		Value string `json:"value" validate:"required"`
	}

	SystemParameterResponse struct {
		Key         string `json:"key" validate:"required"`
		Value       string `json:"value" validate:"required"`
		CreatedBy   string `json:"created-by" validate:"required"`
		CreatedDate string `json:"created-date" validate:"required"`
		UpdatedBy   string `json:"updated-by" validate:"required"`
		UpdateDate  string `json:"update-date" validate:"required"`
	}
)
