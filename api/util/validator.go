package util

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func NewValidator() (*Validator, error) {
	validate := validator.New()
	//TODO:customValidateに引っかかった時のエラーメッセージを指定したい
	if err := validate.RegisterValidation("currency", validCurrency); err != nil {
		return nil, err
	}
	return &Validator{
		validator: validate,
	}, nil
}

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return isSupportedCurrency(currency)
	}
	return false
}

func isSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD:
		return true
	}
	return false
}
