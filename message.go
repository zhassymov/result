package result

import (
	"errors"
)

func Message(err error) any {
	if err == nil {
		return nil
	}
	var cause interface{ Message() any }
	if errors.As(err, &cause) {
		return cause.Message()
	}
	return nil
}

type withMessage struct {
	err error
	msg any
}

func WithMessage(err error, msg any) error {
	return &withMessage{err, msg}
}

func (w *withMessage) Message() any {
	return w.msg
}

// Unwrap provides compatibility for Go 1.13 error chains
func (w *withMessage) Unwrap() error { return w.err }

// Cause provides compatibility for github.com/pkg/errors error chains
func (w *withMessage) Cause() error { return w.err }

func (w *withMessage) Error() string {
	if w.err == nil {
		return ""
	}
	return w.err.Error()
}
