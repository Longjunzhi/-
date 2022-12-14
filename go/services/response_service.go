package services

type Response struct {
	data    any
	message string
	code    int
}

func GetResponse(data any, message string, code int) (resp *Response) {
	resp = &Response{}
	resp.data = data
	resp.message = message
	resp.code = code
	return resp
}
