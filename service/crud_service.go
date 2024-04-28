package service

import "context"

type CrudService interface {
	Create(ctx context.Context, webRequest any) any
	Update(ctx context.Context, webRequest any) any
}
