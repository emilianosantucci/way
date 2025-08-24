package common

import "errors"

//goland:noinspection GoUnusedGlobalVariable
var (
	ErrInvalidID                               = errors.New("id not set or not valid uuid v4")
	ErrApplicationNotFound                     = errors.New("application not found")
	ErrApplicationWithSameNameAndVersionExists = errors.New("application with same name and version already exists")
	ErrRestApiResourceNotFound                 = errors.New("rest api resource not found")
)
