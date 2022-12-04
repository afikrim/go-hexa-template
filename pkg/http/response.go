package httputil

import "fmt"

const (
	StatusSuccess string = "success"
	StatusFailed  string = "failed"

	InternalServerErrorMsg string = "Internal server error"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (Response) NewResponse(status string) *Response {
	return &Response{
		Status: status,
	}
}

func (r *Response) WithInternalError(message string) *Response {
	if r == nil {
		return nil
	}

	fmt.Printf("There is an internal error: %v", message)
	r.Message = InternalServerErrorMsg
	return r
}

func (r *Response) WithMessage(message string) *Response {
	if r == nil {
		return nil
	}

	r.Message = message
	return r
}

func (r *Response) WithData(data interface{}) *Response {
	if r == nil {
		return nil
	}

	r.Data = data
	return r
}

type MapInterfaceByString map[string]interface{}

func (r *Response) WithDataAsMap(keys []string, values []interface{}) *Response {
	if r == nil {
		return nil
	}

	data := MapInterfaceByString{}
	for i, key := range keys {
		data[key] = values[i]
	}

	r.Data = data
	return r
}
