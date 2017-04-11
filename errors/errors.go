package errors

import (
	"github.com/manifoldco/go-manifold/errors"
)

// OAuthErrorType represents a type of OAuthError returned from the Connector
// to the provider.
type OAuthErrorType string

// String returns the string representation of this error type
func (t OAuthErrorType) String() string {
	return string(t)
}

// InvalidRequestErrorType represents an error where the request is missing a
// required parameter, invalid parameter value, or is otherwise malformed.
var InvalidRequestErrorType OAuthErrorType = "invalid_request"

// InvalidClientErrorType represents an error where the authentication has
// failed.
var InvalidClientErrorType OAuthErrorType = "invalid_client"

// InvalidGrantErrorType represents an error where the provided authorization
// grant such as a authorization code or refresh token is invalid, expired, or
// does not match the redirection uri.
var InvalidGrantErrorType OAuthErrorType = "invalid_grant"

// UnsupportedGrantErrorType represents an error where the client has requested
// a grant type not supported by the server.
var UnsupportedGrantErrorType OAuthErrorType = "unsupported_grant_type"

// UnauthorizedClientErrorType represents an error where the client is not
// authorized to use a specific grant type.
var UnauthorizedClientErrorType OAuthErrorType = "unauthorized_client"

// AccessDeniedErrorType represents an error where the server has denied the
// request.
//
// XXX: This error is only ever returned to a user-agent during a request for
// an authorization code.
var AccessDeniedErrorType OAuthErrorType = "access_denied"

// ServerErrorErrorType represents an error where the server encountered an
// unexpected condition that prevented it from fulfilling the request.
//
// Used for relaying "something went horribly wrong" during authorization code
// to token exchange or during a user agent request.
var ServerErrorErrorType OAuthErrorType = "server_error"

// InvalidScopeErrorType represents an error where The requested scope is
// invalid, unknown, malformed, or exceeds the scope granted by the resource
// owner.
var InvalidScopeErrorType OAuthErrorType = "invalid_scope"

// OAuthErrorTypeMap relates an OAuthError to an apierrors ErrorType
var OAuthErrorTypeMap = map[OAuthErrorType]errors.Type{
	InvalidRequestErrorType:     errors.BadRequestError,
	InvalidClientErrorType:      errors.UnauthorizedError,
	InvalidGrantErrorType:       errors.BadRequestError,
	UnsupportedGrantErrorType:   errors.BadRequestError,
	UnauthorizedClientErrorType: errors.BadRequestError,
	ServerErrorErrorType:        errors.BadRequestError,
	InvalidScopeErrorType:       errors.BadRequestError,
}
