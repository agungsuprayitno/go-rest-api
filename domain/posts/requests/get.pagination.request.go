package requests

type GetPaginationRequest struct {
	Page    int `json:"page"  binding:"required"`
	PerPage int `json:"perpage"  binding:"required"`
	SortBy  int `json:"sortby"  binding:"required"`
}
