package test

import (
	"encoding/json"
	"github.com/jarcoal/httpmock"
	"github.com/labstack/gommon/log"
	"net/http"
)

func ResponderJsonResponse(status int, body string) httpmock.Responder {
	var bodyResponse interface{}
	err := json.Unmarshal([]byte(body), &bodyResponse)
	if err != nil {
		log.Error("failed to unmarshal body response from string to json")
		return nil
	}

	return func(req *http.Request) (*http.Response, error) {
		resp, _ := httpmock.NewJsonResponse(status, bodyResponse)
		resp.Header.Set("Content-Type", "application/json")
		return resp, nil
	}

}
