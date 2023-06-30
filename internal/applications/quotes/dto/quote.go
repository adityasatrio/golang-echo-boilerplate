package dto

type QuoteApiResponse struct {
	author string `json:"author,omitempty"`
	quote  string `json:"quote,omitempty"`
}

type QuoteApiRequest struct {
	name   string `json:"name,omitempty"`
	author string `json:"author,omitempty"`
	quote  string `json:"quote,omitempty"`
}
