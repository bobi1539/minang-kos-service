package service

import (
	"context"
	"database/sql"
	"minang-kos-service/exception"
	"minang-kos-service/helper"
	"minang-kos-service/model/web/request"
	"minang-kos-service/model/web/response"
	"minang-kos-service/repository"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type AuthServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewAuthService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) AuthService {
	return &AuthServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *AuthServiceImpl) Login(ctx context.Context, request request.AuthLoginRequest) response.AuthLoginResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	tx := service.beginTransaction()
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByEmail(ctx, tx, request.Email)
	exception.PanicErrorLogin(err)

	validatePassword(user.Password, request.Password)

	jwtToken := helper.GenerateToken(user)
	return helper.ToLoginResponse(jwtToken, helper.GetJwtExpired())
}

func (service *AuthServiceImpl) beginTransaction() *sql.Tx {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	return tx
}

func validatePassword(hashPassword string, password string) {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(password))
	if err != nil {
		exception.PanicErrorLogin(err)
	}
}
