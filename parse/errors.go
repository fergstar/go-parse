package parse

import (
	"fmt"
)

// APIError represents a Parse API Error response.
type APIError struct {
	Code         int    `json:"code"`
	ErrorMessage string `json:"error"`
}

// Message displays the Code and Error as string message.
func (e *APIError) Error() string {

	if len(e.ErrorMessage) > 0 {
		return fmt.Sprintf("parse: %v - %s", e.Code, e.ErrorMessage)
	}

	return ""

}

func releventError(httpError error, apiError *APIError) error {

	if httpError != nil {
		return httpError
	}

	if apiError.ErrorMessage == "" {
		return nil
	}

	return apiError
}
