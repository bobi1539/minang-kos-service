package helper

import (
	"encoding/base64"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

var logger *logrus.Logger
var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func StringQueryLike(value string) string {
	return "%" + strings.ToLower(value) + "%"
}

func StringToInt(value string) int {
	intValue, err := strconv.Atoi(value)
	PanicIfError(err)
	return intValue
}

func StringToInt64(value string) int64 {
	if len(value) == 0 {
		return 0
	}

	intValue, err := strconv.Atoi(value)
	PanicIfError(err)
	return int64(intValue)
}

func GetSqlOffset(page int, size int) int {
	return (page - 1) * size
}

func GetLogger() *logrus.Logger {
	if logger == nil {
		logger = logrus.New()
	}
	return logger
}

func DecodeBase64(value string) []byte {
	decoded, err := base64.StdEncoding.DecodeString(value)
	PanicIfError(err)
	return decoded
}

func GenerateRandomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func GetUnixNano() int {
	return seededRand.Int()
}
