package search

import (
	"minang-kos-service/constant"
	"minang-kos-service/helper"
	"net/http"
)

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

func BuildPageSizeFromRequest(httpRequest *http.Request) PageSize {
	return PageSize{
		Page: helper.GetPageOrSize(httpRequest, constant.PAGE),
		Size: helper.GetPageOrSize(httpRequest, constant.SIZE),
	}
}
