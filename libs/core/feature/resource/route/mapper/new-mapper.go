//go:build !goverter

package mapper

func NewRestDtoMapper() RestDtoMap {
	return &RestDtoMapper{}
}

func NewModelMapper() ModelMap {
	return &ModelMapper{}
}
