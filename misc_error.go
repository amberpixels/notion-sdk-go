package notion

// ErrorCode is a type for Notion API error codes.
type ErrorCode string

// APIError is a type for Notion API errors.
type APIError struct {
	Object  ObjectType `json:"object"`
	Status  int        `json:"status"`
	Code    ErrorCode  `json:"code"`
	Message string     `json:"message"`
}

// Error implements the error interface.
func (e *APIError) Error() string { return e.Message }

// RateLimitedError is a type for rate-limited errors.
type RateLimitedError struct {
	Message string
}

// Error implements the error interface.
func (e *RateLimitedError) Error() string { return e.Message }

// TokenCreateError is a type for token creation errors.
type TokenCreateError struct {
	Code    ErrorCode `json:"error"`
	Message string    `json:"error_description"`
}

// Error implements the error interface.
func (e *TokenCreateError) Error() string { return e.Message }
