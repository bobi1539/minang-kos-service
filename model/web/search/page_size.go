package search

type PageSize struct {
	Page int
	Size int
}

func BuildPageSize(page int, size int) PageSize {
	return PageSize{
		Page: page,
		Size: size,
	}
}
