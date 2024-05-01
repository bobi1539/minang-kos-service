package helper

import (
	"errors"
	"minang-kos-service/constant"
	"minang-kos-service/model/domain"
	"minang-kos-service/model/web/dto"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

var JWT_SIGNING_METHOD = jwt.SigningMethodHS256

func GenerateToken(user domain.User) string {
	jwtClaims := JwtClaims(user)
	token := jwt.NewWithClaims(JWT_SIGNING_METHOD, jwtClaims)
	tokenLogin, err := token.SignedString(getJwtSignatureKey())
	PanicIfError(err)
	return tokenLogin
}

func JwtClaims(user domain.User) dto.JwtClaimsDto {
	return dto.JwtClaimsDto{
		StandardClaims: jwt.StandardClaims{
			Issuer:    constant.APPLICATION_NAME,
			ExpiresAt: GetJwtExpired(),
		},
		UserId: user.Id,
		Email:  user.Email,
		Name:   user.Name,
	}
}

func GetJwtExpired() int64 {
	jwtExpiredInHour := StringToInt64(os.Getenv("JWT_EXPIRED_IN_HOUR"))
	return time.Now().Add(time.Duration(jwtExpiredInHour) * time.Hour).Unix()
}

func ExtractClaims(jwtToken string) (jwt.Claims, error) {
	token, err := jwt.Parse(jwtToken, func(t *jwt.Token) (interface{}, error) {
		method, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New(constant.SIGNING_METHOD_INVALID)
		}
		if method != JWT_SIGNING_METHOD {
			return nil, errors.New(constant.SIGNING_METHOD_INVALID)
		}
		return getJwtSignatureKey(), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, errors.New(constant.SIGNING_METHOD_INVALID)
	}
	return claims, nil
}

func getJwtSignatureKey() []byte {
	return []byte(os.Getenv("JWT_SIGNATURE_KEY"))
}
