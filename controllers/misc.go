package controllers

import (
	"github.com/asaskevich/govalidator"
)

type (
	CustomValidator struct{}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	ok, err := govalidator.ValidateStruct(i)
	if err != nil || !ok {
		return err
	}
	return nil
}
