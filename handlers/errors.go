package handlers

//NOTE yet implemented

//Error represents a handler error. It embends the builtin error and
// provides a method to return an HTTP error response code
type Error interface {
	error
	Status() int
}

// StatusError represents an Http error
type StatusError struct {
	Err     error
	Code    int
	Message string
}

//These Allows StatusError to satisfy the error interface:

//Error the http error
func (se StatusError) Error() string {
	return se.Err.Error()
}

// Status returns the HTTP status code.
func (se StatusError) Status() int {
	return se.Code
}

//StatusMessage returns an message about the error
func (se StatusError) StatusMessage() string {
	return se.Message
}
