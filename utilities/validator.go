package utilities

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
	"mime/multipart"
	"net/url"
	"reflect"
	"strings"
	"ubersnap/static"
)

type ValidationErrors struct {
	Field   string
	Message string
}

func messageValidation(field string, tag string, additionalInfo ...interface{}) string {
	switch tag {
	case "required":
		return fmt.Sprintf(static.REQUEST_REQUIRED_MESSAGE, field)
	case "min":
		return fmt.Sprintf(static.REQUEST_MIN_MESSAGE, field, additionalInfo[0])
	case "max":
		return fmt.Sprintf(static.REQUEST_MAX_MESSAGE, field, additionalInfo[0])
	case "email":
		return static.REQUEST_EMAIL_MESSAGE
	}
	return ""
}

func validateStruct(s interface{}) []error {
	v := validate.New(s)
	v.AddMessages(static.VALIDATOR_MESSAGE)
	v.AddValidator("url_ptr", func(val *string) bool {
		_, err := url.ParseRequestURI(*val)
		if err != nil {
			return false
		}
		return true
	})
	var listErr []error
	if !v.Validate() {
		for _, item := range v.Errors {
			for _, i := range item {
				listErr = append(listErr, errors.New(i))
			}
		}
	}

	return listErr
}

func ValidateForm(in interface{}) error {
	validation := validateStruct(in)

	if len(validation) > 0 {
		return ErrorRequest(validation[0], fiber.StatusBadRequest)
	}

	return nil
}

func ValidateFormFile(in interface{}, ctx *fiber.Ctx) error {
	v := reflect.ValueOf(in)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	for i := 0; i < v.NumField(); i++ {
		typeStruct := v.Type().Field(i).Type.String()
		if typeStruct == "*multipart.FileHeader" || typeStruct == "[]*multipart.FileHeader" {
			tagForm := v.Type().Field(i).Tag.Get("form")
			assignFileIntoStruct(ctx, v.Field(i), tagForm)
		}
	}

	return ValidateForm(in)
}

func assignFileIntoStruct(ctx *fiber.Ctx, value reflect.Value, name string) error {
	form, err := ctx.MultipartForm()
	if err != nil {
		return err
	}

	for formName, fileHeaders := range form.File {
		if formName == name {
			if strings.Contains(name, "[]") {
				value.Set(reflect.ValueOf(fileHeaders))
			} else {
				value.Set(reflect.ValueOf(fileHeaders[0]))
			}
		}
	}

	return nil
}

func MultipleFileValidation(data []*multipart.FileHeader, allowType []string) error {
	for _, i := range data {
		if err := ExternalFileValidation(i, allowType); err != nil {
			return err
		}
	}
	return nil
}

func ExternalFileValidation(data *multipart.FileHeader, allowType []string) error {
	mime := data.Header.Get("Content-Type")
	fileName := data.Filename
	if !ContainsInArray(allowType, mime) {
		return errors.New(fmt.Sprintf("file %s not support, %s", fileName, static.ONLY_IMAGE_ALLOWED))
	}
	return nil
}
