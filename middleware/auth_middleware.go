package middleware

import (
	"context"
	"minang-kos-service/constant"
	"minang-kos-service/endpoint"
	"minang-kos-service/helper"
	"net/http"
	"os"
	"runtime/debug"
	"strings"
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
	if !isApiKeyValid(request) {
		helper.WriteErrorResponse(writer, http.StatusUnauthorized, constant.UNAUTHORIZED)
		return
	}

	path := request.URL.Path
	if strings.Contains(path, endpoint.AUTH) {
		middleware.Handler.ServeHTTP(writer, request)
		return
	}

	authorizationHeader := request.Header.Get(constant.AUTHORIZATION)
	if !strings.Contains(authorizationHeader, constant.BEARER) {
		helper.WriteErrorResponse(writer, http.StatusUnauthorized, constant.UNAUTHORIZED)
		return
	}

	jwtToken := strings.Replace(authorizationHeader, constant.BEARER, "", -1)
	claims, err := helper.ExtractClaims(jwtToken)
	if err != nil {
		log := helper.GetLogger()
		log.Error(err, string(debug.Stack()))

		helper.WriteErrorResponse(writer, http.StatusUnauthorized, constant.UNAUTHORIZED)
		return
	}

	ctx := context.WithValue(context.Background(), constant.KeyContext(constant.USER_INFO), claims)
	request = request.WithContext(ctx)

	middleware.Handler.ServeHTTP(writer, request)
}

func isApiKeyValid(request *http.Request) bool {
	secretApiKey := os.Getenv("API_KEY")
	apiKey := request.Header.Get(constant.X_API_KEY)
	return apiKey == secretApiKey
}
