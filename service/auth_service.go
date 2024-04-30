package service

import (
	"context"
	"minang-kos-service/model/web/request"
	"minang-kos-service/model/web/response"
)

type AuthService interface {
	Login(ctx context.Context, request request.AuthLoginRequest) response.AuthLoginResponse
}
