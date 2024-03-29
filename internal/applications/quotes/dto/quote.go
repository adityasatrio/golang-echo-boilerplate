package dto

type QuoteResponse struct {
	Author string `json:"author,omitempty"`
	Quote  string `json:"quote,omitempty"`
	Custom string `json:"custom,omitempty"`
}

type QuoteRequest struct {
	Name   string `json:"name,omitempty"`
	Author string `json:"author,omitempty"`
	Quote  string `json:"quote,omitempty"`
	Custom string `json:"custom,omitempty"`
}

type QuoteApiResponse struct {
	Author string `json:"author,omitempty"`
	Quote  string `json:"quote,omitempty"`
}

type QuoteApiRequest struct {
	Name   string `json:"name,omitempty"`
	Author string `json:"author,omitempty"`
	Quote  string `json:"quote,omitempty"`
}
