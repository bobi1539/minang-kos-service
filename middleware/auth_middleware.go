package middleware

import (
	"minang-kos-service/constant"
	"minang-kos-service/helper"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler,
	}
}

func (middleware *AuthMiddleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	apiKey := "RAHASIA"
	if apiKey != request.Header.Get("X-API-Key") {
		helper.WriteErrorResponse(writer, http.StatusUnauthorized, constant.UNAUTHORIZED)
	} else {
		middleware.Handler.ServeHTTP(writer, request)
	}
}
