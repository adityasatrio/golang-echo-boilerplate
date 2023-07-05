package dto

type (
	ExampleRequest struct {
		ID int `json:"ID" validate:"required"`
	}

	ExampleResponse struct {
		ID int `json:"ID" validate:"required"`
	}
)
