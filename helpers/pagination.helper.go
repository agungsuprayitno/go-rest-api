package helpers

import (
	"math"
)

type Pagination struct {
	Page         int   `json:"page"`
	PerPage      int   `json:"per_page"`
	TotalPage    int   `json:"total_page"`
	TotalRecord  int64 `json:"total_record"`
	NextPage     bool  `json:"next_page"`
	PreviousPage bool  `json:"previous_page"`
}

func (pagination Pagination) SetPagination(page int, perPage int, totalRecord int64) Pagination {
	totalPage := math.Ceil(float64(totalRecord) / float64(perPage))
	afterPage := int(int(totalPage) - page)
	var isNextPage bool
	var isPrevPage bool

	if afterPage > 0 && totalPage > 1 {
		isNextPage = true
	} else {
		isNextPage = false
	}

	if page > 1 && page <= int(totalPage) {
		isPrevPage = true
	} else {
		isPrevPage = false
	}

	var paginator Pagination

	if page <= int(totalPage) {
		paginator = Pagination{
			Page:         page,
			PerPage:      perPage,
			TotalRecord:  totalRecord,
			TotalPage:    int(totalPage),
			NextPage:     isNextPage,
			PreviousPage: isPrevPage,
		}
	} else {
		paginator = Pagination{}
	}

	return paginator
}
