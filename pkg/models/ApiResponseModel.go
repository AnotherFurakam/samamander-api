package models

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
	PageNumber int
	PageSize   int
	TotalPage  *int
	NextPage   *int
	PrevPage   *int
}
