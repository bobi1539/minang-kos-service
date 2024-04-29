package service

import "context"

type CrudService interface {
	Create(ctx context.Context, webRequest any) any
	Update(ctx context.Context, webRequest any) any
	Delete(ctx context.Context, id int64)
	FindById(ctx context.Context, id int64) any
	FindAllWithPagination(ctx context.Context, searchBy map[string]any) any
	FindAllWithoutPagination(ctx context.Context, searchBy map[string]any) any
}
