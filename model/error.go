package model

import (
	"GolangBookApi/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type Error struct {
	ErrorCode int             `json:"errorCode"`
	ErrorType utils.ErrorType `json:"errorType"`
	Message   string          `json:"message"`
}

func (e *Error) GetError(w http.ResponseWriter, errorCode int, errorType utils.ErrorType, errorMessage string) {
	e.ErrorCode = errorCode
	e.ErrorType = errorType
	e.Message = errorMessage

	w.WriteHeader(http.StatusBadRequest)
	err := json.NewEncoder(w).Encode(e)

	if err != nil {
		fmt.Println(err)
	}
}
