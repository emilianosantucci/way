package common

import (
	"errors"

	"gorm.io/gorm"
)

//goland:noinspection GoUnusedGlobalVariable
var (
	ErrInvalidID                               = errors.New("id not set or not valid uuid v4")
	ErrApplicationNotFound                     = errors.New("application not found")
	ErrApplicationWithSameNameAndVersionExists = errors.New("application with same name and version already exists")
	ErrRestApiResourceNotFound                 = errors.New("rest api resource not found")
	ErrRouteNotFound                           = errors.New("route not found")
)

func GenerateRecordNotFoundError(input error, customErr error) (err error) {
	if err = input; err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = customErr
	}
	return err
}

func GenerateEmptyRowsAffectedError(inputErr error, rowsAffected int, customErr error) (err error) {
	if err = inputErr; err == nil && rowsAffected == 0 {
		err = customErr
	}
	return
}
