// Package errors is the same as github.com/pkg/errors. However, it makes one
// additional guarantee:
//
// - All errors returned by this package are guaranteed to have a stack trace.
//
// More precisely, all returned errors implement the ErrorWithStackTrace
// interface:
//
//    type ErrorWithStackTrace interface{
//        error
//        StackTrace() string
//    }
//
// This allows you to leverage the type system to ensure that all of the errors
// that you return include a stack trace:
//
//    func Foo() ErrorWithStackTrace
//
// This package tries to be a drop-in replacement for github.com/pkg/error.
// However, there are known caveats:
//
// 1. Errors returned by this package do not implement the same StackTracer
// interface as github.com/pkg/errors.
//
// 2. Instances of similar code will no longer compile:
//
//    err := errors.New("foo") // returns errors.ErrorWithStackTrace
//    c, err := net.Listen("...") // returns regular error, won't compile
//
// Note that WithMessage also behaves a bit differently as it will add a stack
// strace to the supplied error if there is not already one.
package errors

import (
	pkgerrors "github.com/pkg/errors"
)

//Cause retrieves the underlying cause of an error
func Cause(err error) error {
	return pkgerrors.Cause(err)
}

//Errorf formats an error
func Errorf(format string, args ...interface{}) ErrorWithStackTrace {
	return new(pkgerrors.Errorf(format, args...))
}

//New creates a new error
func New(message string) ErrorWithStackTrace {
	return new(pkgerrors.New(message))
}

//WithMessage annotates the error with the given message. This implementation
//produces different results than github.com/pkg/errors because it will also
//add a StackTrace to the error if it does not already have one.
func WithMessage(err error, message string) ErrorWithStackTrace {
	// Look for a stack trace
	current := err
	for current != nil {
		if _, ok := current.(ErrorWithStackTrace); ok {
			break
		}
		current = nextCause(current)
	}

	// We couldn't find any stack trace, add one.
	if current == nil {
		err = WithStack(err)
	}

	withMessage := pkgerrors.WithMessage(err, message)

	return new(withMessage)
}

//WithStack annotates the error with a stack trace
func WithStack(err error) ErrorWithStackTrace {
	return new(pkgerrors.WithStack(err))
}

//Wrap wraps an error with a stack trace and a message
func Wrap(err error, message string) ErrorWithStackTrace {
	return new(pkgerrors.Wrap(err, message))
}

//Wrapf wraps an error with a stack trace and formats the message
func Wrapf(err error, format string, args ...interface{}) ErrorWithStackTrace {
	return new(pkgerrors.Wrapf(err, format, args...))
}

//Frame is the same as pkgerrors.Frame
type Frame pkgerrors.Frame

//StackTrace is the same as pkgerrors.Stacktrace
type StackTrace pkgerrors.StackTrace
