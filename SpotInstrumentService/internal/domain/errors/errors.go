package spoterrors

import "errors"

// Ошибки usecase
var (
	Avalible_markets       = errors.New("no markets avalible")
	Unavailable_request_id = errors.New("there is no request ID for caching the server response.")
)
