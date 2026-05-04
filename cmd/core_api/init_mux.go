package coreapi

import "net/http"

func MakeMux() *http.ServeMux {
	return &http.ServeMux{}
}
