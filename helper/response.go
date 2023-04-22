package helper 

import (
	"net/http"
	"strings"
)


type DataResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ResponseFormat(code int, message string, data interface{}) (int, interface{}) {
	res := &DataResponse{
		Code:    code,
		Message: message,
		Data:    data,
	}
	return code, res
}

func SuccessResponse(code int, message string, data ...any) (int, map[string]interface{}) {
	response := make(map[string]interface{})
	response["message"] = message

	switch len(data) {
	case 1:
		response["data"] = data[0]
	case 2:
		response["data"] = data[0]
		response["token"] = data[1]
	}
	return code, response
} 

func ErrorResponse(err error) (int, interface{}) {
	resp := map[string]interface{}{}
	code := http.StatusInternalServerError
	msg := err.Error()

	if msg != "" {
		resp["message"] = msg
	}

	switch true {
	case strings.Contains(msg, "Atoi"):
		resp["message"] = "id must be number, cannot be string"
		code = http.StatusNotFound
	case strings.Contains(msg, "server"):
		code = http.StatusInternalServerError
	case strings.Contains(msg, "format"):
		code = http.StatusBadRequest
	case strings.Contains(msg, "not found"):
		if strings.Contains(msg, "email") {
			resp["message"] = "email not found"
			code = http.StatusNotFound
		}
	case strings.Contains(msg, "access"):
		resp["message"] = "restricted access"
		code = http.StatusInternalServerError
	case strings.Contains(msg, "conflict"):
		code = http.StatusConflict
	case strings.Contains(msg, "Duplicate"):
		if strings.Contains(msg, "username") {
			resp["message"] = "username is already in use"
			code = http.StatusConflict
		} else if strings.Contains(msg, "email") {
			resp["message"] = "email is already in use"
			code = http.StatusConflict
		} else {
			resp["message"] = "Internal server error"
			code = http.StatusInternalServerError
		}
	case strings.Contains(msg, "bad request"):
		code = http.StatusBadRequest
	case strings.Contains(msg, "hashedPassword"):
		resp["message"] = "password do not match"
		code = http.StatusForbidden
	case strings.Contains(msg, "unmarshal"):
		if strings.Contains(msg, "fullname") {
			resp["message"] = "invalid fullname of type string"
			code = http.StatusBadRequest
		} else if strings.Contains(msg, "username") {
			resp["message"] = "invalid username of type string"
			code = http.StatusBadRequest
		} else if strings.Contains(msg, "email") {
			resp["message"] = "invalid email of type string"
			code = http.StatusBadRequest
		} else if strings.Contains(msg, "password") {
			resp["message"] = "invalid password of type string"
			code = http.StatusBadRequest
		}
	case strings.Contains(msg, "upload"):
		code = http.StatusInternalServerError
	}
	return code, resp
}

func ResponseWithData(message string, data any) map[string]any {
	return map[string]any{
		"message": message,
		"data":    data,
	}
}

func Response(message string) map[string]any {
	return map[string]any{
		"message": message,
	}
}
func ResponseFormat(code int, message string, data interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	res["code"] = code
	res["message"] = message
	if data != nil {
		res["data"] = data
	}
	return res
}
