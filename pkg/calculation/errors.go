package calculation

import (
	"errors"
)

var (
	unexpectedServerError        = errors.New("unexpected server error")
	invalidExpressionClientError = errors.New("invaliv expression client error")
)
