package http

import "strings"

func ToString(h HttpMethod) (s string) {
	return strings.ToUpper(h.String())
}

func ToHttpMethod(s string) (method HttpMethod) {
	switch strings.ToLower(s) {
	case "", "all":
		return All
	case "get":
		return Get
	case "post":
		return Post
	case "put":
		return Put
	case "patch":
		return Patch
	case "delete":
		return Delete
	default:
		return Unknown
	}
}
