package validatorconfig

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"reflect"
	"strings"
)

type apiError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func Init(e *echo.Echo) {
	validate := validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("query"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	e.Validator = &CustomValidator{validator: validate}
}

func ErrMsg(err error) interface{} {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]apiError, len(ve))
		for i, fe := range ve {
			out[i] = apiError{fe.Field(), msgForTag(fe, fe.Value())}
		}
		return out
	}

	return nil
}

func msgForTag(fe validator.FieldError, val interface{}) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email"
	case "min":
		return fmt.Sprintf("%d value less than min", val)
	case "max":
		return fmt.Sprintf("%d value greater than max", val)
	}
	return fe.Error() // default error
}
