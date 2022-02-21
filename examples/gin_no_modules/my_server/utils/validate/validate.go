package validate

import (
	"fmt"
	"strings"

	"gopkg.in/go-playground/validator.v9"

	"github.com/gin-gonic/gin"
)

// credits to this --> https://dev.to/koddr/how-to-make-clear-pretty-error-messages-from-the-go-backend-to-your-frontend-21b2

// NewValidator func for create a new validator for struct fields.
func NewValidator() *validator.Validate {
	return validator.New()
}

// ValidatorErrors func for show validation errors for each invalid fields.
func ValidatorErrors(err error) map[string]string {

	errFields := map[string]string{}

	for _, err := range err.(validator.ValidationErrors) {
		structName := strings.Split(err.Namespace(), ".")[0]
		errFields[err.Field()] = fmt.Sprintf(
			"failed '%s' tag check (value '%s' is not valid for %s struct)",
			err.Tag(), err.Value(), structName,
		)
	}
	return errFields
}

// run the sent payload through 2 functions
// 1. the BodyParser will check if the sent json object is correct
// 2 the NewValidator will check if the validate fields for the passed down struct are valid.
func ValidatePayload(c *gin.Context, payload interface{}) (err error) {

	if err := c.ShouldBindJSON(payload); err != nil {
		return err
	}

	if err := NewValidator().Struct(payload); err != nil {
		return err
	}

	return err
}
