package codes

import (
	"net/http"
)

func ToHTTP(code int) int {
	return http.StatusInternalServerError
}
