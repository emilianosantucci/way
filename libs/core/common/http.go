package common

import (
	"database/sql/driver"
	"encoding/json"
	"strings"
)

type HttpMethod struct {
	slug string
}

func (h *HttpMethod) String() string {
	return h.slug
}

func (h *HttpMethod) Value() (driver.Value, error) {
	return strings.ToUpper(h.slug), nil
}

func (h *HttpMethod) Scan(src interface{}) error {
	var valueString string

	if _, ok := src.([]uint8); ok {
		valueString = string(src.([]uint8))
	} else {
		srcString := src.(string)
		valueString = srcString
	}

	h.slug = strings.ToUpper(valueString)

	return nil
}

func (h *HttpMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(h.slug)
}

func (h *HttpMethod) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	h.slug = strings.ToUpper(s)
	return nil
}

var (
	Unknown = HttpMethod{""}
	Get     = HttpMethod{"GET"}
	Post    = HttpMethod{"POST"}
	Put     = HttpMethod{"PUT"}
	Patch   = HttpMethod{"PATCH"}
	Delete  = HttpMethod{"DELETE"}
)
