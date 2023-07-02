package dto

type QuoteApiResponse struct {
	Author string `json:"author,omitempty"`
	Quote  string `json:"quote,omitempty"`
}

type QuoteApiRequest struct {
	Name   string `json:"name,omitempty"`
	Author string `json:"author,omitempty"`
	Quote  string `json:"quote,omitempty"`
}
