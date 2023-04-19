package helper

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
