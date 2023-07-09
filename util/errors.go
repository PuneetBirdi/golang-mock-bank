package util

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type InputError struct {
	index int
	field string
	error string
	message string
}

func ParseError(err error) error {
	var formatted string
	if validationErrs, ok := err.(validator.ValidationErrors); ok {
		for i, e := range validationErrs {
			error := &InputError{
				index: i,
				field: strings.ToLower(e.Field()),
				error: e.Tag(),
				message: "Failure",
			}
			formatted = fmt.Sprintf("%+v", error)
		}
	} else {
		error := &InputError{
			index: 0,
			field: "null",
			error: string(err.Error()),
			message: "None",
		}
		formatted = fmt.Sprintf("%+v", error)
	}

	error := errors.New(string(formatted))
	return error
}

