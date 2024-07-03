package result

import (
	"errors"
	"net/http"
)

const (
	StatusCanceled = 499 // Client Closed Request
	StatusUnknown  = 520 // Web Server Returned an Unknown Error
)

func StatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}
	var cause interface{ StatusCode() int }
	if errors.As(err, &cause) {
		return cause.StatusCode()
	}
	return http.StatusInternalServerError
}

type withStatus struct {
	statusCode int
	err        error
}

func (w *withStatus) StatusCode() int {
	return w.statusCode
}

// Unwrap provides compatibility for Go 1.13 error chains
func (w *withStatus) Unwrap() error { return w.err }

// Cause provides compatibility for github.com/pkg/errors error chains
func (w *withStatus) Cause() error { return w.err }

func (w *withStatus) Error() string {
	if w.err == nil {
		return ""
	}
	return w.err.Error()
}

////////////////////////////////
//     client-side errors     //
////////////////////////////////s

func BadRequest(err error) error {
	return &withStatus{http.StatusBadRequest, err}
}

func Unauthorized(err error) error {
	return &withStatus{http.StatusUnauthorized, err}
}

func PaymentRequired(err error) error {
	return &withStatus{http.StatusPaymentRequired, err}
}

func Forbidden(err error) error {
	return &withStatus{http.StatusForbidden, err}
}

func NotFound(err error) error {
	return &withStatus{http.StatusNotFound, err}
}

func MethodNotAllowed(err error) error {
	return &withStatus{http.StatusMethodNotAllowed, err}
}

func NotAcceptable(err error) error {
	return &withStatus{http.StatusNotAcceptable, err}
}

func ProxyAuthRequired(err error) error {
	return &withStatus{http.StatusProxyAuthRequired, err}
}

func RequestTimeout(err error) error {
	return &withStatus{http.StatusRequestTimeout, err}
}

func Conflict(err error) error {
	return &withStatus{http.StatusConflict, err}
}

func Gone(err error) error {
	return &withStatus{http.StatusGone, err}
}

func LengthRequired(err error) error {
	return &withStatus{http.StatusLengthRequired, err}
}

func PreconditionFailed(err error) error {
	return &withStatus{http.StatusPreconditionFailed, err}
}

func RequestEntityTooLarge(err error) error {
	return &withStatus{http.StatusRequestEntityTooLarge, err}
}

func RequestURITooLong(err error) error {
	return &withStatus{http.StatusRequestURITooLong, err}
}

func UnsupportedMediaType(err error) error {
	return &withStatus{http.StatusUnsupportedMediaType, err}
}

func RequestedRangeNotSatisfiable(err error) error {
	return &withStatus{http.StatusRequestedRangeNotSatisfiable, err}
}

func ExpectationFailed(err error) error {
	return &withStatus{http.StatusExpectationFailed, err}
}

func MisdirectedRequest(err error) error {
	return &withStatus{http.StatusMisdirectedRequest, err}
}

func UnprocessableEntity(err error) error {
	return &withStatus{http.StatusUnprocessableEntity, err}
}

func Locked(err error) error {
	return &withStatus{http.StatusLocked, err}
}

func FailedDependency(err error) error {
	return &withStatus{http.StatusFailedDependency, err}
}

func TooEarly(err error) error {
	return &withStatus{http.StatusTooEarly, err}
}

func UpgradeRequired(err error) error {
	return &withStatus{http.StatusUpgradeRequired, err}
}

func PreconditionRequired(err error) error {
	return &withStatus{http.StatusPreconditionRequired, err}
}

func TooManyRequests(err error) error {
	return &withStatus{http.StatusTooManyRequests, err}
}

func RequestHeaderFieldsTooLarge(err error) error {
	return &withStatus{http.StatusRequestHeaderFieldsTooLarge, err}
}

func UnavailableForLegalReasons(err error) error {
	return &withStatus{http.StatusUnavailableForLegalReasons, err}
}

func Canceled(err error) error {
	return &withStatus{StatusCanceled, err}
}

////////////////////////////////
//     server-side errors     //
////////////////////////////////

func InternalServerError(err error) error {
	return &withStatus{http.StatusInternalServerError, err}
}

func NotImplemented(err error) error {
	return &withStatus{http.StatusNotImplemented, err}
}

func BadGateway(err error) error {
	return &withStatus{http.StatusBadGateway, err}
}

func ServiceUnavailable(err error) error {
	return &withStatus{http.StatusServiceUnavailable, err}
}

func GatewayTimeout(err error) error {
	return &withStatus{http.StatusGatewayTimeout, err}
}

func HTTPVersionNotSupported(err error) error {
	return &withStatus{http.StatusHTTPVersionNotSupported, err}
}

func VariantAlsoNegotiates(err error) error {
	return &withStatus{http.StatusVariantAlsoNegotiates, err}
}

func InsufficientStorage(err error) error {
	return &withStatus{http.StatusInsufficientStorage, err}
}

func LoopDetected(err error) error {
	return &withStatus{http.StatusLoopDetected, err}
}

func NotExtended(err error) error {
	return &withStatus{http.StatusNotExtended, err}
}

func NetworkAuthenticationRequired(err error) error {
	return &withStatus{http.StatusNetworkAuthenticationRequired, err}
}

func Unknown(err error) error {
	return &withStatus{StatusUnknown, err}
}
