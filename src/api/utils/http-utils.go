package utils

import "net/http"

func IsSuccessStatusCode(statusCode int) bool {
	return statusCode == http.StatusOK || statusCode == http.StatusCreated
}

func IsNotFoundStatusCode(statusCode int) bool {
	return statusCode == http.StatusNotFound
}

func IsBadRequestStatusCode(statusCode int) bool {
	return statusCode == http.StatusBadRequest
}
