package errorutil

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	sharedError "github.com/virgiawanly/schedulr-backend/shared/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func MapToGRPCError(err error) error {
	if err == nil {
		return nil
	}

	// Handle AppError (custom error type)
	var appErr *sharedError.AppError
	if errors.As(err, &appErr) {
		return status.Error(appErr.Code, appErr.Message)
	}

	// Handle validation errors
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		var details []string
		for _, fieldErr := range validationErrors {
			if fieldErr.Param() != "" {
				details = append(details, fmt.Sprintf("%s: %s=%s", fieldErr.Field(), fieldErr.Tag(), fieldErr.Param()))
			} else {
				details = append(details, fmt.Sprintf("%s: %s", fieldErr.Field(), fieldErr.Tag()))
			}
		}
		return status.Errorf(codes.InvalidArgument, "validation error: %v", details)
	}

	// Handle standard errors
	switch {
	case errors.Is(err, sharedError.ErrNotFound):
		return status.Error(codes.NotFound, err.Error())
	case errors.Is(err, sharedError.ErrBadRequest):
		return status.Error(codes.InvalidArgument, err.Error())
	case errors.Is(err, sharedError.ErrUnauthorized):
		return status.Error(codes.Unauthenticated, err.Error())
	case errors.Is(err, sharedError.ErrForbidden):
		return status.Error(codes.PermissionDenied, err.Error())
	case errors.Is(err, sharedError.ErrConflict):
		return status.Error(codes.AlreadyExists, err.Error())
	case errors.Is(err, sharedError.ErrInternalServer):
		return status.Error(codes.Internal, err.Error())
	default:
		return status.Error(codes.Unknown, "unknown error occurred")
	}
}
