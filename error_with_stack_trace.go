package errors

import (
	"fmt"

	pkgerrors "github.com/pkg/errors"
)

//ErrorWithStackTrace is an error that has a recorded StackTrace.
type ErrorWithStackTrace interface {
	error
	StackTrace() string
}

func new(err error) ErrorWithStackTrace {
	if err == nil {
		return nil
	}

	e := withStackTrace{
		err,
	}

	return &e
}

type withStackTrace struct {
	error
}

func (err withStackTrace) StackTrace() string {

	type stackTracer interface {
		StackTrace() pkgerrors.StackTrace
	}

	// One of the errors down the chain will be a StackTracer,
	// stop when we find it. This will panic if we are wrong.
	var current error = err
	for current != nil {
		if _, ok := current.(stackTracer); ok {
			break
		}
		current = nextCause(current)
	}

	return fmt.Sprintf("%v", current.(stackTracer).StackTrace())
}

func (err withStackTrace) Cause() error {
	return err.error
}

func nextCause(err error) error {
	type causer interface {
		Cause() error
	}

	if causer, ok := err.(causer); ok {
		return causer.Cause()
	}

	return nil
}
