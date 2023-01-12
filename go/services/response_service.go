package services

type Response struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (r *Response) setData(data any) *Response {
	r.Data = data
	return r
}

func (r *Response) setMessage(message string) *Response {
	r.Message = message
	return r
}

func (r *Response) setCode(code int) *Response {
	r.Code = code
	return r
}
