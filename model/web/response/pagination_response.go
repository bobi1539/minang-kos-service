package response

type PaginationResponse struct {
	Page      int `json:"page"`
	Size      int `json:"size"`
	TotalPage int `json:"totalPage"`
	TotalItem int `json:"totalItem"`
	Data      any `json:"data"`
}
