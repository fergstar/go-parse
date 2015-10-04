package parse

import (
	"fmt"
)

// APIError represents a Parse API Error response.
type APIError struct {
	Code  string `json:"code"`
	Error string `json:"error"`
}

// Message displays the Code and Error as string message.
func (e *APIError) Message() string {

	if len(e.Error) > 0 {
		return fmt.Sprintf("parse: %d %v", e.Code, e.Error)
	}

	return ""

}
