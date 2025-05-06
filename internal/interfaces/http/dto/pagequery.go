package dto

type PagequeryReq struct {
	Page  int64 `json:"page" query:"page"`
	Limit int64 `json:"limit" query:"limit"`
}

type PagequeryResponse[T any] struct {
	Code uint16        `json:"code"`
	Msg  string        `json:"msg"`
	Meta PagequeryMeta `json:"meta"`
	Data T             `json:"data"`
}

type PagequeryMeta struct {
	Page      int64
	Limit     int64
	Total     int64
	TotalPage int64 // (Total + Limit - 1) / Limit
}

func NewPagequeryMeta(page, limit, total int64) PagequeryMeta {
	return PagequeryMeta{
		Page:      page,
		Limit:     limit,
		Total:     total,
		TotalPage: (total + limit - 1) / limit,
	}
}
