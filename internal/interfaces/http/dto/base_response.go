package dto

// base response
type Response[T any | PagequeryData[any]] struct {
	Code uint16 `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data T      `json:"data,omitempty"`
}

func NewResponse[T any](code uint16, msg string, data T) *Response[T] {
	return &Response[T]{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func NewPagequeryResponse[T PagequeryData[any]](code uint16, msg string, page, limit, total int64, list T) *Response[PagequeryData[T]] {
	data := NewPagequeryData(page, limit, total, list)

	return &Response[PagequeryData[T]]{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
