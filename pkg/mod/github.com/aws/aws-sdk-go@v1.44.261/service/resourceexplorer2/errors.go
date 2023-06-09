// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package resourceexplorer2

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// The credentials that you used to call this operation don't have the minimum
	// required permissions.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// The request failed because either you specified parameters that didn’t
	// match the original request, or you attempted to create a view with a name
	// that already exists in this Amazon Web Services Region.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeInternalServerException for service response error code
	// "InternalServerException".
	//
	// The request failed because of internal service error. Try your request again
	// later.
	ErrCodeInternalServerException = "InternalServerException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// You specified a resource that doesn't exist. Check the ID or ARN that you
	// used to identity the resource, and try again.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeServiceQuotaExceededException for service response error code
	// "ServiceQuotaExceededException".
	//
	// The request failed because it exceeds a service quota.
	ErrCodeServiceQuotaExceededException = "ServiceQuotaExceededException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// The request failed because you exceeded a rate limit for this operation.
	// For more information, see Quotas for Resource Explorer (https://docs.aws.amazon.com/arexug/mainline/quotas.html).
	ErrCodeThrottlingException = "ThrottlingException"

	// ErrCodeUnauthorizedException for service response error code
	// "UnauthorizedException".
	//
	// The principal making the request isn't permitted to perform the operation.
	ErrCodeUnauthorizedException = "UnauthorizedException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// You provided an invalid value for one of the operation's parameters. Check
	// the syntax for the operation, and try again.
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":         newErrorAccessDeniedException,
	"ConflictException":             newErrorConflictException,
	"InternalServerException":       newErrorInternalServerException,
	"ResourceNotFoundException":     newErrorResourceNotFoundException,
	"ServiceQuotaExceededException": newErrorServiceQuotaExceededException,
	"ThrottlingException":           newErrorThrottlingException,
	"UnauthorizedException":         newErrorUnauthorizedException,
	"ValidationException":           newErrorValidationException,
}
