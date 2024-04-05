package responses

type PaginationResponse struct {
	Page      int  `json:"page"`
	PerPage   int  `json:"perpage"`
	TotalPage int  `json:"total_page"`
	Next      bool `json:"next"`
	Prev      bool `json:"prev"`
}

func (pagination PaginationResponse) SetPagination(page int, perpage int, totalPage int) PaginationResponse {
	var paginationResponse PaginationResponse
	paginationResponse.Page = page
	paginationResponse.PerPage = page
	paginationResponse.TotalPage = page
	paginationResponse.Next = true
	paginationResponse.Prev = false

	return paginationResponse
}
