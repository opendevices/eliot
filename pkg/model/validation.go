package model

import (
	"log"
	"regexp"

	imageref "github.com/containerd/containerd/reference"
	validator "gopkg.in/go-playground/validator.v9"
)

func getValidator() *validator.Validate {
	once.Do(func() {
		validate = validator.New()

		validate.RegisterValidation("hasName", func(fl validator.FieldLevel) bool {
			return hasValidName(fl.Field().Interface().(Metadata))
		})
		validate.RegisterValidation("imageRef", func(fl validator.FieldLevel) bool {
			return isValidImageReference(fl.Field().Interface().(string))
		})
		validate.RegisterValidation("alphanumOrDash", func(fl validator.FieldLevel) bool {
			return isAlphanumericOrDash(fl.Field().Interface().(string))
		})
		validate.RegisterValidation("empty", func(fl validator.FieldLevel) bool {
			return isEmpty(fl.Field().Interface())
		})
	})
	return validate
}

func hasValidName(metadata Metadata) bool {
	if len(metadata.GetName()) == 0 {
		return false
	}
	return true
}

func isValidImageReference(ref string) bool {
	_, err := imageref.Parse(ref)
	return err == nil
}

func isAlphanumericOrDash(value string) bool {
	match, err := regexp.MatchString("^[A-Za-z0-9]([A-Za-z0-9_-]*[A-Za-z0-9])?$", value)
	if err != nil {
		log.Fatalf("Invalid regexp definition in isAlphanumericOrDash check: %s", err)
	}
	return match
}

func isEmpty(value interface{}) bool {
	switch v := value.(type) {
	case string:
		return value.(string) == ""
	default:
		log.Fatalf("isempty validation supports only string, not %T", v)
	}
	return false
}

// Validate validates given pod definitions
func Validate(pods []Pod) error {
	validate := getValidator()
	for _, pod := range pods {
		err := validate.Struct(pod)
		if err != nil {

			// this check is only needed when your code could produce
			// an invalid value for validation such as interface with nil
			// value most including myself do not usually have code like this.
			if _, ok := err.(*validator.InvalidValidationError); ok {
				return err
			}

			return err.(validator.ValidationErrors)
		}
	}

	return nil
}