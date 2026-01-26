package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	SUCCESS_CODE = 200

	// 鉴权错误 (10000-19999)
	CODE_NOT_TOKEN              = 10001
	CODE_TOKEN_FORMAT_INCORRECT = 10002
	CODE_TOKEN_INVALID          = 10003

	// 参数错误 (20000-29999)
	CODE_INVALID_QUERY_ARGUMENT = 20001
	CODE_INVALID_PATH_ARGUMENT  = 20002
	CODE_INVALID_BODY_ARGUMENT  = 20003

	// 内部错误 (30000-39999)
	CODE_SERVER_ERROR = 30001
)

// base response
type Response[T any | PagequeryRespData[any]] struct {
	Code uint16 `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data T      `json:"data,omitempty"`
}

// PagequeryRespData 分页查询响应数据
type PagequeryRespData[T any] struct {
	Page      int64 `json:"page"`
	PageSize  int64 `json:"page_size"`
	Total     int64 `json:"total"`
	TotalPage int64 `json:"total_page"` // (Total + PageSize - 1) / PageSize
	Items     T     `json:"items,omitempty"`
}

func NewPagequeryRespData[T any](page, pageSize, total int64, items T) PagequeryRespData[T] {
	return PagequeryRespData[T]{
		Page:      page,
		PageSize:  pageSize,
		Total:     total,
		TotalPage: (total + pageSize - 1) / pageSize,
		Items:     items,
	}
}

func NewResponse[T any | PagequeryRespData[any]](code uint16, msg string, data T) *Response[T] {
	return &Response[T]{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func Ok[T any | PagequeryRespData[any]](c *gin.Context, msg string, data T) {
	resp := NewResponse(SUCCESS_CODE, msg, data)
	c.JSON(
		http.StatusOK,
		resp,
	)
}

func Fail(c *gin.Context, code uint16, msg string) {
	resp := NewResponse(code, msg, "")
	c.JSON(
		http.StatusOK,
		resp,
	)
}
