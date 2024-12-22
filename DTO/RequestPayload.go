package DTO

import (
	"encoding/json"
	"io"
)

type RequestPayload struct {
	Message      string  `json:"message"`
	ConnectionID *string `json:"connectionId"`
}

func (payload *RequestPayload) BuildPayloadFromRequest(body io.ReadCloser) error {
	return json.NewDecoder(body).Decode(&payload)
}
