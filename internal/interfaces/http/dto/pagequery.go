package dto

type PagequeryReq struct {
	Page  int64 `json:"page" query:"page"`
	Limit int64 `json:"limit" query:"limit"`
}

type PagequeryData[T any] struct {
	Page      int64 `json:"page"`
	Limit     int64 `json:"limit"`
	Total     int64 `json:"total"`
	TotalPage int64 `json:"total_page"` // (Total + Limit - 1) / Limit
	List      T     `json:"list,omitempty"`
}

func NewPagequeryData[T any](page, limit, total int64, list T) PagequeryData[T] {
	return PagequeryData[T]{
		Page:      page,
		Limit:     limit,
		Total:     total,
		TotalPage: (total + limit - 1) / limit,
		List:      list,
	}
}
