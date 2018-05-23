# errors

[![godoc reference](https://godoc.org/github.com/aviau/errors?status.svg)](http://godoc.org/github.com/aviau/errors)
[![Build Status](https://travis-ci.com/aviau/errors.svg?branch=master)](https://travis-ci.com/aviau/errors)
[![go report card](https://goreportcard.com/badge/github.com/aviau/errors)](https://goreportcard.com/report/github.com/aviau/errors)

Package errors is the same as github.com/pkg/errors. However, it makes one
additional guarantee:

- All errors returned by this package are guaranteed to have a stack trace.

More precisely, all returned errors implement the ErrorWithStackTrace
interface:

```golang
   type ErrorWithStackTrace interface{
       error
       StackTrace() string
   }
```

This allows you to leverage the type system to ensure that all of the errors
that you return include a stack trace:

```golang
   func Foo() ErrorWithStackTrace
```

This package tries to be a drop-in replacement for github.com/pkg/error.
However, there are known caveats:

1. Errors returned by this package do not implement the same StackTracer
interface as github.com/pkg/errors.

2. Instances of similar code will no longer compile:

```golang
   err := errors.New("foo") // returns errors.ErrorWithStackTrace
   c, err := net.Listen("...") // returns regular error, won't compile
```

Note that WithMessage also behaves a bit differently as it will add a stack
strace to the supplied error if there is not already one.
