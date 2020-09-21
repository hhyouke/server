package models

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// APIMessage the rest api response message
type APIMessage struct {
	HTTPCode        int         `json:"-"`
	Code            string      `json:"code"`
	RequestID       string      `json:"-"`
	Message         string      `json:"message"`
	Payload         interface{} `json:"payload"`
	InternalError   error       `json:"-"`
	InternalMessage string      `json:"-"`
}

func (e *APIMessage) Error() string {
	if e.InternalMessage != "" {
		return e.InternalMessage
	}
	return fmt.Sprintf("%d: %s: %s", e.HTTPCode, e.Code, e.Message)
}

// Cause returns the root cause error
func (e *APIMessage) Cause() error {
	if e.InternalError != nil {
		return e.InternalError
	}
	return e
}

// WithInternalError adds internal error information to the error
func (e *APIMessage) WithInternalError(err error) *APIMessage {
	e.InternalError = err
	return e
}

// WithInternalMessage adds internal message information to the error
func (e *APIMessage) WithInternalMessage(fmtString string, args ...interface{}) *APIMessage {
	e.InternalMessage = fmt.Sprintf(fmtString, args...)
	return e
}

// NewHTTPError construct an api error
func NewHTTPError(httpCode int, code string, msgFmt string, args ...interface{}) *APIMessage {
	return &APIMessage{
		HTTPCode: httpCode,
		Code:     code,
		Message:  fmt.Sprintf(msgFmt, args...),
	}
}

// NewAPIError construct result for a successful api call, but the result is invalid
func NewAPIError(code string, msgFmt string, args ...interface{}) *APIMessage {
	return &APIMessage{
		HTTPCode: http.StatusOK,
		Code:     code,
		Message:  fmt.Sprintf(msgFmt, args...),
	}
}

//NewAPIResult construct result for a successful api call, but the result is valid
func NewAPIResult(code string, obj interface{}, msgFmt string, args ...interface{}) *APIMessage {
	return &APIMessage{
		HTTPCode: http.StatusOK,
		Code:     code,
		Message:  fmt.Sprintf(msgFmt, args...),
		Payload:  obj,
	}
}

// SendJSON handle rest api returns
func SendJSON(w http.ResponseWriter, apiMsg *APIMessage) error {
	w.Header().Set("Content-Type", "application/json")

	b, err := json.Marshal(apiMsg)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Error encoding json response: %v", apiMsg))
	}
	w.WriteHeader(apiMsg.HTTPCode)
	_, err = w.Write(b)
	return err
}
