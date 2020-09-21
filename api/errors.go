package api

import (
	"fmt"
	"net/http"
	"os"
	"runtime/debug"

	"github.com/hhyouke/server/models"
	"go.uber.org/zap"
)

// // HTTPError is an error with a message and an HTTP status code.
// type HTTPError struct {
// 	Code            int    `json:"code"`
// 	Message         string `json:"msg"`
// 	InternalError   error  `json:"-"`
// 	InternalMessage string `json:"-"`
// 	ErrorID         string `json:"error_id,omitempty"`
// }

// // Error prints the code and message
// func (e *HTTPError) Error() string {
// 	if e.InternalMessage != "" {
// 		return e.InternalMessage
// 	}
// 	return fmt.Sprintf("%d: %s", e.Code, e.Message)
// }

// // Cause returns the root cause error
// func (e *HTTPError) Cause() error {
// 	if e.InternalError != nil {
// 		return e.InternalError
// 	}
// 	return e
// }

// // WithInternalError adds internal error information to the error
// func (e *HTTPError) WithInternalError(err error) *HTTPError {
// 	e.InternalError = err
// 	return e
// }

// // WithInternalMessage adds internal message information to the error
// func (e *HTTPError) WithInternalMessage(fmtString string, args ...interface{}) *HTTPError {
// 	e.InternalMessage = fmt.Sprintf(fmtString, args...)
// 	return e
// }

// func httpError(code int, fmtString string, args ...interface{}) *HTTPError {
// 	return &HTTPError{
// 		Code:    code,
// 		Message: fmt.Sprintf(fmtString, args...),
// 	}
// }

// Recoverer is a middleware that recovers from panics, logs the panic (and a
// backtrace), and returns a HTTP 500 (Internal Server Error) status if
// possible. Recoverer prints a request ID if one is provided.
func recoverer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			// log.Printf("[#%d]", goid.Get())
			if rvr := recover(); rvr != nil {
				logEntry := getLogEntry(r)
				if logEntry != nil {
					logEntry.Panic(rvr, debug.Stack())
				} else {
					fmt.Fprintf(os.Stderr, "Panic: %+v\n", rvr)
					debug.PrintStack()
				}
				se := models.NewHTTPError(http.StatusInternalServerError, models.ErrInternal, http.StatusText(http.StatusInternalServerError))
				handleError(se, w, r)
			}
		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func handleError(err error, w http.ResponseWriter, r *http.Request) {
	logger := getLogEntry(r)
	errorID := GetRequestID(r.Context())
	switch e := err.(type) {
	case *models.APIMessage:
		if e.HTTPCode >= http.StatusInternalServerError {
			e.RequestID = errorID
			logger.Logger.Error(e.Error(), zap.Any("cause", e.Cause()))
		} else {
			logger.Logger.Info(e.Error(), zap.Any("cause", e.Cause()))
		}
		if jsonErr := models.SendJSON(w, e); jsonErr != nil {
			handleError(jsonErr, w, r)
		}
	default:
		logger.Logger.Error(e.Error(), zap.Any("err", e))
		// hide real error details from response to prevent info leaks
		w.WriteHeader(http.StatusInternalServerError)
		if _, writeErr := w.Write([]byte(`{"code":500,"message":"Internal server error","error_id":"` + errorID + `"}`)); writeErr != nil {
			logger.Logger.Error("Error writing generic error message", zap.Any("err", writeErr))
		}
	}
}

// func badRequestError(fmtString string, args ...interface{}) *HTTPError {
// 	return httpError(http.StatusBadRequest, fmtString, args...)
// }

// func internalServerError(fmtString string, args ...interface{}) *HTTPError {
// 	return httpError(http.StatusInternalServerError, fmtString, args...)
// }

// func notFoundError(fmtString string, args ...interface{}) *HTTPError {
// 	return httpError(http.StatusNotFound, fmtString, args...)
// }

// func unauthorizedError(fmtString string, args ...interface{}) *HTTPError {
// 	return httpError(http.StatusUnauthorized, fmtString, args...)
// }
