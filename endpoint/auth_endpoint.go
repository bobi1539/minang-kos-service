package endpoint

import (
	"database/sql"
	"minang-kos-service/controller"
	"minang-kos-service/service"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

const AUTH = "/api/auth"
const AUTH_LOGIN = AUTH + "/login"

func SetAuthEndpoint(router *httprouter.Router, db *sql.DB, validate *validator.Validate) {
	authController := getAuthController(db, validate)

	router.POST(AUTH_LOGIN, authController.Login)
}

func getAuthController(db *sql.DB, validate *validator.Validate) controller.AuthController {
	authService := service.NewAuthService(getUserRepository(), db, validate)
	return controller.NewAuthController(authService)
}
