package handler

import (
	"encoding/json"
	"net/http"
	"todo/model"
)

func SendJSONWithStatus(w http.ResponseWriter, v interface{}, status int) {
	w.Header().Set("Content-Type", "application/json; charset=\"utf-8\"")
	w.WriteHeader(status)

	if status == http.StatusNoContent {
		return
	}

	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		panic(err)
	}
}

func SendJSONOk(w http.ResponseWriter, v interface{}) {
	SendJSONWithStatus(w, v, http.StatusOK)
}

func SendJSONError(w http.ResponseWriter, message string, code int) {
	err := &model.Error{Message: message, Code: code}
	SendJSONWithStatus(w, err, code)
}

func ParseJSON(r *http.Request, v interface{}) error {
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(v)
}
