package pkgModel

type PaginationQuery struct {
	PageSize   int `validate:"required,min=5"`
	PageNumber int `validate:"required,min=1"`
}

type PaginationResponse struct {
	PageNumber int
	PageSize   int
	TotalPage  *int
	NextPage   *int
	PrevPage   *int
}
