package errorssx

import "reflect"

// Who - interface to determine who the error originated from
type Who interface {
	Name() string
}

// Error - Package defined error struct to find panic-initiated error statements
type Error struct {
	who       Who
	operation string
	goerr     error
}

func (err *Error) Error() string {
	return reflect.TypeOf(err).PkgPath() + " " + err.operation + "\nError: " + err.goerr.Error()
}

// New - returns errorssx error type
func New(who Who, operation string, err error) *Error {
	return &Error{who: who, operation: operation, goerr: err}
}
