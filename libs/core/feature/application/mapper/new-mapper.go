//go:build !goverter

package mapper

func NewDtoMapper() RestDtoMap {
	return &RestDtoMapper{}
}

func NewModelMapper() ModelMap {
	return &ModelMapper{}
}
