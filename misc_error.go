package notion

type ErrorCode string

type APIError struct {
	Object  ObjectType `json:"object"`
	Status  int        `json:"status"`
	Code    ErrorCode  `json:"code"`
	Message string     `json:"message"`
}

func (e *APIError) Error() string {
	return e.Message
}

type RateLimitedError struct {
	Message string
}

func (e *RateLimitedError) Error() string {
	return e.Message
}

type TokenCreateError struct {
	Code    ErrorCode `json:"error"`
	Message string    `json:"error_description"`
}

func (e *TokenCreateError) Error() string {
	return e.Message
}
