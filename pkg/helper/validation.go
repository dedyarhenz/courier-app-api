package helper

import (
	custErr "final-project-backend/pkg/errors"
	"fmt"

	"github.com/go-playground/validator/v10"
)

func ValidateError(err error) string {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return custErr.ErrInvalidRequest.Error()
	}

	for _, errs := range errs {
		errMsg := fmt.Sprintf("Error field %s condition %s", errs.Field(), errs.Tag())
		return errMsg
	}

	return ""
}
