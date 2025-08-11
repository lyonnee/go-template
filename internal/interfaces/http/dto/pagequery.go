package dto

type PagequeryReq struct {
	Page     int64  `json:"page" query:"page"`
	PageSize int64  `json:"page_size" query:"page_size"`
	OrderBy  string `json:"order_by" query:"order_by"` // 排序字段
	Order    string `json:"order" query:"order"`       // 是否降序
	Filter   string `json:"filter" query:"filter"`     // 过滤条件
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
