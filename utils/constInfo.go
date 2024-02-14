package utils

type ErrorType string

const (
	StatusBadRequest ErrorType = "StatusBadRequest"
	StatusNotFound   ErrorType = "StatusNotFound"
)

const (
	DEFAULT_PORT = "8080"
)
