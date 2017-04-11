package connector

import (
	"net/http"
	"strings"

	"github.com/go-openapi/runtime"

	"github.com/manifoldco/go-connector/errors"
	"github.com/manifoldco/go-manifold"
	merrors "github.com/manifoldco/go-manifold/errors"
)

// NewOAuthError returns an OAuthError which satisfies the Error, HTTPError and
// swag.Error interfaces.
//
// description is optional and will be omitted from the response if its an
// empty string.
func NewOAuthError(t errors.OAuthErrorType, description string) *OAuthError {
	return &OAuthError{
		Type:        errors.OAuthErrorTypeMap[t],
		ErrorType:   t,
		Description: description,
	}
}

// ToOAuthError receives an error and mutates it into an OAuthError based on
// the concrete type of the Error
func ToOAuthError(err error) manifold.HTTPError {
	switch e := err.(type) {
	case *OAuthError:
		return e
	case *manifold.Error:
		switch e.Type {
		case merrors.BadRequestError:
			return NewOAuthError(errors.InvalidRequestErrorType, strings.Join(e.Messages, ","))
		case merrors.UnauthorizedError:
			return NewOAuthError(errors.AccessDeniedErrorType, strings.Join(e.Messages, ","))
		default:
			return NewOAuthError(errors.ServerErrorErrorType, "")
		}
	default:
		return NewOAuthError(errors.ServerErrorErrorType, "")
	}
}

// OAuthError represents an error that follows the OAuth 2.0
// specification. An OAuthError occurs during a component of an
// OAuth authentication flow.
//
// These errors are only surfaced to a provider.
type OAuthError struct {
	Type        merrors.Type          `json:"-"`
	ErrorType   errors.OAuthErrorType `json:"error"`
	Description string                `json:"error_description,omitempty"`
}

// Error returns a string equivalent of this Error. It also
// completes the error interface.
func (e *OAuthError) Error() string {
	return e.ErrorType.String() + ": " + e.Description
}

// StatusCode returns the http status code for this Error. It completes the
// HTTPError interface.
func (e *OAuthError) StatusCode() int {
	return e.Type.Code()
}

// WriteResponse completes the interface for a HTTPError; enabling an error to
// be returned as a middleware.Responder from go-openapi/runtime
//
// A panic will occur if the given producer errors.
func (e *OAuthError) WriteResponse(rw http.ResponseWriter, pr runtime.Producer) {
	if e.ErrorType == errors.InvalidClientErrorType {
		rw.Header().Set("WWW-Authenticate", "Basic")
	}

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(e.StatusCode())
	if err := pr.Produce(rw, e); err != nil {
		panic(err)
	}
}
