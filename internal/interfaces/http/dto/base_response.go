package dto

// base response
type Response[T any] struct {
	Code uint16 `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

func New[T any](code uint16, msg string, data T) *Response[T] {
	return &Response[T]{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}
