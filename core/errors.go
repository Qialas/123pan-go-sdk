package core

import "fmt"

type APIError struct {
	StatusCode int
	Code       int
	Message    string
	TraceID    string
}

func (e *APIError) Error() string {
	if e == nil {
		return ""
	}
	if e.TraceID != "" {
		return fmt.Sprintf("123pan api error: status=%d code=%d message=%q traceID=%s", e.StatusCode, e.Code, e.Message, e.TraceID)
	}
	return fmt.Sprintf("123pan api error: status=%d code=%d message=%q", e.StatusCode, e.Code, e.Message)
}
