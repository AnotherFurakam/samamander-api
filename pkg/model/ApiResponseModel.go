package model

type ApiResponse[T any] struct {
	Message    string `json:"message"`
	Data       T      `json:"data"`
	Success    bool   `json:"success"`
	StatusCode int    `json:"statusCode"`
}

type PaginatedApiResponse[T any] struct {
	Message    string `json:"message"`
	Data       T      `json:"data"`
	Success    bool   `json:"success"`
	StatusCode int    `json:"statusCode"`
	PageNumber int    `json:"pageNumber"`
	PageSize   int    `json:"pageSize"`
	TotalPage  *int   `json:"totalPage"`
	NextPage   *int   `json:"nextPage"`
	PrevPage   *int   `json:"prevPage"`
}
