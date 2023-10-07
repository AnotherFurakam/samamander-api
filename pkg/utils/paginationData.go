package utils

func CalculatePaginationData(pageNumber int, pageSize int, totalRecords int64) (totalPage *int, nextPage *int, prevPage *int) {

	totalPageValue := int((totalRecords + int64(pageSize) - 1) / int64(pageSize))
	totalPage = &totalPageValue

	nextPageValue := pageNumber + 1
	if nextPageValue > totalPageValue {
		nextPage = nil
	} else {
		nextPage = &nextPageValue
	}

	prevPageValue := pageNumber - 1
	if prevPageValue <= 0 {
		prevPage = nil
	} else {
		prevPage = &prevPageValue
	}

	return totalPage, nextPage, prevPage
}
