package page

import "math"

type Page struct {
	Page      int64 `json:"Page"`
	PageSize  int64 `json:"PageSize"`
	Total     int64 `json:"Total"`
	PrevPage  int64 `json:"PrevPage"`
	NextPage  int64 `json:"NextPage"`
	TotalPage int64 `json:"TotalPage"`
	Limit     int64 `json:"-"`
	Offset    int64 `json:"-"`
}

func Paginator(page, pageSize, total int64) *Page {

	//根据nums总数，和prePage每页数量 生成分页总数
	totalPage := int64(math.Ceil(float64(total) / float64(pageSize))) //page总数
	if page > totalPage {
		page = totalPage
	}
	if page <= 0 {
		page = 1
	}
	prevPage := page - 1
	if prevPage <= 0 {
		prevPage = 1
	}
	nextPage := page + 1
	if nextPage > totalPage {
		nextPage = totalPage
	}
	return &Page{
		Page:      page,
		PageSize:  pageSize,
		Total:     total,
		PrevPage:  prevPage,
		NextPage:  nextPage,
		TotalPage: totalPage,
		Limit:     pageSize,
		Offset:    (page - 1) * pageSize,
	}
}
