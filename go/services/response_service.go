package services

type Response struct {
	Data     any    `json:"data"`
	Message  string `json:"message"`
	Code     int    `json:"code"`
	Current  int    `json:"current"`
	PageSize int    `json:"page_size"`
	Total    int    `json:"total"`
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

func (r *Response) setCurrent(current int) *Response {
	r.Current = current
	return r
}

func (r *Response) setPageSize(pageSize int) *Response {
	r.PageSize = pageSize
	return r
}

func (r *Response) setTotal(total int) *Response {
	r.Total = total
	return r
}
