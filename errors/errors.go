// Package errors contains common business errors to be user all over the server.
package errors

import "errors"

const (
	E_BAD_REQUEST           string = "bad_request"
	E_FORBIDDEN_REQUEST     string = "access_denied"
	E_GATEWAY_TIMEOUT       string = "gateway_timeout"
	E_INTERNAL_SERVER_ERROR string = "internal_server_error"
	E_NOT_FOUND             string = "not_found"
	E_SERVICE_UNAVAILABLE   string = "service_unavailable"
	E_UNAUTHORIZED_REQUEST  string = "unauthorized_access"
)

var (
	ErrAuthenticationFailed   = errors.New("Authentication failed.")
	ErrBadRequest             = errors.New("The request was invalid or cannot be served.")
	ErrClientNotFound         = errors.New("Client not found.")
	ErrClientNotFoundOnTeam   = errors.New("Client not found on this team.")
	ErrInvalidTokenFormat     = errors.New("Invalid token format.")
	ErrLoginRequired          = errors.New("Invalid or expired token. Please log in with your Backstage credentials.")
	ErrOnlyOwnerHasPermission = errors.New("Only the owner has permission to perform this operation.")
	ErrServiceNotFound        = errors.New("Service not found.")
	ErrTeamNotFound           = errors.New("Team not found.")
	ErrTokenNotFound          = errors.New("Token not found.")
	ErrUserNotInTeam          = errors.New("You do not belong to this team!")
	ErrConfirmationPassword   = errors.New("Your new password and confirmation password do not match.")
)

// The ValidationError type indicates that any validation has failed.
type ValidationError struct {
	Payload string
}

func (err *ValidationError) Error() string {
	return err.Payload
}

// The ForbiddenError type indicates that the user does not have
// permission to perform some operation.
type ForbiddenError struct {
	Payload string
}

func (err *ForbiddenError) Error() string {
	return err.Payload
}

type NotFoundError struct {
	Payload string
}

func (err *NotFoundError) Error() string {
	return err.Payload
}
