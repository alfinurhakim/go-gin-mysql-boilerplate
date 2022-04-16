package helper

import (
	"encoding/json"
	"go-gin-mysql-boilerplate/responses"
	"net/http"
	"strings"
)

func Responses(w http.ResponseWriter, code int, msg string, payload interface{}) {
	var result responses.Responses

	if code >= 400 {
		result.Error = true
		result.Code = code
		result.Message = msg
	} else {
		result.Error = false
		result.Code = code
		if msg != "" {
			result.Message = msg
		} else {
			result.Message = "success"
		}
	}
	result.Data = payload
	response, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func ErrorCustomStatus(w http.ResponseWriter, code int, msg string) {
	Responses(w, code, msg, map[string]string{"error": msg})
}

func ToStringField(field string) string {
	var myField []string
	_ = json.Unmarshal([]byte(field), &myField)
	field = strings.Join(myField, ", ")
	return field
}

func OK(done <-chan bool) bool {
	select {
	case ok := <-done:
		if ok {
			return true
		}
	}
	return false
}
