package calculation

import (
	"errors"
)

var (
	errFooUnexpectedServerError        = errors.New("unexpected server error")
	errFooInvalidExpressionClientError = errors.New("invaliv expression client error")
)
