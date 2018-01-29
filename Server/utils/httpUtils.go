package utils

import (
	"encoding/json"
	"net/http"
)

// ContentType - ContentType
const ContentType = "Content-Type"

// JSONApplicationContent - JSONApplicationContent
const JSONApplicationContent = "application/json"

// JSONSuccessResponse - Writes success message and json output
func JSONSuccessResponse(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Header().Add(ContentType, JSONApplicationContent)
	json.NewEncoder(w).Encode(data)
}

// JSONBadRequestResponse - JSONErrorResponse
func JSONBadRequestResponse(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusBadRequest)
	if data != nil {
		w.Header().Set(ContentType, JSONApplicationContent)
		json.NewEncoder(w).Encode(data)
	}
}

// JSONUnAuthorizedResponse - JSONUnAuthorizedResponse
func JSONUnAuthorizedResponse(w http.ResponseWriter, msg string) {
	http.Error(w, msg, http.StatusUnauthorized)
}

// JSONInternalServerErrorResponse - JSONInternalServerErrorResponse
func JSONInternalServerErrorResponse(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusInternalServerError)
	if data != nil {
		w.Header().Set(ContentType, JSONApplicationContent)
		json.NewEncoder(w).Encode(data)
	}
}

func FailOnServerError(w http.ResponseWriter, msg string) {
	http.Error(w, msg, http.StatusInternalServerError)
}
