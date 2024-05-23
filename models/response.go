package models

type Response struct {
	Status  int         `json:"status"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type Data struct {
	Message string `json:"message"`
}

func NewSuccesResponse(status int, data interface{}) *Response {
	return &Response{
		Status:  status,
		Success: true,
		Data:    data,
	}
}

func NewFailedResponse(status int, data interface{}) *Response {
	return &Response{
		Status:  status,
		Success: false,
		Data:    data,
	}
}
