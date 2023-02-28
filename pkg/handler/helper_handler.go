package handler

import (
	"encoding/json"
	"net/http"
)

func HandleJsonResponse(w http.ResponseWriter, code int, structData interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(structData)
}

func CreateBodyDecoder(r *http.Request) *json.Decoder {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder
}
