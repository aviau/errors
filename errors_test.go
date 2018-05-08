package errors_test

import (
	"testing"

	goerrors "errors"
	"github.com/stretchr/testify/assert"
	"github.com/aviau/errors"
)

func assertContainsStackTrace(t *testing.T, err errors.ErrorWithStackTrace) {
	assert.NotNil(t, err)
	assert.Contains(t, err.StackTrace(), "errors.go")
	assert.Contains(t, err.StackTrace(), "errors_test.go")
	assert.Contains(t, err.StackTrace(), "testing.go")
}

func TestErrorf(t *testing.T) {
	err := errors.Errorf("test")

	assertContainsStackTrace(t, err)
	assert.Equal(t, "test", err.Error())
}

func TestNew(t *testing.T) {
	err := errors.New("test")

	assertContainsStackTrace(t, err)
	assert.Equal(t, "test", err.Error())
}

func TestWithMessage(t *testing.T) {
	base := errors.New("test")
	err := errors.WithMessage(base, "msg")

	assertContainsStackTrace(t, err)
	assert.Equal(t, "msg: test", err.Error())
}

func TestWithMessageGoError(t *testing.T) {
	base := goerrors.New("test")
	err := errors.WithMessage(base, "msg")

	assertContainsStackTrace(t, err)
	assert.Equal(t, "msg: test", err.Error())
}

func TestWithStackGoError(t *testing.T) {
	base := goerrors.New("test")
	err := errors.WithStack(base)

	assertContainsStackTrace(t, err)
	assert.Equal(t, "test", err.Error())
}

func TestWrapGoError(t *testing.T) {
	base := goerrors.New("test")
	err := errors.Wrap(base, "msg")

	assertContainsStackTrace(t, err)
	assert.Equal(t, "msg: test", err.Error())
}

func TestWrapfGoError(t *testing.T) {
	base := goerrors.New("test")
	err := errors.Wrapf(base, "msg")

	assertContainsStackTrace(t, err)
	assert.Equal(t, "msg: test", err.Error())
}
